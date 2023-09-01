package gravity

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  int     = 980
	WINDOW_HEIGHT int     = 640
	START_POS     float32 = 32
)

var (
	bodies    []*Body
	simSpeed  float32 = 1
	pastSpeed float32 = 0
)

func toggleFullScreen() {
	if rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
		rl.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	} else {
		monitor := rl.GetCurrentMonitor()
		rl.SetWindowSize(rl.GetMonitorWidth(monitor), rl.GetMonitorHeight(monitor))
		rl.ToggleFullscreen()
	}
}

func inputLogic(cam *rl.Camera3D) {
	// Toggle Cursor
	if rl.IsKeyPressed(rl.KeyOne) {
		rl.EnableCursor()
	} else if rl.IsKeyPressed(rl.KeyTwo) {
		rl.DisableCursor()
	}

	// Toggle Fullscreen
	if rl.IsKeyPressed(rl.KeyEscape) {
		toggleFullScreen()
	}

	// Fix Camera
	if rl.IsKeyPressed(rl.KeyR) {
		cam.Up = rl.NewVector3(0, 1, 0)
		cam.Target = rl.NewVector3(cam.Target.X, cam.Position.Y, cam.Target.Z)
	}

	// Return to the Starting Position
	if rl.IsKeyPressed(rl.KeyT) {
		cam.Position = rl.NewVector3(START_POS, START_POS, START_POS)
		cam.Target = rl.Vector3Zero()
		cam.Up = rl.NewVector3(0, 1, 0)
	}

	// Speed Up/Down
	if rl.IsKeyDown(rl.KeyThree) && simSpeed > 0 {
		simSpeed -= 0.05
	} else if rl.IsKeyDown(rl.KeyFour) && simSpeed < 5 {
		simSpeed += 0.05
	}

	// To fix a small bug
	if simSpeed < 0 {
		simSpeed = 0
	} else if simSpeed > 5 {
		simSpeed = 5
	}

	// Pause
	if rl.IsKeyPressed(rl.KeySpace) && simSpeed != 0 {
		pastSpeed = simSpeed
		simSpeed = 0
	} else if rl.IsKeyPressed(rl.KeySpace) {
		simSpeed = pastSpeed
		pastSpeed = 0
	}
}

func SimInit() rl.Camera {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	monitor := rl.GetCurrentMonitor()

	rl.InitWindow(
		int32(rl.GetMonitorWidth(monitor)), int32(rl.GetMonitorHeight(monitor)),
		"Gravity",
	)
	rl.ToggleFullscreen()

	rl.SetWindowMinSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	rl.SetTargetFPS(60)

	rl.DisableCursor()
	rl.SetExitKey(0)

	cam := rl.NewCamera3D(
		rl.NewVector3(START_POS, START_POS, START_POS), rl.Vector3Zero(),
		rl.NewVector3(0, 1, 0), 60, rl.CameraPerspective,
	)

	bodies = append(
		bodies, NewBody(float32(math.Pow10(12)), 5, rl.Yellow, rl.Vector3Zero(), rl.Vector3Zero()),
	)

	bodies = append(
		bodies, NewBody(float32(math.Pow10(8)), 1, rl.DarkGreen, rl.NewVector3(0, 2, 24), rl.NewVector3(1.25, 0, 0)),
	)

	return cam
}

func SimQuit() {
	rl.CloseWindow()
}

func Update(cam *rl.Camera3D) {
	// Input Logic
	inputLogic(cam)

	// Updates
	rl.UpdateCamera(cam, rl.CameraFree)

	for _, body := range bodies {
		body.Update(bodies, simSpeed)
	}
}

func Draw(cam rl.Camera3D) {
	rl.BeginDrawing()
	defer rl.EndDrawing()
	rl.BeginMode3D(cam)

	rl.ClearBackground(rl.Black)

	for _, body := range bodies {
		body.Draw()
	}

	rl.EndMode3D()

	// UI
	rl.DrawRectangle(0, 0, 320, 270, rl.NewColor(255, 255, 255, 25))
	rl.DrawText(
		fmt.Sprintf(
			"Esc: Toggle Fullscreen\nR: Look Straight\nT: Return to Original Position\n"+
				"1: Enable Mouse\n2: Disable Mouse\n"+
				"3: Speed Down\n4: Speed Up\nSpace: Pause/Resume\nCurrent Speed: %.2fx", simSpeed,
		),
		10, 5, 20, rl.RayWhite,
	)
}
