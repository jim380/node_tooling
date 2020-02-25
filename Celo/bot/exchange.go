package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
	"github.com/shopspring/decimal"
)

func exchangeUSDRun(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string, perct uint) string {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Exchange of all gold from "+role+" was requested"))
	if role == "group" {
		usd, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
		if perct == 100 {
			output := usdToGold(bot, msg, usd, role, 100)
			botSendMsg(bot, msg, output)
		} else if perct == 50 {
			output := usdToGold(bot, msg, usd, role, 50)
			botSendMsg(bot, msg, output)
		} else if perct == 25 {
			output := usdToGold(bot, msg, usd, role, 25)
			botSendMsg(bot, msg, output)
		} else if perct == 75 {
			output := usdToGold(bot, msg, usd, role, 75)
			botSendMsg(bot, msg, output)
		}
		valGrBalance := valGrGetBalance(msg)
		msgPiece := `gold: ` + valGrBalance.balance.gold + "\n" + `usd: ` + valGrBalance.balance.usd
		msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
	} else if role == "validator" {
		usd, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
		if perct == 100 {
			output := usdToGold(bot, msg, usd, role, 100)
			botSendMsg(bot, msg, output)
		} else if perct == 50 {
			output := usdToGold(bot, msg, usd, role, 50)
			botSendMsg(bot, msg, output)
		} else if perct == 25 {
			output := usdToGold(bot, msg, usd, role, 25)
			botSendMsg(bot, msg, output)
		} else if perct == 75 {
			output := usdToGold(bot, msg, usd, role, 75)
			botSendMsg(bot, msg, output)
		}
		valBalance := valGetBalance(msg)
		msgPiece := `gold: ` + valBalance.balance.gold + "\n" + `usd: ` + valBalance.balance.usd
		msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
	}
	return msg.Text
}

func usdToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, role string, perct uint) string {
	amountUsd := cmd.AmountAvailable(target, "usd")
	switch perct {
	case 25:
		usdToGoldValidate(bot, msg, amountUsd, "4", role)
	case 50:
		usdToGoldValidate(bot, msg, amountUsd, "2", role)
	case 75:
		usdToGoldValidate(bot, msg, amountUsd, "1.333333", role)
	case 100:
		usdToGoldValidate(bot, msg, amountUsd, "", role)
	}
	return msg.Text
}

func usdToGoldValidate(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount interface{}, div string, role string) {
	if div == "" {
		amountUsdValue, _ := decimal.NewFromString(fmt.Sprintf("%v", amount))
		zeroValue, _ := decimal.NewFromString("0")
		if amountUsdValue.Cmp(zeroValue) == 1 {
			toExchange := amountUsdValue.String()
			output := usdToGoldExecute(bot, msg, toExchange, role)
			msg.Text = output
		} else {
			msg.Text = warnText("Don't bite more than you can chew! You only have " + amountUsdValue.String() + " usd available")
		}
	} else {
		dividend, _ := decimal.NewFromString(div)
		amountUsdValue, _ := decimal.NewFromString(fmt.Sprintf("%v", amount))
		amountUsdValue = amountUsdValue.DivRound(dividend, 18)
		zeroValue, _ := decimal.NewFromString("0")
		if amountUsdValue.Cmp(zeroValue) == 1 {
			toExchange := amountUsdValue.String()
			output := usdToGoldExecute(bot, msg, toExchange, role)
			msg.Text = output
		} else {
			msg.Text = warnText("Don't bite more than you can chew! You only have " + amountUsdValue.String() + " usd available")
		}
	}
}

func usdToGoldExecute(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
	if role == "group" {
		botSendMsg(bot, msg, boldText("Exchanging "+amount+" usd from validator group"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_ADDRESS --value "+amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		}
	} else if role == "validator" {
		botSendMsg(bot, msg, boldText("Exchanging "+amount+" usd from validator"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value "+amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		}
	}
	return msg.Text
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
