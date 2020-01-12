package action

import "github.com/node_tooling/Celo/util"

func ChainDataDel(target string) {
	util.TitlePrint("delete", target)
	switch target {
	case "local":
		util.ChangeDir("/Documents/celo-accounts-node")
		util.ExecuteCmd("sudo rm -rf geth*")
	case "validator":
		util.ChangeDir("/Documents/celo-validator-node")
		util.ExecuteCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	case "proxy":
		util.ChangeDir("/Documents/celo-proxy-node")
		util.ExecuteCmd("mv geth/nodekey nodekey")
		util.ExecuteCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
		util.ExecuteCmd("mkdir geth")
		util.ExecuteCmd("mv nodekey geth/nodekey")
	case "attestation":
		util.ChangeDir("/Documents/celo-attestations-node")
		util.ExecuteCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	}
}
