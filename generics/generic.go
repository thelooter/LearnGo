package generics

import "fmt"

func newGenericFunc[age int | float64](myAge age) {
	fmt.Println("my age is", myAge)
}

func main() {
	newGenericFunc(27)
	newGenericFunc(27.5)
}
