package plane

import (
	"testing"
)

func TestNewPlane(t *testing.T) {
	pl := NewPlane(3, 3)

	pl.GetCell(1, 1).attraction = 1

	po := NewPoint(pl)

	po.SetLocation(0, 0)
	po.SetLocation(2, 2)

	if len(pl.cells[0][0].points) != 1 {
		t.Errorf("failed to set point in location 0,0")
	}

	pl.Converge()

	if len(pl.cells[0][0].points) == 1 {
		t.Errorf("failed to remove point from location 0,0")
	}

	if len(pl.cells[1][1].points) != 2 {
		t.Errorf("failed to move point to location 1,1")
	}
}
