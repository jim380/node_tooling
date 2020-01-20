package setup

import (
	"os"

	"github.com/node_tooling/Celo/cmd"
	"github.com/node_tooling/Celo/util"
)

// ChainDataDel clearss data stored on the local node
func ChainDataDel(target string) {
	util.TitlePrint("delete", target)
	switch target {
	case "local":
		workingDir := os.Getenv("CELO_ACCOUNT_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("sudo rm -rf geth*")
	case "validator":
		workingDir := os.Getenv("CELO_VALIDATOR_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("sudo rm -rf geth*")
		cmd.ExecuteCmd("sudo rm static-nodes.json")
	case "proxy":
		workingDir := os.Getenv("CELO_PROXY_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("mv geth/nodekey nodekey")
		cmd.ExecuteCmd("sudo rm -rf geth*")
		cmd.ExecuteCmd("sudo rm static-nodes.json")
		cmd.ExecuteCmd("mkdir geth")
		cmd.ExecuteCmd("mv nodekey geth/nodekey")
	case "attestation":
		workingDir := os.Getenv("CELO_ATTESTATION_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("sudo rm -rf geth*")
		cmd.ExecuteCmd("sudo rm static-nodes.json")
	}
}
