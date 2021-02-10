package gamemap

type Tile struct {
	Blocked     bool // Blocks Movement
	BlocksSight bool // Blocks Sight
}

type Map struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

// Setup game map
func (m *Map) InitializeMap() {
	m.Tiles = make([][]*Tile, m.Width)

	// Make walls
	for i := range m.Tiles {
		m.Tiles[i] = make([]*Tile, m.Height)
	}

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			if x == 0 || x == m.Width-1 || y == 0 || y == m.Height-1 {
				m.Tiles[x][y] = &Tile{true, true}
			} else {
				m.Tiles[x][y] = &Tile{false, false}
			}
		}
	}
}

func (m *Map) IsBlocked(x int, y int) bool {
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}
