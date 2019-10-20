package main

import (
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	triangle = []float32{
		-0.5, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right

		-0.5, 0.5, 0, // top
		0.5, 0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

const (
	RenderPolys = iota
	RenderPoints
	RenderLines
	RenderMax
)

var renderStyle = RenderPolys

func main() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	})

	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	program := initOpenGL()

	logrus.Info("starting loop")
	vao := makeVao(triangle)
	for !window.ShouldClose() {
		t := time.Now()
		draw(vao, window, program)
		time.Sleep(time.Second/time.Duration(60) - time.Since(t))
	}
	logrus.Info("closing")
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	switch renderStyle {
	case RenderPoints:
		gl.DrawArrays(gl.POINTS, 0, int32(len(triangle)))
	case RenderPolys:
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle))/3)
	case RenderLines:
		gl.DrawArrays(gl.LINE_STRIP, 0, int32(len(triangle)))
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
