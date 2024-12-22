package main

import (
	"os"
	"time"

	"example.com/mocking/countdown"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper, 5, "Go!")
}
