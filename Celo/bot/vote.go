package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
	"github.com/shopspring/decimal"
)

func (v *validatorGr) vote(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Casting of all non-voting gold from validator group was requested"))
	nonvotingGoldAvailable, _ := decimal.NewFromString(v.nonVoting)
	zeroValue, _ := decimal.NewFromString("0")
	if nonvotingGoldAvailable.Cmp(zeroValue) == 1 {
		toVote := nonvotingGoldAvailable.String()
		botSendMsg(bot, msg, boldText("Casting "+toVote+" votes from validator group"))
		output, _ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value "+toVote, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		}
		botSendMsg(bot, msg, msg.Text)
	} else {
		msg.Text = warnText("Don't bite more than you can chew! You only have " + nonvotingGoldAvailable.String() + " non-voting gold available")
		botSendMsg(bot, msg, msg.Text)
	}
}

func (v *validator) vote(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Casting of all non-voting gold from validator was requested"))
	nonvotingGoldAvailable, _ := decimal.NewFromString(v.nonVoting)
	zeroValue, _ := decimal.NewFromString("0")
	if nonvotingGoldAvailable.Cmp(zeroValue) == 1 {
		toVote := nonvotingGoldAvailable.String()
		botSendMsg(bot, msg, boldText("Casting "+toVote+" votes from validator"))
		output, _ := botExecCmdOut("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value "+toVote, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		}
		botSendMsg(bot, msg, msg.Text)
	} else {
		msg.Text = warnText("Don't bite more than you can chew! You only have " + nonvotingGoldAvailable.String() + " non-voting gold available")
		botSendMsg(bot, msg, msg.Text)
	}
}

func voteRun(p perform, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	p.vote(bot, msg)
}
