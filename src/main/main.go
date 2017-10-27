package main

import (
	"main/elements"
)

func main() {
	system := elements.System{}
	system.Init()
	messages := make([]int, elements.TACTS_COUNT)
	noopMessage := elements.Message{}
	handler := elements.Handler{0.5, false, 0, &system, noopMessage}
	breakHandler := elements.BreakHandler{0.48, false, &handler, 0, &system, noopMessage}
	query := elements.Queue{2, 0, messages,&breakHandler, 0}
	generator := elements.Generator{2, false, &query, &system}

	system.AddElement(&generator)
	system.AddElement(&query)
	system.AddElement(&breakHandler)
	system.AddElement(&handler)

	for i := 0; i < elements.TACTS_COUNT; i++ {
		system.SystemTact()
	}
	system.GetStatics()
	for _, element := range system.SystemElements {
		element.GetStatistics()
	}
}