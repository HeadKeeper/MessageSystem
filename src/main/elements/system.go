package elements

import "fmt"

type System struct {
	SystemElements []Executable
	CurrentState string
	CurrentStates map[string]int
}

func (system *System) AddElement(element Executable) {
	system.SystemElements = append(system.SystemElements, element)
}

func (system *System) executeForeachAtLast() {
	index := len(system.SystemElements) - 1
	for i := index; i >= 0; i-- {
		system.SystemElements[i].Execute()
	}
}

func (system *System) saveCurrentState() {
	var resultState string
	for _, element := range system.SystemElements {
		resultState += element.GetElementState()
	}
	system.CurrentState = resultState
	fmt.Println("State: " + resultState)
	// if state exist ...
	if value, key := system.CurrentStates[resultState]; key {
		value++
	} else {
		value = 1;
	}
}

func (system *System) getStatics() {
	// ADD
}

func (system *System) SystemTact() {
	system.saveCurrentState()
	system.executeForeachAtLast()
}

// TODO: ADD SYSTEM STATISTICS STRUCT