package main

import "fmt"

func main() {
	var map1 map[string]string
	map2 := map[string]int{}
	map3 := map[int]string{
		1: "one",
		2: "two",
		3: "two",
		4: "two",
	}
	map4 := make(map[string]string)
	map5 := make(map[string]string, 120)
	value1, ok1 := map3[3]
	fmt.Println(value1, ok1)
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map4)
	fmt.Println(map5)
	delete(map3, 3)
	fmt.Println(map3)

	for key, value := range map3 {
		fmt.Println(key, value)
	}

}
