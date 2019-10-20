package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/sirupsen/logrus"
)

const (
	width  = 500
	height = 500
)

func initGLFW() *glfw.Window {

	logrus.Debug("glfw: initializing")

	if err := glfw.Init(); err != nil {
		panic("failed to init glfw")
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	logrus.Debug("glfw: creating window")
	window, err := glfw.CreateWindow(500, 500, "Triangles!!!", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(keyboardHandler)
	window.SetSizeCallback(onWindowResize)
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	logrus.Debug("opengl: initializing")
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	logrus.Debugf("opengl: version: %s", version)

	logrus.Debug("opengl: compiling shader: vertexShader")
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	logrus.Debug("opengl: compiling shader: fragmentShader")
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	logrus.Debug("opengl: attaching shaders")
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func errorCallback(err glfw.ErrorCode, desc string) {
	logrus.Errorf("%v: %v\n", err, desc)
}

func keyboardHandler(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	logrus.WithFields(logrus.Fields{
		"key":      key,
		"scancode": scancode,
		"action":   action,
		"mods":     mods,
	}).Debugf("key: %v", key)
	switch key {
	case glfw.KeyQ, glfw.KeyEscape:
		w.SetShouldClose(true)
	case glfw.KeyR:
		if action == 1 {
			toggleRenderType()
		}
	}
}

func onWindowResize(w *glfw.Window, height int, width int) {
	logrus.Debug("onWindowResize")
	gl.Clear(gl.COLOR_BUFFER_BIT)
	w.SwapBuffers()
}

func toggleRenderType() {
	renderStyle++
	if renderStyle == RenderMax {
		renderStyle = 0
	}
}
