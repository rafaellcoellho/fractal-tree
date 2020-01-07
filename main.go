package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func createFractalTreeByLenght(basePoint pixel.Vec, len float64, angle float64, imd *imdraw.IMDraw) {
	imd.Push(basePoint)
	nextPoint := basePoint.Add(
		pixel.V(
			len*math.Cos(angle),
			len*math.Sin(angle),
		),
	)
	if len > 4 {
		createFractalTreeByLenght(nextPoint, len*0.67, angle-(math.Pi/6), imd)
		createFractalTreeByLenght(nextPoint, len*0.67, angle+(math.Pi/6), imd)
	}
	imd.Line(1)
}

func createFractalTreeByDepth(basePoint pixel.Vec, depth int, angle float64, imd *imdraw.IMDraw) {
	if depth > 0 {
		nextPoint := basePoint.Add(
			pixel.V(
				math.Cos(angle)*float64(depth)*10,
				math.Sin(angle)*float64(depth)*10,
			),
		)
		imd.Push(basePoint, nextPoint)
		imd.Line(1)
		createFractalTreeByDepth(nextPoint, depth-1, angle-(math.Pi/6), imd)
		createFractalTreeByDepth(nextPoint, depth-1, angle+(math.Pi/6), imd)
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

	createFractalTreeByLenght(pixel.V(200, 0), 100, math.Pi/2, fractalTree)
	// createFractalTreeByDepth(pixel.V(200, 0), 7, math.Pi/2, fractalTree)

	for !win.Closed() {
		win.Clear(colornames.Black)
		fractalTree.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
