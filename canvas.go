package canvas

// Canvas .
type Canvas [][]rune

const (
	width  = 54
	height = 13
)

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
