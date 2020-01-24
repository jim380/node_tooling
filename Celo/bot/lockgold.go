package bot

import (
    "fmt"
	"strconv"
    "log"
    "github.com/node_tooling/Celo/cmd"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func allLockedGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
	msg.ParseMode = "Markdown"
    // msg.Text = `*Locking all gold from* ` + "*" + role + "*" + ` *was requested*`
	msg.Text = boldText("Locking all gold from " + role + " was requested")
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
    if role == "group" {
        gold, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := lockGold(bot, msg, gold, "all", role)
        msg.Text = output
    } else if role == "validator" {
        gold, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := lockGold(bot, msg, gold, "all", role)
        msg.Text = output
    }
    return msg.Text
}

// lockGold locks a specific amount of gold available
func lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, amount string, role string) string {
	amountGold := cmd.AmountAvailable(target, "gold")
	switch amount {
		case "all":
			// msg.Text = "Locking all " + fmt.Sprintf("%v", amountGold) + "gold has been requested"
			toLock := fmt.Sprintf("%v", amountGold)
			msg.Text = lockGoldAmount(bot, msg, toLock, role)
		// case "amount":
        // TODO 
		// 	break
		default:
			panic("Invalid input")
	}
	return msg.Text
}

func lockGoldAmount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
	toLock, _ := strconv.ParseFloat(amount, 64)
	toLock = toLock - 1000000000000000000
	if toLock <= 0 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved.")
	} else {
		if role == "group" {
			msg.Text = boldText("Locking " + fmt.Sprintf("%f", toLock) + " gold from validator group")
			if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        	}
			output,_ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + fmt.Sprintf("%f", toLock), msg)
			outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		} else if role == "validator" {
			msg.Text = boldText("Locking " + fmt.Sprintf("%f", toLock) + " gold from validator")
			if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        	}
			output,_ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value " + fmt.Sprintf("%f", toLock), msg)
			outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
		}
	}
	return msg.Text
}