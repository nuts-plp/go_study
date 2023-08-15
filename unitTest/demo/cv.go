package demo

//函数的单元功能测试
func getArea(width, height int64) int64 {
	return width * height
}

//方法的单元功能测试
type Args struct {
	width  float64
	height float64
}

func NewArgs(a, b float64) *Args {
	return &Args{a, b}
}
func (a *Args) GetMatrix() float64 {
	return a.width * a.height
}
