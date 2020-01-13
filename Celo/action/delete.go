package action

import (
	"os"
	"github.com/node_tooling/Celo/util"
)

func ChainDataDel(target string) {
	util.TitlePrint("delete", target)
	switch target {
	case "local":
		workingDir := os.Getenv("CELO_ACCOUNT_DIR")
		util.ChangeDir(workingDir)
		util.ExecuteCmd("sudo rm -rf geth*")
	case "validator":
		workingDir := os.Getenv("CELO_VALIDATOR_DIR")
		util.ChangeDir(workingDir)
		util.ExecuteCmd("sudo rm -rf geth*")
		util.ExecuteCmd("sudo rm static-nodes.json")
	case "proxy":
		workingDir := os.Getenv("CELO_PROXY_DIR")
		util.ChangeDir(workingDir)
		util.ExecuteCmd("mv geth/nodekey nodekey")
		util.ExecuteCmd("sudo rm -rf geth*")
		util.ExecuteCmd("sudo rm static-nodes.json")
		util.ExecuteCmd("mkdir geth")
		util.ExecuteCmd("mv nodekey geth/nodekey")
	case "attestation":
		workingDir := os.Getenv("CELO_ATTESTATION_DIR")
		util.ChangeDir(workingDir)
		util.ExecuteCmd("sudo rm -rf geth*")
		util.ExecuteCmd("sudo rm static-nodes.json")
	}
}
