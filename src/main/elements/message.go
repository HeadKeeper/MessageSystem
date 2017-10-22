package elements

type Message struct {
	startTact int
	endTact int
}

func (message *Message) GetMessageInSystemTime() int {
	return message.endTact - message.startTact
}
