package main

import (
	"main/elements"
)

func main() {
	system := elements.System{}
	handler := elements.Handler{0.5, false}
	breakHandler := elements.BreakHandler{0.55, false, &handler, 0}
	query := elements.Query{2, 0, &breakHandler}
	generator := elements.Generator{2, false, &query}

	system.AddElement(&generator)
	system.AddElement(&query)
	system.AddElement(&breakHandler)
	system.AddElement(&handler)

	for i := 0; i < 1000; i++ {
		system.SystemTact()
	}
}