package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondsInClock = 60
	minutesInClock = 60
	hoursInClock   = 12
)

func secondHandPoint(t time.Time) Point {
	return radiansToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / ((secondsInClock / 2) / float64(t.Second()))
}

func minuteHandPoint(t time.Time) Point {
	return radiansToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / ((minutesInClock / 2) / float64(t.Minute())))
}

func hourHandPoint(t time.Time) Point {
	return radiansToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / ((hoursInClock / 2) / (float64(t.Hour() % hoursInClock))))
}

func radiansToPoint(radians float64) Point {
	return Point{X: math.Sin(radians), Y: math.Cos(radians)}
}
