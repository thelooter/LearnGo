package deferpanicrecover

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	panic("panic")

}
