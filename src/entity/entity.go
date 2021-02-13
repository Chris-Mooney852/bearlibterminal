package entity

import (
	blt "bearlibterminal"
)

type GameEntity struct {
	X     int
	Y     int
	Layer int
	Char  string
	Color string
}

// Move the entity by the amount (dx, dy)
func (e *GameEntity) Move(dx int, dy int) {
	e.X += dx
	e.Y += dy
}

// Draw the entity to its location
func (e *GameEntity) Draw(mapX int, mapY int) {
	blt.Layer(e.Layer)
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(mapX, mapY, e.Char)
}

// Remove entity from screen
func (e *GameEntity) Clear(mapX int, mapY int) {
	blt.Layer(e.Layer)
	blt.Print(mapX, mapY, " ")
}
