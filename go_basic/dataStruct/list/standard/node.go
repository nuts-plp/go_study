package standard

type Person struct {
	Data string
	Next *Person
}

func NewPerson(data string) *Person {
	return &Person{
		Data: data,
	}
}
