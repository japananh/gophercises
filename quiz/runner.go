package quiz

import (
	"errors"
	"fmt"
	"time"
)

func RunQuiz() (err error) {
	input := readFlags("./quiz/problem.csv", 5)
	records := readCsvFile(input.path)
	qs := filterValidQs(records)

	if len(qs) == 0 {
		return errors.New("no question")
	}

	count := runQuiz(qs, input.timePerQsInSeconds)
	time.Sleep(time.Second)
	fmt.Printf("Result: %d/%d\n", count, len(qs))
	return
}
