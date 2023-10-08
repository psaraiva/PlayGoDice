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

	start := time.Now()
	ch := make(chan int)
	for i := 1; i <= qtd; i++ {
		go playDiceChannel(dice, ch)
	}

	for i := 1; i <= qtd; i++ {
		number := <-ch
		aggregate[number-1]++
	}

	fmt.Println("---------")
	fmt.Println(".:Table:.")
	fmt.Println("---------")
	for idx, val := range aggregate {
		fmt.Print(aggregateByIndexToText(idx+1, val))
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func loadDice(quantity int) []int {
	var dice = make([]int, quantity)
	for i := 0; i < quantity; i++ {
		dice[i] = i + 1
	}

	return dice
}

func loadAggregate(quantity int) []int {
	return make([]int, quantity)
}

func playDiceChannel(dice []int, ch chan int) {
	time.Sleep(500 * time.Millisecond)
	number := dice[rand.Intn(len(dice))]
	fmt.Print(playDiceToText(number))
	ch <- dice[rand.Intn(len(dice))]
}

func playDiceToText(value int) string {
	return fmt.Sprintf("Dice: %d\n", value)
}

func aggregateByIndexToText(face int, value int) string {
	return fmt.Sprintf("Face: %d Qtd: %d\n", face, value)
}
