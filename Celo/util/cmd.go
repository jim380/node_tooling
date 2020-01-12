package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

func ExecuteCmd(cmd string) {
	// setEnv()
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linux, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\nExecuting ", cmdString)
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
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
}
