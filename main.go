package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// Read quiz.csv file
    // Open file
	if fObj, err := os.Open(fileName); err == nil {
		// Create a new reader
		csvR := csv.NewReader(fObj)
		// Read file contents
		if clines, err := csvR.ReadAll(); err == nil {
            return parseProblem(clines), nil
        } else {
            return nil, fmt.Errorf("failed to read data in csv" + "format from %s file; %s", fileName, err.Error())
        }
	} else {
        return nil, fmt.Errorf("error opening %s file; %s", fileName, err.Error())
    }

	//
}

func main() {
	// Input the file name
	fName := flag.String("f", "quiz.csv", "./quiz.csv")
	// Set duration timer
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	// Pulling the function from the file
	problems, err := problemPuller(*fName)
	// Handling errors
	if err != nil {
		exit(fmt.Sprint("Something went wrong:%s", err.Error()))
	}
	// Create a variable to count our current answers
	correctAns := 0
	// Initailize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	// Loop through the problems and get answers
	// Calculate and print out results
}

func parseProblem(lines [][]string) []problem {
    // Go through the lines with the porblems
    r := make([]problem, len(lines))
    for i := 0 ; i < len(lines) ; i++ {
        r[i] = problem{q: lines[i][0], a: lines[i][1]}
    }
}

type problem struct {
	q string
	a string
}
