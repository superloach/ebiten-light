package main

import "image/color"

type Point [2]int

type Polygon struct {
	Color color.Color

	Vtc []Point

	Edge map[Point]struct{}
	In   map[Point]struct{}
}

func NewPolygon(col color.Color, vtc []Point) *Polygon {
	if len(vtc) < 3 {
		panic("poly needs >=3 Vtc")
	}

	p := &Polygon{}

	p.Color = col
	p.Vtc = vtc

	p.MakeEdge()
	p.MakeIn()

	return p
}

func (p *Polygon) MakeEdge() {
	p.Edge = make(map[Point]struct{})

	vtc := make([]Point, len(p.Vtc)+1)
	copy(vtc, p.Vtc)
	vtc[len(vtc)-1] = vtc[0]

	for i := 0; i < len(vtc)-1; i++ {
		for _, pt := range line(vtc[i], vtc[i+1]) {
			p.Edge[pt] = struct{}{}
		}
	}
}

func (p *Polygon) MakeIn() {
	p.In = make(map[Point]struct{})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := Point{x, y}
			if p.Contains(c) {
				p.In[c] = struct{}{}
			}
		}
	}
}

func (p *Polygon) Contains(c Point) bool {
	if c[0] < 0 || c[1] < 0 || c[0] > width || c[1] > height {
		return false
	}

	b := false
	i, j := 0, len(p.Vtc)-1
	for i < len(p.Vtc) {
		if (((p.Vtc[i][1] <= c[1]) && (c[1] < p.Vtc[j][1])) || ((p.Vtc[j][1] <= c[1]) && (c[1] < p.Vtc[i][1]))) && (c[0] < (p.Vtc[j][0]-p.Vtc[i][0])*(c[1]-p.Vtc[i][1])/(p.Vtc[j][1]-p.Vtc[i][1])+p.Vtc[i][0]) {
			b = !b
		}

		j = i
		i++
	}

	return b
}
