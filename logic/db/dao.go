package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Dao struct {
	transactor *sql.DB
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDao() (*Dao, error) {
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, errors.New("no DB_NAME found in .env file")
	}

	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("no DB_USER found in .env file")
	}

	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return nil, errors.New("no DB_PASSWORD found in .env file")
	}

	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return nil, errors.New("no DB_HOST found in .env file")
	}

	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("no DB_PORT found in .env file")
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")

	return &Dao{transactor: db}, nil
}

func (dao *Dao) GetVkId(userId int) int {
	statement := fmt.Sprintf(`select vk_id from users where id = %d`, userId)

	var vkId int
	dao.transactor.QueryRow(statement).Scan(&vkId)

	return vkId
}

func (dao *Dao) GetTelegramId(userId int) int {
	statement := fmt.Sprintf(`select telegram_id from users where id = %d`, userId)

	var telegramId int
	dao.transactor.QueryRow(statement).Scan(&telegramId)

	return telegramId
}

func (dao *Dao) GetByTelegramId(telegramId int) (int, error) {
	statement := fmt.Sprintf(`select id from users where telegram_id = %d`, telegramId)

	var id int
	err := dao.transactor.QueryRow(statement).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) InsertNewTelegramUser(telegramId int) error {
	statement := fmt.Sprintf(`insert into users (telegram_id) values (%d)`, telegramId)

	_, err := dao.transactor.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}
