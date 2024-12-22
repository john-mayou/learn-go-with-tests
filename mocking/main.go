package main

import (
	"os"

	"example.com/mocking/countdown"
)

func main() {
	countdown.Countdown(os.Stdout)
}
