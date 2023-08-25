package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const ShellToUse = "sh"

// start system process
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// print system exec command
func cmdVerbose(cmd string, verbose bool) {
	if verbose {
		fmt.Println()
		fmt.Println(cmd)
		fmt.Println()
	}
}

// print system exec response
func execVerbose(err error, out, errout string, verbose bool) {
	if verbose {
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
	}
}
