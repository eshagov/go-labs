package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int32 {3, 1, 4 , 1 , 5, 9}
	//sort.Ints(intSlice)
	fmt.Println(intSlice)

	//intSlice32 := []int {3, 1, 4 , 1 , 5, 9}

	//sort.Slice(intSlice32, func(i int, j int) bool {
	//	return intSlice[i] > intSlice[j]
	//} )

	sort.Slice(intSlice, func(i int, j int) bool {
		return intSlice[i] > intSlice[j]
	} )

	fmt.Println(intSlice)
}