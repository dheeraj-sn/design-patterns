package prototype

type Inode interface {
	Clone() Inode
	Show(string)
}
