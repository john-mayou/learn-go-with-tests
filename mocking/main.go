package main

import (
	"os"
	"time"

	"example.com/mocking/countdown"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	countdown.Countdown(os.Stdout, &DefaultSleeper{}, 5, "Go!")
}
