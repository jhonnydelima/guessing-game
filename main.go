package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A random number will be generated. Try to guess it. The number is an integer between 0 and 100.")
	number := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}
	for i := range attempts {
		fmt.Println("What's your guess?")
		scanner.Scan()
		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)
		attemptInt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("Your guess must to be an integer number.")
			return
		}
		switch {
		case attemptInt < number:
			fmt.Println("You loose! The number is greater than", attemptInt)
		case attemptInt > number:
			fmt.Println("You loose! The number is less than", attemptInt)
		case attemptInt == number:
			fmt.Printf(
				"Good job! The number was: %d\n"+
					"You've got it in %d attempts\n"+
					"These were your attempts: %v\n",
				number, i+1, attempts[:i],
			)
			return
		}
		attempts[i] = attemptInt
	}
	fmt.Printf(
		"You didn't get the number =(\n"+
			"The number was: %d\n"+
			"These were your attempts: %v\n",
		number, attempts,
	)
}
