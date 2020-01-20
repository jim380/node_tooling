package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// lockGold locks a specific amount of gold available
func lockGold(target []byte, asset string) {
	amountGold := amoutAvailable(target, "gold")
	message := "\nHow much would you like to lock?\n1) All\n2) A specific amount"
	fmt.Printf(message)
	fmt.Printf("\n=> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "1":
			fmt.Println("\nLocking", amountGold, "gold has been requested.")
			toLock := fmt.Sprintf("%v", amountGold)
			lockGoldAmount(toLock)
		case "2":
			fmt.Printf("\nHow much would you like to lock?")
			fmt.Printf("\n=> ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				toLock := scanner.Text()
				toLockValue, _ := strconv.Atoi(toLock)
				amountGoldValue := amountGold.(int)
				if toLockValue < amountGoldValue {
					fmt.Println("\nLocking", toLock, "gold has been requested.")
					lockGoldAmount(toLock)
				} else {
					fmt.Println("\n==> Don't bite more than you can chew!")
					fmt.Println("    You only have " + fmt.Sprintf("%v", amountGold) + " gold available")
				}
				break
			}
		}
		break
	}
}

func lockGoldAmount(amount string) {
	toLock, _ := strconv.Atoi(amount)
	toLock = toLock - 1000000000000000000
	if toLock <= 0 {
		// fmt.Printf("%v is of the type %T", toLockAfter, toLockAfter)
		fmt.Println("\n==> Not enough gold to set aside 1 gold for fees." + " Must have at least 1 gold reserved.")
	} else {
		fmt.Println("Locking", toLock, "gold")
		ExecuteCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + strconv.Itoa(toLock))
		// ExecuteCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value 10000000000000000000000")
	}
}