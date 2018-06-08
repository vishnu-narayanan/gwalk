package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

	var (
		cmdOut []byte
		err2    error
	)

	cmdName := "git"
	cmdArgs := []string{"log", "--pretty=%h"}
	if cmdOut, err2 = exec.Command(cmdName, cmdArgs...).Output(); err2 != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err2)
		os.Exit(1)
	}
	
	commits := strings.Split(string(cmdOut),"\n")
	for i,c := range commits {
			fmt.Println(i,c)
		}

	fmt.Println("Replaying history...")
	fmt.Println("Checking out initial commit")

	cmd = exec.Command("git", "checkout", commits[0])
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}