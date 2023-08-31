package main

import (
	src "newgravity/gravity/src"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	cam := src.SimInit()
	defer src.SimQuit()

	for !rl.WindowShouldClose() {
		// Update
		src.Update(&cam)

		// Draw
		src.Draw(cam)
	}
}
