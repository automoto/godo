package main

import (
	"flag"
)

func main() {
	itemsFlagPtr := flag.String("td", "", "checklist items")
	pathFlagPtr := flag.String("p", "", "file path to save your checklist")
	titleFlagPtr := flag.String("t", "", "title for your checklist")

	flag.Parse()

	parsedTodos := getItems(*itemsFlagPtr)
	outputPath := getOutputPath(*pathFlagPtr)

	outputMd := generateMd(parsedTodos, *titleFlagPtr)
	exportMd(outputMd, outputPath)
}


