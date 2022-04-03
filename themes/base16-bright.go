package themes

import "image/color"

func init() {
	Base16["Bright"] = Theme{
		Name: "Bright",
		Author: "Chris Kempson (http://chriskempson.com)",
		Color00: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color01: color.RGBA{R: 48, G: 48, B: 48, A: 255},
		Color02: color.RGBA{R: 80, G: 80, B: 80, A: 255},
		Color03: color.RGBA{R: 176, G: 176, B: 176, A: 255},
		Color04: color.RGBA{R: 208, G: 208, B: 208, A: 255},
		Color05: color.RGBA{R: 224, G: 224, B: 224, A: 255},
		Color06: color.RGBA{R: 245, G: 245, B: 245, A: 255},
		Color07: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Color08: color.RGBA{R: 251, G: 1, B: 32, A: 255},
		Color09: color.RGBA{R: 252, G: 109, B: 36, A: 255},
		Color0A: color.RGBA{R: 253, G: 163, B: 49, A: 255},
		Color0B: color.RGBA{R: 161, G: 198, B: 89, A: 255},
		Color0C: color.RGBA{R: 118, G: 199, B: 183, A: 255},
		Color0D: color.RGBA{R: 111, G: 179, B: 210, A: 255},
		Color0E: color.RGBA{R: 211, G: 129, B: 195, A: 255},
		Color0F: color.RGBA{R: 190, G: 100, B: 60, A: 255},
	}
}
