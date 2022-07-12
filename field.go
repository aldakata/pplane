package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Polynomial struct {
	original   string
	XPow, YPow []int     // [1, 2, 1, 0], [0, 0, 2, 3]
	Coeffs     []float64 // [1, 5, 7, 0]
}

type Field struct {
	XField, YField *Polynomial
}

// Example: [x, 5x2, 7xy2, dy3]
// default grade is 2
func ParseField(lines []string, verbose bool) *Field {
	f := new(Field)
	f.XField = ParsePolynomial(strings.Split(lines[0], "=")[1], verbose)
	f.YField = ParsePolynomial(strings.Split(lines[1], "=")[1], verbose)
	return f
}

// From string to Polnomial struct.
// String: 1+2x^2y^3 -> polynomial{XPow : {0 2}, YPow : {0 3}, Coeffs : {1 2}}
func ParsePolynomial(rawPolynomial string, verbose bool) *Polynomial {
	pol := new(Polynomial)
	pol.XPow = make([]int, 0, 5)
	pol.YPow = make([]int, 0, 5)
	pol.Coeffs = make([]float64, 0, 5)
	pol.original = rawPolynomial
	trimmed := strings.Replace(rawPolynomial, " ", "", -1)
	for _, element := range strings.Split(trimmed, "+") {
		if verbose {
			fmt.Println("ELEMENT", element)
		}
		for i, negative_element := range strings.Split(element, "-") {
			if verbose {
				fmt.Println("\tNEGATIVE", i, negative_element)
			}
			sign := 1.
			if len(negative_element) == 0 {
				continue
			}
			if i > 0 {
				sign = -1.
			}
			XPow, YPow, coeff := ParseElement(negative_element, verbose)
			pol.XPow = append(pol.XPow, XPow)
			pol.YPow = append(pol.YPow, YPow)
			pol.Coeffs = append(pol.Coeffs, coeff*sign)
		}
	}
	fmt.Println()
	return pol
}

// 24y^23x^32 -> 24 32 23
func ParseElement(element string, verbose bool) (int, int, float64) {
	yIndex := strings.Index(element, "y")
	xIndex := strings.Index(element, "x")
	n := len(element)

	endX := n
	endY := n
	powX := 0
	powY := 0

	if xIndex < yIndex {
		endX = yIndex
	} else {
		endY = xIndex
	}

	if yIndex == -1 {
		yIndex = n
	} else {
		if yIndex+1 < n && element[yIndex+1] == '^' {
			powY, _ = strconv.Atoi(element[yIndex+2 : endY])
		} else {
			powY = 1
		}
	}
	if xIndex == -1 {
		xIndex = n
	} else {
		if xIndex+1 < n && element[xIndex+1] == '^' {
			powX, _ = strconv.Atoi(element[xIndex+2 : endX])
		} else {
			powX = 1
		}
	}
	coeff, _ := strconv.ParseFloat(element[0:Min(xIndex, yIndex)], 64)
	if verbose {
		fmt.Println("\t\tCOEFF", coeff)
	}
	return powX, powY, math.Max(coeff, 1)
}

// Eval polynomial pol with parameters parameters
func (f Field) Eval(x float64, y float64) [2]float64 {
	totalSum := [2]float64{f.XField.Eval(x, y), f.YField.Eval(x, y)}
	return totalSum
}

// Eval polynomial pol with parameters parameters
func (pol Polynomial) Eval(x float64, y float64) float64 {
	var totalSum float64 = 0
	for i, coeff := range pol.Coeffs {
		totalSum += coeff * math.Pow(x, float64(pol.XPow[i])) * math.Pow(y, float64(pol.YPow[i]))
	}

	return totalSum
}

func (f Field) String() string {
	return fmt.Sprintf("XField: %+v\nYField: %+v\n", *f.XField, *f.YField)
}
