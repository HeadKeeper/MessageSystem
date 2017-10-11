package elements

import "fmt"

type Generator struct {
	TactsToNewMessage int
	IsBlocked bool
	NextElement *Executable
}

func (generator Generator) Execute()  {
	fmt.Println("LOL")
}

type PGenerator struct {
	// TODO: ADD PROBABILITY GENERATOR
}
