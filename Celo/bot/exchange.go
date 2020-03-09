package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
	"github.com/shopspring/decimal"
)

func (v *validatorGr) exchanegUSDToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Exchange USD to Gold from validator group was requested"))
	usdAvailable, _ := decimal.NewFromString(v.usd)
	if perct == 100 {
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 50 {
		dividend, _ := decimal.NewFromString("2")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 25 {
		dividend, _ := decimal.NewFromString("4")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 75 {
		dividend, _ := decimal.NewFromString("1.333333")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	}
}

func (v *validator) exchanegUSDToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Exchange USD to Gold from validator was requested"))
	usdAvailable, _ := decimal.NewFromString(v.usd)
	if perct == 100 {
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 50 {
		dividend, _ := decimal.NewFromString("2")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 25 {
		dividend, _ := decimal.NewFromString("4")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 75 {
		dividend, _ := decimal.NewFromString("1.333333")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = v.exchangeCmdExecute(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	}
}

func (v *validatorGr) exchangeCmdExecute(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toExchange := amount.String()
		botSendMsg(bot, msg, boldText("Exchanging "+toExchange+" usd from validator group"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_ADDRESS --value "+toExchange, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func (v *validator) exchangeCmdExecute(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toExchange := amount.String()
		botSendMsg(bot, msg, boldText("Exchanging "+toExchange+" usd from validator"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value "+toExchange, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func getExchangeRate(msg tgbotapi.MessageConfig) string {
	output, _ := botExecCmdOut("celocli exchange:show", msg)
	cGold := cmd.ParseCmdOutput(output, "string", "(\\d.*) cGLD =>", 1)
	cUsd := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) cUSD", 1)
	cGoldDecimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cGold))
	cUsdDecimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cUsd))
	goldToUsdRatio := cUsdDecimal.DivRound(cGoldDecimal, 18)

	cUsd2 := cmd.ParseCmdOutput(output, "string", "(\\d.*) cUSD =>", 1)
	cGold2 := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) cGLD", 1)
	cUsd2Decimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cUsd2))
	cGold2Decimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cGold2))
	usdToGoldRatio := cGold2Decimal.DivRound(cUsd2Decimal, 18)
	msgPiece1 := "1 cGLD = " + goldToUsdRatio.String() + " cUSD\n"
	msgPiece2 := "1 cUSD = " + usdToGoldRatio.String() + " cGLD"
	return msgPiece1 + msgPiece2
}

func exchangeUSDToGoldRun(p perform, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16) {
	p.exchanegUSDToGold(bot, msg, perct)
}
