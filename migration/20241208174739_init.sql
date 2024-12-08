-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
	id serial PRIMARY KEY,
	vk_id int,
	telegram_id int
);
-- +goose StatementEnd
