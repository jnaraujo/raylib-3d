package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	width  = 1280
	height = 720
)

func main() {
	rl.InitWindow(width, height, "Raylib test 3d")
	defer rl.CloseWindow()

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	camera := rl.NewCamera3D(rl.NewVector3(6, 6, 6), rl.NewVector3(0, 1, 0), rl.NewVector3(0, 1, 0), 45, rl.CameraPerspective)

	model := rl.LoadModel("./resources/ship-ocean-liner.obj")
	defer rl.UnloadModel(model)
	position := rl.NewVector3(0, 0, 0)

	water := rl.LoadModelFromMesh(rl.GenMeshPlane(20, 20, 4, 3))

	rl.DisableCursor()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		rl.BeginMode3D(camera)
		rl.DrawModel(water, position, 1, rl.Blue)
		rl.DrawModel(model, position, 0.4, rl.White)
		rl.DrawGrid(14, 1)
		rl.EndMode3D()

		rl.EndDrawing()
	}
}
