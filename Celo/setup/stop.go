package setup

import (
	"github.com/node_tooling/Celo/cmd"
	"github.com/node_tooling/Celo/util"
)

// NodeStop stops and deletes the docker containers on the local node
func NodeStop(target string) {
	util.TitlePrint("stop", target)
	switch target {
	case "local":
		cmd.ExecuteCmd("docker stop celo-accounts && docker rm celo-accounts")
	case "validator":
		cmd.ExecuteCmd("docker stop celo-validator && docker rm celo-validator")
	case "proxy":
		cmd.ExecuteCmd("docker stop celo-proxy && docker rm celo-proxy")
	case "attestation":
		cmd.ExecuteCmd("docker stop celo-attestations && docker rm celo-attestations")
		cmd.ExecuteCmd("docker stop celo-attestation-service && docker rm celo-attestation-service")
	}
}
