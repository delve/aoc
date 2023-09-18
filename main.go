package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"test/pkg/common"
	"time"

	"github.com/sirupsen/logrus"
)

type problem struct {
	q, a string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "CSV file formated question,answer")
	timeLimit := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	common.Check(err)

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	common.Check(err)

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case time := <-timer.C:
			fmt.Println()
			fmt.Printf("%v", time)
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Score %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return ret
}

func exit(msg string) {
	logrus.Info(msg)
	os.Exit(1)
}
