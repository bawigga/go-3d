package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func keyboardHandler(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyQ, glfw.KeyEscape:
		w.SetShouldClose(true)
	}
}
func onWindowResize(w *glfw.Window, height int, width int) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	w.SwapBuffers()
}

func main() {

	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Can't get glfw!")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Triangles!!!", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(keyboardHandler)
	window.SetSizeCallback(onWindowResize)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	gl.ClearColor(.5, 0, 0, .2)

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()

		glfw.PollEvents()
	}
}
