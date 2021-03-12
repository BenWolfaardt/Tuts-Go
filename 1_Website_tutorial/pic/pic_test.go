  // Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*package pic_test

import "golang.org/x/tour/pic"

func ExampleShow() {
	f := func(dx, dy int) [][]uint8 {
		ss := make([][]uint8, dy)
		for y := 0; y < dy; y++ {
			s := make([]uint8, dx)
			for x := 0; x < dx; x++ {
				s[x] = uint8((x + y) / 2)
			}
			ss[y] = s
		}
		return ss
	}

	return 0

	pic.Show(f)

	// Output:
	// IMAGE:
	}*/