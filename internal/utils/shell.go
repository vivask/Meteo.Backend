package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

const ShellToUse = "sh"

type shell struct {
	stdout string
	stderr string
	err    error
}

type Shell struct {
	command string
	out     chan shell
}

func NewShell(cmd string) *Shell {
	return &Shell{
		command: cmd,
		out:     make(chan shell),
	}
}

func (s *Shell) Run(seconds uint) (error, string, string) {
	timeout := time.After(time.Duration(seconds) * time.Second)
	go func() {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd := exec.Command(ShellToUse, "-c", s.command)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		out := shell{
			err:    cmd.Run(),
			stdout: stdout.String(),
			stderr: stderr.String(),
		}
		s.out <- out
	}()
	select {
	case o := <-s.out:
		return o.err, strings.Replace(o.stdout, "\n", "", -1), o.stderr
	case <-timeout:
		return fmt.Errorf("command timed out"), "", ""
	}
}

// start system process
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
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
