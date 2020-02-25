package bot

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
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
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 25%", "valGrOneForthUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 50%", "valGrHalfUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 75%", "valGrFourThirdsUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAllUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator 25%", "valOneFourthUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator 50%", "valHalfUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator 75%", "valFourThirdsUsd"),
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
	gold       string
	usd        string
	lockedGold string
	nonVoting  string
	total      string
}
type Validator struct {
	balance Balance
}
type ValidatorGr struct {
	balance Balance
}

func Run() {
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
			chatID := update.CallbackQuery.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatID, "")
			msg.ParseMode = "Markdown"
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
				msg.Text = allExchangeUsd(bot, msg, "group", 100)
				break
			case "valGrHalfUsd":
				msg.Text = allExchangeUsd(bot, msg, "group", 50)
				break
			case "valGrOneForthUsd":
				msg.Text = allExchangeUsd(bot, msg, "group", 25)
				break
			case "valGrFourThirdsUsd":
				msg.Text = allExchangeUsd(bot, msg, "group", 75)
				break
			case "valAllUsd":
				msg.Text = allExchangeUsd(bot, msg, "validator", 100)
				break
			case "valHalfUsd":
				msg.Text = allExchangeUsd(bot, msg, "validator", 50)
				break
			case "valOneForthUsd":
				msg.Text = allExchangeUsd(bot, msg, "validator", 25)
				break
			case "valFourThirdsUsd":
				msg.Text = allExchangeUsd(bot, msg, "validator", 75)
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
		chatID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatID, "")
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
			command, _ := botExecCmdOut("celocli node:synced", msg)
			msg.Text = string(command)
		case "balance":
			valGrBalance := valGrGetBalance(msg)
			valBalance := valGetBalance(msg)
			msgPiece1 := `*gold*: ` + valGrBalance.balance.gold + "\n" + `*lockedGold*: ` + valGrBalance.balance.lockedGold + "\n" + `*usd*: ` + valGrBalance.balance.usd + "\n" + `*non-voting*: ` + valGrBalance.balance.nonVoting + "\n" + `*total*: ` + valGrBalance.balance.total + "\n"
			msgPiece2 := `*gold*: ` + valBalance.balance.gold + "\n" + `*lockedGold*: ` + valBalance.balance.lockedGold + "\n" + `*usd*: ` + valBalance.balance.usd + "\n" + `*non-voting*: ` + valBalance.balance.nonVoting + "\n" + `*total*: ` + valBalance.balance.total + "\n"
			msg.Text = "Validator Group\n\n" + msgPiece1 + "--------------\n" + "Validator\n\n" + msgPiece2
		case "status":
			command, _ := botExecCmdOut("celocli validator:status --validator $CELO_VALIDATOR_ADDRESS", msg)
			words := cmd.ParseCmdOutput(command, "string", "(true|false)\\s*(true|false)\\s*(\\d*)\\s*(\\d*.)", 0)
			wordsSplit := strings.Fields(fmt.Sprintf("%v", words))
			ifElected := wordsSplit[0] + "\n"
			ifFrontRunner := wordsSplit[1] + "\n"
			numProposed := wordsSplit[2] + "\n"
			perctSigned := wordsSplit[3] + "\n"
			message := `*Elected*: ` + ifElected + `*Frontrunner*: ` + ifFrontRunner + `*Proposed*: ` + numProposed + `*Signatures*: ` + perctSigned
			msg.Text = message
		case "score":
			command, _ := botExecCmdOut("celocli validator:show $CELO_VALIDATOR_ADDRESS", msg)
			words := cmd.ParseCmdOutput(command, "string", "score: (\\d.\\d*)", 1)
			msg.Text = `*Score: *` + fmt.Sprintf("%v", words)
		case "lockgold":
			valGrBalance := valGrGetBalance(msg)
			valBalance := valGetBalance(msg)
			msgPiece1 := boldText("Gold Available\n") + "Validator Group: " + valGrBalance.balance.gold + "\n"
			msgPiece2 := "Validator: " + valBalance.balance.gold + "\n"
			msgPiece3 := "\nHow much would you like to lock?"
			msg.Text = msgPiece1 + msgPiece2 + msgPiece3
			msg.ReplyMarkup = lockGoldKeyboard
		case "exchange":
			valGrBalance := valGrGetBalance(msg)
			valBalance := valGetBalance(msg)
			msgPiece1 := boldText("USD Available\n") + "Validator Group: " + valGrBalance.balance.usd + "\n"
			msgPiece2 := "Validator: " + valBalance.balance.usd + "\n"
			msgPiece3 := "\nHow much would you like to exchange?\n"
			msg.Text = msgPiece1 + msgPiece2 + msgPiece3
			msg.ReplyMarkup = exchangeUsdKeyboard
		case "vote":
			valGrBalance := valGrGetBalance(msg)
			valBalance := valGetBalance(msg)
			if valGrBalance.balance.nonVoting == "" && valBalance.balance.nonVoting == "" {
				msg.Text = "You have no non-voting lockedGold available"
			} else {
				msgPiece1 := boldText("Non-voting Locked Gold Available\n") + "Validator Group: " + valGrBalance.balance.nonVoting + "\n"
				msgPiece2 := "Validator: " + valBalance.balance.nonVoting + "\n"
				msgPiece3 := "\nHow much would you like to cast?\n"
				msg.Text = msgPiece1 + msgPiece2 + msgPiece3
				msg.ReplyMarkup = electionVoteKeyboard
			}
		case "signing":
			_, output := botExecCmdOut("celocli validator:signed-blocks --signer $CELO_VALIDATOR_SIGNER_ADDRESS", msg)
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
		msg.Text = fmt.Sprint(err) + ": " + string(output)
	} else {
		if string(output) != "" {
			msg.Text = string(output)
		}
	}
	return output, msg.Text
}

func valGrGetBalance(msg tgbotapi.MessageConfig) ValidatorGr {
	var valGr ValidatorGr
	return valGr.getBalance(msg)
}

func valGetBalance(msg tgbotapi.MessageConfig) Validator {
	var val Validator
	return val.getBalance(msg)
}

func (v *Validator) getBalance(msg tgbotapi.MessageConfig) Validator {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
	gold := cmd.AmountAvailable(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.AmountAvailable(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.AmountAvailable(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.AmountAvailable(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.AmountAvailable(target, "total")
	totalVal := isZero(total, "totalVal")
	res := Validator{balance: Balance{gold: goldVal, usd: usdVal, lockedGold: lockedGoldVal, nonVoting: nonVotingLockedGoldVal, total: totalVal}}
	return res
}

func (vgr *ValidatorGr) getBalance(msg tgbotapi.MessageConfig) ValidatorGr {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
	// TO-DO extract the logic for checking if zero
	gold := cmd.AmountAvailable(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.AmountAvailable(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.AmountAvailable(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.AmountAvailable(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.AmountAvailable(target, "total")
	totalVal := isZero(total, "totalVal")
	res := ValidatorGr{balance: Balance{gold: goldVal, usd: usdVal, lockedGold: lockedGoldVal, nonVoting: nonVotingLockedGoldVal, total: totalVal}}
	return res
}
