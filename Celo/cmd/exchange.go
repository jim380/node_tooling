package cmd

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"github.com/shopspring/decimal"
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
            amountUsdValue, _ := decimal.NewFromString(fmt.Sprintf("%v", amountUsd))
            zeroValue, _ := decimal.NewFromString("0")
			if amountUsdValue.Cmp(zeroValue) == 1 {
                fmt.Println("\nExchange of", amountUsd, "usd has been requested.")
			    toExchange := fmt.Sprintf("%v", amountUsd)
			    UsdToGoldExecute(toExchange, role)
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
				toExchangeValue, _ := decimal.NewFromString(toExchange)
				amountUsdValue, _ := decimal.NewFromString(fmt.Sprintf("%v", amountUsd))
				if toExchangeValue.Cmp(amountUsdValue) == -1 {
					fmt.Println("\nExchange of", toExchange, "usd has been requested.")
					UsdToGoldExecute(toExchange, role)
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

func UsdToGoldExecute(amount string, role string) {
	//toExchange, _ := strconv.Atoi(amount)
    if role == "group" {
        fmt.Println("\nExchanging", amount, "usd from validator group")
	    ExecuteCmd("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + amount)
    } else if role == "validator" {
        fmt.Println("\nExchanging", amount, "usd from validator")
	    ExecuteCmd("celocli exchange:dollars --from $CELO_VALIDATOR_ADDRESS --value " + amount)
    }
}