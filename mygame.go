package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number")


	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(50) 

	var guess int

	var trialcount int = 0

	for{
		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if trialcount > 4{
			fmt.Println("You lose")
			break
		}else{
			if guess > secretNumber {
				fmt.Println("Too Big")
			} else if guess < secretNumber {
				fmt.Println("Too small")
			} else {
				fmt.Println("You win!")
				break
			}
		}
		trialcount += 1
	}
}