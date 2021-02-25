package composite

type Component interface {
	// 针对场景实现具体Operation方法
	GetNameOperation() string
	GetSizeOperation() int

	// 必选方法
	Add(c Component) bool
	Remove(index int) bool
	GetChild(index int) Component
}

// 文件及目录必带name，size
type FileItem struct {
	name string
	size int
}

func (f *FileItem) GetNameOperation() string {
	return f.name
}

func (f *FileItem) GetSizeOperation() int {
	return f.size
}