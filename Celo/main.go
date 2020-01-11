package main

//                                                  jim380 <admin@cyphercore.io>
//  ============================================================================
//
//  Copyright (C) 2020 jim380
//
//  Permission is hereby granted, free of charge, to any person obtaining
//  a copy of this software and associated documentation files (the
//  "Software"), to deal in the Software without restriction, including
//  without limitation the rights to use, copy, modify, merge, publish,
//  distribute, sublicense, and/or sell copies of the Software, and to
//  permit persons to whom the Software is furnished to do so, subject to
//  the following conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
//  CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
//  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
//  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//  ============================================================================

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	var machine string
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setEnv()
	message := "Which machine are you on:\n\n1) Local\n2) Validator\n3)" +
		" " + "Proxy\n4) Attestation\n\nEnter down below (e.g. \"1\" or \"Local\"): "
	machine = inputReader(message, machine)
	nodeStop(machine)
	keyCheck(machine)
	chainDataDel(machine)
	nodeRun(machine)
	accountReg(machine)

	// args := []string{"HHHOME", "TEST"}
	// envExists(args)
}

// needs a rewrite. too tired to do it today
func setEnv() {
	celoImage := os.Getenv("CELO_IMAGE")
	os.Setenv("CELO_IMAGE", celoImage)
	networkID := os.Getenv("NETWORK_ID")
	os.Setenv("NETWORK_ID", networkID)
	celoValAddr := os.Getenv("CELO_VALIDATOR_ADDRESS")
	os.Setenv("CELO_VALIDATOR_ADDRESS", celoValAddr)
	celoValGroupAddr := os.Getenv("CELO_VALIDATOR_GROUP_ADDRESS")
	os.Setenv("CELO_VALIDATOR_GROUP_ADDRESS", celoValGroupAddr)
	celoValSignerAddr := os.Getenv("CELO_VALIDATOR_SIGNER_ADDRESS")
	os.Setenv("CELO_VALIDATOR_SIGNER_ADDRESS", celoValSignerAddr)
	celoValSignerPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_PUBLIC_KEY")
	os.Setenv("CELO_VALIDATOR_SIGNER_PUBLIC_KEY", celoValSignerPubKey)
	celoValSignerSig := os.Getenv("CELO_VALIDATOR_SIGNER_SIGNATURE")
	os.Setenv("CELO_VALIDATOR_SIGNER_SIGNATURE", celoValSignerSig)
	celoValSignerBlsPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY")
	os.Setenv("CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY", celoValSignerBlsPubKey)
	celoValSignerBlsSig := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_SIGNATURE")
	os.Setenv("CELO_VALIDATOR_SIGNER_BLS_SIGNATURE", celoValSignerBlsSig)
	celoProxyEnode := os.Getenv("PROXY_ENODE")
	os.Setenv("PROXY_ENODE", celoProxyEnode)
	celoProxyExternalIP := os.Getenv("PROXY_EXTERNAL_IP")
	os.Setenv("PROXY_EXTERNAL_IP", celoProxyExternalIP)
	celoProxyInternalIP := os.Getenv("PROXY_INTERNAL_IP")
	os.Setenv("PROXY_INTERNAL_IP", celoProxyInternalIP)
	celoValName := os.Getenv("VALIDATOR_NAME")
	os.Setenv("VALIDATOR_NAME", celoValName)

	if celoImage == "" || networkID == "" {
		log.Fatal("Missing fields in .env file")
	}
	// fmt.Println(celoImage)
}

func inputReader(msg string, target string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(msg)
	fmt.Print("\n-> ")
	target, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read string.")
	}
	switch strings.TrimSpace(target) {
	case "1", "Local", "local":
		target = "local"
		fmt.Println(target, "machine selected")
	case "2", "Validator", "validator":
		target = "validator"
		fmt.Println(target, "machine selected")
	case "3", "Proxy", "proxy":
		target = "proxy"
		fmt.Println(target, "machine selected")
	case "4", "Attestation", "attestation":
		target = "attestation"
		fmt.Println(target, "machine selected")
	default:
		t := strings.TrimSpace(target)
		fmt.Println(t, "is not a valid input")
	}
	return target
}

func keyCheck(target string) {
	switch target {
	case "local":
		fmt.Printf("\nChecking keys on %s machine", target)
		executeCmd("ls ~/Documents/celo-accounts-node/keystore")
	case "validator":
		fmt.Printf("\nChecking keys on %s machine", target)
		executeCmd("ls ~/Documents/celo-validator-node/keystore")
	case "proxy":
		fmt.Printf("\nSkip: no keys are stored on %s machine.", target)
	case "attestation":
		fmt.Printf("\nChecking keys on %s machine", target)
		executeCmd("ls ~/Documents/celo-attestations-node/keystore")
	}
}

