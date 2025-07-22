package main
import "fmt"

func GetName() string {
	name := ""
	fmt.Println("welocme to casino!")
	fmt.Println("Enter your name")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welocme %s, lets play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Entr your bet or 0 to quit(balance = $%d) :", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("bet cant be larger than balance")
		} else {
			break
		}
	}
	return bet
}