package bot

import (
	"fmt"
    "log"
    "github.com/node_tooling/Celo/cmd"
    "github.com/shopspring/decimal"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func allVote(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
    msg.ParseMode = "Markdown"
    msg.Text = boldText("Casting of all non-voting gold from " + role + " was requested")
    if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
    if role == "group" {
        nonvotingGold, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
        // msg.Text = voteOutput
        // if _, err := bot.Send(msg); err != nil {
        //     log.Panic(err)
        // }
        output := allVoteValidate(bot, msg, nonvotingGold, role)
        msg.Text = output
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        target,_ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
		voteUpdate := botUpdateBalance(target)
		msgPiece := `total: ` + voteUpdate.balance.total + "\n" + `non-voting: ` + voteUpdate.balance.nonVoting
        msg.Text = boldText("Validator group lockedGold after voting") + "\n\n" + msgPiece
    } else if role == "validator" {
        nonvotingGold, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
        // msg.Text = voteOutput
        // if _, err := bot.Send(msg); err != nil {
        //     log.Panic(err)
        // }
        output := allVoteValidate(bot, msg, nonvotingGold, role)
        msg.Text = output
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        target,_ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
		voteUpdate := botUpdateBalance(target)
		msgPiece := `total: ` + voteUpdate.balance.total + "\n" + `non-voting: ` + voteUpdate.balance.nonVoting
        msg.Text = boldText("Validator lockedGold after voting") + "\n\n" + msgPiece
    }    
    return msg.Text
}

func allVoteValidate(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, role string) string {
    nonvotingLockedGold := cmd.AmountAvailable(target, "nonvotingLockedGold")
    nonvotingLockedGoldValue, _ := decimal.NewFromString(fmt.Sprintf("%v", nonvotingLockedGold))
    zeroValue, _ := decimal.NewFromString("0")
    if nonvotingLockedGoldValue.Cmp(zeroValue) == 1 {
	    toVote := nonvotingLockedGoldValue.String()
        // 
	    output := allVoteExecution(bot, msg, toVote, role)
        msg.Text = output
    } else {
        msg.Text = warnText("Don't bite more than you can chew! You only have " + nonvotingLockedGoldValue.String() + " non-voting gold available")
	}
    return msg.Text
}

func allVoteExecution(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
    if role == "group" {
        msg.Text = boldText("Casting " + amount + " votes from validator group")
        if _, err := bot.Send(msg); err != nil {
	        log.Panic(err)
	    }
        // --- Display parsed success/fail output --- //
        output,_ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
        if outputParsed == nil {
            msg.Text = successText("Success")
        } else {
		    msg.Text = errText(fmt.Sprintf("%v", outputParsed))
        }
        // _,output := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount, msg)
		// outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		// msg.Text = output
    } else if role == "validator" {
        msg.Text = boldText("Casting " + amount + " votes from validator")
        if _, err := bot.Send(msg); err != nil {
	        log.Panic(err)
	    }
        output,_ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
        if outputParsed == nil {
            msg.Text = successText("Success")
        } else {
		    msg.Text = errText(fmt.Sprintf("%v", outputParsed))
        }
        // _,output := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount, msg)
		// outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		// msg.Text = output
    }
    return msg.Text
}