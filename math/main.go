package main

import (
	"os"
	"time"

	"learn-go-with-tests/math/clockface"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
