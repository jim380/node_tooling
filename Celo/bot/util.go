package bot

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
)
func boldText(str string) string {
    return "*" + str + "*"
}

func warnText(str string) string {
    return "\xE2\x9A\xA0 " + str
}

func errText(str string) string {
    return "\xE2\x9D\x8C " + str
}

func successText(str string) string {
    return "\xE2\x9C\x94 " + str
}

func ifNil(str string) string {
    var result string
    if str == "" {
        result = "0"
    } else {
        result = str
    }
    return result
}

func botSendMsg(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, msgTxt string) {
    msg.Text = msgTxt
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}