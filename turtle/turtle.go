package turtle

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/golang-collections/collections/stack"

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
	Color     int
	Left      float64 // 1 or -1
}

// Turtle allows drawing on a canvas
type Turtle struct {
	state      State
	stateStack *stack.Stack
	system     *l.System
	palette    GradientTable
}

// NewTurtle returns a new turtle centered at pos
// rotate controls rotation of the entire drawing by n°. Positive values rotate counterclockwise,
// negative values rotate clockwise. With the default of 0, the turtle begins pointing up.
// For example, to start with the turtle pointing to the right, use rotate 90.
func NewTurtle(lsystem *l.System, state State, rotate float64, palette GradientTable) *Turtle {
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

// State returns the pointer to the current state attached to the turtle
func (t *Turtle) State() *State {
	return &t.state
}

// shiftColor returns the color n steps after the current color and sets it in state
func (t *Turtle) shiftColor(n int) color.Color {
	i := (t.state.Color + n) % 256
	t.state.Color = i
	// log.Printf("n is: %v; i is: %v", n, i)

	c := t.palette.GetInterpolatedColorFor(float64(i))
	// log.Println(c)
	return c
}

// setColor returns color n and sets it in state
func (t *Turtle) setColor(n int) color.Color {
	i := n % 255
	t.state.Color = i

	c := t.palette.GetInterpolatedColorFor(float64(i))
	// log.Printf("setting: %v -> %v", i, c)
	return c
}

// ShowDocs prints out usage documentation
func ShowDocs() {
	fmt.Println()
	fmt.Println("> Use the  '+' and '-' keys to zoom in and out.")
	fmt.Println("> Use the arrow keys (or w,s,a,d) to move render around.")
	fmt.Println("> Use 'n' to add another step.")
	fmt.Println("> Use 'e' and 'q' to rotate right and left.")
	fmt.Println()
}

// Draw makes the turtle draw on the canvas based on the state in the system
func (t *Turtle) Draw(cv *canvas.Canvas, w, h float64, zoom float64, xoffset, yoffset float64) {

	unitPixel := 10.0 * zoom

	// https://cgjennings.ca/articles/l-systems/
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
	cv.SetStrokeStyle(t.setColor(0))

	// save initial state for redraw
	oldstate := t.state

	// set turtle position based on screen size
	t.state.Position.X = t.state.Position.X*w + xoffset
	t.state.Position.Y = t.state.Position.Y*h + yoffset

	for e := t.System().State().Front(); e != nil; e = e.Next() {
		i := e.Value.(string)

		cv.BeginPath()
		cv.MoveTo(t.state.Position.X, t.state.Position.Y)

		switch {
		case i == "F":
			dirR := t.state.Direction * (math.Pi / 180)
			x := t.state.Position.X + t.state.StepSize*unitPixel*math.Sin(dirR)
			y := t.state.Position.Y + t.state.StepSize*unitPixel*math.Cos(dirR)

			cv.LineTo(x, y)
			t.state.Position.X = x
			t.state.Position.Y = y
		case i == "G":
			dirR := t.state.Direction * (math.Pi / 180)
			x := t.state.Position.X + t.state.StepSize*unitPixel*math.Sin(dirR)
			y := t.state.Position.Y + t.state.StepSize*unitPixel*math.Cos(dirR)

			cv.MoveTo(x, y)
			t.state.Position.X = x
			t.state.Position.Y = y

		case i == "-":
			t.state.Direction = t.state.Direction + t.state.Left*(360/t.state.Angle)
		case i == "+":
			t.state.Direction = t.state.Direction + (-t.state.Left)*(360/t.state.Angle)
		case i == "!":
			t.state.Left = -t.state.Left
		case i == "|":
			// Turns the turtle around (as close to 180° as the angle value allows).
			turnby := 0.0
			for turnby < 180 {
				turnby = turnby + 360/t.state.Angle
			}
			t.state.Direction = math.Mod(t.state.Direction+turnby, 360)
		case i[0] == '@':
			s := 0.6
			var err error

			index := 1
			reciprocal := false
			squareroot := false

			if len(i) > 1 {
				switch i[1] {
				case 'I':
					index++
					reciprocal = true
				case 'Q':
					index++
					squareroot = true
				}

				s, err = strconv.ParseFloat(i[index:], 64)
				if err != nil {
					panic(err)
				}
			}
			if reciprocal {
				s = 1 / s
			}
			if squareroot {
				s = math.Sqrt(s)
			}

			t.state.StepSize = t.state.StepSize * s
			t.state.BrushSize = t.state.BrushSize * s
		case i == "[":
			// push state
			t.stateStack.Push(t.state)
		case i == "]":
			// pop state
			t.state = (t.stateStack.Pop()).(State)

			c := t.palette.GetInterpolatedColorFor(float64(t.state.Color))
			cv.SetStrokeStyle(c)
			cv.MoveTo(t.state.Position.X, t.state.Position.Y)
		case i[0] == '<':
			n := 1
			if len(i) > 1 {
				var err error
				n, err = strconv.Atoi(i[1:])
				if err != nil {
					panic(err)
				}
			}
			cv.SetStrokeStyle(t.shiftColor(-n))
		case i[0] == '>':
			n := 1
			if len(i) > 1 {
				var err error
				n, err = strconv.Atoi(i[1:])
				if err != nil {
					panic(err)
				}
			}
			cv.SetStrokeStyle(t.shiftColor(n))
		case i[0] == '%':
			n := 0
			if len(i) > 1 {
				var err error
				n, err = strconv.Atoi(i[1:])
				if err != nil {
					panic(err)
				}
			}
			cv.SetStrokeStyle(t.setColor(n))
		}

		cv.SetLineWidth(t.state.BrushSize)
		cv.Stroke()
	}

	t.state = oldstate
}
