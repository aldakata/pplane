package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fileName := os.Args[1]
	fmt.Println("CONGRATS!! AT LEAST IT RUNS")

	// file, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }

	// content := string(file)
	// lines := strings.Split(content, "\n")
	// f := *ParseField(lines[0:2], false)
	figRectangle := image.Rect(0, 0, 200, 200) // ParseRectangle(lines[2:4])
	// resolution := ParseResolution(lines[4])
	m := image.NewRGBA(figRectangle)
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	out, _ := os.Create("test.png")
	l := Line{
		P0: Point{0.0, 0.0},
		P1: Point{200.0, 200.0},
	}
	l1 := Line{
		P0: Point{0.0, 200.0},
		P1: Point{200.0, 0.0},
	}
	l2 := Line{
		P0: Point{100.0, 0.0},
		P1: Point{100.0, 200.0},
	}
	l3 := Line{
		P0: Point{0.0, 100.0},
		P1: Point{200.0, 100.0},
	}
	lines := []Line{l, l1, l2, l3}
	DrawLines(m, lines)
	png.Encode(out, m)
	out.Close()
}

// func PrintTests() {
// 	fmt.Println("TESTING PARSE:", ("24y^23x^32 +xy +2x +2"), parsePolynomial("24y^23x^32 + xy+2x+2"))
// 	fmt.Println("TESTING PARSE:", ("24y^23x^32 -xy -2x +3xy^3 -1"), parsePolynomial("24y^23x^32 - xy-2x+3xy^3-1"))
// 	fmt.Println("TESTING PARSE:", ("5 -12 -12 -12"), parsePolynomial("5-12-12-12"))
// 	fmt.Println("TESTING PARSE:", ("5 +-12 -12 -12"), parsePolynomial("5+-12-12-12"))
// }

func PrintDummy() {
	dummyPolynomial := new(Polynomial) //{XPow: make([]int, 2), YPow: make([]int, 2), Coeffs: make([]float64, 2)}
	dummyPolynomial.XPow = make([]int, 2)
	dummyPolynomial.XPow[0] = 2
	dummyPolynomial.XPow[1] = 0
	dummyPolynomial.YPow = make([]int, 2)
	dummyPolynomial.YPow[0] = 1
	dummyPolynomial.YPow[1] = 0
	dummyPolynomial.Coeffs = make([]float64, 2)
	dummyPolynomial.Coeffs[0] = 1.24
	dummyPolynomial.Coeffs[1] = 3.14

	fmt.Println("dummy pol", dummyPolynomial)
}

func ParseRectangle(lines []string) image.Rectangle {
	minBound := strings.Split(lines[0], ",")
	maxBound := strings.Split(lines[1], ",")
	x0, _ := strconv.Atoi(minBound[0])
	y0, _ := strconv.Atoi(minBound[1])
	x1, _ := strconv.Atoi(maxBound[0])
	y1, _ := strconv.Atoi(maxBound[1])
	return image.Rect(
		x0,
		y0,
		x1,
		y1,
	)
}

func ParseResolution(resolution string) int {
	payload, _ := strconv.Atoi(resolution)
	return payload
}

//
// Objectiu: Representar el camp vectorial d'una equacio diferencial
// H(f,g):R2->R2; f:R2->R, g:R2->R
//
// Requirements:
// 		- Representar polinomis de dues variables.
//
// Evaluar en un vector de punts equidistants
// Generar PNG amb les linies
