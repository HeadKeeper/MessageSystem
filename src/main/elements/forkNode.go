package elements

type ForkNode struct {
	FirstFork Executable
	SecondFork Executable
}

func (forkNode *ForkNode) Execute() {
	forkNode.FirstFork.Execute()
	forkNode.SecondFork.Execute()
}

func (forkNode *ForkNode) GetStatistics() string {

	return ""
}

func (forkNode *ForkNode) GetElementState() string {
	var forkState string
	forkState += forkNode.FirstFork.GetElementState()
	forkState += forkNode.SecondFork.GetElementState()
	return forkState
}

func (forkNode *ForkNode) CanAcceptMessage() bool {
	if !(forkNode.FirstFork.CanAcceptMessage()) && !(forkNode.SecondFork.CanAcceptMessage()) {
		return false
	}
	return true
}

func (forkNode *ForkNode) AcceptMessage() {
	if forkNode.FirstFork.CanAcceptMessage() {
		forkNode.FirstFork.AcceptMessage()
	} else if forkNode.SecondFork.CanAcceptMessage() {
		forkNode.SecondFork.AcceptMessage()
	}
}