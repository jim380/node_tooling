package bot

import (
    "fmt"
	"strconv"
    "log"
    "github.com/node_tooling/Celo/cmd"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func allLockedGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
    msg.Text = "Locking all gold from " + role + " was requested"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
    if role == "group" {
        gold, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := botLockGold(msg, gold, "all", role)
        msg.Text = output
    } else if role == "validator" {
        gold, balanceOutput := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
        msg.Text = balanceOutput
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
        output := botLockGold(msg, gold, "all", role)
        msg.Text = output
    }
    return msg.Text
}

// lockGold locks a specific amount of gold available
func botLockGold(msg tgbotapi.MessageConfig, target []byte, amount string, role string) string {
	amountGold := cmd.AmountAvailable(target, "gold")
	switch amount {
		case "all":
			// msg.Text = "Locking all " + fmt.Sprintf("%v", amountGold) + "gold has been requested"
			toLock := fmt.Sprintf("%v", amountGold)
			msg.Text = botLockGoldAmount(msg, toLock, role)
		// case "amount":
        // TODO 
		// 	break
		case "back":
			break
		default:
			panic("Invalid input")
	}
	return msg.Text
}

func botLockGoldAmount(msg tgbotapi.MessageConfig, amount string, role string) string {
	toLock, _ := strconv.ParseFloat(amount, 64)
	toLock = toLock - 1000000000000000000
	if toLock <= 0 {
		// fmt.Printf("%v is of the type %T", toLockAfter, toLockAfter)
		msg.Text = "Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved."
	} else {
		if role == "group" {
			// msg.Text = "Locking " + fmt.Sprintf("%f", toLock) + "gold from validator group"
			_, msg.Text = botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + fmt.Sprintf("%f", toLock), msg)
		} else if role == "validator" {
			// msg.Text = "Locking " + fmt.Sprintf("%f", toLock) + "gold from validator"
			_, msg.Text = botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value " + fmt.Sprintf("%f", toLock), msg)
		}
		// ExecuteCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value 10000000000000000000000")
	}
	return msg.Text
}