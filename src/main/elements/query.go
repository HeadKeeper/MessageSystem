package elements

import (
	"strconv"
)

type Query struct {
	Size int
	MessagesInQuery int
	NextElement Executable
}

func (query *Query) Execute() {
	if (query.NextElement.CanAcceptMessage()) && (query.MessagesInQuery > 0) {
		query.MessagesInQuery--
		query.NextElement.AcceptMessage()
	}
}

func (query *Query) GetElementState() string {
	return strconv.Itoa(query.MessagesInQuery)
}

func (query *Query) AcceptMessage() {
	query.MessagesInQuery++
	if query.NextElement.CanAcceptMessage() {
		query.MessagesInQuery--
		query.NextElement.AcceptMessage()
	}
}

func (query *Query) CanAcceptMessage() bool {
	if query.MessagesInQuery == query.Size {
		return false
	} else {
		return true
	}
}