package main

import (
	"fmt"
	"os"
)

func main() {

	compressedString := "wq3ewew"
	uncompressedString := ""

	uncompressedString = FuncUncompress(compressedString)
	if len(uncompressedString) == 0 {
		err := fmt.Errorf("%s", "String couldn't be uncopressed")
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", uncompressedString)
	os.Exit(0)
}

func FuncUncompress(compressedString string) string {
	const ZeroSymbol = int('0')
	uncompressedString := ""

	var accumulator int

	for i := 0; i < len(compressedString); i++ {
		if compressedString[i] >= '0' && compressedString[i] <= '9' {
			if i == 0 {
				return ""
			}
			accumulator = 10*accumulator + int(compressedString[i]) - ZeroSymbol
		} else if accumulator > 0 {
			uncompressedString = multiplicator(accumulator, uncompressedString)
			uncompressedString = uncompressedString + string(compressedString[i])
			accumulator = 0
		} else {
			uncompressedString = uncompressedString + string(compressedString[i])
		}
	}
	if accumulator > 0 {
		uncompressedString = multiplicator(accumulator, uncompressedString)
	}
	return uncompressedString
}

func multiplicator(multiplier int, uncompressedString string) string {
	uncompressedStringInternal := string(uncompressedString)
	for j := 1; j < multiplier; j++ {
		suffix := string(uncompressedStringInternal[len(uncompressedStringInternal)-1])
		uncompressedStringInternal = uncompressedStringInternal + suffix
	}
	return uncompressedStringInternal
}
