package cmd

import (
	"bufio"
	"fmt"
	"os"

	// "strconv"
	"github.com/shopspring/decimal"
)

// lockGold locks a specific amount of gold available
func elecVote(target []byte, role string) {
	nonVoting := ParseAmount(target, "nonVoting")
	message := "\nHow many votes would you like to cast?\n1) All\n2) A specific amount\n3) Move on"
	fmt.Printf(message)
	fmt.Printf("\n=> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "1":
			fmt.Println("\nCasting of ", nonVoting, "votes has been requested.")
			toVote := fmt.Sprintf("%v", nonVoting)
			voteAmount(toVote, role)
		case "2":
			fmt.Printf("\nHow many votes would you like to cast?")
			fmt.Printf("\n=> ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				toVote := scanner.Text()
				toVoteValue, _ := decimal.NewFromString(toVote)
				nonvotingLockedGoldValue, _ := decimal.NewFromString(fmt.Sprintf("%v", nonVoting))

				if toVoteValue.Cmp(nonvotingLockedGoldValue) == -1 {
					fmt.Println("\nCasting", toVote, "votes has been requested.")
					voteAmount(toVote, role)
				} else {
					fmt.Println("\n==> Don't bite more than you can chew!")
					fmt.Println("    You only have " + fmt.Sprintf("%v", nonVoting) + " non-voting gold available")
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
	toVote, _ := decimal.NewFromString(amount)
	if role == "group" {
		fmt.Println("\nCasting", toVote.String(), "votes from validator group to validator group")
		ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + toVote.String())
	} else if role == "validator" {
		fmt.Println("\nCasting", toVote.String(), "votes from validator to validator group")
		ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value " + toVote.String())
	}
}
