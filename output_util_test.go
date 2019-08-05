package main

import (
	"bytes"
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
	testMd := "- [ ] Do it"
	b := &bytes.Buffer{}
	exportMd(testMd, b)
	assert.Equal(t, testMd, string(b.Bytes()))
	assert.CallerInfo()
}

func TestGenerateFileName(t *testing.T) {
	testResult := generateFileName()
	assert.Contains(t, testResult, "checklist_")
	assert.Contains(t, testResult, ".md")
}
