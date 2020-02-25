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
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linux, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\nExecuting ", cmdString)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
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
	match := regexp.MustCompile(reg).FindStringSubmatch(string(output))
	var result interface{}
	if parseType == "int" {
		if match != nil {
			if i, err := strconv.Atoi(match[matchGr]); err == nil {
				result = i
			}
		}
	} else if parseType == "float" {
		if match != nil {
			if i, err := strconv.ParseFloat(match[matchGr], 64); err == nil {
				result = i
			}
		}
	} else if parseType == "string" {
		if match != nil {
			result = match[matchGr]
		}
	}
	return result
}

func AmountAvailable(target []byte, asset string) interface{} {
	var result interface{}
	switch asset {
	case "gold":
		result = ParseCmdOutput(target, "float", "gold: (\\d.\\d*.+)", 1)
		if result == nil {
			result = "0"
			fmt.Printf("\nYou have no gold available\n")
		} else {
			fmt.Printf("\nYou have %v gold available to lock\n", result)
		}
	case "usd":
		result = ParseCmdOutput(target, "float", "usd: (\\d.\\d*.+)", 1)
		if result == nil {
			result = "0"
			fmt.Printf("\nYou have no usd available\n")
		} else {
			fmt.Printf("\nYou have %v usd available to exchange\n", result)
		}
	case "lockedGold":
		result = ParseCmdOutput(target, "float", "lockedGold: (\\d.\\d*.+)", 1)
		if result == nil {
			result = "0"
			fmt.Printf("\nYou have no lockedGold available\n")
		} else {
			fmt.Printf("\nYou have %v lockedGold\n", result)
		}
	case "total":
		result = ParseCmdOutput(target, "float", "total: (\\d.\\d*.+)", 1)
		if result == nil {
			result = "0"
			fmt.Printf("\nYou have no gold available\n")
		} else {
			fmt.Printf("\nYou have %v gold in total\n", result)
		}
	case "nonVoting":
		result = ParseCmdOutput(target, "float", "nonvoting: (\\d.\\d*)", 1)
		if result == nil {
			result = "0"
			fmt.Printf("\nYou have no nonvoting lockedGold available\n")
		} else {
			fmt.Printf("\nYou have %v nonvoting lockedGold\n", result)
		}
	}
	return result
}
