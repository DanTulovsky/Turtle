package turtle

import (
	"image/color"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/golang-collections/collections/stack"
	"github.com/lucasb-eyer/go-colorful"

	"github.com/DanTulovsky/L-System/l"
	"github.com/tfriedel6/canvas"
)

// Point is a point on a 2D plane
type Point struct {
	X, Y float64
}

// NewPoint returns a new point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// State is the turtle state
type State struct {
	Position  Point
	Direction float64 // in degrees
	StepSize  float64
	BrushSize float64
	Angle     float64 // Sets the number of turns that make up a complete circle to n. (Each turn will be by 360°/n.)
	color     int
}

// Turtle allows drawing on a canvas
type Turtle struct {
	state      State
	stateStack *stack.Stack
	system     *l.System
	palette    []colorful.Color
}

// NewTurtle returns a new turtle centered at pos
// rotate controls rotation of the entire drawing by n°. Positive values rotate counterclockwise,
// negative values rotate clockwise. With the default of 0, the turtle begins pointing up.
// For example, to start with the turtle pointing to the right, use rotate 90.
func NewTurtle(lsystem *l.System, state State, rotate float64, palette []colorful.Color) *Turtle {
	t := &Turtle{
		state:      state,
		system:     lsystem,
		stateStack: stack.New(),
		palette:    palette,
	}

	t.state.Direction = math.Mod(t.state.Direction+rotate, 360)
	return t
}

// Step makes the turtle take n steps
func (t *Turtle) Step(n int, delay time.Duration) {
	log.Println("Calculating system...")

	for i := 0; i < n; i++ {
		t.system.Step(delay)

	}
	log.Println("Finished calculating system...")
}

// System returns the system attached to the turtle
func (t *Turtle) System() *l.System {
	return t.system
}

// updateColor returns the color n steps after the current color and sets it in state
func (t *Turtle) updateColor(n int) color.Color {
	i := (t.state.color + n) % len(t.palette)
	t.state.color = i
	return t.palette[i]
}

// Draw makes the turtle draw on the canvas based on the state in the system
func (t *Turtle) Draw(cv *canvas.Canvas, w, h float64) {

	unitPixel := 50.0

	// F: move forward one step with pen down
	// G: Moves the turtle forward 1 step with the pen up, leaving no mark.
	// -: turn right 45
	// +: turn left 45
	// @: change the step size by 0.6  // TODO: Take arbitrary number after this to multiply by
	// [: write current state to stack
	// ]: pop last state from stack

	// clear screen
	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, w, h)
	cv.SetStrokeStyle(t.palette[t.state.color])

	lstate := t.state
	// set turtle position based on screen size
	lstate.Position.X = lstate.Position.X * w
	lstate.Position.Y = lstate.Position.Y * h

	for e := t.System().State().Front(); e != nil; e = e.Next() {
		i := e.Value.(string)

		cv.BeginPath()
		cv.MoveTo(lstate.Position.X, lstate.Position.Y)
		switch {
		case i == "F":
			dirR := lstate.Direction * (math.Pi / 180)
			x := lstate.Position.X + lstate.StepSize*unitPixel*math.Sin(dirR)
			y := lstate.Position.Y + lstate.StepSize*unitPixel*math.Cos(dirR)

			cv.LineTo(x, y)
			lstate.Position.X = x
			lstate.Position.Y = y
		case i == "G":
			dirR := lstate.Direction * (math.Pi / 180)
			x := lstate.Position.X + lstate.StepSize*unitPixel*math.Sin(dirR)
			y := lstate.Position.Y + lstate.StepSize*unitPixel*math.Cos(dirR)

			cv.MoveTo(x, y)
			lstate.Position.X = x
			lstate.Position.Y = y

		case i == "-":
			lstate.Direction = lstate.Direction - 360/lstate.Angle
		case i == "+":
			lstate.Direction = lstate.Direction + 360/lstate.Angle
		case i == "@":
			lstate.StepSize = lstate.StepSize * 0.6
			lstate.BrushSize = lstate.BrushSize * 0.6
		case i == "[":
			// push state
			t.stateStack.Push(lstate)
		case i == "]":
			// pop state
			lstate = (t.stateStack.Pop()).(State)
			cv.MoveTo(lstate.Position.X, lstate.Position.Y)
		case i[0] == '<':
			n := 1
			if len(i) > 1 {
				var err error
				n, err = strconv.Atoi(i[1:])
				if err != nil {
					panic(err)
				}
			}
			cv.SetStrokeStyle(t.updateColor(n))
		case i[0] == '>':
			n := 1
			if len(i) > 1 {
				var err error
				n, err = strconv.Atoi(i[1:])
				if err != nil {
					panic(err)
				}
			}
			cv.SetStrokeStyle(t.updateColor(n))
		}

		cv.SetLineWidth(lstate.BrushSize)
		cv.Stroke()
	}
}
