package elements

type Executable interface {
	Execute()
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