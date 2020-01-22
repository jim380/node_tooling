package bot

import (
	"log"
    "os"
    "os/exec"
    "fmt"
    "runtime"
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
		tgbotapi.NewKeyboardButton("/empty"),
		tgbotapi.NewKeyboardButton("/close"),
	),
)

var lockGoldKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "all"),
		tgbotapi.NewInlineKeyboardButtonData("specific amount", "lock a specific amount"),
		tgbotapi.NewInlineKeyboardButtonData("back", "go back to main menu"),
	),
	// tgbotapi.NewKeyboardButtonRow(
	// 	tgbotapi.NewKeyboardButton("/lockgold"),
	// 	tgbotapi.NewKeyboardButton("/empty"),
	// 	tgbotapi.NewKeyboardButton("/close"),
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
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
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
            command := botExecCmd("celocli node:synced")
			msg.Text = string(command)
        case "balance":
			valBalance := botExecCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
            valB := botGetValBalance(valBalance)
            valGrBalance := botExecCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
            valGrB := botGetValBalance(valGrBalance)
			msgPiece1 := `*gold*: ` + valGrB.balance.gold + "\n" + `*usd*: ` + valGrB.balance.usd + "\n"
            msgPiece2 := `*gold*: ` + valB.balance.gold + "\n" + `*usd*: ` + valB.balance.usd + "\n"
            msg.Text = "Validator Group\n" + msgPiece1 + "--------------\n" + "Validator\n" + msgPiece2
		case "status":
            command := botExecCmd("celocli validator:status --validator $CELO_VALIDATOR_ADDRESS")
			words := cmd.ParseCmdOutput(command, "string", "(true|false)\\s*(true|false)\\s*(\\d*)\\s*(\\d*.)", 0)
			wordsSplit := strings.Fields(fmt.Sprintf("%v", words))
			ifElected := wordsSplit[0] + "\n"
			ifFrontRunner := wordsSplit[1] + "\n"
			numProposed := wordsSplit[2] + "\n"
			perctSigned := wordsSplit[3] + "\n"
			message := `*Elected*: ` + ifElected + `*Frontrunner*: ` + ifFrontRunner + `*Proposed*: ` + numProposed + `*Signatures*: ` + perctSigned
			msg.Text =  message
		case "score":
			command := botExecCmd("celocli validator:show $CELO_VALIDATOR_ADDRESS")
			words := cmd.ParseCmdOutput(command, "string", "score: (\\d.\\d*)", 1)
			msg.Text = `*Score: *` + fmt.Sprintf("%v", words)
		case "lockgold":
			command := botExecCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
			amountGold := cmd.AmountAvailable(command, "gold")
			msgPiece1 := "You have " + fmt.Sprintf("%v", amountGold) + " gold available.\n"
			msgPiece2 := "How much would you like to lock?\n"
			msg.Text = msgPiece1 + msgPiece2
			msg.ReplyMarkup = lockGoldKeyboard
			// if update.CallbackQuery != nil {
			// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
    		// 	// msg.ReplyToMessageID = update.Message.MessageID
			// 	msg.Text = "Locking all gold"
			// 	bot.Send(msg)
			// }
		case "all":
			msg.Text = "Locking all gold"
		case "specific_amount":
			msg.Text = "Locking a specific amount of gold"
		default:
			msg.Text = "Command not yet supported"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// botExecCmd is a wrapper for os.exec()
func botExecCmd(cmd string) []byte {
	// setEnv()
	if runtime.GOOS == "windows" {
		fmt.Println("You need to switch to Linux, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\nExecuting ", cmdString)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println("\n", fmt.Sprint(err)+": "+string(output))
	} else {
		if string(output) != "" {
			fmt.Println(string(output))
		}
	}
	return output
}

func botGetValBalance(target []byte) Validator{
    gold := cmd.AmountAvailable(target, "gold")
    goldVal := fmt.Sprintf("%v", gold)
    usd := cmd.AmountAvailable(target, "usd")
    usdVal := fmt.Sprintf("%v", usd)
    b := Validator {balance: Balance{gold: goldVal, usd: usdVal},}
    return b
}