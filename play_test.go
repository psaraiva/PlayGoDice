package main

import "testing"

func TestLoadDice6(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6}
	result := loadDice(6)

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("resultado '%v', esperado '%v'", result, expected)
		}
	}
}

func TestLoadAggregate6(t *testing.T) {
	expected := []int{0, 0, 0, 0, 0, 0}
	result := loadAggregate(6)

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}
}

func TestPlayDice6(t *testing.T) {
	dice := []int{1, 2, 3, 4, 5, 6}
	result := playDice(dice)

	if result < 1 || result > 6 {
		t.Errorf("result '%v', out of range 1-6.", result)
	}

}

func TestPlayDiceToText(t *testing.T) {
	expected := "[10] Dice: 6\n"
	result := playDiceToText(10, 6)

	if result != expected {
		t.Errorf("result '%v', expected '%v'", result, expected)
	}
}

func TestAggregateByIndexToText(t *testing.T) {
	expected := "Face: 10 Qtd: 6\n"
	result := aggregateByIndexToText(10, 6)

	if result != expected {
		t.Errorf("result '%v', expected '%v'", result, expected)
	}
}
