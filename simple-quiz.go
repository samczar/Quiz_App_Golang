package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//You can put this before the main function
type csvLine struct {
	Questions string
	Answers   string
}

var score int

func processQuiz(lines [][]string) int {
	var total int
	for _, line := range lines {
		total++
		data := csvLine{
			Questions: line[0],
			Answers:   line[1],
		}
		fmt.Println(data.Questions)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter correct option: ")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error occured with answer input", err)
		}
		processResult(response, data.Answers)
	}
	return total
}

func increment() {
	score = score + 1
}

func processResult(response string, answer string) {
	res := strings.ToUpper(strings.TrimSpace(response))
	ans := strings.ToUpper(strings.TrimSpace(answer))
	if res == ans {
		score++
	}
}

func main() {

	filename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	var totalScore int
	var totalScorePercentage int

	lines := processFile(filename)
	totalScore = processQuiz(lines)
	totalScorePercentage = (score / totalScore) * (100)
	fmt.Println("Score: ", totalScorePercentage, "%")

	/*Save score to a folder */

	f, err := os.OpenFile("scoreslog.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(strconv.Itoa(totalScorePercentage) + ",\n"); err != nil {
		panic(err)
	}

	/*  Compare user score with current score */

}

func processFile(filename *string) [][]string {
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
