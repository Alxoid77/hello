package main

import (
	"log"

	pkg "github.com/Alxoid77/hello/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//chatgpt "github.com/golang-infrastructure/go-ChatGPT"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6262169744:AAE1_U3PkQaBDqQkveQ5xzpVXc_AlSSTvDg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			text := update.Message.Text
			log.Printf("[%s] %s", update.Message.From.UserName, text)

			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyMarkup = pkg.GetMainMenuKeyboard()
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