func changeDir(dir string) {
	homeDir := os.Getenv("HOME")
	fullDir := homeDir + dir
	fmt.Println("\n------------------------------------------")
	fmt.Printf("Changing directory to \"%s\"", fullDir)
	fmt.Println("\n------------------------------------------")
	if err := os.Chdir(fullDir); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Ran successfully")
	}
}

func executeCmd(cmd string) {
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linus, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\n------------------------------------------")
	fmt.Printf("Executing %s", cmdString)
	fmt.Println("\n------------------------------------------")
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	if string(output) != "" {
		fmt.Printf("Output: %s\n", output)
	}
	fmt.Println("Ran successfully")

}

func chainDataDel(target string) {
	switch target {
	case "local":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		changeDir("/Documents/celo-accounts-node")
		executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	case "validator":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		changeDir("/Documents/celo-validator-node")
		executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	case "proxy":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		changeDir("/Documents/celo-proxy-node")
		executeCmd("mv geth/nodekey nodekey")
		executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
		executeCmd("mkdir geth")
		executeCmd("mv nodekey geth/nodekey")
	case "attestation":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		changeDir("/Documents/celo-attestations-node")
		executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	}
}

func nodeRun(target string) {
	executeCmd("docker pull $CELO_IMAGE")
	switch target {
	case "local":
		fmt.Printf("\nStarting node on %s machine", target)
		executeCmd("docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json")
		executeCmd("docker run -v $PWD:/root/.celo --rm -it --entrypoint cp $CELO_IMAGE /celo/static-nodes.json /root/.celo/")
		executeCmd("docker run --name celo-accounts -dt --restart always -p 127.0.0.1:8545:8545 -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --rpc --rpcaddr 0.0.0.0 --rpcapi eth,net,web3,debug,admin,personal")

	case "validator":
		fmt.Printf("\nStarting node on %s machine", target)
		changeDir("/Documents/celo-validator-node")
		executeCmd("docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json")
		executeCmd("docker run --name celo-validator -dt --restart always -p 30303:30303 -p 30303:30303/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --mine --istanbul.blockperiod=5 --istanbul.requesttimeout=3000 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --nodiscover --proxy.proxied --proxy.proxyenodeurlpair=enode://$PROXY_ENODE@$PROXY_INTERNAL_IP:30503\\;enode://$PROXY_ENODE@$PROXY_EXTERNAL_IP:30303  --unlock=$CELO_VALIDATOR_SIGNER_ADDRESS --password /root/.celo/.password --ethstats=$VALIDATOR_NAME@baklava-ethstats.celo-testnet.org")
	case "proxy":
		fmt.Printf("\nStarting node on %s machine", target)
		changeDir("/Documents/celo-proxy-node")
		executeCmd("docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json")
		executeCmd("export BOOTNODE_ENODES=`docker run --rm --entrypoint cat $CELO_IMAGE /celo/bootnodes`")
		executeCmd("docker run --name celo-proxy -dt --restart always -p 30303:30303 -p 30303:30303/udp -p 30503:30503 -p 30503:30503/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --proxy.proxy --proxy.proxiedvalidatoraddress $CELO_VALIDATOR_SIGNER_ADDRESS --proxy.internalendpoint :30503 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --bootnodes $BOOTNODE_ENODES --ethstats=$VALIDATOR_NAME-proxy@baklava-ethstats.celo-testnet.org")
	case "attestation":
		url := "https://docs.celo.org/getting-started/baklava-testnet/running-a-validator#running-the-attestation-service"
		fmt.Printf("\nPlease check here ($%s) for instructions on setting up an %s machine", url, target)
	}
}

// needs debugging
func envExists(envs []string) bool {
	for _, v := range envs {
		fmt.Printf(v)
		// result := os.Getenv(v)
		// if result == "" {
		// 	fmt.Printf("$%s is missing\n", v)
		// 	return false
		// } else {
		// 	fmt.Printf("$%s is set to %s\n", v, result)
		// 	return true
		// }
	}
	return false
}

func nodeStop(target string) {
	switch target {
	case "local":
		fmt.Printf("\nStoping node on %s machine", target)
		executeCmd("docker stop celo-accounts && docker rm celo-accounts")
	case "validator":
		fmt.Printf("\nStoping node on %s machine", target)
		executeCmd("docker stop celo-validator && docker rm celo-validator")
	case "proxy":
		fmt.Printf("\nStoping node on %s machine", target)
		executeCmd("docker stop celo-proxy && docker rm celo-proxy")
	case "attestation":
		fmt.Printf("\nStoping node on %s machine", target)
		executeCmd("docker stop celo-attestations && docker rm celo-attestations")
		executeCmd("docker stop celo-attestation-service && docker rm celo-attestation-service")
	}
}
