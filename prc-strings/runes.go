package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	A = iota
	B
	C
	_
	_
	D = iota - 5
	E
)

func main() {
	simpleString := "öMy first string in Русском and Suomiö"
	fmt.Printf("String is: %s \n", simpleString)
	runeInString, size := utf8.DecodeRuneInString(simpleString)
	fmt.Printf("%c %d\n", runeInString, size)

	inString := utf8.RuneCountInString(simpleString)
	fmt.Printf("%d\n", inString)

	for index, symbol := range simpleString {
		fmt.Printf("%d - %c\n", index, symbol)
	}

	if strings.Contains(simpleString, "first") {
		fmt.Println("Contains!")
	}

	if strings.Contains(simpleString, "second") {
		fmt.Println("Contains!")
	} else {
		fmt.Println("Not contain!")
	}

	hasPrefix := strings.HasPrefix(simpleString, "My")
	fmt.Println(hasPrefix)

	myStrtings := []string {"First string", "Second string"}
	joinResult := strings.Join(myStrtings, ", ")
	fmt.Println(joinResult)

	splitResult := strings.Split(joinResult, ", ")
	fmt.Println(len(splitResult))

	//String builder
	var sb strings.Builder
	sb.WriteString("Hello ")
	sb.WriteRune('A')
	fmt.Println(sb.String())

	fmt.Println(A, B, C, D, E)

}
