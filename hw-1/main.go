package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	PrintTime(err, time)
	os.Exit(0)
}

func PrintTime(err error, time time.Time) {
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", time)
}