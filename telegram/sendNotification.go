package telegram

import (
	"log"

	mongodb "github.com/shashank404error/draco/mongo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendNotificationToAllListners(apiToken string, body string) {
	bot, _ := newBot(apiToken, false)
	users := mongodb.GetByField("telegram-notification", "username", bot.Self.UserName)

	for _, user := range users {
		log.Printf("[TELEGRAM PUSH] bot: %s body: %s", bot.Self.UserName, body)
		msg := tgbotapi.NewMessage(user["chatID"].(int64), body)
		bot.Send(msg)
	}
}
