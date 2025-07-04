package grid

type Position struct {
	Row int
	Col int
}

type Grid [][]rune

func (g Grid) DeepCopy() Grid {
	newGrid := make([][]rune, len(g))
	for i := range g {
		newGrid[i] = make([]rune, len(g[i]))
		copy(newGrid[i], g[i])
	}
	return newGrid
}

func (g Grid) CheckOutOfBounds(position Position) bool {
	// Check row bounds
	if position.Row < 0 || position.Row >= len(g) {
		return true
	}

	// Check column bounds
	if position.Col < 0 || position.Col >= len(g[position.Row]) {
		return true
	}

	return false
}

func NewGrid(rows, cols int, defaultValue rune) Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = defaultValue
		}
	}
	return grid
}
