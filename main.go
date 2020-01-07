package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func createFractaTree(basePoint pixel.Vec, len float64, angle float64, imd *imdraw.IMDraw) {
	imd.Push(basePoint)
	nextPoint := basePoint.Add(
		pixel.V(
			len*math.Cos(angle),
			len*math.Sin(angle),
		),
	)
	if len > 4 {
		createFractaTree(
			nextPoint,
			len*0.67,
			angle-(math.Pi/4),
			imd,
		)
		createFractaTree(
			nextPoint,
			len*0.67,
			angle+(math.Pi/4),
			imd,
		)
	}
}

func run() {
	const (
		WindowWidth  = 400
		WindowHeight = 400
	)

	cfg := pixelgl.WindowConfig{
		Title:  "fractal tree",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fractalTree := imdraw.New(nil)
	fractalTree.Color = colornames.White

	createFractaTree(
		pixel.V(200, 0),
		100,
		(math.Pi / 2),
		fractalTree,
	)

	fractalTree.Line(1)

	for !win.Closed() {
		win.Clear(colornames.Black)
		fractalTree.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
