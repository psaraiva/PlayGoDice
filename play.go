package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const quantity_dice_face = 6

	var qtd int
	dice := loadDice(quantity_dice_face)
	aggregate := loadAggregate(quantity_dice_face)

	fmt.Print("Enter an integer: ")
	fmt.Scan(&qtd)
	for i := 1; i <= qtd; i++ {
		number := playDice(dice)
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

func loadDice(quantity int) []int {
	var dice = make([]int, quantity)
	for i := 0; i < quantity; i++ {
		dice[i] = i + 1
	}

	return dice
}

func loadAggregate(quantity int) []int {
	var dice = make([]int, quantity)
	for i := 0; i < quantity; i++ {
		dice[i] = 0
	}

	return dice
}

func playDice(dice []int) int {
	time.Sleep(500 * time.Millisecond)
	return dice[rand.Intn(len(dice))]
}

func playDiceToText(turn int, value int) string {
	return fmt.Sprintf("[%d] Dice: %d\n", turn, value)
}

func aggregateByIndexToText(face int, value int) string {
	return fmt.Sprintf("Face: %d Qtd: %d\n", face, value)
}
