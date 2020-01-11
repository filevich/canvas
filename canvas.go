package canvas

import (
	"strconv"
	"strings"
)

// Canvas .
type Canvas [][]rune

const (
	width  = 75
	height = 19
)

// Raw .
func Raw(obj string) string {
	return strings.Trim(obj, "\n ")
}

// Replace .
func Replace(this, that, here string) string {
	return strings.Replace(here, this, that, 1)
}

// Debug .
func (canvas Canvas) Debug(p Point) (rune, string) {
	var c rune = canvas[p.X][p.Y]
	return c, strconv.QuoteRune(c)
}

// Render .
func (canvas Canvas) Render() string {
	// el string va a tener el mismo tamano que
	// el canvas + los \n que no estan
	res := ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			res += string(canvas[x][y])
		}
		res += "\n"
	}
	return res
}

// DrawAt .
func (canvas Canvas) DrawAt(p Point, obj string) {
	canvas.Draw(p.X, p.Y, obj)
}

// Draw .
func (canvas Canvas) Draw(fromX, fromY int, obj string) {
	var (
		x = fromX
		y = fromY
	)
	for _, char := range obj {
		if char == '\n' {
			y++
			x = fromX
		} else {
			canvas[x][y] = char
			x++
		}
	}
}

// NewCanvas crea un canvas vacio
func NewCanvas() Canvas {
	var canvas Canvas
	canvas = make([][]rune, width)
	for i := range canvas {
		canvas[i] = make([]rune, height)
	}
	// cargo con ' '
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			canvas[x][y] = ' '
		}
	}
	return canvas
}
