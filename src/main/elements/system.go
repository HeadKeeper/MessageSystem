package elements

type System struct {
	systemElements []Executable
	CurrentState string
	currentStates map[string]int
}

func (system System) AddElement(element *Executable) {
	system.systemElements = append(system.systemElements, element)
}

func (system System) executeForeachAtLast() {
	index := len(system.systemElements)
	for i := index; i > 0; i-- {
		system.systemElements[i].Execute()
	}
}

func (system System) saveCurrentState() {
	var resultState string
	for _, element := range system.systemElements {
		resultState += element.GetElementState()
	}
	system.CurrentState = resultState
	// if state exist ...
	if value, key := system.currentStates[resultState]; key {
		value++
	} else {
		value = 1;
	}
}

func (system System) getStatics() {
	// ADD
}

func (system System) SystemTact() {
	system.executeForeachAtLast()
	system.saveCurrentState()
}

// TODO: ADD SYSTEM STATISTICS STRUCT