package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/DanTulovsky/L-System/l"
	"github.com/DanTulovsky/Turtle/turtle"
)

func dragonCurve() (string, float64, l.Rules) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add("X", "-FX++FY-")
	rules.Add("Y", "+FX--FY+")
	rules.Add("F", "")

	return axiom, 8, rules
}

func tree1() (string, float64, l.Rules) {
	axiom := "+++FX"
	rules := l.NewRules()
	rules.Add("X", "@[>4-FY]+FX")
	rules.Add("Y", "FX+FY-FX")

	return axiom, 12, rules
}

func tree3() (string, float64, l.Rules) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "F-[[>6X]+X]+F[>6+FX]->X")
	rules.Add("F", "FF")

	return axiom, 16, rules
}
func ytree() (string, float64, l.Rules) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add("X", "@[-FX]+FX")

	return axiom, 8, rules
}

func sierpinkskiTriangle() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "FXF")
	rules.Add("X", "+FXRF-FLXR<F-F>LXR<F+")
	// rules.Add("L", ">6")
	// rules.Add("R", "<6")
	rules.Add("L", "")
	rules.Add("R", "")

	return axiom, 3, rules
}

func sierpinkskiCarpet() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "<F+F-F-F-G+F+F+F-F")
	rules.Add("G", "GGG")

	return axiom, 4, rules
}

func kochCurve() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "<F+F--F+F")

	return axiom, 6, rules
}
func bush() (string, float64, l.Rules) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "FF-[>5-F+F+F]+[>5+F-F-F]")

	return axiom, 16, rules
}
func circular() (string, float64, l.Rules) {
	axiom := "X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X"

	rules := l.NewRules()
	rules.Add("X", "<[<F+F+F+F[<---X-Y]+++++F++++++++F-F-F-F]")
	rules.Add("Y", "[F+F+F+F[---Y]+++++F++++++++F-F-F-F]")

	return axiom, 24, rules
}
func fasscurve2() (string, float64, l.Rules) {
	axiom := "-L"

	rules := l.NewRules()
	rules.Add("L", "<LFLF+RFR+FLFL-FRF-LFL-FR+F+RF-LFL-FRFRFR+")
	rules.Add("R", "<-LFLFLF+RFR+FL-F-LF+RFR+FLF+RFRF-LFL-FRFR")

	return axiom, 4, rules
}
func lawninspring() (string, float64, l.Rules) {
	axiom := "%40X"

	rules := l.NewRules()
	rules.Add("X", "[+++++F-F-FZ]GX++++GY")
	rules.Add("Y", "[+++++F-F-FZ]GX----GY")
	rules.Add("Z", "W")
	rules.Add("W", "U")
	rules.Add("U", "[@.3[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G]Z")
	rules.Add("F", "")
	rules.Add("G", "")

	return axiom, 16, rules
}
func fractalplant() (string, float64, l.Rules) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "F+[[X]-X]-F[-FX]+X")
	rules.Add("F", "FF")

	return axiom, 14.4, rules
}
func simple() (string, float64, l.Rules) {
	axiom := "Y"
	rules := l.NewRules()
	rules.Add("Y", "XYX")

	return axiom, 8, rules
}
func main() {
	runtime.LockOSThread()
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	width, height := 1024, 768

	axiom, angle, rules := bush()
	lexer := l.NewDefaultLexer(rules)
	system := l.NewSystem(axiom, rules, lexer)

	order := 4
	// turtle lives in 0,0 -> 1,1 space; top left is 0,0
	xstart, ystart := 0.2, 1.0

	palette := colorful.FastWarmPalette(9)
	initialState := turtle.State{
		Position:  turtle.NewPoint(xstart, ystart),
		Direction: 180, // up
		StepSize:  0.2,
		BrushSize: 1,
		Angle:     angle,
	}
	rotate := 0.0
	t := turtle.NewTurtle(system, initialState, rotate, palette)

	// Execute the steps
	delay := 0 * time.Millisecond
	// TODO: Put in own thread and add locking as needed
	go func() {
		// time.Sleep(1 * time.Second)
		t.Step(order, delay)
	}()

	// Display results
	turtle.Run(t, width, height)
}
