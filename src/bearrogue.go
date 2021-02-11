package main

import (
	"C"
	blt "bearlibterminal"
	"strconv"

	"bearrogue/src/entity"
	"bearrogue/src/gamemap"
)

const (
	WindowSizeX = 100
	WindowSizeY = 35
	MAP_WIDTH   = WindowSizeX
	MAP_HEIGHT  = WindowSizeY
	TITLE       = "BearRogue"
)

var (
	player   *entity.GameEntity
	entities []*entity.GameEntity
	gameMap  *gamemap.Map
)

func init() {
	blt.Open()

	// Setup window config settings
	size := "size=" + strconv.Itoa(WindowSizeX) + "x" + strconv.Itoa(WindowSizeY)
	title := "title='" + TITLE + "', "
	window := "window: " + title + size

	blt.Set(window)
	blt.Clear()

	// Create player
	player = &entity.GameEntity{X: 1, Y: 1, Layer: 1, Char: "@", Color: "white"}

	// Create NPC's
	npc := &entity.GameEntity{X: 10, Y: 10, Layer: 0, Char: "N", Color: "red"}

	// Save entities
	entities = append(entities, player, npc)

	// Create GameMap
	gameMap = &gamemap.Map{Width: MAP_WIDTH, Height: MAP_HEIGHT}
	gameMap.InitializeMap()

	playerX, playerY := gameMap.GenerateCavern()
	player.X = playerX
	player.Y = playerY
}

func main() {

	renderAll()

	for {
		blt.Refresh()

		key := blt.Read()

		clearEntities()

		if key != blt.TK_CLOSE {
			handleInput(key, player)
		} else {
			break
		}

		renderAll()
	}

	// If the user hits the clone button, close the window and exit
	blt.Close()
}

// Handle character input
func handleInput(key int, player *entity.GameEntity) {
	var (
		dx, dy int
	)

	switch key {
	case blt.TK_RIGHT:
		dx, dy = 1, 0
	case blt.TK_LEFT:
		dx, dy = -1, 0
	case blt.TK_UP:
		dx, dy = 0, -1
	case blt.TK_DOWN:
		dx, dy = 0, 1
	}

	player.Move(dx, dy)
}

// Clear entities from the screen
func clearEntities() {
	for _, e := range entities {
		e.Clear()
	}
}

// Render all game elements
func renderAll() {
	renderMap()
	renderEntities()
}

// Draw entities to the screen
func renderEntities() {
	for _, e := range entities {
		e.Draw()
	}
}

func renderMap() {
	for x := 0; x < gameMap.Width; x++ {
		for y := 0; y < gameMap.Height; y++ {
			if gameMap.Tiles[x][y].Blocked == true {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x, y, "#")
			} else {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x, y, ".")
			}
		}
	}
}
