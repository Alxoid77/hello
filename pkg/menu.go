package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			//tgbotapi.NewKeyboardButton("голосовалка"),
			tgbotapi.NewKeyboardButton("є світло?"),
			tgbotapi.NewKeyboardButton("тривоги"),
			tgbotapi.NewKeyboardButtonLocation("де я?"),
		),
	)
}
