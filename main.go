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

var commits []string

func main() {
	fmt.Println("Initializing gwalk v0.1.0")

	branch := flag.String("b", "master", "branch to use")
	flag.Parse()
	fmt.Println("Checking out branch ", *branch)
	cmd := exec.Command("git", "checkout", *branch)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	history := generateHistory()
	commits = strings.Split(string(history), "\n")
	lengthOfCommits := len(commits) - 2

	fmt.Printf("%d commits found ", lengthOfCommits)
	fmt.Println("Available actions - init,next,prev,exit")
	fmt.Println(commits[0])
	i := 0

	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		switch strings.TrimRight(text, "\n") {
		case "init":
			fmt.Println("Checking out initial commit")
			i = 0
			checkout(i, commits)
		case "next":
			fmt.Println("Moving forward by one commit")
			if i == lengthOfCommits {
				fmt.Println("Cannot move forward. Already at the latest commit.")
				break
			}
			i++
			checkout(i, commits)
		case "prev":
			fmt.Println("Moving backward by one commit")
			if i == 0 {
				fmt.Println("Cannot move backward. Already at first commit.")
				break
			}
			i--
			checkout(i, commits)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(".....")
		}

	}
}

func checkout(i int, commit []string) {
	cmd := exec.Command("git", "checkout", commits[i])
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		os.Exit(1)
	}
	return
}

func generateHistory() []byte {

	var (
		cmdOut []byte
		err    error
	)

	cmdName := "git"
	cmdArgs := []string{"log", "--reverse", "--pretty=%h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git log command: ", err)
		os.Exit(1)
	}

	return cmdOut
}
