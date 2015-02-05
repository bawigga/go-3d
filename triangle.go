package main

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
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

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	gl.ClearColor(1, 0, 0)

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
