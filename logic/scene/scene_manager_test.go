package scene

import (
	"context"
	"testing"
)

func TestSceneManager(t *testing.T) {
	sceneManager := MakeSceneManager(1)
	sceneManager.AddScene(&Scene{1, func(ctx context.Context) {}, func(ctx context.Context) int { return 2 }})
	sceneManager.AddScene(&Scene{2, func(ctx context.Context) {}, func(ctx context.Context) int { return 1 }})

	sceneManager.Handle("some string", context.Background())
	if sceneManager.state != 2 {
		t.Fail()
	}
}
