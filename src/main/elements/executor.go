package elements

type Executable interface {
	Execute()
	Stateble
}

type Stateble interface {
	GetElementState() string
}
