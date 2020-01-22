package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// lockGold locks a specific amount of gold available
func elecVote(target []byte, role string) {
	nonvotingLockedGold := AmountAvailable(target, "nonvotingLockedGold")
	message := "\nHow many votes would you like to cast?\n1) All\n2) A specific amount\n3) Move on"
	fmt.Printf(message)
	fmt.Printf("\n=> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "1":
			fmt.Println("\nCasting of ", nonvotingLockedGold, "votes has been requested.")
			toVote := fmt.Sprintf("%v", nonvotingLockedGold)
			voteAmount(toVote, role)
		case "2":
			fmt.Printf("\nHow many votes would you like to cast?")
			fmt.Printf("\n=> ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				toVote := scanner.Text()
				toVoteValue, _ := strconv.ParseFloat(toVote, 64)
				nonvotingLockedGoldValue := nonvotingLockedGold.(float64)
				if toVoteValue <= nonvotingLockedGoldValue {
					fmt.Println("\nCasting", toVote, "votes has been requested.")
					voteAmount(toVote, role)
				} else {
					fmt.Println("\n==> Don't bite more than you can chew!")
					fmt.Println("    You only have " + fmt.Sprintf("%v", nonvotingLockedGold) + " non-voting gold available")
				}
				break
			}
		case "3":
			return
		default:
			panic("Invalid input")
		}
		break
	}
}

func voteAmount(amount string, role string) {
	toVote, _ := strconv.ParseFloat(amount, 64)
		if role == "group" {
			fmt.Println("\nCasting", toVote, "votes from validator group to validator group")
			ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + fmt.Sprintf("%f", toVote))
		} else if role == "validator" {
			fmt.Println("\nCasting", toVote, "votes from validator to validator group")
			ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + fmt.Sprintf("%f", toVote))
		}
}