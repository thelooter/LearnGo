package main

import "fmt"

func main() {

	number := 50

	guess := 0

	if number < guess {
		fmt.Println("true")
	} else if number > guess {
		fmt.Println("false")
	} else {
		fmt.Println("equal")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	var i interface{} = 42

	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}

	if population, ok := statePopulations["Ohio"]; ok {
		fmt.Println(population)
	}

}
