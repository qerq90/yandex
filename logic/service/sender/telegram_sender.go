package sender

import (
	"qerq90/yandex/logic/client"
	"qerq90/yandex/logic/db"
)

type TelegramSender struct {
	telegramClient *client.TelegramClient
	dao            *db.Dao
}

func MakeTelegramSender(dao *db.Dao, telegram *client.TelegramClient) *TelegramSender {
	return &TelegramSender{telegram, dao}
}

func (sender *TelegramSender) Send(id int, message string) {
	telegramId := sender.dao.GetTelegramId(id)
	sender.telegramClient.SendMessage(telegramId, message)
}
