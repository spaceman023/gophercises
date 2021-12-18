package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var timelimit = flag.Int("timelimit", 15, "The total time for finishing the quiz.")
	flag.Parse()

	limit := *timelimit
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	quizdata := make(map[string]string)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quizdata[record[0]] = record[1]
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Quiz Game")
	fmt.Println("---------------------")
	score := 0
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\n You're out of time! \n")
		fmt.Printf("You scored %v out of %v! \n", score, len(quizdata))
		os.Exit(0)
	}()

	for q, a := range quizdata {
		fmt.Printf("What is %v? \n", q)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(a, text) == 0 {
			fmt.Printf("Great job! \n")
			score++
		} else {
			fmt.Printf("You're a moron! \n")
		}
	}
	fmt.Printf("Great job you scored %v out of %v", score, len(quizdata))
}
