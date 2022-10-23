package main

import "fmt"

func main() {

	const myConst int = 42

	const (
		catSpecialist = iota
		dogSpecialist
		snakeSpecialist
	)

	fmt.Println(myConst)

	fmt.Printf("%v\n", dogSpecialist)

}
