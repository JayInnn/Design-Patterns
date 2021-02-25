package composite

type File struct {
	FileItem
}

func NewFile(name string) Component {
	return &File{
		FileItem: FileItem{
			name: name,
			size: 0,
		},
	}
}

func (f *File) Add(c Component) bool {
	// panic()
	return true
}

func (f *File) Remove(index int) bool {
	// panic()
	return true
}

func (f *File) GetChild(index int) Component {
	// panic()
	return nil
}
