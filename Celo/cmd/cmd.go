package cmd

import (
	"fmt"
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

func ParseCmdOutput(output []byte, parseType string, reg string, matchGr int) interface{} {
	// regexp.MustCompile("lockedGold: (\\d+)").FindStringSubmatch(string(output))
	match := regexp.MustCompile(reg).FindStringSubmatch(string(output))
	var result interface{}
	if parseType == "int" {
		if match != nil {
			if i, err := strconv.Atoi(match[matchGr]); err == nil {
				result = i
			}
		}
		// fmt.Println("Test output: ", result)
	} else if parseType == "float" {
		if match != nil {
			if i, err := strconv.ParseFloat(match[matchGr], 64); err == nil {
				result = i
			}
		}
		// fmt.Println("Test output: ", result)
	} else if parseType == "string" {
		if match != nil {
			// if i, err := string(match[1]); err == nil {
			// 	result = i
			// }
			result = match[matchGr]

		}
		// fmt.Println("Test output: ", result)
	}
	return result
}

func AmountAvailable(target []byte, asset string) interface{} {
	// gold := ExecuteCmd("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
	var result interface{}
	switch asset {
	case "gold":
		result = ParseCmdOutput(target, "float", "gold: (\\d.\\d*.+)", 1)
		if result == nil {
			fmt.Printf("\nYou have no gold available\n")
		} else {
			fmt.Printf("\nYou have %v gold available to lock\n", result)
		}
	case "nonvotingLockedGold":
		result = ParseCmdOutput(target, "float", "nonvoting: (\\d.\\d*)", 1)
		if result == nil {
			fmt.Printf("\nYou have no nonvoting lockedGold available\n")
		} else {
			fmt.Printf("\nYou have %v nonvoting lockedGold\n", result)
		}
	case "usd":
		result = ParseCmdOutput(target, "float", "usd: (\\d.\\d*.+)", 1)
		if result == nil {
			fmt.Printf("\nYou have no usd available\n")
		} else {
			fmt.Printf("\nYou have %v usd available to exchange\n", result)
		}
	}
	return result
}
