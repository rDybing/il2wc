/***********************************************************************************************************************

  windCalc.go
  License: MIT
  Copyright (c) 2021 Roy Dybing

  github	: rDybing
  Linked In	: Roy Dybing
  Twitter	: @DybingRoy

  Full license text in README.md

***********************************************************************************************************************/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type directionT byte

const (
	left directionT = iota
	right
)

type entityT byte

const (
	plane entityT = iota
	wind
)

type bearingT struct {
	plane     int
	wind      int
	offset    int
	direction directionT
}

func main() {
	var quit bool
	var b bearingT
	for {
		if quit = b.getInput(plane); quit {
			break
		}
		if quit = b.getInput(wind); quit {
			break
		}
		b.calcWind()
		b.showOffset()
	}
}

func (b *bearingT) getInput(what entityT) bool {
	var whatStr string
	switch what {
	case plane:
		whatStr = "plane"
	case wind:
		whatStr = "wind"
	}
	var input string
	var quit bool
	var done bool
	for !done {
		fmt.Printf("Enter %s bearing (q to quit):\n", whatStr)
		fmt.Scanf("%s\n", &input)
		input = stripNewline(input)
		if input == "q" || input == "Q" {
			quit = true
			done = true
		} else {
			var doOver bool
			dir, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Numbers only!")
				doOver = true
			}
			if dir > 359 {
				fmt.Println("Must be between 0 and 359!")
				doOver = true
			}
			if !doOver {
				switch what {
				case plane:
					b.plane = dir
				case wind:
					b.wind = dir
				}
				done = true
			}
		}
	}
	return quit
}

func (b *bearingT) calcWind() {
	tmp := b.wind - b.plane
	offset := tmp
	if tmp > 180 {
		offset = tmp - 360
	} else if tmp < -180 {
		offset = 360 + tmp
	}
	b.offset = offset
	if b.offset < 0 {
		b.direction = left
		b.offset = b.offset - (b.offset * 2)
	} else {
		b.direction = right
	}
}

func (b bearingT) showOffset() {
	var dStr string
	switch b.direction {
	case left:
		dStr = "Left"
	case right:
		dStr = "Right"
	}
	if b.offset == 180 {
		dStr = "Left or Right"
	}
	fmt.Printf("Windage offset is: %d° %s\n", b.offset, dStr)
}

func stripNewline(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}
