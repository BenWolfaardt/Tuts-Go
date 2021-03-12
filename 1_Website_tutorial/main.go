package main

import (
	"golang.org/x/tour/pic"
	//"golang.org/1_Website_tutorial/pic"
)

// image is dy (256) long and dx (256) "high"
func Pic(dx, dy int) [][]uint8 {
	
	// Create the initial slice. It must have dy elements of type []uint8
	pic := make([][]uint8, dy)
	
	// Loop the slices
	for y := range pic {
		// Add dx slices to each element of the initial slice
		for x := 0; x < dx; x++ {
			pic[y] = append(pic[y], uint8(x * y))
		}
	}
	
	return pic 
}

func main() {
	pic.Show(Pic)
}