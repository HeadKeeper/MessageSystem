package elements

import "fmt"

type Query struct {
	Size int
	MessagesInQuery int
	NextElement *Executable
}

func (query Query) Execute() {
	fmt.Println("AHAHAHA")
}