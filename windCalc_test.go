package main

import (
	"fmt"
	"testing"
)

func TestCalcwind(t *testing.T) {
	fmt.Println("Testing calcWind")

	tt := []struct {
		name      string
		offset    int
		direction directionT
		b         bearingT
	}{
		{name: "p:20 - w:40 - o:20R", offset: 20, direction: right, b: bearingT{plane: 20, wind: 40}},
		{name: "p:40 - w:20 - o:20L", offset: 20, direction: left, b: bearingT{plane: 40, wind: 20}},
		{name: "p:180 - w:90 - o:90L", offset: 90, direction: left, b: bearingT{plane: 180, wind: 90}},
		{name: "p:180 - w:270 - o:90R", offset: 90, direction: right, b: bearingT{plane: 180, wind: 270}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.b.calcWind()
			if tc.b.offset != tc.offset || tc.b.direction != tc.direction {
				t.Fatalf("Expected offset %d direction %d, got offset %d direction %d\n",
					tc.offset, tc.direction, tc.b.offset, tc.b.direction)
			}
		})
	}
}
