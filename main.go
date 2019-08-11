package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	itemsDesc := "(required) TODO items"
	pathDesc := "(optional) file path to save your checklist. Can also be set via the TODO_PATH environment variable."
	titleDesc := "(optional) title for your checklist."
	itemsFlagPtr := flag.String("td", "", itemsDesc)
	pathFlagPtr := flag.String("p", "",pathDesc)
	titleFlagPtr := flag.String("t", "", titleDesc)

	flag.Parse()

	checkInput(os.Args)

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

func printHelp() {
	flag.PrintDefaults()
}

func exitGracefully() {
	os.Exit(0)
}

func checkInput(args []string) {
	if len(args) < 2{
		fmt.Println("No input detected, please include some TODO items")
		printHelp()
		exitGracefully()
	}
}
