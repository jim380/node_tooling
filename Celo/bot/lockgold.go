package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
	"github.com/shopspring/decimal"
)

func allLockedGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, role string) string {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Locking all gold from "+role+" was requested"))
	if role == "group" {
		gold, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
		output := lockGold(bot, msg, gold, "all", role)
		botSendMsg(bot, msg, output)
		var valGr validatorGr
		UpdateBalance(&valGr, msg)
		msgPiece := `gold: ` + valGr.balance.gold + "\n" + `lockedGold: ` + valGr.balance.lockedGold
		msg.Text = boldText("Validator Group Balance After Locking") + "\n\n" + msgPiece
	} else if role == "validator" {
		gold, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
		output := lockGold(bot, msg, gold, "all", role)
		botSendMsg(bot, msg, output)
		var val validator
		UpdateBalance(&val, msg)
		msgPiece := `gold: ` + val.balance.gold + "\n" + `lockedGold: ` + val.balance.lockedGold
		msg.Text = boldText("Validator Balance After Locking") + "\n\n" + msgPiece
	}
	return msg.Text
}

// lockGold locks a specific amount of gold available
func lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, target []byte, amount string, role string) string {
	amountGold := cmd.ParseAmount(target, "gold")
	switch amount {
	case "all":
		toLock := fmt.Sprintf("%v", amountGold)
		msg.Text = lockGoldAmount(bot, msg, toLock, role)
	default:
		panic("Invalid input")
	}
	return msg.Text
}

func lockGoldAmount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount string, role string) string {
	toLock, _ := decimal.NewFromString(fmt.Sprintf("%v", amount))
	reserve, _ := decimal.NewFromString("1000000000000000000")
	toLock = toLock.Sub(reserve)
	zeroValue, _ := decimal.NewFromString("0")
	if toLock.Cmp(zeroValue) == -1 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved.")
	} else {
		if role == "group" {
			botSendMsg(bot, msg, boldText("Locking "+toLock.String()+" gold from validator group"))
			output, _ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value "+toLock.String(), msg)
			outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
			if outputParsed == nil {
				msg.Text = successText("Success")
			} else {
				msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			}
		} else if role == "validator" {
			botSendMsg(bot, msg, boldText("Locking "+toLock.String()+" gold from validator"))
			output, _ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value "+toLock.String(), msg)
			outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
			if outputParsed == nil {
				msg.Text = successText("Success")
			} else {
				msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			}
		}
	}
	return msg.Text
}
