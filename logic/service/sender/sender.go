package sender

type Sender interface {
	Send(id int, message string)
}
