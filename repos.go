package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var dir = flag.String("dir", ".", "root directory to find git repositories in")

func repositories() {
	fileSystem := os.DirFS(*dir)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() && d.Name() == ".git" {
			fmt.Println(path)
			return fs.SkipDir
		}

		return nil
	})

}

func main() {
	flag.Parse()
	repositories()
}
