package main

import (
	"github.com/shashank404error/draco/mongo"
	"github.com/shashank404error/draco/telegram"
)

var (
	TelegramAPIToken = "Token"
)

func main() {
	mongo.MongoInit()
	telegram.AddListners(TelegramAPIToken)
}
