package turtle

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

// for cursor position
var mx, my float64

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

	// initialize canvas with zero size, since size is set in main loop
	cv := canvas.New(backend)

	for !window.ShouldClose() {
		window.MakeContextCurrent()
		glfw.PollEvents()

		// set canvas size
		ww, wh := window.GetFramebufferSize()
		backend.SetBounds(0, 0, ww, wh)

		// call the run function to do all the drawing
		t.Draw(cv, float64(ww), float64(wh))

		// swap back and front buffer
		window.SwapBuffers()
	}
}
