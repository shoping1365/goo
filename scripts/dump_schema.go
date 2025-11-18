package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	root := filepath.Join("my-go-backend", "internal", "database", "backup")
	createRe := regexp.MustCompile(`(?i)CREATE\s+TABLE`)
	alterRe := regexp.MustCompile(`(?i)ALTER\s+TABLE`)

	f, err := os.Create("schema_dump.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			if createRe.MatchString(line) || alterRe.MatchString(line) {
				fmt.Fprintf(f, "%s:%d:%s\n", path, lineNum, line)
			}
		}
		return scanner.Err()
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}

	fmt.Println("schema_dump.txt generated")
}
