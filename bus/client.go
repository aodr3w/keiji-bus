package bus

func (mq *MessageQueue) Push(msg map[string]string) {
	mq.messages <- msg
}

func (mq *MessageQueue) Pull() map[string]string {
	return <-mq.messages
}
