package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"log", "--pretty=%h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	fmt.Println(cmdOut)
	fmt.Println(reflect.TypeOf(cmdOut))
}
