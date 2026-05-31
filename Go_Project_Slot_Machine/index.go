package main

import (
	"fmt"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	winAmount := []uint{}
	for _, row := range spin {
		checkSymbol := row[0]
		win := true
		for _, colSymbol := range row[1:] {
			if colSymbol != checkSymbol {
				win = false
				break
			}
		}
		if win {
			winAmount = append(winAmount, multipliers[checkSymbol])
		}
	}
	return winAmount
}

func main() {

	symbols := map[string]uint{
		"A": 4,
		"B": 8,
		"C": 12,
		"D": 20,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	symbolArr := GenSymbolArr(symbols)

	GetName()

	balance := uint(100)

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winAmount := checkWin(spin, multipliers)
		for i, amount := range winAmount {
			win := amount * bet
			balance += win
			if amount > 0 {
				fmt.Printf("🎉 You won $%d (%dx) from line %d\n", win, amount, i+1)
			}
		}
	}

	fmt.Println("🎉 Game over ! You left with balance: $", balance)

}
