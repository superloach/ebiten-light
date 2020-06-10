package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	width  = 100
	height = 100
	scale  = 4
	glow   = width / 5
	edge   = glow * 3
)

var id = 0

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	light = color.RGBA{0, 255, 255, 63}
)

type Game struct {
	Shapes []*Polygon
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cx, cy := ebiten.CursorPosition()

	for x, shape := range g.Shapes {
		if shape.Contains(Point{cx, cy}) {
			screen.Set(x, 0, green)
		}

		for i, _ := range shape.In {
			d := (glow - math.Sqrt(math.Pow(float64(i[1]-cy), 2)+math.Pow(float64(i[0]-cx), 2))) / glow * 255
			if d < 0 {
				d = 0
			}
			ud := uint8(d)
			col := color.RGBA{0, ud, ud, 255}
			screen.Set(i[0], i[1], col)
		}

		for e, _ := range shape.Edge {
			d := (edge - math.Sqrt(math.Pow(float64(e[1]-cy), 2)+math.Pow(float64(e[0]-cx), 2))) / edge * 255
			if d < 0 {
				d = 0
			}
			ud := uint8(d)
			col := color.RGBA{0, ud, ud, 255}
			screen.Set(e[0], e[1], col)
		}
	}

	screen.Set(cx, cy, red)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %0.2f\n(%d, %d)",
		ebiten.CurrentFPS(),
		cx, cy,
	))
}

func (g *Game) Layout(ow, oh int) (int, int) {
	return width, height
}

func main() {
	g := &Game{}

	g.Shapes = append(g.Shapes,
		NewPolygon(red, []Point{
			{width / 4, height / 4},
			{3 * width / 4, height / 4},
			{2 * width / 3, 3 * height / 4},
			{width / 3, 3 * height / 4},
		}),
		NewPolygon(red, []Point{
			{2 * width / 5, 2 * height / 5},
			{3 * width / 5, 2 * height / 5},
			{3 * width / 5, 3 * height / 5},
			{2 * width / 5, 3 * height / 5},
		}),
	)

	ebiten.SetWindowSize(width*scale, height*scale)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
