package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	pkg "github.com/Alxoid77/hello/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//sqlite "github.com/mattn/go-sqlite3"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
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
			coord := ""
			text := ""
			if update.Message.Location == nil {
				text = update.Message.Text
				if text == MENU_ELECTRICITY {
					text = pkg.GetElectricityStatus()
				} else if text == MENU_RAID {
					text = "https://alerts.com.ua/map.png"
				} else if strings.Split(text, "клює") != nil {
					text = "ооооотакого ляща піймав !"
				}
			} else {
				lat := update.Message.Location.Latitude
				lon := update.Message.Location.Longitude
				coord = "https://maps.google.com/maps?q=" + strconv.FormatFloat(lat, 'f', -1, 64) + "," + strconv.FormatFloat(lon, 'f', -1, 64)
			}
			log.Printf("[%s] %s", update.Message.From.UserName, text+coord)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text+coord)
			msg.ReplyMarkup = pkg.GetMainMenuKeyboard()
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		}
	}
}
