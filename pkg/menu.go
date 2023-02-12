package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	MENU_POLL        = "голосовалка"
	MENU_ELECTRICITY = "є світло?"
	MENU_RAID        = "тривоги"
	MENU_GEOLOC      = "де я?"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			//tgbotapi.NewKeyboardButton(MENU_POLL),
			tgbotapi.NewKeyboardButton(MENU_ELECTRICITY),
			tgbotapi.NewKeyboardButton(MENU_RAID),
			tgbotapi.NewKeyboardButtonLocation(MENU_GEOLOC),
		),
	)
}
