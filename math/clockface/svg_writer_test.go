package clockface

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	tests := []struct {
		time time.Time
		want Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 15), Line{150, 150, 240, 150}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
		{simpleTime(0, 0, 45), Line{150, 150, 60, 150}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s returns a second hand line of %+v", tt.time.Format(time.TimeOnly), tt.want)
		t.Run(testname, func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tt.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			assert.True(t, containsLine(tt.want, svg.Line), fmt.Sprintf("got %+v", svg.Line))
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	tests := []struct {
		time time.Time
		want Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
		{simpleTime(0, 15, 0), Line{150, 150, 230, 150}},
		{simpleTime(0, 30, 0), Line{150, 150, 150, 230}},
		{simpleTime(0, 45, 0), Line{150, 150, 70, 150}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s returns a minute hand line of %+v", tt.time.Format(time.TimeOnly), tt.want)
		t.Run(testname, func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tt.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			assert.True(t, containsLine(tt.want, svg.Line), fmt.Sprintf("got %+v", svg.Line))
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	tests := []struct {
		time time.Time
		want Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 100}},
		{simpleTime(3, 0, 0), Line{150, 150, 200, 150}},
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
		{simpleTime(9, 0, 0), Line{150, 150, 100, 150}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s returns a hour hand line of %+v", tt.time.Format(time.TimeOnly), tt.want)
		t.Run(testname, func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tt.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			assert.True(t, containsLine(tt.want, svg.Line), fmt.Sprintf("got %+v", svg.Line))
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}
