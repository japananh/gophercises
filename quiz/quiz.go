package quiz

import (
	"bufio"
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
	path      *string
	timePerQs *int
}

func readFlags() *input {
	path := flag.String("path", "./problem.csv", "csv file path")
	timePerQs := flag.Int("time", 30, "timer per question per second")

	flag.Parse()
	return &input{path: path, timePerQs: timePerQs}
}

func readCsvFile(filePath string) [][]string {
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

	return records
}

func filterQs(qs [][]string) (validQs [][]string) {
	re := regexp.MustCompile(`\d+\+\d+`)

	for _, qs := range qs {
		question := re.FindString(qs[0])
		answer, err := strconv.Atoi(qs[1])
		if err != nil {
			return
		}

		arr := strings.Split(question, "+")
		var sum int
		for _, s := range arr {
			if num, err := strconv.Atoi(s); err == nil {
				sum += num
			}
		}

		if sum == answer {
			validQs = append(validQs, qs)
		}
	}

	return
}

func getAnsFromInput(qs [][]string, timePerQs int) (ans [][2]int) {
	fmt.Printf("The game will start after 3 seconds.\nYou have %d seconds to answer each question.", timePerQs)
	fmt.Print("\n3...")
	time.Sleep(time.Second)
	fmt.Print("2...")
	time.Sleep(time.Second)
	fmt.Println("1...")
	time.Sleep(time.Second)
	fmt.Println("-------------------")

	for i, row := range qs {
		fmt.Printf("-------------------\n* Questions %d: %s\n", i+1, row[0])
		correctAnswer, _ := strconv.Atoi(row[1])

		reader := bufio.NewReader(os.Stdin)

		go func(reader *bufio.Reader, i int) {
			fmt.Println("Press Enter to save the answer.")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if text, _ := strconv.Atoi(text); text == correctAnswer {
				ans = append(ans, [2]int{i, 1})
			} else {
				ans = append(ans, [2]int{i, 0})
			}
			fmt.Printf("Your answer for question %d: %s", i+1, text)
		}(reader, i)

		// stop := make(chan bool, 1)
		// TODO: Press enter to continue loop
		// go func(reader *bufio.Reader, stop chan bool) {
		// 	// reader := bufio.NewReader(os.Stdin)
		// 	for {
		// 		char, _ := reader.ReadByte()
		// 		fmt.Print(char, "------")
		// 		if char == 49 {
		// 			fmt.Println(char, "charrrrr")
		// 			stop <- true
		// 			close(stop)
		// 		}
		// 	}
		// }(reader, stop)

		time.Sleep(time.Second * time.Duration(timePerQs))

		if i <= len(qs)-2 {
			fmt.Println("\nTime end. Next question...")
		}
	}

	fmt.Println("\n------------------\nEnd game.")

	return
}

func countCorrectAns(ans [][2]int) (count int) {
	for _, answer := range ans {
		if answer[1] == 1 {
			count++
		}
	}
	return
}
