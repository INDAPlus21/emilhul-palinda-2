// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s: hmmm...\n", star)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	go handleQuestions(questions, answers)
	go makePredictions(answers)
	go handleAnswers(answers)
	return questions
}

func handleQuestions(questions chan string, answers chan string) {
	for question := range questions {
		go prophecy(question, answers)
	}
}

func handleAnswers(answers <-chan string) {
	for answer := range answers {
		fmt.Printf("\r%s: ", star)
		for _, char := range answer {
			fmt.Printf("%c", char)
			time.Sleep(time.Duration(rand.Intn(80)) * time.Millisecond)
		}
		fmt.Printf("\n> ")
	}
}

func makePredictions(answers chan<- string) {
	prefixes := []string{
		"I know what you're thinking",
		"Maybe you'd like to know",
		"No I won't tell you",
	}
	questions := []string{
		"What to watch",
		"Where's my refund",
		"How you like that",
		"What is my ip address",
		"How many ounces",
		"What time is it",
		"How I met your mother",
		"How to screenshot on mac",
		"Where am i",
		"How to lose weight fast",
	}
	rambling := []string{
		"hmm... ",
		"maybe... ",
		"maybe... no that's not it ",
		"Eureka!     Wait nono",
		"The meaning of life is.......... 42?",
	}

	for {
		time.Sleep(time.Duration(10+4*rand.Intn(4)) * time.Second)
		if rand.Intn(4) != 0 {
			answers <- prefixes[rand.Intn(len(prefixes))] + "... " + questions[rand.Intn(len(questions))]
		} else {
			answers <- rambling[rand.Intn(len(rambling))]
		}
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
