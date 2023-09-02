package gravity

import (
	"math"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Id       int
	Mass     float32
	Radius   float32
	Color    rl.Color
	Position rl.Vector3
	Velocity rl.Vector3
	Force    rl.Vector3
}

var (
	currId int = 0
)

func NewBody(mass float32, radius float32, color rl.Color, pos rl.Vector3, v rl.Vector3) *Body {
	body := &Body{
		Id:       currId,
		Mass:     mass,
		Radius:   radius,
		Color:    color,
		Position: pos,
		Velocity: v,
		Force:    rl.Vector3Zero(),
	}
	currId++
	return body
}

func (body *Body) Update(bodies []*Body, simSpeed float32) {
	dTime := rl.GetFrameTime() * simSpeed
	var distance rl.Vector3
	var force float32 = 0

	for _, b := range bodies {
		if body.Id == b.Id {
			continue
		}

		distance = rl.Vector3Multiply(rl.Vector3Subtract(body.Position, b.Position), -1)
		force += 6.67 * float32(math.Pow10(-11)) * body.Mass * b.Mass / float32(math.Pow(float64(rl.Vector3Length(distance)), 2))
		body.Force = rl.Vector3Multiply(rl.Vector3Normalize(distance), force)
	}

	body.Velocity = rl.Vector3Add(body.Velocity, rl.Vector3Multiply(body.Force, 1/body.Mass*dTime))
	body.Position = rl.Vector3Add(rl.Vector3Scale(body.Velocity, dTime), body.Position)
}

func (body *Body) ConcurrentUpdate(bodies []*Body, simSpeed float32, wg *sync.WaitGroup) {
	defer wg.Done()
	body.Update(bodies, simSpeed)
}

func (body *Body) Draw() {
	rl.DrawSphereEx(body.Position, body.Radius, 20, 20, body.Color)
	rl.DrawSphereWires(body.Position, body.Radius, 20, 20, rl.RayWhite)
}
