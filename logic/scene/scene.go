package scene

import "context"

type Scene struct {
	id int

	init func(context.Context)

	handle func(context.Context) int
}

func MakeScene(id int, init func(context.Context), handle func(context.Context) int) *Scene {
	return &Scene{id, init, handle}
}
