package elements

import (
	"math/rand"
	"strconv"
	"fmt"
)

// STANDARD HANDLER
type Handler struct {
	StayProbability float64
	IsHandling      bool
	HandledMessages int
	ParentSystem    *System
	HandledMessage	Message
}

func (handler *Handler) Execute() {
	if handler.IsHandling {
		randomValue := rand.Float64()
		fmt.Printf("HAND: %f\n",randomValue)
		if randomValue >= handler.StayProbability {
			handler.IsHandling = false
			handler.HandledMessage.endTact = handler.ParentSystem.GetCurrentTact()
			handler.ParentSystem.HandledMessages = append(handler.ParentSystem.HandledMessages, handler.HandledMessage)
			handler.HandledMessages++
		}
	}
}

func (handler *Handler) CanAcceptMessage() bool {
	return !handler.IsHandling
}

func (handler *Handler) AcceptMessage() {
	handler.HandledMessage = handler.ParentSystem.getLastUnhandledMessage()
	handler.IsHandling = true
}

func (handler *Handler) GetElementState() string {
	var state int
	if handler.IsHandling {
		state = 1
	} else {
		state = 0
	}
	return strconv.Itoa(state)
}

func (handler *Handler) GetStatistics() string {
	result := float64(handler.HandledMessages) / TACTS_COUNT
	fmt.Printf("A = %f\n", result)
	return "\n"
}

// HANDLER WITH 'BREAK MESSAGE IF NEXT IS HANDLING' BEHAVIOUR
type BreakHandler struct {
	StayProbability     float64
	IsHandling          bool
	NextElement         Executable
	BreakedMessageCount int
	ParentSystem        *System
	HandledMessage		Message
}

func (handler *BreakHandler) Execute() {
	if handler.IsHandling {
		randomValue := rand.Float64()
		fmt.Println(randomValue)
		if randomValue >= handler.StayProbability {
			handler.IsHandling = false
			if handler.NextElement.CanAcceptMessage() {
				handler.ParentSystem.returnUnhandledMessage(handler.HandledMessage)
				handler.NextElement.AcceptMessage()
			} else {
				handler.BreakedMessageCount++
			}
		}
	}
}

func (handler *BreakHandler) CanAcceptMessage() bool {
	return !handler.IsHandling
}

func (handler *BreakHandler) AcceptMessage() {
	handler.HandledMessage = handler.ParentSystem.getLastUnhandledMessage()
	handler.IsHandling = true
}

func (handler *BreakHandler) GetElementState() string {
	var state int
	if handler.IsHandling {
		state = 1
	} else {
		state = 0
	}
	return strconv.Itoa(state)
}

func (handler *BreakHandler) GetStatistics() string {
	return "\n"
}

// HANDLER WITH 'LOCK MESSAGE IF NEXT IS HANDLING' BEHAVIOUR
type LockHandler struct {
	// ADD LOCK HANDLER
}
