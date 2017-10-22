package main

import (
	"main/elements"
	"fmt"
)

func main() {
	system := elements.System{}
	system.Init()
	messages := make([]int, 100000)
	noopMessage := elements.Message{}
	handler := elements.Handler{0.55, false, 0, &system, noopMessage}
	breakHandler := elements.BreakHandler{0.5, false, &handler, 0, &system, noopMessage}
	query := elements.Query{2, 0, messages,&breakHandler, 0}
	generator := elements.Generator{2, false, &query, &system}

	system.AddElement(&generator)
	system.AddElement(&query)
	system.AddElement(&breakHandler)
	system.AddElement(&handler)

	for i := 0; i < 100000; i++ {
		system.SystemTact()
	}
	system.GetStatics()
	for _, element := range system.SystemElements {
		element.GetStatistics()
	}
	var timeInSystem int
	for _, element := range system.HandledMessages {
		timeInSystem += element.GetMessageInSystemTime()
	}
	fmt.Printf("Wc = %f", float64(timeInSystem) / float64(len(system.HandledMessages)))
}