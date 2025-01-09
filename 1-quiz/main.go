package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "problems.csv", "a filename")

	var timeInSeconds int
	flag.IntVar(&timeInSeconds, "time", 30, "time to wait")

	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading records")
	}

	totalQuestions := len(records)
	totalCorrect := 0

	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("Press enter to start timer")
	readUserInput(buf)

	timer := time.NewTimer(time.Duration(timeInSeconds) * time.Second)

	done := make(chan bool)

	go func() {
		<-timer.C
		fmt.Println("\nTime's up! Exiting...")
		fmt.Printf("Score: %d/%d\n", totalCorrect, totalQuestions)
		os.Exit(0)
	}()

	go func() {
		for i, record := range records {
			fmt.Printf("Question #%d: %s = ", i+1, record[0])
			userAnswer := readUserInput(buf)
			if userAnswer == record[1] {
				totalCorrect++
			}
		}
	}()

	<-done
}

func readUserInput(buf *bufio.Reader) string {
	userInput, err := buf.ReadBytes('\n')
	if err != nil {
		log.Fatal("error reading user answer")
	}

	// Remove last \n
	return string(userInput[:len(userInput)-1])
}
