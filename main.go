package main

import (
	"runtime"

	"github.com/DanTulovsky/L-System/l"
	"github.com/DanTulovsky/Turtle/turtle"
)

func dragonCurve() (string, float64, l.Rules) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add('X', "-FX++FY-")
	rules.Add('Y', "+FX--FY+")
	rules.Add('F', "")

	return axiom, 8, rules
}

func tree1() (string, float64, l.Rules) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add('X', "@[+FX][-FX]")

	return axiom, 8, rules
}

func ytree() (string, float64, l.Rules) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add('X', "@[-FX]+FX")

	return axiom, 8, rules
}

func sierpinkskiTriangle() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add('F', "FXF")
	rules.Add('X', "+FXRF-FLXR<F-F>LXR<F+")
	// rules.Add('L', ">6")
	// rules.Add('R', "<6")
	rules.Add('L', "")
	rules.Add('R', "")

	return axiom, 3, rules
}

func sierpinkskiCarpet() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add('F', "<F+F-F-F-G+F+F+F-F")
	rules.Add('G', "GGG")

	return axiom, 4, rules
}

func kochCurve() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add('F', "<F+F--F+F")

	return axiom, 6, rules
}
func brush() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add('F', "FF-[>5-F+F+F]+[>5+F-F-F]")

	return axiom, 16, rules
}
func main() {
	runtime.LockOSThread()
	width, height := 1300, 1000

	// axiom, rules := dragonCurve()
	axiom, angle, rules := brush()
	system := l.NewSystem(axiom, rules)

	order := 4
	// xstart, ystart := float64(width)/2.0, float64(height)/2.0
	xstart, ystart := float64(width)/2.0, float64(height)
	// xstart, ystart := float64(width), float64(height)
	initialState := turtle.State{
		Position:  turtle.NewPoint(xstart, ystart),
		Direction: 180, // up
		StepSize:  0.1,
		BrushSize: 1,
		Angle:     angle,
	}
	rotate := 0.0
	t := turtle.NewTurtle(system, initialState, rotate)

	// Execute the steps
	t.Step(order)

	// Display results
	turtle.Run(t, width, height)
}
