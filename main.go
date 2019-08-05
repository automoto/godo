package main

import (
	"flag"
	"os"
)

func main() {
	itemsFlagPtr := flag.String("td", "", "checklist items")
	pathFlagPtr := flag.String("p", "", "file path to save your checklist")
	titleFlagPtr := flag.String("t", "", "title for your checklist")

	flag.Parse()

	parsedTodos := getItems(*itemsFlagPtr)
	outputPath := getOutputPath(*pathFlagPtr)

	outputMd := generateMd(parsedTodos, *titleFlagPtr)
	if len(outputPath) > 1 {
		file, err := os.Create(getFullOutputPath(outputPath))
		handleError(err)
		exportMd(outputMd, file)
	} else {
		printOutput(outputMd)
	}
}


