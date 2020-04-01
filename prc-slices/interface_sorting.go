package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type MyInt struct {
	Item int
}

type ByMyInt []MyInt

func (a ByMyInt) Len() int           { return len(a) }
func (a ByMyInt) Less(i, j int) bool { return a[i].Item < a[j].Item }
func (a ByMyInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }


// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}

	numbers := []MyInt {
		{3},
		{1},
		{4},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]

	sort.Sort(ByMyInt(numbers))
	fmt.Println(numbers) // [{Eve 2} {Alice 23} {Bob 25}]
}
