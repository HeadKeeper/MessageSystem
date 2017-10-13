package elements

import (
	"math/rand"
	"strconv"
)

// STANDARD HANDLER
type Handler struct {
	StayProbability float64
	IsHandling bool
}

func (handler *Handler) Execute() {
	if handler.IsHandling {
		randomValue := rand.Float64()
		if randomValue > handler.StayProbability {
			handler.IsHandling = false
		}
	}
}

func (handler *Handler) CanAcceptMessage() bool {
	return !handler.IsHandling
}

func (handler *Handler) AcceptMessage() {
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

// HANDLER WITH 'BREAK MESSAGE IF NEXT IS HANDLING' BEHAVIOUR
type BreakHandler struct {
	StayProbability float64
	IsHandling bool
	NextElement Executable
	BreakedMessageCount int
}

func (handler *BreakHandler) Execute() {
	if handler.IsHandling {
		randomValue := rand.Float64()
		if randomValue > handler.StayProbability {
			handler.IsHandling = false
			if handler.NextElement.CanAcceptMessage() {
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

// HANDLER WITH 'LOCK MESSAGE IF NEXT IS HANDLING' BEHAVIOUR
type LockHandler struct {
	// ADD LOCK HANDLER
}
