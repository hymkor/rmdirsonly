package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main1(arg string) (bool, error) {
	dirEntries, err := os.ReadDir(arg)
	if err != nil {
		return false, fmt.Errorf("%s: os.ReadDir: %w", arg, err)
	}
	allEmpty := true
	for _, dirEntry1 := range dirEntries {
		if !dirEntry1.IsDir() {
			allEmpty = false
			continue
		}
		name := dirEntry1.Name()
		if name == ".." {
			continue
		}
		fullPath := filepath.Join(arg, name)
		isEmpty, err := main1(fullPath)
		if err != nil {
			return false, err
		}
		if !isEmpty {
			allEmpty = false
		}
	}
	if allEmpty {
		err := os.Remove(arg)
		if err != nil {
			return false, fmt.Errorf("%s: os.Remove: %w", arg, err)
		}
		fmt.Printf("rmdir \"%s\"\n", arg)
	}
	return allEmpty, nil
}

func main() {
	for _, arg1 := range os.Args[1:] {
		if _, err := main1(arg1); err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
