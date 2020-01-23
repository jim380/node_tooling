package bot

import (
	"log"
    "os"
    "os/exec"
    "fmt"
	"strings"
    "github.com/node_tooling/Celo/cmd"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/balance"),
		tgbotapi.NewKeyboardButton("/synced"),
		tgbotapi.NewKeyboardButton("/status"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/score"),
		tgbotapi.NewKeyboardButton("/lockgold"),
		tgbotapi.NewKeyboardButton("/signing"),
	),
		tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/close"),
		// tgbotapi.NewKeyboardButton("/lockgold"),
		// tgbotapi.NewKeyboardButton("/signing"),
	),
)

var lockGoldKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valAll"),
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAll"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Amount", "valAmount"),
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Group Amount", "valGrAmount"),
	// ),
)

type Balance struct {
    gold           string
    usd            string
}
type Validator struct {
    balance Balance
}
type ValidatorGr struct {
    balance Balance
}

func BotRun() {
    botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore any non-Message Updates
		// if update.Message == nil {
		// 	continue
		// }

		// ignore any non-command Messages
		// if !update.Message.IsCommand() { 
		// 	continue
		// }

		if update.CallbackQuery != nil {
			// msg := update.CallbackQuery.Message
			data := update.CallbackQuery.Data
			chatId := update.CallbackQuery.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatId, "")
			switch data {
				case "valGrAll":
					msg.Text = allLockedGold(bot, msg, "group")
					break
				case "valAll":
					msg.Text = allLockedGold(bot, msg, "validator")
					break
				case "valAmount":
					msg.Text = "Locking a specific amount from validator was requested"
					break
				case "valGrAmount":
					msg.Text = "Locking a specific amount from validator group was requested"
					break
				case "cancel": 
					msg.Text = "Back to back menu"
				break
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue 
		}
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		chatId := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatId, "")
		msg.ParseMode = "Markdown"
		// msg.ReplyToMessageID = update.Message.MessageID
		// bot.Send(msg)
		switch update.Message.Command() {
			case "help":
				msg.Text = "type /balance or /status."
			case "test":
				msg.Text = "I'm ok."
			case "open":
				msg.Text = "What would you like to query?"
				msg.ReplyMarkup = mainKeyboard
			case "close":
				msg.Text = "keyboard closed. Type /open to reopen"
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case "synced":
				command,_ := botExecCmdOut("celocli node:synced", msg)
				msg.Text = string(command)
			case "balance":
				valBalance,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
				valB := botGetValBalance(valBalance)
				valGrBalance,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
				valGrB := botGetValBalance(valGrBalance)
				msgPiece1 := `*gold*: ` + valGrB.balance.gold + "\n" + `*usd*: ` + valGrB.balance.usd + "\n"
				msgPiece2 := `*gold*: ` + valB.balance.gold + "\n" + `*usd*: ` + valB.balance.usd + "\n"
				msg.Text = "Validator Group\n" + msgPiece1 + "--------------\n" + "Validator\n" + msgPiece2
			case "status":
				command,_ := botExecCmdOut("celocli validator:status --validator $CELO_VALIDATOR_ADDRESS", msg)
				words := cmd.ParseCmdOutput(command, "string", "(true|false)\\s*(true|false)\\s*(\\d*)\\s*(\\d*.)", 0)
				wordsSplit := strings.Fields(fmt.Sprintf("%v", words))
				ifElected := wordsSplit[0] + "\n"
				ifFrontRunner := wordsSplit[1] + "\n"
				numProposed := wordsSplit[2] + "\n"
				perctSigned := wordsSplit[3] + "\n"
				message := `*Elected*: ` + ifElected + `*Frontrunner*: ` + ifFrontRunner + `*Proposed*: ` + numProposed + `*Signatures*: ` + perctSigned
				msg.Text =  message
			case "score":
				command,_ := botExecCmdOut("celocli validator:show $CELO_VALIDATOR_ADDRESS", msg)
				words := cmd.ParseCmdOutput(command, "string", "score: (\\d.\\d*)", 1)
				msg.Text = `*Score: *` + fmt.Sprintf("%v", words)
			case "lockgold":
				command,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
				amountGold := cmd.AmountAvailable(command, "gold")
				msgPiece1 := "You have " + fmt.Sprintf("%v", amountGold) + " gold available.\n"
				msgPiece2 := "How much would you like to lock?\n"
				msg.Text = msgPiece1 + msgPiece2
				msg.ReplyMarkup = lockGoldKeyboard
			case "signing":
				_,output := botExecCmdOut("celocli validator:signed-blocks --signer $CELO_VALIDATOR_SIGNER_ADDRESS", msg)
				msg.Text = output
			default:
				msg.Text = "Command not yet supported"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// botExecCmdOut executes commands and returns command outputs
func botExecCmdOut(cmd string, msg tgbotapi.MessageConfig) ([]byte, string) {
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		msg.Text = fmt.Sprint(err)+": "+string(output)
	} else {
		if string(output) != "" {
			msg.Text = string(output)
		}
	}
	return output, msg.Text
}

func botGetValBalance(target []byte) Validator{
    gold := cmd.AmountAvailable(target, "gold")
    goldVal := fmt.Sprintf("%v", gold)
    usd := cmd.AmountAvailable(target, "usd")
    usdVal := fmt.Sprintf("%v", usd)
    b := Validator {balance: Balance{gold: goldVal, usd: usdVal},}
    return b
}