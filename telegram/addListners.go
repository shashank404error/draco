package telegram

import (
	"log"

	mongodb "github.com/shashank404error/draco/mongo"
	"go.mongodb.org/mongo-driver/bson"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddListners(apiToken string) {
	bot, _ := newBot(apiToken, false)

	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[TELEGRAM UPDATE]: Added %d to %s", update.Message.Chat.ID, bot.Self.UserName)
		// Store the chat ID along with the TelegramBotUsername
		filter := bson.M{"username": bot.Self.UserName}
		update := bson.M{"$set": bson.M{"username": bot.Self.UserName, "chatID": update.Message.Chat.ID}}
		_, _ = mongodb.Update("telegram-notification", filter, update, true)
	}
}

func newBot(token string, debug bool) (*tg.BotAPI, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Println("[Telegram Bot Error] Failed to get new bot: ", err)
		return bot, err
	}
	bot.Debug = debug
	return bot, err
}
