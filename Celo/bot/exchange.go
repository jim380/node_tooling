package bot

import (
	"fmt"
    "log"
    "github.com/node_tooling/Celo/cmd"
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
    amountUsdValue := amountUsd.(float64)
    if amountUsdValue > 0 {
	    toExchange := fmt.Sprintf("%v", amountUsd)
	    output := usdToGoldAmount(bot, msg, toExchange, role)
        msg.Text = output
    } else {
        msg.Text = warnText("Don't bite more than you can chew! You only have " + fmt.Sprintf("%v", amountUsd) + " usd available")
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
		msg.Text = errText(fmt.Sprintf("%v", outputParsed))
    } else if role == "validator" {
        msg.Text = boldText("Exchanging " + amount + " usd from validator")
        if _, err := bot.Send(msg); err != nil {
	        log.Panic(err)
	    }
	    output,_ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value " + amount, msg)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		msg.Text = errText(fmt.Sprintf("%v", outputParsed))
    }
    return msg.Text
}