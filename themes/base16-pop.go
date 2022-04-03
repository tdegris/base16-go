package themes

import "image/color"

func init() {
	Base16["Pop"] = Theme{
		Name: "Pop",
		Author: "Chris Kempson (http://chriskempson.com)",
		Color00: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color01: color.RGBA{R: 32, G: 32, B: 32, A: 255},
		Color02: color.RGBA{R: 48, G: 48, B: 48, A: 255},
		Color03: color.RGBA{R: 80, G: 80, B: 80, A: 255},
		Color04: color.RGBA{R: 176, G: 176, B: 176, A: 255},
		Color05: color.RGBA{R: 208, G: 208, B: 208, A: 255},
		Color06: color.RGBA{R: 224, G: 224, B: 224, A: 255},
		Color07: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Color08: color.RGBA{R: 235, G: 0, B: 138, A: 255},
		Color09: color.RGBA{R: 242, G: 147, B: 51, A: 255},
		Color0A: color.RGBA{R: 248, G: 202, B: 18, A: 255},
		Color0B: color.RGBA{R: 55, G: 179, B: 73, A: 255},
		Color0C: color.RGBA{R: 0, G: 170, B: 187, A: 255},
		Color0D: color.RGBA{R: 14, G: 90, B: 148, A: 255},
		Color0E: color.RGBA{R: 179, G: 30, B: 141, A: 255},
		Color0F: color.RGBA{R: 122, G: 45, B: 0, A: 255},
	}
}
