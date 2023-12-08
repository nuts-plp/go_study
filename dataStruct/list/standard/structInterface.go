package standard

type listStandard interface {
	Add(data string) error
	Search(index int) (data interface{}, err error)
	Delete(index int) error
	Modify(index int, data string) (err error)
	GetSize() int
}
