package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Initializing gwalk v0.1.0")

	branch := flag.String("b", "master", "branch to use")
	// action := flag.String("a", "init", "action to perform. Available actions: init,next,prev")
	flag.Parse()
	cmd := exec.Command("git", "checkout", *branch)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Checking out", *branch)

	history := generateHistory()
	commits := strings.Split(string(history), "\n")

	for {
		fmt.Println("Available actions - init,next,prev,exit")
		fmt.Println("Enter action: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		i := 0

		switch text {
		case "init":
			fmt.Println("Checking out initial commit")
			i = 0
			checkout(i)
		case "next":
			fmt.Println("Moving forward by one commit")
			i++
			checkout(i)
		case "prev":
			i--
			checkout(i)
			fmt.Println("Moving backward by one commit")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Checking out initial commit")
			i = 0
			checkout(i)

		}

	}
}

func checkout(i int,commit []string) {
	cmd = exec.Command("git", "checkout", commits[i])
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func generateHistory() []byte {

	var (
		cmdOut []byte
		err    error
	)

	cmdName := "git"
	cmdArgs := []string{"log", "--pretty=%h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git log command: ", err)
		os.Exit(1)
	}

	return cmdOut
}
