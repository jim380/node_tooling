package action

import "github.com/node_tooling/Celo/util"

func NodeStop(target string) {
	util.TitlePrint("stop", target)
	switch target {
	case "local":
		util.ExecuteCmd("docker stop celo-accounts && docker rm celo-accounts")
	case "validator":
		util.ExecuteCmd("docker stop celo-validator && docker rm celo-validator")
	case "proxy":
		util.ExecuteCmd("docker stop celo-proxy && docker rm celo-proxy")
	case "attestation":
		util.ExecuteCmd("docker stop celo-attestations && docker rm celo-attestations")
		util.ExecuteCmd("docker stop celo-attestation-service && docker rm celo-attestation-service")
	}
}
