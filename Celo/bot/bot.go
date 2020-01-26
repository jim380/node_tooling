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
		tgbotapi.NewKeyboardButton("/synced"),
		tgbotapi.NewKeyboardButton("/balance"),
		tgbotapi.NewKeyboardButton("/status"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/score"),
		tgbotapi.NewKeyboardButton("/signing"),
		tgbotapi.NewKeyboardButton("/exchange_rate"),
	),
		tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/lockgold"),
		tgbotapi.NewKeyboardButton("/exchange"),
		tgbotapi.NewKeyboardButton("/vote"),
	),
		tgbotapi.NewKeyboardButtonRow(
		// tgbotapi.NewKeyboardButton("/exchange"),
		// tgbotapi.NewKeyboardButton("/exchange_rate"),
		tgbotapi.NewKeyboardButton("/close"),
	),
)

var lockGoldKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAll"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valAll"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Amount", "valAmount"),
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Group Amount", "valGrAmount"),
	// ),
)

var exchangeUsdKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAllUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valAllUsd"),
	),
)

var electionVoteKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAllVote"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valAllVote"),
	),
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
				case "valGrAllUsd":
					msg.Text = allExchangeUsd(bot, msg, "group")
					break
				case "valAllUsd":
					msg.Text = allExchangeUsd(bot, msg, "validator")
					break
				case "valGrAllVote":
					msg.Text = allVote(bot, msg, "group")
					break
				case "valAllVote":
					msg.Text = allVote(bot, msg, "validator")
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
				commandValGr,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
				commandVal,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
				amountGoldValGr := cmd.AmountAvailable(commandValGr, "gold")
				amountGoldVal := cmd.AmountAvailable(commandVal, "gold")
				msgPiece1 := boldText("Gold Available\n") + "Validator Group: " + fmt.Sprintf("%v", amountGoldValGr) + "\n"
				msgPiece2 := "Validator: " + fmt.Sprintf("%v", amountGoldVal) + "\n"
				msgPiece3 := "\nHow much would you like to lock?"
				msg.Text = msgPiece1 + msgPiece2 + msgPiece3
				msg.ReplyMarkup = lockGoldKeyboard
			case "exchange":
				commandValGr,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
				commandVal,_ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
				amountUsdValGr := cmd.AmountAvailable(commandValGr, "usd")
				amountUsdVal := cmd.AmountAvailable(commandVal, "usd")
				msgPiece1 := boldText("USD Available\n") + "Validator Group: " + fmt.Sprintf("%v", amountUsdValGr) + "\n"
				msgPiece2 := "Validator: " + fmt.Sprintf("%v", amountUsdVal) + "\n"
				msgPiece3 := "\nHow much would you like to exchange?\n"
				msg.Text = msgPiece1 + msgPiece2 + msgPiece3
				msg.ReplyMarkup = exchangeUsdKeyboard
			case "vote":
				commandValGr,_ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
				commandVal,_ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
				amountNonvotingGoldValGr := cmd.AmountAvailable(commandValGr, "nonvotingLockedGold")
				amountNonvotingGoldVal := cmd.AmountAvailable(commandVal, "nonvotingLockedGold")
				if amountNonvotingGoldValGr == nil && amountNonvotingGoldVal == nil{
					msg.Text = "You have no non-voting lockedGold available"
				} else {
					msgPiece1 := boldText("Non-voting Locked Gold Available\n") + "Validator Group: " + fmt.Sprintf("%v", amountNonvotingGoldValGr) + "\n"
					msgPiece2 := "Validator: " + fmt.Sprintf("%v", amountNonvotingGoldVal) + "\n"
					msgPiece3 := "\nHow much would you like to cast?\n"
					msg.Text = msgPiece1 + msgPiece2 + msgPiece3
					msg.ReplyMarkup = electionVoteKeyboard	
				}
			case "signing":
				_,output := botExecCmdOut("celocli validator:signed-blocks --signer $CELO_VALIDATOR_SIGNER_ADDRESS", msg)
				msg.Text = output
			case "exchange_rate":
				msg.Text = getExchangeRate(msg)
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