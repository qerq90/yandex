package bot

import (
	"database/sql"
	"log"
	"qerq90/yandex/logic/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	api *tgbotapi.BotAPI
	dao *db.Dao
}

func MakeTelegramBot(api *tgbotapi.BotAPI, dao *db.Dao) *TelegramBot {
	return &TelegramBot{api, dao}
}

func (bot *TelegramBot) getUpdates() tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.api.GetUpdatesChan(updateConfig)

	return updates
}

func (bot *TelegramBot) isNewUser(telegramId int) bool {
	_, err := bot.dao.GetByTelegramId(telegramId)
	if err == sql.ErrNoRows {
		return true
	}

	if err != nil {
		log.Default().Println(err)
	}

	return false
}

func (bot *TelegramBot) Run() {
	updates := bot.getUpdates()

	for update := range updates {
		if update.Message == nil && !update.Message.Chat.IsPrivate() {
			continue
		}

		userId := int(update.Message.Chat.ID)

		isNewUser := bot.isNewUser(userId)
		if isNewUser {
			bot.dao.InsertNewTelegramUser(userId)
		}

		msg := tgbotapi.NewMessage(int64(userId), update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.api.Send(msg); err != nil {
			log.Default().Println(err)
		}
	}
}
