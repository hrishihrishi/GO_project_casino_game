package main

import (
	"fmt"
	
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _,symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	
	multipliers := map[string]uint{
		"A":20,
		"B":10,
		"C":5,
		"D":2,
	}

	sa := Generate(symbols)
	
	//fmt.Println(sa)
	//fmt.Println(spin)
	

	balance := uint(200)

	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		//fmt.Println(bet)
		balance -= bet
		spin := GetSpin(sa, 3, 3)
		PrintSpin(spin)
		winningLines := checkWin(spin, multipliers)
		for i, m := range winningLines {
			win := m*bet
			balance += win
			if m >0 {
				fmt.Printf("Won $%d, (%dx) on line #%d\n", win,m,i+1)
			}
		}
	}

	fmt.Printf("You left with $%d\n", balance)
}
