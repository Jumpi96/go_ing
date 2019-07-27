package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a correct csv file") // Used to parse CLI args and adds it to the -h of the tool
	flag.Parse()
	csvFile, err := os.Open(*csvFilename)
	if err == nil {
		r := csv.NewReader(csvFile)
		correct := 0
		total := 0

		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Couldn't read the question.")
				os.Exit(1)
			}

			total++
			problem := parseLine(record)

			fmt.Printf("Question: %s Answer: ", problem.question)
			var answer string
			fmt.Scanf("%s\n", &answer)

			if answer == problem.answer {
				correct++
			}
		}
		fmt.Printf("Correct answers: %d. Incorrect answers: %d\n", correct, total-correct)
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

type problem struct {
	question string
	answer   string
}
