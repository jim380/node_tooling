package action

import (
	"fmt"

	"github.com/node_tooling/Celo/util"
)

func Register(target string) {
	if target == "local" {
		fmt.Printf("\nChecking balance on %s machine", target)
		util.ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
		util.ExecuteCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nUnlocking accounts on %s machine", target)
		util.ExecuteCmd("celocli account:unlock --account $CELO_VALIDATOR_GROUP_ADDRESS")
		util.ExecuteCmd("celocli account:unlock --account $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nRegistering accounts on %s machine", target)
		util.ExecuteCmd("celocli account:register --from $CELO_VALIDATOR_GROUP_ADDRESS --name $VALIDATOR_NAME")
		util.ExecuteCmd("celocli account:register --from $CELO_VALIDATOR_ADDRESS --name $VALIDATOR_NAME")

		fmt.Printf("\nChecking post-registration balance on %s machine", target)
		util.ExecuteCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
		util.ExecuteCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nLocking up gold on %s machine", target)
		util.ExecuteCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")
		util.ExecuteCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value 10000000000000000000000")

		fmt.Printf("\nChecking balance on %s machine after gold lockup", target)
		util.ExecuteCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
		util.ExecuteCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nRunning election on %s machine", target)
		util.ExecuteCmd("celocli account:authorize --from $CELO_VALIDATOR_ADDRESS --role validator --signature 0x$CELO_VALIDATOR_SIGNER_SIGNATURE --signer 0x$CELO_VALIDATOR_SIGNER_ADDRESS")

		fmt.Printf("\nChecking authorized validator signer on %s machine", target)
		util.ExecuteCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nRegistering validator group on %s machine", target)
		util.ExecuteCmd("celocli validatorgroup:register --from $CELO_VALIDATOR_GROUP_ADDRESS --commission 0.1")

		fmt.Printf("\nChecking validator group on %s machine", target)
		util.ExecuteCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")

		fmt.Printf("\nChecking validator on %s machine", target)
		util.ExecuteCmd("celocli validator:register --from $CELO_VALIDATOR_ADDRESS --ecdsaKey $CELO_VALIDATOR_SIGNER_PUBLIC_KEY --blsKey $CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY --blsSignature $CELO_VALIDATOR_SIGNER_BLS_SIGNATURE")

		fmt.Printf("\nAffiliating validator with the validator group on %s machine", target)
		util.ExecuteCmd("celocli validator:affiliate $CELO_VALIDATOR_GROUP_ADDRESS --from $CELO_VALIDATOR_ADDRESS")

		fmt.Printf("\nAccepting affiliation on %s machine", target)
		util.ExecuteCmd("celocli validatorgroup:member --accept $CELO_VALIDATOR_ADDRESS --from $CELO_VALIDATOR_GROUP_ADDRESS")

		fmt.Printf("\nChecking if validator is now a member of the validator group on %s machine", target)
		util.ExecuteCmd("celocli validator:show $CELO_VALIDATOR_ADDRESS")
		util.ExecuteCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")

		fmt.Printf("\nVoting for validator group on %s machine", target)
		util.ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")
		util.ExecuteCmd("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")

		fmt.Printf("\nChecking if votes were cast successfully on %s machine", target)
		util.ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --group")
		util.ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --voter")
		util.ExecuteCmd("celocli election:show $CELO_VALIDATOR_ADDRESS --voter")
		fmt.Println("Open a new screen and execute the following:")
		fmt.Println("celocli election:activate --from $CELO_VALIDATOR_ADDRESS --wait && celocli election:activate --from $CELO_VALIDATOR_GROUP_ADDRESS --wait")
	} else {
		fmt.Printf("You are on %s machine, please switch to your local machine and try again.", target)
	}
}
