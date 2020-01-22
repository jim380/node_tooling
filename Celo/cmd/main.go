package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// CmdAll contains all actions can be performed
// when the s-cmd flag is provided
func OptionsAll() {
	// input = InputReader(message, input)
	var ifContinue = true
	for ifContinue {
		message := "\nWhat would you like?\n\n1) Election Show\n2) Account Balance\n" +
			"3) Account Show\n4) Lockgold Show\n5) Validator Show\n6) Validator Status\n" +
			"7) Get Metadata\n8) Node Synced\n" +
			"\nEnter down below (e.g. \"1\"): "
		fmt.Println(message)
		fmt.Printf("=> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			switch input {
			case "1":
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --group")
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --voter")
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_ADDRESS --voter")
			case "2":
				valGrGold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
				lockGold(valGrGold, "group")
				valGold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
				lockGold(valGold, "validator")

				valGrUsd := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
				UsdToGold(valGrUsd, "group")
				valUsd := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
				UsdToGold(valUsd, "validator")
			case "3":
				ExecuteCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
				ExecuteCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")
			case "4":
				valGrVote := ExecuteCmd("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS")
				elecVote(valGrVote, "group")
				valVote := ExecuteCmd("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS")
				elecVote(valVote, "validator")
			case "5":
				ExecuteCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")
				ExecuteCmd("celocli validator:show $CELO_VALIDATOR_ADDRESS")
			case "6":
				ExecuteCmd("celocli validator:status --validator $CELO_VALIDATOR_ADDRESS")
			case "7":
				ExecuteCmd("celocli account:get-metadata $CELO_VALIDATOR_ADDRESS")
			case "8":
				ExecuteCmd("celocli node:synced")
			default:
				panic("invalid input")
			}
			break
		}
		///////////////////
		message1 := "\nWould you like to continue (y or n) => "
		fmt.Printf(message1)
		scanner1 := bufio.NewScanner(os.Stdin)
		for scanner1.Scan() {
			input := scanner1.Text()
			if input == "n" || input == "N" {
				return
			} else if input == "y" || input == "Y" {
				break
			} else {
				panic("invalid input")
			}
		}
		ifContinue = true
	}
}