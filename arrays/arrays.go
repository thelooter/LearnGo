package main

import "fmt"

func main() {

	grades := [3]int{97, 85, 93}

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[3:6]
	fmt.Printf("Grades: %v", grades)
	fmt.Println(b)

}
