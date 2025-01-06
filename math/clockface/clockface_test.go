package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecondHandPoint(t *testing.T) {
	tests := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assertPointInDelta(t, tt.want, secondHandPoint(tt.time), 1e-7)
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	tests := []struct {
		time time.Time
		want float64
	}{
		{simpleTime(0, 0, 0), 0.0},
		{simpleTime(0, 0, 15), 0.5 * math.Pi},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), 1.5 * math.Pi},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assert.InDelta(t, tt.want, secondsInRadians(tt.time), 1e-7)
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	tests := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assertPointInDelta(t, tt.want, minuteHandPoint(tt.time), 1e-7)
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	tests := []struct {
		time time.Time
		want float64
	}{
		{simpleTime(0, 0, 0), 0.0},
		{simpleTime(0, 15, 0), 0.5 * math.Pi},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), 1.5 * math.Pi},
		{simpleTime(0, 45, 30), (1.5 + (1 / 60.0)) * math.Pi},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assert.InDelta(t, tt.want, minutesInRadians(tt.time), 1e-7)
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	tests := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
		{simpleTime(18, 0, 0), Point{0, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assertPointInDelta(t, tt.want, hourHandPoint(tt.time), 1e-7)
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	tests := []struct {
		time time.Time
		want float64
	}{
		{simpleTime(0, 0, 0), 0.0},
		{simpleTime(3, 0, 0), 0.5 * math.Pi},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(9, 0, 0), 1.5 * math.Pi},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(18, 30, 30), (1 + (1 / 12.0) + (1 / 720.0)) * math.Pi},
	}

	for _, tt := range tests {
		t.Run(tt.time.Format(time.TimeOnly), func(t *testing.T) {
			assert.InDelta(t, tt.want, hoursInRadians(tt.time), 1e-7)
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2000, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func assertPointInDelta(t testing.TB, want, got Point, delta float64) {
	assert.InDelta(t, want.X, got.X, delta)
	assert.InDelta(t, want.Y, got.Y, delta)
}
