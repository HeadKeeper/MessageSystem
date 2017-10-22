package elements

import (
	"strconv"
	"fmt"
)

type Query struct {
	Size            int
	MessagesInQuery int
	Messages        []int
	NextElement     Executable
	Index           int
}

func (query *Query) Execute() {
	if (query.NextElement.CanAcceptMessage()) && (query.MessagesInQuery > 0) {
		query.MessagesInQuery--
		query.NextElement.AcceptMessage()
	}
}

func (query *Query) GetElementState() string {
	query.Messages[query.Index] = query.MessagesInQuery
	query.Index++
	return strconv.Itoa(query.MessagesInQuery)
}

func (query *Query) AcceptMessage() {
	query.MessagesInQuery++
	if query.NextElement.CanAcceptMessage() {
		query.NextElement.AcceptMessage()
		query.MessagesInQuery--
	}
}

func (query *Query) CanAcceptMessage() bool {
	if query.MessagesInQuery == query.Size {
		return false
	} else {
		return true
	}
}

func (query *Query) GetStatistics() string {
	var allMessagesInQuery int
	var index int
	for i := 0; i < 100000; i++ {
		if (query.Messages[i] != 0) {
			allMessagesInQuery += query.Messages[i]
			index++
		}
	}
	fmt.Printf("MEDIUM QUEUE LENGTH (L) = %f\n", float64(allMessagesInQuery) / float64(100000));
	return "/n"
}