package main

import (
	"fmt"
	"math/rand"
)

func getRandomIdx(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("🧧 Enter your bet amount or 0 to stop (Balance: %d): ", balance)
		fmt.Scan(&bet)
		if bet == 0 {
			break
		}

		if bet > balance {
			fmt.Println("❌ Bet cannot be more than balance.")
		} else {
			break
		}
	}
	return bet
}

func GenSymbolArr(symbols map[string]uint) (symbolArr []string) {
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return
}

func GetSpin(reel []string, rows uint, cols uint) (result [][]string) {
	for i := uint(0); i < rows; i++ {
		result = append(result, []string{})
	}

	for col := uint(0); col < cols; col++ {
		selected := map[int]bool{}
		// Implementation for spinning the reels
		for row := uint(0); row < rows; row++ {
			for true {
				symbolIdx := getRandomIdx(0, len(reel))
				_, exists := selected[symbolIdx]
				if !exists {
					selected[symbolIdx] = true
					result[row] = append(result[row], reel[symbolIdx])
					break
				}
			}
		}
	}
	return
}

func PrintSpin(spin [][]string) {
	for _, row := range spin {
		for j, col := range row {
			fmt.Print(col)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}
