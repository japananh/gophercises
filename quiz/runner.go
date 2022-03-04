package quiz

import (
	"fmt"
	"log"
	"time"
)

func RunQuiz() {
	input := readFlags()
	path := input.path
	timePerQs := input.timePerQs

	records := readCsvFile(*path)
	qs := filterQs(records)

	if len(qs) == 0 {
		log.Fatal("Invalid question.")
	}

	ans := getAnsFromInput(qs, *timePerQs)
	count := countCorrectAns(ans)

	time.Sleep(time.Duration(*timePerQs))

	fmt.Println("Total questions are", len(qs))
	fmt.Println("\nThe number of correct questions are", count)
}
