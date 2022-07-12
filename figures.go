package main

import (
	"image"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	P0, P1 Point
}

type Rectangle struct {
	Min, Max Point
}

func DrawLine(img *image.RGBA, line Line) {
	p0 := image.Point{
		int(math.Floor(line.P0.X)),
		int(math.Floor(line.P0.Y)),
	}
	p1 := image.Point{
		int(math.Floor(line.P1.X)),
		int(math.Floor(line.P1.Y)),
	}
	iHeight := line.P0.Y
	slope := line.Slope()

	firstPoint := p0
	lastPoint := p1
	if p0.X > p1.X {
		firstPoint = p1
		lastPoint = p0
	}
	for pixel := firstPoint; ; {
		nextHeight := iHeight + float64(pixel.X+1)*slope
		img.Set(pixel.X, pixel.Y, color.Black)
		if pixel == lastPoint {
			// fmt.Println("BREAKING")
			break
		}
		if math.Inf(-1) < slope && slope < math.Inf(1) {
			pixel.X++
		}
		if slope == 0 {
		} else if nextHeight >= float64(pixel.Y+1) {
			// fmt.Println("Y++")
			pixel.Y++
		} else if nextHeight <= float64(pixel.Y) {
			// fmt.Println("X++ Y++")
			pixel.Y--
		}
	}
}

func DrawLines(img *image.RGBA, lines []Line) {
	for _, l := range lines {
		DrawLine(img, l)
	}
}

func (p Point) scale(rectangle Rectangle) Point {
	height := rectangle.Min.Y - rectangle.Max.Y
	width := rectangle.Min.X - rectangle.Max.X
	q := Point{p.X*width + rectangle.Min.X, p.Y*height + rectangle.Min.Y}
	return q
}

func (l Line) Slope() float64 {
	xDiff := l.P1.X - l.P0.X
	yDiff := l.P1.Y - l.P0.Y
	if xDiff == 0 {
		if yDiff > 0 {
			return math.Inf(1)
		} else if yDiff > 0 {
			return math.Inf(int(-1))
		} else {
			return 0
		}
	}
	return yDiff / xDiff
}

// func DrawLines(rectangle image.Rectangle, lines []Line) image.Image {
// 	image := new(image.Image)
// 	DrawLines()
// }
