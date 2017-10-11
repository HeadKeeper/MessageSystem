package elements

type Handler struct {
	StayProbability float64
	IsHandling bool
	NextElement *Executable
}
