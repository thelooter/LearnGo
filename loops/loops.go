package loops

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}

	population := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}

	for key, value := range population {
		fmt.Println(key, value)
	}

}
