package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/tuckersGo/puzzle/global"
	"github.com/tuckersGo/puzzle/scenemanager"
	"github.com/tuckersGo/puzzle/scenes"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "Puzzle")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
