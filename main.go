package main

import (
	"fmt"
	"os/exec"
    "log"
)

func main() {
	cmd := exec.Command("git","checkout","master")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Checking out master")
}
