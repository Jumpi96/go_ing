package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a correct csv file") // Used to parse CLI args and adds it to the -h of the
	duration := flag.Int("duration", 30, "time in seconds to complete the test")
	random := flag.Bool("random", false, "add a random order to the problems")
	flag.Parse()

	csvFile, err := os.Open(*csvFilename)
	if err == nil {
		problems := readProblems(csvFile)
		if *random {
			shuffle(problems)
		}

		correct := 0
		fmt.Printf("Press enter to start the quiz. You'll have %d seconds to complete it.", *duration)
		fmt.Scanf("%s")
		timer := time.NewTimer(time.Duration(*duration) * time.Second)

		for _, problem := range problems {
			fmt.Printf("Question: %s Answer: ", problem.question)
			answerCh := make(chan string)
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerCh <- answer
			}()
			select {
			case <-timer.C:
				fmt.Printf("\nYour time has ended. You've completed %d out of %d questions.\n", correct, len(problems))
				return
			case answer := <-answerCh:
				if answer == problem.answer {
					correct++
				}
			}
		}
		fmt.Printf("You've completed %d out of %d questions.\n", correct, len(problems))
	} else {
		fmt.Printf("Couldn't open the csv file.")
	}
}

func parseLine(line []string) problem {
	return problem{
		question: line[0],
		answer:   strings.TrimSpace(line[1]),
	}
}

func shuffle(array []problem) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		j := r1.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func readProblems(csvFile *os.File) []problem {
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Couldn't read the problems.")
		os.Exit(1)
	}
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = parseLine(record)
	}
	return problems
}

type problem struct {
	question string
	answer   string
}
