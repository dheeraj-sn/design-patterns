package prototype

import "fmt"

type Folder struct {
	Name     string
	Children []Inode
}

func (f *Folder) Show(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, child := range f.Children {
		child.Show(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	for _, child := range f.Children {
		cloneFolder.Children = append(cloneFolder.Children, child.Clone())
	}
	return cloneFolder
}
