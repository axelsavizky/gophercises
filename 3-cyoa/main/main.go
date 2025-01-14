package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	cyoa "gophercises/3-cyoa"
)

func storyHandler(tmpl *template.Template, story cyoa.Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storyArcName := r.URL.Path[1:]
		if storyArcName == "" {
			storyArcName = "intro"
		}
		storyArc, ok := story[storyArcName]
		if !ok {
			http.Error(w, "Could not find story arc", http.StatusBadRequest)
			log.Println("Could not find story arc:", storyArcName)
		}

		storyArc.JoinedStory = strings.Join(storyArc.Story, "")
		// Execute the template with data
		err := tmpl.Execute(w, storyArc)
		if err != nil {
			http.Error(w, "Could not execute template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	}
}

func main() {
	_, b, _, _ := runtime.Caller(0) // Get the path of this file
	basePath := filepath.Dir(b)     // Get the directory of this file

	// Open the JSON file
	file, err := os.Open(filepath.Join(basePath, "../gopher.json"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON into a Story struct
	var story cyoa.Story
	err = json.NewDecoder(file).Decode(&story)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Parse the template file
	tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "../cyoa.html")))

	http.HandleFunc("/", storyHandler(tmpl, story))

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
