package core

type Message map[string]string

type MessageQueue struct {
	messages chan Message
	cap      int
}

func newMessage(cmd string, taskID string) Message {
	msg := make(map[string]string)
	msg["cmd"] = cmd
	msg["taskID"] = taskID
	return msg
}
func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		messages: make(chan Message, size),
		cap:      size,
	}
}

func (mq *MessageQueue) Push(msg map[string]string) {
	mq.messages <- msg
}

func (mq *MessageQueue) Pull() map[string]string {
	return <-mq.messages
}
