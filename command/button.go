package command

type Button struct {
	C Command
}

func (b *Button) Press() {
	b.C.Execute()
}
