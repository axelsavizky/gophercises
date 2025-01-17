package main

import (
	"path/filepath"
	"runtime"

	"gophercises/7-taskcli/cmd"
	"gophercises/7-taskcli/db"
)

func main() {
	_, b, _, _ := runtime.Caller(0) // Get the path of this file
	basePath := filepath.Dir(b)     // Get the directory of this file

	dbPath := filepath.Join(basePath, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	err = cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
