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
	// keyCheck(machine)
	executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")

}

func setEnv() {
	celoImage := os.Getenv("CELO_IMAGE")
	networkID := os.Getenv("NETWORK_ID")
	// celoValAddr := os.Getenv("CELO_VALIDATOR_ADDRESS")
	// celoValGroupAddr := os.Getenv("CELO_VALIDATOR_GROUP_ADDRESS")
	// celoValSignerAddr := os.Getenv("CELO_VALIDATOR_SIGNER_ADDRESS")
	// celoValSignerPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_PUBLIC_KEY")
	// celoValSignerSig := os.Getenv("CELO_VALIDATOR_SIGNER_SIGNATURE")
	// celoValSignerBlsPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY")
	// celoValSignerBlsSig := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_SIGNATURE")
	// celoProxyEnode := os.Getenv("PROXY_ENODE")
	// celoProxyExternalIP := os.Getenv("PROXY_EXTERNAL_IP")
	// celoProxyInternalIP := os.Getenv("PROXY_INTERNAL_IP")
	// celoValName := os.Getenv("VALIDATOR_NAME")

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

func executeCmd(cmd string) {
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linus, stoopid!")
	}

	command := exec.Command("sh", "-c", cmd)
	cmdString := "\"$ " + cmd + "\""
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	fmt.Println("\n------------------------------------------")
	fmt.Printf("Executing %s", cmdString)
	fmt.Println("\n------------------------------------------")
	err := command.Run()
	if err != nil {
		log.Fatalf("%s failed with %s\n", cmdString, err)
		return
	}
	fmt.Println("Success!")
}
func chainDataDel(target string) {
	switch target {
	case "local":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		executeCmd("cd ~/Documents/celo-accounts-node")
		executeCmd("sudo rm -rf geth* && sudo rm static-nodes.json")
	case "validator":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		executeCmd("ls ~/Documents/celo-validator-node/keystore")
	case "proxy":
		fmt.Printf("\nDeleting chain data on %s machine", target)
	case "attestation":
		fmt.Printf("\nDeleting chain data on %s machine", target)
		executeCmd("ls ~/Documents/celo-attestations-node/keystore")
	}
}
