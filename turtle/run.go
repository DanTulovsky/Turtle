package turtle

import (
	"log"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

var (
	// for cursor position
	mx, my float64

	// framebuffer widht and height
	ww, wh int

	// zoom controls zooming in and out
	zoom float64
	// control how much to change the zoom
	zoomfactor float64

	// offsets control the relative initial position of the render on the screeen
	xoffset, yoffset float64
)

// Run runs the simulation with the given turtle
func Run(t *Turtle, width, height int) {

	// init GLFW
	err := glfw.Init()
	if err != nil {
		log.Fatalf("Error initializing GLFW: %v", err)
	}
	defer glfw.Terminate()

	// the stencil size setting is required for the canvas to work
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	// glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// create window
	window, err := glfw.CreateWindow(width, height, "Turtles away!", nil, nil)
	if err != nil {
		log.Fatalf("Error creating window: %v", err)
	}
	window.MakeContextCurrent()

	// init GL
	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	// set vsync on, enable multisample (if available)
	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	// load GL backend
	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		mx, my = xpos, ypos
	})

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, heigh int) {
		ww, wh = window.GetFramebufferSize()
		backend.SetBounds(0, 0, ww, wh)
	})

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press, glfw.Repeat:
			switch key {
			case glfw.KeyKPAdd, glfw.KeyEqual:
				zoom = zoom + zoomfactor
			case glfw.KeyKPSubtract, glfw.KeyMinus:
				if zoom < zoomfactor*2 {
					zoomfactor = zoomfactor / 10
				}
				zoom = zoom - zoomfactor
				log.Println(zoom)
			case glfw.KeyLeft, glfw.KeyA:
				xoffset = xoffset - 10
			case glfw.KeyRight, glfw.KeyD:
				xoffset = xoffset + 10
			case glfw.KeyUp, glfw.KeyW:
				yoffset = yoffset - 10
			case glfw.KeyDown, glfw.KeyS:
				yoffset = yoffset + 10
			case glfw.KeyN:
				delay := 0 * time.Millisecond
				t.System().Step(delay)
			case glfw.KeyE:
				t.State().Direction = t.State().Direction + t.State().Left*10
			case glfw.KeyQ:
				t.State().Direction = t.State().Direction - t.State().Left*10
			}
		}
	})

	ww, wh = window.GetFramebufferSize()
	backend.SetBounds(0, 0, ww, wh)
	cv := canvas.New(backend)
	zoom = 1.0
	zoomfactor = 0.1 // change zoom by this much every time

	for !window.ShouldClose() {
		window.MakeContextCurrent()
		glfw.PollEvents()

		// call the run function to do all the drawing
		t.Draw(cv, float64(ww), float64(wh), zoom, xoffset, yoffset)

		// swap back and front buffer
		window.SwapBuffers()
	}
}
