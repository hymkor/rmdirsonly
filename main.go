package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func rmdirsonly(arg string) bool {
	dirEntries, err := os.ReadDir(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: os.ReadDir: %s\n", arg, err.Error())
		return false
	}
	allEmpty := true
	for _, dirEntry1 := range dirEntries {
		if !dirEntry1.IsDir() {
			allEmpty = false
			continue
		}
		name := dirEntry1.Name()
		if name == "." || name == ".." {
			continue
		}
		fullPath := filepath.Join(arg, name)
		isEmpty := rmdirsonly(fullPath)
		if !isEmpty {
			allEmpty = false
		}
	}
	if allEmpty {
		err := os.Remove(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: os.Remove: %s\n", arg, err.Error())
			return false
		}
		fmt.Printf("rmdir \"%s\"\n", arg)
	}
	return allEmpty
}

func main() {
	for _, arg1 := range os.Args[1:] {
		files, err := filepath.Glob(arg1)
		if err != nil {
			rmdirsonly(arg1)
		} else {
			for _, fn := range files {
				rmdirsonly(fn)
			}
		}
	}
}
