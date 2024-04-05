package utils

import (
	"fmt"
	"log"
	"marky/utils/files"
	"os"

	"github.com/pkg/browser"
)

func OutputWebpage(path string, html []byte) {
	createDirectory(path)
	createIndexHTML(path, html)
	createCSS(path)

	fmt.Println("Success!")
	browser.OpenURL("file://" + path + "/index.html")
}

func createDirectory(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Println("Directory already exists...")
	}
}

func createIndexHTML(path string, html []byte) {
	formattedHTML := files.FormatHTML(html, "Test")
	htmlPath := path + "/index.html"

	// create 'index.html'
	if htmlFile, err := os.Create(htmlPath); err != nil {
		fmt.Println("Index.html already exists...")
	} else {
		defer htmlFile.Close()
	}

	// write to 'index.html'
	if err := os.WriteFile(htmlPath, formattedHTML, 0666); err != nil {
		log.Fatal("Error writing to 'index.html':", err)
	}
}

func createCSS(path string) {
	css := files.GetCSS()
	cssPath := path + "/nano.css"

	// create 'nano.css'
	if htmlFile, err := os.Create(cssPath); err != nil {
		fmt.Println("Index.html already exists...")
	} else {
		defer htmlFile.Close()
	}

	// write to 'nano.css'
	if err := os.WriteFile(cssPath, css, 0666); err != nil {
		log.Fatal("Error writing to 'index.html':", err)
	}
}
