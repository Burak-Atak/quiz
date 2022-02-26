package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var score = 0

// Takes input and returns it
func take_input() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return name
}

// Asks question and if answer is true increase the score
func ask_question(r [][]string) int {
	for _, i := range r {

		fmt.Printf("What is the answer of %v operation?: ", i[0])
		a := take_input()
		if strings.ReplaceAll(a, " ", "") == i[1] {
			score += 1
		}
	}
	return score
}

// reads csv file, shuffles the question list, when time is up prints last score
func main() {

	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println(err)
	}
	r, _ := csv.NewReader(file).ReadAll()
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	fmt.Println(r)
	go ask_question(r)

	time.Sleep(30 * time.Second)

	fmt.Printf("\nYour %d answer were correct out of %d questions.", score, len(r))
}
