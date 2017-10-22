package elements

import (
	"strconv"
)

const (
	TACTS_TO_NEW_MESSAGE = 2
)

type Generator struct {
	TactsToNewMessage int
	IsBlocked         bool
	NextElement       Executable
	ParentSystem      *System
}

func (generator *Generator) Execute() {
	if !generator.IsBlocked {
		if generator.TactsToNewMessage > 0 {
			generator.TactsToNewMessage--
		}
		if generator.TactsToNewMessage == 0 {
			if generator.NextElement.CanAcceptMessage() {
				generator.resetGenerator()
				generator.ParentSystem.UnhandledMessages = append(generator.ParentSystem.UnhandledMessages, Message{
					generator.ParentSystem.GetCurrentTact(),
					0,
				})
				generator.NextElement.AcceptMessage() // SEND MESSAGE
			} else {
				generator.lockGenerator()
				return
			}
		} else {
			if generator.NextElement.CanAcceptMessage() {
				// NOOP
				return
			}
		}
	} else if generator.NextElement.CanAcceptMessage() {
		generator.resetGenerator()
		generator.ParentSystem.UnhandledMessages = append(generator.ParentSystem.UnhandledMessages, Message{
			generator.ParentSystem.GetCurrentTact(),
			0,
		})
		generator.NextElement.AcceptMessage()
	}
}

func (generator *Generator) GetStatistics() string {
	return "\n"
}

func (generator *Generator) GetElementState() string {
	return strconv.Itoa(generator.TactsToNewMessage)
}

func (generator *Generator) CanAcceptMessage() bool {
	return false
}

func (generator *Generator) AcceptMessage() {
	// NOOP
	return
}

func (generator *Generator) lockGenerator() {
	generator.IsBlocked = true
}

func (generator *Generator) resetGenerator() {
	generator.TactsToNewMessage = TACTS_TO_NEW_MESSAGE
	generator.IsBlocked = false
}

type PGenerator struct {
	// TODO: ADD PROBABILITY GENERATOR
}