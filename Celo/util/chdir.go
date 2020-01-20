package util

import (
	"fmt"
	"log"
	"os"
)

// ChangeDir changes the working dir to a specified dir
func ChangeDir(dir string) {
	// homeDir := os.Getenv("HOME")
	// fullDir := homeDir + dir
	fmt.Println("\nChanging directory to \"", dir, "\"")
	if err := os.Chdir(dir); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\n\u2713\u2713\u2713\u2713\u2713\u2713Ran successfully\u2713\u2713\u2713\u2713\u2713\u2713")
		fmt.Println("-----")
	}
}
