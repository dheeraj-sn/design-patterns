package prototype

import "fmt"

type File struct {
	Name string
}

func (f *File) Show(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_clone"}
}
