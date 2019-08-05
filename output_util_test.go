package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetItems(t *testing.T) {
	testItems := "create a template, add some images, deploy the page"
	expectedOutput := []string{"create a template", " add some images", " deploy the page"}
	assert.Equal(t, len(getItems(testItems)), 3)
	assert.Equal(t, getItems(testItems), expectedOutput)
}

func TestMakeDo(t *testing.T) {
	todo := "Get lunch"
	testOutput := makeDo(todo)
	expectedOutput := "- [ ] Get lunch\n"
	assert.Equal(t, expectedOutput, testOutput)
}

func TestExportMd(t *testing.T) {

}
