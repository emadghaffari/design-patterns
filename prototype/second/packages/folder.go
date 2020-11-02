package packages

import "fmt"

type Folder struct {
	Childrens []Inode
	Name      string
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, item := range f.Childrens {
		item.Print(s + s)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	Childrens := []Inode{}
	for _, i := range f.Childrens {
		clone := i.Clone()
		Childrens = append(Childrens,clone)
	}
	cloneFolder.Childrens = Childrens
	return cloneFolder
}
