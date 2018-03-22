package plane

// A cell is a section of a plane
type Cell struct {
	X, Y       int
	Attraction float64
	Points     []Point
	plane      *Plane
}

// A point is a thing in a cell on a plane
type Point interface {
	SetLocation(x, y int)
}

// A plane is a collection of cells
type Plane struct {
	cells [][]*Cell
}

func (p *Plane) ShakeOut() {
	var adjusted bool
	for _, rows := range p.cells {
		for _, cell := range rows {
			newx, newy := cell.GetBestNeighbor()
			if newx == -1 && newy == -1 {
				continue // there is no better neighbor
			}

			adjusted = true
			for _, point := range cell.Points {
				p.cells[newx][newy].Points = append(p.cells[newx][newy].Points, point)
				point.SetLocation(newx, newy)
			}
			cell.Points = nil // They all moved.
		}
	}

	if adjusted {
		// Recurse until we are no longer adjusting
		p.ShakeOut()
	}
}

// If the cell is the most attractive thing in the neighboorhood (or is tied) it will
// return -1, -1
func (c *Cell) GetBestNeighbor() (x, y int) {
	// Look at all the neighbours
	// + + +
	// + ! +
	// + + +
	// And choose the most attractive one

	xs := []int{c.X - 1, c.X, c.X + 1}
	ys := []int{c.Y - 1, c.Y, c.Y + 1}

	bestX, bestY := -1, -1
	mostAttractive := c.Attraction

	for _, x := range xs {
		if x < 0 || x >= len(c.plane.cells) {
			// If we are off the grid, ignore
			continue
		}
		for _, y := range ys {
			if y < 0 || y >= len(c.plane.cells[x]) {
				continue
			}

			comparison := c.plane.cells[x][y]

			if comparison.Attraction > mostAttractive {
				bestX = comparison.X
				bestY = comparison.Y
				mostAttractive = comparison.Attraction
			}
		}
	}

	return bestX, bestY

}
