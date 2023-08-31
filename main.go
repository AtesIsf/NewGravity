package main

import (
	src "newgravity/pkg/src"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	src.SimInit()
	defer src.SimQuit()

	for !rl.WindowShouldClose() {
		// Update
		src.Update()

		// Draw
		src.Draw()
	}
}
