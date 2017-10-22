package elements

import (
	"fmt"
)

type System struct {
	SystemElements []Executable
	CurrentState string
	CurrentStates map[string]int
	UnhandledMessages []Message
	HandledMessages []Message
	LastHandledMessageIndex int // DEPRECATED
	LastBreakedMessageIndex int // DEPRECATED
	currentTact int
}

/* ADD ELEMENT TO SYSTEM */
func (system *System) AddElement(element Executable) {
	system.SystemElements = append(system.SystemElements, element)
}

/* INITIAL METHOD */
func (system *System) Init() {
	system.CurrentStates = make(map[string]int)
	system.LastBreakedMessageIndex = 1
}

/* EXECUTE ALL SYSTEM ELEMENTS AT LAST */
func (system *System) executeForeachAtLast() {
	index := len(system.SystemElements) - 1
	for i := index; i >= 0; i-- {
		system.SystemElements[i].Execute()
	}
}

/* RETURN MESSAGE ON FIRST POSITION */
func (system *System) returnUnhandledMessage(message Message) {
	var newMessages []Message
	newMessages = append(newMessages, message)
	for _, unhandledMessage := range system.UnhandledMessages {
		newMessages = append(newMessages, unhandledMessage)
	}
	system.UnhandledMessages = newMessages
}

/* GET LAST UNHANDLED MESSAGES FROM SYSTEM QUEUE */
func (system *System) getLastUnhandledMessage() Message {
	result := system.UnhandledMessages[0]
	system.UnhandledMessages = system.UnhandledMessages[1:]
	return result
}

/* BYPASS ALL SYSTEM ELEMENTS, GET THEIR STATE, AND CREATE ALL-SYSTEM STATE */
func (system *System) saveCurrentState() {
	var resultState string
	for _, element := range system.SystemElements {
		resultState += element.GetElementState()
	}
	system.CurrentState = resultState
	fmt.Println("State: " + resultState)
	// if state exist ...
	if _, key := system.CurrentStates[resultState]; key {
		system.CurrentStates[resultState] += 1
	} else {
		system.CurrentStates[resultState] = 1
	}
}

/*
* RETURN -1 IF ALL MESSAGES HANDLED
* DEPRECATED
*/
func (system *System) getLastNonHandledMessageIndex() int {
	/*for i := system.LastHandledMessageIndex + 1; i < len(system.Messages); i++ {
		if system.Messages[i].endTact == 0 {
			return i
		}
	}
	if (len(system.Messages) - 1) == 0 {
		return 0
	}*/
	return -1;
}

func (system *System) GetStatics() {
	for key, element := range system.CurrentStates {
		fmt.Printf("%s : %f\n", key, float64(float64(element) / 100000))
	}
	// ADD
}

/*	RETURN CURRENT SYSTEM TACT */
func (system *System) GetCurrentTact() int {
	return system.currentTact
}

/* EXECUTE ALL SYSTEM ELEMENTS ON CURRENT TACT */
func (system *System) SystemTact() {
	system.currentTact++
	system.saveCurrentState()
	system.executeForeachAtLast()
}

// TODO: ADD SYSTEM STATISTICS STRUCT