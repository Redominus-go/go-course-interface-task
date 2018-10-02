package main

import "testing"

func TestZeroAreaShapes(*testing.T) {
	t := triangle{
		height: 0,
		base:   0,
	}
	s := square{
		sideLength: 0,
	}
	var ta, sa float64
	ta = t.getArea()
	sa = s.getArea()
	if ta != 0 {
		t.Error("Expected 0, got ", ta)
	}
	if sa != 0 {
		t.Error("Expected 0 got ", sa)
	}
}
