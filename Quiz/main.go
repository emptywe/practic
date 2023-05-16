package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const waitTime = time.Second * 30

// readQuizFile - reading csv file by given filepath, return ordered questions map
func readQuizFile(path string) (map[int][]string, error) {
	var (
		questions = make(map[int][]string)
		order     int
	)

	quizFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer quizFile.Close()

	csvReader := csv.NewReader(quizFile)

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		questions[order] = record
		order++
	}

	return questions, nil
}

// processInput - wait users input from Stdin and sends it to given string channel
func processInput(answerChan chan string) {
	var input string

	if _, err := fmt.Fscan(os.Stdin, &input); err != nil {
		fmt.Println(err)
		return
	}

	answerChan <- input
}

// processQuestion - send question to Stdout and waits for users answer return 1 with correct answer and 0 with wrong answer or timeout
func processQuestion(withTimer *bool, questionNum int, question []string, answerChan chan string) int {
	var (
		answer string
		timer  = new(time.Timer)
	)

	if *withTimer {
		timer = time.NewTimer(waitTime)
	}

	fmt.Fprintf(os.Stdout, "Question %d: %s: ", questionNum, question[0])

	go processInput(answerChan)

	select {
	case <-timer.C:
		fmt.Fprintln(os.Stdout)
		return 0
	case answer = <-answerChan:
		break
	}

	if strings.TrimSpace(strings.ToLower(question[len(question)-1])) == strings.TrimSpace(strings.ToLower(answer)) {
		return 1
	}

	return 0
}

func main() {
	var (
		rightAnswers int
		quizFilePath string
		answerChan   = make(chan string)
	)

	flag.StringVar(&quizFilePath, "src", "problems.csv", "specifies the quiz filename")
	withTimer := flag.Bool("timer", false, "specifies if 30 seconds timer is on")
	withShuffle := flag.Bool("shuffle", false, "specifies if questions start in random order")
	flag.Parse()

	questions, err := readQuizFile("./source/" + quizFilePath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Problem reading quiz file: %v", err)
		return
	}

	if *withShuffle {
		counter := 0
		for _, line := range questions {
			rightAnswers += processQuestion(withTimer, counter+1, line, answerChan)
			counter++
		}
	} else {
		for i := 0; i < len(questions); i++ {
			rightAnswers += processQuestion(withTimer, i+1, questions[i], answerChan)
		}
	}

	fmt.Fprintf(os.Stdout, "Right Answers: %d, Wrong Answers: %d\n", rightAnswers, len(questions)-rightAnswers)

}
