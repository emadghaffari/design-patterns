package main

import (
	"fmt"

	"github.com/emadghaffari/design-patterns/prototype/second/packages"
)

func main() {
	file1 := packages.File{Name: "file1"}
	file2 := packages.File{Name: "file2"}
	file3 := packages.File{Name: "file3"}

	folder1 := packages.Folder{Name: "folder1", Childrens: []packages.Inode{&file1}}

	folder2 := packages.Folder{Name: "folder2", Childrens: []packages.Inode{&file1, &folder1, &file2, &file3}}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
