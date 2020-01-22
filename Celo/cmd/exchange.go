package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UsdToGold(target []byte, role string) {
    amountUsd := AmountAvailable(target, "usd")
    message := "\nHow much would you like to exchange?\n1) All\n2) A specific amount\n3) Move on"
	fmt.Printf(message)
	fmt.Printf("\n=> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "1":
            amountUsdValue := amountUsd.(float64)
            if amountUsdValue > 0 {
                fmt.Println("\nExchange of", amountUsd, "usd has been requested.")
			    toExchange := fmt.Sprintf("%v", amountUsd)
			    UsdToGoldAmount(toExchange, role)
            } else {
					fmt.Println("\n==> Don't bite more than you can chew!")
					fmt.Println("    You only have " + fmt.Sprintf("%v", amountUsd) + " usd available")
			}
		case "2":
			fmt.Printf("\nHow much would you like to exchange?")
			fmt.Printf("\n=> ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				toExchange := scanner.Text()
				toExchangeValue, _ := strconv.ParseFloat(toExchange, 64)
				amountUsdValue := amountUsd.(float64)
				if toExchangeValue <= amountUsdValue {
					fmt.Println("\nExchange of", toExchange, "usd has been requested.")
					UsdToGoldAmount(toExchange, role)
				} else {
					fmt.Println("\n==> Don't bite more than you can chew!")
					fmt.Println("    You only have " + fmt.Sprintf("%v", amountUsd) + " usd available")
				}
				break
			}
		case "3":
			return
        default:
			panic("invalid input")
		}
		break
	}
}

func UsdToGoldAmount(amount string, role string) {
	//toExchange, _ := strconv.Atoi(amount)
    if role == "group" {
        fmt.Println("\nExchanging", amount, "usd from validator group")
	    ExecuteCmd("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount)
    } else if role == "validator" {
        fmt.Println("\nExchanging", amount, "usd from validator")
	    ExecuteCmd("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value " + amount)
    }
}