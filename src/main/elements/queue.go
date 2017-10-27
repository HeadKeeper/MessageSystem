package elements

import (
	"strconv"
	"fmt"
)

type Queue struct {
	Size            int
	MessagesInQueue int
	Messages        []int
	NextElement     Executable
	Index           int
}

func (query *Queue) Execute() {
	if (query.NextElement.CanAcceptMessage()) && (query.MessagesInQueue > 0) {
		query.MessagesInQueue--
		query.NextElement.AcceptMessage()
	}
}

func (query *Queue) GetElementState() string {
	query.Messages[query.Index] = query.MessagesInQueue
	query.Index++
	return strconv.Itoa(query.MessagesInQueue)
}

func (query *Queue) AcceptMessage() {
	query.MessagesInQueue++
	if query.NextElement.CanAcceptMessage() {
		query.NextElement.AcceptMessage()
		query.MessagesInQueue--
	}
}

func (query *Queue) CanAcceptMessage() bool {
	if query.MessagesInQueue == query.Size {
		return false
	} else {
		return true
	}
}

func (query *Queue) GetStatistics() string {
	var allMessagesInQueue int
	var index int
	for i := 0; i < TACTS_COUNT; i++ {
		if (query.Messages[i] != 0) {
			allMessagesInQueue += query.Messages[i]
			index++
		}
	}
	fmt.Printf("MEDIUM QUEUE LENGTH (L) = %f\n", float64(allMessagesInQueue) / float64(TACTS_COUNT));
	return "/n"
}