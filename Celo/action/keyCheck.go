package action

import (
	"fmt"

	"github.com/node_tooling/Celo/util"
)

func KeyCheck(target string) {
	switch target {
	case "local":
		fmt.Printf("\nChecking keys on %s machine", target)
		util.ExecuteCmd("sudo ls ~/Documents/celo-accounts-node/keystore")
	case "validator":
		fmt.Printf("\nChecking keys on %s machine", target)
		util.ExecuteCmd("sudo ls ~/Documents/celo-validator-node/keystore")
	case "proxy":
		fmt.Printf("\nSkip: no keys are stored on %s machine.", target)
	case "attestation":
		fmt.Printf("\nChecking keys on %s machine", target)
		util.ExecuteCmd("sudo ls ~/Documents/celo-attestations-node/keystore")
	}
}
