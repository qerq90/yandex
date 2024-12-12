package scene

import (
	"context"
	"fmt"
)

type SceneManager struct {
	state  int // scene id
	scenes map[int]*Scene
}

func MakeSceneManager(state int) *SceneManager {
	map_ := make(map[int]*Scene)
	return &SceneManager{state, map_}
}

func (sc *SceneManager) Handle(input string, c context.Context) {
	scene := sc.scenes[sc.state]
	c = context.WithValue(c, "message", input)
	nextSceneId := scene.handle(c)

	sc.state = nextSceneId
	newScene := sc.scenes[sc.state]
	newScene.init(c)
}

func (sc *SceneManager) AddScene(scene *Scene) error {
	for _, sc := range sc.scenes {
		if scene.id == sc.id {
			return fmt.Errorf("Scene with that id already exist - id[%d]", scene.id)
		}
	}

	sc.scenes[scene.id] = scene
	return nil
}
