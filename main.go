package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	branch := flag.String("b", "master", "branch to use")
	//action := flag.String("a", "init", "action to perform. Available actions: init,next,prev")
	flag.Parse()

	cmd := exec.Command("git", "checkout", *branch)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Checking out", *branch)
}
