package functions

import "fmt"

func main() {
	m := MyStruct{A: 10, B: 5}
	fmt.Println(m.add())
	fmt.Println(m.sub())
	fmt.Println(m.mul())
	fmt.Println(m.div())
	fmt.Println(m.mod())
}

type MyStruct struct {
	A int
	B int
}

func (m *MyStruct) add() int {
	return m.A + m.B
}

func (m *MyStruct) sub() int {
	return m.A - m.B
}

func (m *MyStruct) mul() int {
	return m.A * m.B
}

func (m *MyStruct) div() int {
	return m.A / m.B
}

func (m *MyStruct) mod() int {
	return m.A % m.B
}
