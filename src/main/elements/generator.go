package elements

import (
	"strconv"
	"fmt"
)

type Generator struct {
	TactsToNewMessage int
	IsBlocked bool
	NextElement Executable
}

func (generator *Generator) Execute() {
	if !generator.IsBlocked {
		if generator.TactsToNewMessage > 0 {
			generator.TactsToNewMessage--
		}
		if generator.TactsToNewMessage == 0 {
			if generator.NextElement.CanAcceptMessage() {
				generator.resetGenerator()
				generator.NextElement.AcceptMessage() // SEND MESSAGE
			} else {
				generator.lockGenerator()
				return
			}
		} else {
			if generator.NextElement.CanAcceptMessage() {
				// NOOP?
			}
		}
	} else if generator.NextElement.CanAcceptMessage() {
		generator.resetGenerator()
	}
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
	fmt.Println("LOCK!!!")
	generator.IsBlocked = true
}

func (generator *Generator) resetGenerator() {
	generator.TactsToNewMessage = 2
	generator.IsBlocked = false
}

type PGenerator struct {
	// TODO: ADD PROBABILITY GENERATOR
}