package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Must provide target directory, output directory, and extension(s) to clean up")
	}

	root := os.Args[1]
	destination := os.Args[2]
	var extensions []string

	for i, arg := range os.Args {
		if i > 2 {
			extensions = append(extensions, arg)
		}
	}

	fmt.Println("Cleaning up...")

	createDir(destination)
	files := findFiles(os.DirFS(root), extensions)
	moveFiles(files, "./"+destination)
}
