package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var qtd int
	dice := [6]int{1, 2, 3, 4, 5, 6}
	aggregate := [6]int{0, 0, 0, 0, 0, 0}

	fmt.Print("Enter an integer: ")
	fmt.Scan(&qtd)
	for i := 1; i <= qtd; i++ {
		number := playDice(i, dice)
		aggregate[number-1]++
		fmt.Print(playDiceToText(i, number))
	}

	fmt.Println("---------")
	fmt.Println(".:Table:.")
	fmt.Println("---------")
	for idx, val := range aggregate {
		fmt.Print(aggregateByIndexToText(idx+1, val))
	}
}

func playDice(i int, dice [6]int) int {
	time.Sleep(500 * time.Millisecond)
	return dice[rand.Intn(len(dice))]
}

func playDiceToText(i int, number int) string {
	return fmt.Sprintf("[%d] Dice: %d\n", i, number)
}

func aggregateByIndexToText(index int, value int) string {
	return fmt.Sprintf("Index: %d Qtd: %d\n", index, value)
}
