package composite

// component 组件是接口，文件和文件夹实现了该接口
type component interface {
	search(string)
}
