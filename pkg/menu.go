package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("голосовалка"),
			tgbotapi.NewKeyboardButton("Menu 2"),
			tgbotapi.NewKeyboardButton("коли днюха ?"),
			tgbotapi.NewKeyboardButtonLocation("де я?"),
		),
	)
}
