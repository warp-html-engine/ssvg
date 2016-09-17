package ssvg

import (
	"testing"
)

func TestWriteSvg(t *testing.T) {
	svg := new(Svg)

	svg.Add(&Line{X1: 0, Y1: 0, X2: 80, Y2: 80})
	svg.Add(&Text{X: 0, Y: -0, Text: "Test"})

	svg.CurrentFrame().KeepVisible = true

	for i := 10; i <= 100; i += 10 {
		svg.NextFrame().Duration = 200
		svg.Add(&Circle{
			Cx: 0,
			Cy: 0,
			R:  float64(i),
			Style: Style{
				StrokeColor: "red",
			}})
	}

	svg.WriteFile("test-output.svg", 220)
}
