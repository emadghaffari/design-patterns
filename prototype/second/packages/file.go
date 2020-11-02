package packages

import "fmt"

type File struct {
	Name string
}

func (f *File) Print(s string) {
	fmt.Println(s + f.Name + "_clone")
}

func (f *File) Clone() Inode {
	return &File{Name: f.Name}
}
