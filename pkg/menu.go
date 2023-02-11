package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Menu 1"),
			tgbotapi.NewKeyboardButton("Menu 2"),
			tgbotapi.NewKeyboardButton("Menu 3"),
			tgbotapi.NewKeyboardButtonLocation("geoloc"),
		),
	)
}
