package main
import (
	"fmt"
	"math/rand"
)

func Generate(symbols map[string]uint) []string {
	symbolArr := []string{}
	for s, c := range symbols {
		for i := uint(0); i < c; i++ {
			symbolArr = append(symbolArr, s)
		}
	}
	return symbolArr
}

func GetRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				randomIndex := GetRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func PrintSpin(spin [][]string) {
	for _,row := range spin {
		for j,symbol := range row {
			fmt.Printf(symbol)
			if j != len(row) - 1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}