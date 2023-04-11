package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type number in, just press ENTER when ready."

func main() {
	// one way - declare, then assign (two steps)
	// var firstNumber int
	// firstNumber = 2

	// another way, declare type and name and assign value
	// var secondNumber = 5

	// one step variable: declare name, assign value and let Go figure out type
	// subtraction := 7

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// var firstNumber = 2
	// var secondNumber = 5
	// var subtraction = 7
	// var answer int

	// rand generates a number between 0 and whatever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and
	// no greater than 10 (multiplying by 1 makes the game a bit silly)
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Gamne")
	fmt.Println("------------------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them through the games
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
	reader.ReadString('\n')
}
