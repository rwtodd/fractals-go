// Package fractals-go/algo contains the algorithms
// for computing the pixel values of various fractal
// images.
package algo

import (
	"fmt"
)

// Fractal is the interface that decouples the driver program
// from the particular fractal being rendered.
type Fractal interface {
	fmt.Stringer

	// Intensity determines the pixel value of the given
	// coordinates.
	Intensity(x, y float64) uint8

	// ArgHelp provides a help string for use in the UI,
	// telling the user what the arguments to that fractal
	// represent.
	ArgHelp() string
}

// Norm computes the norm of a complex variable
func norm(x complex128) float64 {
	rl, im := real(x), imag(x)
	return rl*rl + im*im
}
