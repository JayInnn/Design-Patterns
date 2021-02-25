package composite

type Directory struct {
	FileItem
	children []Component
}

func NewDirectory(name string) Component {
	return &Directory{
		FileItem: FileItem{
			name: name,
			size: 0,
		},
	}
}

func (d *Directory) Add(c Component) bool {
	d.children = append(d.children, c)
	return true
}

func (d *Directory) Remove(index int) bool {
	d.children = append(d.children[:index], d.children[index+1:]...)
	return true
}

func (d *Directory) GetChild(index int) Component {
	return d.children[index]
}
