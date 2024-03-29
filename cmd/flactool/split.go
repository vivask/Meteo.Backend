package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

// split flac, ape, wav files according to cue by directories
func SplitApeOrFlac(shntool, dir string, parallel uint, remove, verbose bool) error {
	files, err := getSplitFilesFromDir(dir, ".flac", ".ape", ".wav")
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	if len(files) == 0 {
		return fmt.Errorf("files not found")
	}

	for _, file := range files {
		fmt.Println(file)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Split this files?")
	fmt.Println("[Yes/no]?")
	t, _ := reader.ReadString('\n')
	t = strings.ToLower(t)
	if t == "y\n" || t == "yes\n" || t == "\n" {
		StartSpinner()
		g, _ := errgroup.WithContext(context.Background())
		g.SetLimit(int(parallel))
		for _, file := range files {
			//prepare shell command
			file := file
			cue := replaceExtToCue(file)
			out := filepath.Dir(file)
			cmd := shntool + " split -f \"" + cue + "\" -o flac -t \"%n %t\" " + "\"" + file + "\" -d " + "\"" + out + "\""
			cmdVerbose(cmd, verbose)
			g.Go(func() error {
				//split file
				stdout, errout, err := Shellout(cmd)
				execVerbose(err, stdout, errout, verbose)
				if err != nil {
					return err
				}
				//remove source
				if remove {
					err = os.Remove(file)
					if err != nil {
						return err
					}
				}
				return err
			})
		}
		err = g.Wait()
		StopSpinner()
		return err
	}

	return nil
}
