package quiz

import (
	"fmt"
	"log"
	"time"
)

func RunQuiz() {
	input := readFlags("./quiz/problem.csv", 5)
	records := readCsvFile(input.path)
	qs := filterValidQs(records)

	if len(qs) == 0 {
		log.Fatal("No question.")
	}

	count := runQuiz(qs, input.timePerQsInSeconds)
	time.Sleep(time.Second)
	fmt.Printf("Result: %d/%d\n", count, len(qs))
}
