package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// LockGold locks a specific amount of gold available
func LockGold(target []byte, asset string) {
	amountGold := amoutAvailable(target, "gold")
	message := "\nHow much would you like to lock?\n1) All\n2) A specific amount"
	fmt.Printf(message)
	fmt.Printf("\n=> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "1":
			fmt.Println("Locking all", amountGold, "gold")
		// todo: lock all gold here
		case "2":
			fmt.Printf("\nHow much would you like to lock?")
			fmt.Printf("\n=> ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input := scanner.Text()
				fmt.Println("Locking", input, "gold")
				// todo: lock $input amount of gold here
				break
			}
		}
		break
	}
}
