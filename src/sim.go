package gravity

import rl "github.com/gen2brain/raylib-go/raylib"

func SimInit() {
	rl.InitWindow(980, 640, "Window")
	rl.SetTargetFPS(60)
}

func SimQuit() {
	rl.CloseWindow()
}

func Update() {

}

func Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.EndDrawing()
}
