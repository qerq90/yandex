package sender

type Sender interface {
	send(id int, message string)
}
