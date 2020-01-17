package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputReader(msg string, target string) string {
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
		//fmt.Println("\n", target, "machine selected")
	case "2", "Validator", "validator":
		target = "validator"
		//fmt.Println("\n", target, "machine selected")
	case "3", "Proxy", "proxy":
		target = "proxy"
		//fmt.Println("\n", target, "machine selected")
	case "4", "Attestation", "attestation":
		target = "attestation"
		//fmt.Println("\n", target, "machine selected")
	default:
		//t := strings.TrimSpace(target)
		fmt.Println("Not a valid input")
	}
	return target
}

