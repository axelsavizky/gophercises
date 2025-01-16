package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	linkparser "gophercises/4-linkparser"
)

func main() {
	_, b, _, _ := runtime.Caller(0) // Get the path of this file
	basePath := filepath.Dir(b)     // Get the directory of this file

	// Open the JSON file
	file, err := os.ReadFile(filepath.Join(basePath, "ex4.html"))
	if err != nil {
		panic(err)
	}

	links := linkparser.GetLinksFromPage(file)

	fmt.Println(links)
}
