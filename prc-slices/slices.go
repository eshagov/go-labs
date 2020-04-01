package main

import "fmt"

func main() {
	slice1 := []string{}
	slice2 := make([]string, 3, 4)
	//slice3 := make([]string, 6, 5)

	//Output
	fmt.Println(slice1)
	fmt.Println(slice2)
	//fmt.Println(slice3)

	//Checks 1
	slice4 := make([]int, 0, 1<<16)
	for i := 1; i < 128; i++ {
		slice4 = append(slice4, i)
		fmt.Printf("i=%3d len=%3d cap=%3d addr=%0x\n", i, len(slice4), cap(slice4), &slice4[0])
	}

	//Check 2
	slice5 := make([]int, 3, 64)
	fmt.Println(slice5)
	fmt.Println(slice5[1:])
	fmt.Println(slice5[:1])
	fmt.Println(slice5[:30])
	fmt.Println(slice5[27:30])

	//fmt.Println(slice5[70:75]) panic: runtime error: slice bounds out of range [:75] with capacity 64
	//fmt.Println(slice5[30])  //panic: runtime error: index out of range [30] with length 3
	//fmt.Println(slice5[30:]) //panic: runtime error: slice bounds out of range [30:3]

	slice6 := make([]int, 2, 4)
	slice7 := slice6
	for i := 1; i < 710; i++ {
		slice6 = append(slice6, i)
		fmt.Printf("i=%d addr6=%0x addr7=%0x\n", i, &slice6[0], &slice7[0])
	}

	slice8 := []int{1,2,3}
	slice9 := make([]int, len(slice8))
	copy(slice9, slice8)
	fmt.Printf("addr8=%0x addr9=%0x\n", &slice8[0], &slice9[0])
}
