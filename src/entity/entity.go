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
func (e *GameEntity) Draw() {
	blt.Layer(e.Layer)
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(e.X, e.Y, e.Char)
}

// Remove entity from screen
func (e *GameEntity) Clear() {
	blt.Layer(e.Layer)
	blt.Print(e.X, e.Y, " ")
}
