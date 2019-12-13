package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/DanTulovsky/L-System/l"
	"github.com/DanTulovsky/Turtle/turtle"
)

func dragonCurve() (string, float64, l.Rules, int) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add("X", "-FX++FY-")
	rules.Add("Y", "+FX--FY+")
	rules.Add("F", "")

	return axiom, 8, rules, 8
}

func tree1() (string, float64, l.Rules, int) {
	axiom := "+++FX"
	rules := l.NewRules()
	rules.Add("X", "@[>4-FY]+FX")
	rules.Add("Y", "FX+FY-FX")

	return axiom, 12, rules, 8
}
func tree2() (string, float64, l.Rules, int) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "F[>8+X][>8-X]FX")
	rules.Add("F", "FF")

	return axiom, 14, rules, 6
}

func tree3() (string, float64, l.Rules, int) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "F-[[>6X]+X]+F[>6+FX]->X")
	rules.Add("F", "FF")

	return axiom, 16, rules, 7
}
func ytree() (string, float64, l.Rules, int) {
	axiom := "FX"
	rules := l.NewRules()
	rules.Add("X", "@[-FX]+FX")

	return axiom, 8, rules, 4
}

func sierpinkskiTriangle() (string, float64, l.Rules, int) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "FXF")
	rules.Add("X", "+FXRF-FLXR<F-F>LXR<F+")
	// rules.Add("L", ">6")
	// rules.Add("R", "<6")
	rules.Add("L", "")
	rules.Add("R", "")

	return axiom, 3, rules, 4
}

func sierpinkskiCarpet() (string, float64, l.Rules, int) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "<F+F-F-F-G+F+F+F-F")
	rules.Add("G", "GGG")

	return axiom, 4, rules, 4
}

func kochCurve() (string, float64, l.Rules, int) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "<F+F--F+F")

	return axiom, 6, rules, 4
}
func bush() (string, float64, l.Rules, int) {
	axiom := "F"

	rules := l.NewRules()
	rules.Add("F", "FF-[>5-F+F+F]+[>5+F-F-F]")

	return axiom, 16, rules, 4
}
func circular() (string, float64, l.Rules, int) {
	axiom := "X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X+X"

	rules := l.NewRules()
	rules.Add("X", "<[<F+F+F+F[<---X-Y]+++++F++++++++F-F-F-F]")
	rules.Add("Y", "[F+F+F+F[---Y]+++++F++++++++F-F-F-F]")

	return axiom, 24, rules, 4
}
func fasscurve2() (string, float64, l.Rules, int) {
	axiom := "-L"

	rules := l.NewRules()
	rules.Add("L", "<LFLF+RFR+FLFL-FRF-LFL-FR+F+RF-LFL-FRFRFR+")
	rules.Add("R", "<-LFLFLF+RFR+FL-F-LF+RFR+FLF+RFRF-LFL-FRFR")

	return axiom, 4, rules, 4
}
func lawninspring() (string, float64, l.Rules, int) {
	axiom := "%40X"

	rules := l.NewRules()
	rules.Add("X", "[+++++F-F-FZ]GX++++GY")
	rules.Add("Y", "[+++++F-F-FZ]GX----GY")
	rules.Add("Z", "W")
	rules.Add("W", "U")
	rules.Add("U", "[@.3[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G++[+++F]G]Z")
	rules.Add("F", "")
	rules.Add("G", "")

	return axiom, 16, rules, 8
}
func fractalplant() (string, float64, l.Rules, int) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "F+[[X]-X]-F[-FX]+X")
	rules.Add("F", "FF")

	return axiom, 14.4, rules, 6
}
func sphinx() (string, float64, l.Rules, int) {
	axiom := "X"
	rules := l.NewRules()
	rules.Add("X", "+FF-YFF+FF--FFF|X|F--YFFFYFFF|")
	rules.Add("Y", "-FF+XFF-FF++FFF|Y|F++XFFFXFFF|")
	rules.Add("F", "GG")
	rules.Add("G", "G>G")

	return axiom, 6, rules, 4
}
func pentaplexity() (string, float64, l.Rules, int) {
	axiom := "F++F++F++F++F"
	rules := l.NewRules()
	rules.Add("F", "F++F++F|F-F++F")

	return axiom, 10, rules, 4
}
func pentagrams() (string, float64, l.Rules, int) {
	axiom := "FX++FX++FX++FX++FX"
	rules := l.NewRules()
	rules.Add("X", "[++++@I1.618033989F@.618033989F!X!@I.618033989F]")

	return axiom, 10, rules, 4
}

