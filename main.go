package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"marky/utils"
)

func main() {
	var html []byte
	path := os.Args[1]

	// check if input file is markdown
	if filepath.Ext(path) != ".md" {
		log.Fatal("File must be in Markdown format")
	}

	if contents, err := os.ReadFile(path); err != nil {
		log.Fatal("Error reading file:", err)
	} else {
		html = utils.ParseMarkdown(contents)
	}

	outputPath := strings.Split(path, ".")[0]

	utils.OutputWebpage(outputPath, html)
}
