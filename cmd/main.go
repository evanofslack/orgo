package main

import (
	"find"
	"fmt"
	"os"
)

func main() {
	count := find.Files(os.DirFS(os.Args[1]))
	fmt.Println(count)
}
