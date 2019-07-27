package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	csvfile, err := os.Open("problems.csv")
	if err == nil {
		r := csv.NewReader(csvfile)
		correct := 0
		incorrect := 0

		for {
			record, err := r.Read()
			
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Couldn't read the question.")
			}

			fmt.Printf("Question: %s Answer: ", record[0])

			var answer int
			fmt.Scanf("%d", &answer)
			
			correct_answer, _ := strconv.Atoi(record[1])

			if answer == correct_answer {
				correct = correct + 1
			} else {
				incorrect = incorrect + 1
			}
		}
		fmt.Printf("Correct answers: %s. Incorrect answers: %s\n", strconv.Itoa(correct), strconv.Itoa(incorrect))
	} else {
		fmt.Printf("Couldn't open the csv file.")
	}
}
