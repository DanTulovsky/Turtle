package turtle

import (
	"testing"
	"time"

	"github.com/DanTulovsky/L-System/l"
)

func customtest1() (string, float64, l.Rules, int) {
	axiom := "F"
	rules := l.NewRules()
	rules.Add("F", "F+F--")

	return axiom, 8, rules, 4
}
func defaultTestState(position Point, stepSize, angle float64) State {
	return State{
		Position:  position,
		Direction: 180, // up
		StepSize:  stepSize,
		BrushSize: 2,
		Angle:     angle,
		Color:     0,
		Left:      -1.0,
	}
}
func BenchmarkTurtle(b *testing.B) {

	axiom, angle, rules, order := customtest1()
	lexer := l.NewDefaultLexer(rules)
	system := l.NewSystem(axiom, rules, lexer)

	// order override for testing
	order = 12

	// turtle lives in 0,0 -> 1,1 space; top left is 0,0
	position := NewPoint(0.5, 0.5)

	palette := NewPalette()
	initialState := defaultTestState(position, 1, angle)
	rotate := 0.0 // rotate turtle by this many degrees initially
	trt := NewTurtle(system, initialState, rotate, palette)

	// Execute the steps, keep track of min and max x,y coordinates for scaling
	delay := 0 * time.Millisecond
	for i := 0; i < order; i++ {
		trt.System().Step(delay)
	}
}
