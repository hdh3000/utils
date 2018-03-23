package plane

// A cell is a section of a plane
type Cell struct {
	x, y         int
	attraction   float64
	points       []Point
	plane        *Plane
	bestNeighbor *struct{ x, y int }
}

// A point is a thing in a cell on a plane
type Point interface {
	SetLocation(x, y int)
}

// A plane is a collection of cells
type Plane struct {
	cells [][]Cell
}

func NewPlane(width, height int) *Plane {
	p := &Plane{}
	for x := 0; x < width; x++ {
		p.cells = append(p.cells, make([]Cell, height))
		for y := range p.cells[x] {
			p.cells[x][y].plane = p
			p.cells[x][y].x = x
			p.cells[x][y].y = y
		}
	}
	return p
}

func NewPoint(p *Plane) Point {
	return &point{x: -1, y: -1, plane: p}
}

type point struct {
	plane *Plane
	x, y  int
}

func (p *point) SetLocation(x, y int) {
	if p.x == x && p.y == y {
		return
	}
	p.x = x
	p.y = y
	p.plane.cells[x][y].points = append(p.plane.cells[x][y].points, p)
}

func (p *Plane) Converge() {
	var adjusted bool
	for x := range p.cells {
		for y := range p.cells[x] {
			cell := p.GetCell(x, y)
			newx, newy := cell.GetBestNeighbor()
			if newx == -1 && newy == -1 {
				continue // there is no better neighbor
			}

			if len(cell.points) > 0 {
				// The cell has a better neighbor, AND it has points in it that need moving!
				adjusted = true
				newCell := p.GetCell(newx, newy)
				newCell.points = append(newCell.points, cell.points...)
				cell.points = nil // They all moved.
			}
		}
	}

	if adjusted {
		// Recurse until we are no longer adjusting
		p.Converge()
	}
}

func (p *Plane) GetCell(x, y int) *Cell {
	return &p.cells[x][y]
}

// If the cell is the most attractive thing in the neighborhood (or is tied) best neighbor will return -1, -1
func (c *Cell) GetBestNeighbor() (x, y int) {
	if c.bestNeighbor != nil {
		// Don't compute it twice
		return c.bestNeighbor.x, c.bestNeighbor.y
	}
	// Look at all the neighbours
	// + + +
	// + ! +
	// + + +
	// And choose the most attractive one

	xs := []int{c.x - 1, c.x, c.x + 1}
	ys := []int{c.y - 1, c.y, c.y + 1}

	bestX, bestY := -1, -1
	mostAttractive := c.attraction

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

			if comparison.attraction > mostAttractive {
				bestX = comparison.x
				bestY = comparison.y
				mostAttractive = comparison.attraction
			}
		}
	}

	c.bestNeighbor = &struct {
		x, y int
	}{
		x: bestX,
		y: bestY,
	}

	return bestX, bestY

}
