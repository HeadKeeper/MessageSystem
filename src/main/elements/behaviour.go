package elements

type Executable interface {
	Execute()
	GetStatistics() string
	Stateble
	MessageAcceptable
}

type Stateble interface {
	GetElementState() string
}

type MessageAcceptable interface {
	CanAcceptMessage() bool
	AcceptMessage()
}