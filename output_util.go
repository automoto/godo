package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func handleError(e error) {
	if e!= nil {
		panic(e)
	}
}

func getItems(s string) []string{
	parsed := strings.Split(s, ",")
	return parsed
}

func getOutputPath(p string) string{
	if p != ""{
		return p
	}
	return os.Getenv("TODO_PATH")
}

func getFullOutputPath(p string) string {
	return fmt.Sprintf("%s/%s", p, generateFileName())
}

func generateFileName() string{
	now := time.Now()
	return fmt.Sprintf("checklist_%d_%d_%d_%d_%d_%d.md", now.Day(), now.Month(), now.Year(), now.Hour(),
		now.Minute(), now.Second())
}

func makeDo(todo string) string{
	return fmt.Sprintf("- [ ] %s\n", strings.Trim(todo, " "))
}

func generateMd(parsedTodos []string, title string) string{
	var output string

	if len(title) > 1{
		output += fmt.Sprintf("**%s**\n", title)
	}
	for _, todo := range parsedTodos {
		output += fmt.Sprintf("%s", makeDo(todo))
	}
	return output
}

func printOutput(todos string) {
	fmt.Println(todos)
}

func exportMd(todos string, file io.Writer) {
	writtenBytes, err := fmt.Fprint(file, todos)
	handleError(err)
	log.Printf("Generated Markdown File Succesfully of Size: %d bytes\n", writtenBytes)
}
