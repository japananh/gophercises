package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type input struct {
	path               string
	timePerQsInSeconds int
}

type question struct {
	text       string
	correctAns int
}

func readFlags(defaultPath string, defaultTimeInSeconds int) *input {
	path := flag.String("path", defaultPath, "csv file path")
	timePerQsInSeconds := flag.Int("time", defaultTimeInSeconds, "time per question in second")
	flag.Parse()
	return &input{path: *path, timePerQsInSeconds: *timePerQsInSeconds}
}

func readCsvFile(filePath string) (result []question) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, row := range records {
		correctAns, _ := strconv.Atoi(row[1])
		result = append(result, question{row[0], correctAns})
	}

	return
}

func filterValidQs(qs []question) (validQs []question) {
	re := regexp.MustCompile(`\d+\+\d+`)

	for _, qs := range qs {
		question := re.FindString(qs.text)
		if question == "" {
			break
		}

		arr := strings.Split(question, "+")
		sum := 0
		for _, str := range arr {
			if num, err := strconv.Atoi(str); err == nil {
				sum += num
			}
		}

		if sum == qs.correctAns {
			validQs = append(validQs, qs)
		}
	}

	return
}

func runQuiz(qs []question, seconds int) (correctAnsCount int) {
	fmt.Printf("The game will start after 3 seconds.\nYou have %d seconds to answer each question.", seconds)
	fmt.Print("\n3...")
	time.Sleep(time.Second)
	fmt.Print("2...")
	time.Sleep(time.Second)
	fmt.Println("1...")
	time.Sleep(time.Second)
	fmt.Println()

	userInput := make(chan int)
	go readInput(userInput)

	for i, row := range qs {
		fmt.Printf("%d - Question: %s\n", i+1, row.text)
		fmt.Print("Press Enter to save the answer, otherwise your answer won't be recorded.\n-> ")

		select {
		case userAnswer := <-userInput:
			if userAnswer == row.correctAns {
				fmt.Println("Good jobs! Correct answer is", userAnswer)
				correctAnsCount++
			} else {
				fmt.Println("Wrong answer!")
			}
			fmt.Println()
		case <-time.After(time.Duration(seconds) * time.Second):
			fmt.Println("\nTime is over!")
		}
	}

	return
}

func readInput(input chan<- int) {
	for {
		var u int
		if _, err := fmt.Scanf("%d\n", &u); err != nil {
			fmt.Print("Oops. Invalid input. Please try again.\n-> ")
		} else {
			input <- u
		}
	}
}
