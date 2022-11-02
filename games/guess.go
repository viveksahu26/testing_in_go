package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the target number
	fmt.Print("I've choosen random number in between 1-100 ")
	target := getTargetNumber()
	fmt.Println("Target: ", target)

	// General Information
	fmt.Println("You will be given total 5 chances to guess exact number...")

	totalChances, remaingChances := 4, 4
	fmt.Println("totalChances: ", totalChances)

	success := false

	for guesses := 0; guesses < 5; guesses++ {
		remaingChances = totalChances - guesses
		fmt.Printf("You have %d guesses left \n ", remaingChances)

		fmt.Print("Guess the Number: ")
		userInput := GuessInput()

		if userInput > target {
			fmt.Println("Oops!! You are FAR")
			continue
		} else if userInput < target {
			fmt.Println("Yeah !! You are CLOSE")
			continue
		} else {
			fmt.Println("GUESSED correctly. Hurray !!!")
			success = true
			break
		}
	}

	if !success {
		fmt.Printf("Sorry, You can not guess my number. My number was %d", target)
	}
}

func getTargetNumber() int {
	seconds := time.Now()
	rand.Seed(seconds.Unix())
	return rand.Intn(100) + 1
}

// take input from user
func GuessInput() int {
	reader := bufio.NewReader(os.Stdin)
	userInputInStringWithNewLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInputInString := strings.TrimSpace(userInputInStringWithNewLine)
	userInput, err := strconv.Atoi(userInputInString)
	if err != nil {
		log.Fatal(err)
	}
	return userInput
}
