package packages

type Inode interface {
	Print(string)
	Clone() Inode
}
