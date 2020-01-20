package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

// ExecuteCmd is a wrapper for os.exec()
func ExecuteCmd(cmd string) []byte {
	// setEnv()
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linux, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\nExecuting ", cmdString)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	// if string(output) != "" {
	// 	fmt.Printf("Output: %s\n", output)
	// }
	if err != nil {
		// switch err.Error() {
		// case "Error response from daemon: No such container: celo-accounts":
		// 	fmt.Printf("error has occurred.")
		// default:
		// 	log.Fatal(err)
		// }
		//fmt.Println("Error:", err.Error())
		//log.Fatal(err)
		fmt.Println("\n", fmt.Sprint(err)+": "+string(output))
	} else {
		if string(output) != "" {
			fmt.Println("\nOutput=>", string(output))
		}
		fmt.Println("\n\u2713\u2713\u2713\u2713\u2713\u2713Ran successfully\u2713\u2713\u2713\u2713\u2713\u2713")
		fmt.Println("-----")
	}
	return output
}

func parseCmdOutput(output []byte, parseType string, reg string) interface{} {
	// regexp.MustCompile("lockedGold: (\\d+)").FindStringSubmatch(string(output))
	match := regexp.MustCompile(reg).FindStringSubmatch(string(output))
	var result interface{}
	if parseType == "int" {
		// var result int
		if match != nil {
			if i, err := strconv.Atoi(match[1]); err == nil {
				result = i
			}
		}
		// fmt.Println("Test output: ", result)
	} else if parseType == "float" {
		// var result float64
		if match != nil {
			if i, err := strconv.ParseFloat(match[1], 64); err == nil {
				result = i
			}
		}
		// fmt.Println("Test output: ", result)
	}
	return result
}

func amoutAvailable(target []byte, asset string) interface{} {
	// gold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
	var result interface{}
	switch asset {
	case "gold":
		result = parseCmdOutput(target, "int", "gold: (\\d+)")
		fmt.Printf("\nYou have %v gold available to lock\n", result)
	case "lockedGold":
		result = parseCmdOutput(target, "float", "lockedGold: (\\d+\\.\\d+e+\\+22)")
		fmt.Printf("\nYou have %v lockedGold\n", result)
	case "usd":
		result = parseCmdOutput(target, "int", "usd: (\\d+)")
		fmt.Printf("\nYou have %v usd available to exchange\n", result)
	}
	return result
}

// CmdAll contains all actions can be performed
// when the s-cmd flag is provided
func OptionsAll() {
	// input = InputReader(message, input)
	var ifContinue = true
	for ifContinue {
		message := "\nWhat would you like?\n\n1) Election Show\n2) Account Balance\n" +
			"3) Account Show\n4) Lockgold Show\n5) Validator Show\n6) Validator Status\n" +
			"7) Get Metadata\n8) Node Synced\n" +
			"\nEnter down below (e.g. \"1\"): "
		fmt.Println(message)
		fmt.Printf("=> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			switch input {
			case "1":
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --group")
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_GROUP_ADDRESS --voter")
				ExecuteCmd("celocli election:show $CELO_VALIDATOR_ADDRESS --voter")
			case "2":
				valGrGold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
				lockGold(valGrGold, "gold")
				valGold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
				lockGold(valGold, "gold")

				valGrUsd := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
				amoutAvailable(valGrUsd, "usd")
				valUsd := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_ADDRESS")
				amoutAvailable(valUsd, "usd")
			case "3":
				ExecuteCmd("celocli account:show $CELO_VALIDATOR_GROUP_ADDRESS")
				ExecuteCmd("celocli account:show $CELO_VALIDATOR_ADDRESS")
			case "4":
				ExecuteCmd("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS")
				ExecuteCmd("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS")
			case "5":
				ExecuteCmd("celocli validatorgroup:show $CELO_VALIDATOR_GROUP_ADDRESS")
				ExecuteCmd("celocli validator:show $CELO_VALIDATOR_ADDRESS")
			case "6":
				ExecuteCmd("celocli validator:status --validator $CELO_VALIDATOR_ADDRESS")
			case "7":
				ExecuteCmd("celocli account:get-metadata $CELO_VALIDATOR_ADDRESS")
			case "8":
			ExecuteCmd("celocli node:synced")
			}
			break
		}
		///////////////////
		message1 := "\nWould you like to continue (y or n) => "
		fmt.Printf(message1)
		scanner1 := bufio.NewScanner(os.Stdin)
		for scanner1.Scan() {
			input := scanner1.Text()
			if input == "n" || input == "N" {
				return
			} else if input == "y" || input == "Y" {
				break
			} else {
				panic("invalid input")
			}
		}
		ifContinue = true
	}
}