func kitesdarts() (string, float64, l.Rules, int) {
	axiom := "%160WG+XG+WG+XG+WG+XG+WG+XG+WG+X"
	rules := l.NewRules()
	rules.Add("W", "[F][++@1.618033989F][++G---@.618033989G|X-Y|G|W]")
	rules.Add("X", "[F+++@1.618033989F][++@.618033989GZ|X|-G|W]")
	rules.Add("Y", "[+F][@1.618033989F][+G@.618033989|Y+X]")
	rules.Add("Z", "[-F][@1.618033989F][@.618033989G--WG|+Z]")
	rules.Add("F", "")

	return axiom, 10, rules, 4
}

func penrose() (string, float64, l.Rules, int) {
	axiom := "+WF--XF---YF--ZF"
	rules := l.NewRules()
	rules.Add("W", "YF++ZF----XF[-YF----WF]++")
	rules.Add("X", "+YF--ZF[---WF--XF]+")
	rules.Add("Y", "-WF++XF[+++YF++ZF]-")
	rules.Add("Z", "--YF++++WF[+ZF++++XF]--XF")
	rules.Add("F", "")

	return axiom, 10, rules, 4
}
func doublepenrose() (string, float64, l.Rules, int) {
	axiom := "%105[X][Y]++[X][Y]++[X][Y]++[X][Y]++[X][Y]"
	rules := l.NewRules()
	rules.Add("W", "YF++ZF----XF[-YF----WF]++")
	rules.Add("X", "+YF--ZF[---WF--XF]+")
	rules.Add("Y", "-WF++XF[+++YF++ZF]-")
	rules.Add("Z", "--YF++++WF[+ZF++++XF]--XF")
	rules.Add("F", ">")

	return axiom, 10, rules, 3
}
func spiral() (string, float64, l.Rules, int) {
	axiom := "X++X++X++X++|G|X++X++X++X"
	rules := l.NewRules()
	rules.Add("X", "[>12FX+++++@.7653668647>12F@I.7653668647[-----Y]+++++>12F]")
	rules.Add("Y", "[>12F+++++@.7653668647F@I.7653668647[-----Y]+++++>12F]")

	return axiom, 16, rules, 4
}
func island() (string, float64, l.Rules, int) {
	axiom := "%90F+F+F+F"
	rules := l.NewRules()
	rules.Add("F", "FFFF-F+F+F-F[>8-GFF+F+FF+F]FF")
	rules.Add("G", "@8G@I8")

	return axiom, 4, rules, 2
}
func simple() (string, float64, l.Rules, int) {
	axiom := "Y"
	rules := l.NewRules()
	rules.Add("Y", "XYX")

	return axiom, 8, rules, 4
}

func defaultState(position turtle.Point, stepSize, angle float64) turtle.State {
	return turtle.State{
		Position:  position,
		Direction: 180, // up
		StepSize:  stepSize,
		BrushSize: 2,
		Angle:     angle,
		Color:     0,
		Left:      -1.0,
	}
}

func main() {
	runtime.LockOSThread()
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	width, height := 1024, 768

	axiom, angle, rules, order := kitesdarts()
	lexer := l.NewDefaultLexer(rules)
	system := l.NewSystem(axiom, rules, lexer)

	// order override for testing
	order = 1

	// turtle lives in 0,0 -> 1,1 space; top left is 0,0
	position := turtle.NewPoint(0.5, 1.0)

	palette := turtle.NewPalette()
	initialState := defaultState(position, 1, angle)
	rotate := 0.0 // rotate turtle by this many degrees initially
	t := turtle.NewTurtle(system, initialState, rotate, palette)

	// Execute the steps, keep track of min and max x,y coordinates for scaling
	delay := 0 * time.Millisecond
	for i := 0; i < order; i++ {
		t.System().Step(delay)
	}

	// run the last step in its own thread for possible animation
	// func() {
	// 	time.Sleep(1*time.Second)
	// 	t.System().Step(delay)
	// }()

	// display some directions
	turtle.ShowDocs()

	// Display results
	// delay in drawing each segment of each step
	delay = 10 * time.Millisecond
	turtle.Run(t, width, height, delay)

}
