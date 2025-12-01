package day1

import "testing"

func TestTurnDial(t *testing.T) {
	tests := []struct {
		name           string
		currentDial    int
		rotation       int
		expectedDial   int
		expectedClicks int
	}{
		{"right from 50", 50, 10, 60, 0},
		{"left from 50", 50, -10, 40, 0},
		{"wrap around right", 95, 10, 5, 1},
		{"wrap around left", 5, -10, 95, 1},
		{"reach zero", 50, 50, 0, 1},
		{"99+1 -> (0,1)", 99, 1, 0, 1},
		{"99+2 -> (1,1)", 99, 1, 0, 1},
		{"1-1 -> (0,1)", 1, -1, 0, 1},
		{"1-2 -> (99,1)", 1, -2, 99, 1},
		{"99+100 -> (99,1)", 99, 100, 99, 1},
		{"99+101 -> (0,2)", 99, 101, 0, 2},
		{"1-101 -> (0,2)", 1, -101, 0, 2},
		{"0-1 -> (99,0)", 0, -1, 99, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultDial, resultClicks := turnDial(tt.currentDial, tt.rotation)
			if resultDial != tt.expectedDial || resultClicks != tt.expectedClicks {
				t.Errorf("turnDial(%d, %d) = (%d, %d); want (%d, %d)",
					tt.currentDial, tt.rotation, resultDial, resultClicks, tt.expectedDial, tt.expectedClicks)
			}
		})
	}
}
