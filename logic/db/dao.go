package db

import (
	"database/sql"
	"fmt"
)

type Dao struct {
	transactor *sql.DB
}

func (dao *Dao) GetVkId(id int) int {
	statement := fmt.Sprintf(`select vk_id from users where id = %d`, id)

	var vkId int
	dao.transactor.QueryRow(statement).Scan(&vkId)

	return vkId
}

func (dao *Dao) GetTelegramId(id int) int {
	statement := fmt.Sprintf(`select telegram_id from users where id = %d`, id)

	var telegramId int
	dao.transactor.QueryRow(statement).Scan(&telegramId)

	return telegramId
}
