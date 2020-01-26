package bot

import (
	"fmt"
    "log"
    // "strconv"
    "github.com/node_tooling/Celo/cmd"
    "github.com/shopspring/decimal"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func allExchangeUsd(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
    msg.ParseMode = "Markdown"
    msg.Text = boldText("Exchange of all gold from " + role + " was requested")
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
    if role == "group" {
        usd, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := usdToGold(bot, msg, usd, role)
        msg.Text = output
    } else if role == "validator" {
        usd, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := usdToGold(bot, msg, usd, role)
        msg.Text = output
    }
    return msg.Text
}

func usdToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, role string) string {
    amountUsd := cmd.AmountAvailable(target, "usd")
    // msg.Text = "Exchange of all usd from " + role + " was requested"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
    amountUsdValue, _ := decimal.NewFromString(fmt.Sprintf("%v", amountUsd))
    zeroValue, _ := decimal.NewFromString("0")
    if amountUsdValue.Cmp(zeroValue) == 1 {
	    toExchange := amountUsdValue.String()
	    output := usdToGoldAmount(bot, msg, toExchange, role)
        msg.Text = output
    } else {
        msg.Text = warnText("Don't bite more than you can chew! You only have " + amountUsdValue.String() + " usd available")
	}

    return msg.Text
}

func usdToGoldAmount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
	//toExchange, _ := strconv.Atoi(amount)
    if role == "group" {
        msg.Text = boldText("Exchanging " + amount + " usd from validator group")
        if _, err := bot.Send(msg); err != nil {
	        log.Panic(err)
	    }
        output,_ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
        if outputParsed == nil {
            msg.Text = successText("Success")
        } else {
            msg.Text = errText(fmt.Sprintf("%v", outputParsed))
        }
    } else if role == "validator" {
        msg.Text = boldText("Exchanging " + amount + " usd from validator")
        if _, err := bot.Send(msg); err != nil {
	        log.Panic(err)
	    }
	    output,_ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value " + amount, msg)
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
    output,_ := botExecCmdOut("celocli exchange:show", msg)
    // msg.Text = string(output)
    // if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	cGold := cmd.ParseCmdOutput(output, "string", "(\\d.*) cGLD =>", 1)
    cUsd := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) cUSD", 1)
    cGoldDecimal,_ := decimal.NewFromString(fmt.Sprintf("%v", cGold))
    cUsdDecimal,_ := decimal.NewFromString(fmt.Sprintf("%v", cUsd))
    goldToUsdRatio := cUsdDecimal.DivRound(cGoldDecimal, 18)

    cUsd2 := cmd.ParseCmdOutput(output, "string", "(\\d.*) cUSD =>", 1)
    cGold2 := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) cGLD", 1)
    cUsd2Decimal,_ := decimal.NewFromString(fmt.Sprintf("%v", cUsd2))
    cGold2Decimal,_ := decimal.NewFromString(fmt.Sprintf("%v", cGold2))
    usdToGoldRatio := cGold2Decimal.DivRound(cUsd2Decimal, 18)
    msgPiece1 := "1 cGLD = " + goldToUsdRatio.String() + " cUSD\n" 
    msgPiece2 := "1 cUSD = " + usdToGoldRatio.String() + " cGLD"
    // msgPiece3 := boldText("\n\nResults are truncated")
    return msgPiece1 + msgPiece2
}