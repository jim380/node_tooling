package setup

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/node_tooling/Celo/cmd"
	"github.com/node_tooling/Celo/util"
)

// NodeRun starts the docker containers on the local node
func NodeRun(target string) {
	util.TitlePrint("start", target)
	cmd.ExecuteCmd("docker pull $CELO_IMAGE")
	switch target {
	case "local":
		cmd.ExecuteCmd("docker run -v $PWD:/root/.celo --rm -i $CELO_IMAGE init /celo/genesis.json")
		cmd.ExecuteCmd("docker run -v $PWD:/root/.celo --rm -i --entrypoint cp $CELO_IMAGE /celo/static-nodes.json /root/.celo/")
		cmd.ExecuteCmd("docker run --name celo-accounts -d --restart always -p 127.0.0.1:8545:8545 -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --rpc --rpcaddr 0.0.0.0 --rpcapi eth,net,web3,debug,admin,personal")

	case "validator":
		workingDir := os.Getenv("CELO_VALIDATOR_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("docker run -v $PWD:/root/.celo --rm -i $CELO_IMAGE init /celo/genesis.json")
		cmd.ExecuteCmd("docker run --name celo-validator -d --restart always -p 30303:30303 -p 30303:30303/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --mine --istanbul.blockperiod=5 --istanbul.requesttimeout=3000 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --nodiscover --proxy.proxied --proxy.proxyenodeurlpair=enode://$PROXY_ENODE@$PROXY_INTERNAL_IP:30503\\;enode://$PROXY_ENODE@$PROXY_EXTERNAL_IP:30303  --unlock=$CELO_VALIDATOR_SIGNER_ADDRESS --password /root/.celo/.password --ethstats=$VALIDATOR_NAME@baklava-ethstats.celo-testnet.org")

	case "proxy":
		workingDir := os.Getenv("CELO_PROXY_DIR")
		util.ChangeDir(workingDir)
		cmd.ExecuteCmd("docker run -v $PWD:/root/.celo --rm -i $CELO_IMAGE init /celo/genesis.json")
		cmd.ExecuteCmd("export BOOTNODE_ENODES=`docker run --rm --entrypoint cat $CELO_IMAGE /celo/bootnodes`")
		cmd.ExecuteCmd("docker run --name celo-proxy -d --restart always -p 30303:30303 -p 30303:30303/udp -p 30503:30503 -p 30503:30503/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --proxy.proxy --proxy.proxiedvalidatoraddress $CELO_VALIDATOR_SIGNER_ADDRESS --proxy.internalendpoint :30503 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --bootnodes $BOOTNODE_ENODES --ethstats=$VALIDATOR_NAME-proxy@baklava-ethstats.celo-testnet.org")
		validatorReg(target)
	case "attestation":
		url := "https://docs.celo.org/getting-started/baklava-testnet/running-a-validator#running-the-attestation-service"
		fmt.Printf("\nPlease check here ($%s) for instructions on setting up an %s machine", url, target)
	}
}

func validatorReg(target string) {
	message := "Would you like to proceed and register your validator? (Y or N)"
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(message)
	fmt.Print("\n-> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read string.")
	}
	switch strings.TrimSpace(input) {
	case "Y", "y", "yes":
		cmd.Register(target)
	case "N", "n", "no":
		break
	default:
		t := strings.TrimSpace(target)
		fmt.Println(t, "is not a valid input")
		break
	}
}
