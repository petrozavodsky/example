package student

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age int, grade int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
