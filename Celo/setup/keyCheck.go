package setup

import (
	"fmt"
	"os"

	"github.com/node_tooling/Celo/cmd"
	"github.com/node_tooling/Celo/util"
)

// KeyCheck checks if all env variables are set
func KeyCheck(target string) {
	util.TitlePrint("keyCheck", target)
	switch target {
	case "local":
		workingDir := os.Getenv("CELO_ACCOUNT_DIR")
		util.ChangeDir(workingDir)
		fmt.Printf("\nChecking keys on %s machine", target)
		cmd.ExecuteCmd("sudo ls keystore")
	case "validator":
		workingDir := os.Getenv("CELO_VALIDATOR_DIR")
		util.ChangeDir(workingDir)
		fmt.Printf("\nChecking keys on %s machine", target)
		cmd.ExecuteCmd("sudo ls keystore")
	case "proxy":
		fmt.Printf("\nSkip: no keys are stored on %s machine.\n", target)
	case "attestation":
		workingDir := os.Getenv("CELO_ATTESTATION_DIR")
		util.ChangeDir(workingDir)
		fmt.Printf("\nChecking keys on %s machine", target)
		cmd.ExecuteCmd("sudo ls keystore")
	}
}
