package main

import "fmt"

func accountReg(target string) {
	if target == "local" {
		fmt.Printf("\nChecking balance on %s machine", target)
		executeCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
		executeCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nUnlocking accounts on %s machine", target)
		executeCmd("celocli account:unlock --account $CELO_VALIDATOR_GROUP_ADDRESS")
		executeCmd("celocli account:unlock --account $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nRegistering accounts on %s machine", target)
		executeCmd("celocli account:register --from $CELO_VALIDATOR_GROUP_ADDRESS --name $VALIDATOR_NAME")
		executeCmd("celocli account:register --from $CELO_VALIDATOR_ADDRESS --name $VALIDATOR_NAME")
		fmt.Printf("\nChecking balance on %s machine after registration", target)
		executeCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
		executeCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nLocking up gold on %s machine", target)
		executeCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")
		executeCmd("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value 10000000000000000000000")
		fmt.Printf("\nChecking balance on %s machine after gold lockup", target)
		executeCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
		executeCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nRunning election on %s machine", target)
		executeCmd("celocli account:authorize --from $CELO_VALIDATOR_ADDRESS --role validator --signature 0x$CELO_VALIDATOR_SIGNER_SIGNATURE --signer 0x$CELO_VALIDATOR_SIGNER_ADDRESS")
		fmt.Printf("\nChecking authorized validator signer on %s machine", target)
		executeCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nRegistering validator group on %s machine", target)
		executeCmd("celocli validatorgroup:register --from $CELO_VALIDATOR_GROUP_ADDRESS --commission 0.1")
		fmt.Printf("\nChecking validator group on %s machine", target)
		executeCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")
		fmt.Printf("\nChecking validator on %s machine", target)
		executeCmd("celocli validator:register --from $CELO_VALIDATOR_ADDRESS --ecdsaKey $CELO_VALIDATOR_SIGNER_PUBLIC_KEY --blsKey $CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY --blsSignature $CELO_VALIDATOR_SIGNER_BLS_SIGNATURE")
		fmt.Printf("\nAffiliating validator with validator group on %s machine", target)
		executeCmd("celocli validator:affiliate $CELO_VALIDATOR_GROUP_ADDRESS --from $CELO_VALIDATOR_ADDRESS")
		fmt.Printf("\nAccepting affiliation on %s machine", target)
		executeCmd("celocli validatorgroup:member --accept $CELO_VALIDATOR_ADDRESS --from $CELO_VALIDATOR_GROUP_ADDRESS")
		fmt.Printf("\nChecking if validator is now a member of validator group on %s machine", target)
		executeCmd("celocli validator:show $CELO_VALIDATOR_ADDRESS")
		executeCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")
		fmt.Printf("\nVoting for validator group on %s machine", target)
		executeCmd("celocli election:vote --from $CELO_VALIDATOR_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")
		executeCmd("celocli election:vote --from $CELO_VALIDATOR_GROUP_ADDRESS --for $CELO_VALIDATOR_GROUP_ADDRESS --value 10000000000000000000000")
		fmt.Printf("\nChecking if votes were cast successfully on %s machine", target)
		executeCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --group")
		executeCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --voter")
		executeCmd("celocli election:show $CELO_VALIDATOR_ADDRESS --voter")
		fmt.Println("Open a new screen and execute the following:")
		fmt.Println("celocli election:activate --from $CELO_VALIDATOR_ADDRESS --wait && celocli election:activate --from $CELO_VALIDATOR_GROUP_ADDRESS --wait")
	}
}