package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
	"github.com/shopspring/decimal"
)

func allVote(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Casting of all non-voting gold from "+role+" was requested"))
	if role == "group" {
		nonvotingGold, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
		output := allVoteValidate(bot, msg, nonvotingGold, role)
		botSendMsg(bot, msg, output)
		valGrBalance := valGrGetBalance(msg)
		// TO-DO validation needs fixed
		if valGrBalance.balance.nonVoting == "" {
			msgPiece := `non-voting: 0`
			msg.Text = boldText("Validator group lockedGold after voting") + "\n\n" + msgPiece
		}
		msgPiece := `non-voting: ` + valGrBalance.balance.nonVoting
		msg.Text = boldText("Validator group lockedGold after voting") + "\n\n" + msgPiece
	} else if role == "validator" {
		nonvotingGold, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
		output := allVoteValidate(bot, msg, nonvotingGold, role)
		botSendMsg(bot, msg, output)
		valBalance := valGetBalance(msg)
		// TO-DO validation needs fixed
		if valBalance.balance.nonVoting == "" {
			msgPiece := `non-voting: 0`
			msg.Text = boldText("Validator group lockedGold after voting") + "\n\n" + msgPiece
		}
		msgPiece := `non-voting: ` + valBalance.balance.nonVoting
		msg.Text = boldText("Validator lockedGold after voting") + "\n\n" + msgPiece
	}
	return msg.Text
}

func allVoteValidate(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, role string) string {
	nonVoting := cmd.AmountAvailable(target, "nonVoting")
	nonvotingLockedGoldValue, _ := decimal.NewFromString(fmt.Sprintf("%v", nonVoting))
	zeroValue, _ := decimal.NewFromString("0")
	if nonvotingLockedGoldValue.Cmp(zeroValue) == 1 {
		toVote := nonvotingLockedGoldValue.String()
		output := allVoteExecute(bot, msg, toVote, role)
		msg.Text = output
	} else {
		msg.Text = warnText("Don't bite more than you can chew! You only have " + nonvotingLockedGoldValue.String() + " non-voting gold available")
	}
	return msg.Text
}

func allVoteExecute(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
	if role == "group" {
		botSendMsg(bot, msg, boldText("Casting "+amount+" votes from validator group"))
		output, _ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value "+amount, msg)
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
		botSendMsg(bot, msg, boldText("Casting "+amount+" votes from validator"))
		output, _ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value "+amount, msg)
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
