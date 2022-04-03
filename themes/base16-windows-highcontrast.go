package themes

import "image/color"

func init() {
	Base16["Windows High Contrast"] = Theme{
		Name: "Windows High Contrast",
		Author: "Fergus Collins (https://github.com/C-Fergus)",
		Color00: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color01: color.RGBA{R: 28, G: 28, B: 28, A: 255},
		Color02: color.RGBA{R: 56, G: 56, B: 56, A: 255},
		Color03: color.RGBA{R: 84, G: 84, B: 84, A: 255},
		Color04: color.RGBA{R: 162, G: 162, B: 162, A: 255},
		Color05: color.RGBA{R: 192, G: 192, B: 192, A: 255},
		Color06: color.RGBA{R: 222, G: 222, B: 222, A: 255},
		Color07: color.RGBA{R: 252, G: 252, B: 252, A: 255},
		Color08: color.RGBA{R: 252, G: 84, B: 84, A: 255},
		Color09: color.RGBA{R: 128, G: 128, B: 0, A: 255},
		Color0A: color.RGBA{R: 252, G: 252, B: 84, A: 255},
		Color0B: color.RGBA{R: 84, G: 252, B: 84, A: 255},
		Color0C: color.RGBA{R: 84, G: 252, B: 252, A: 255},
		Color0D: color.RGBA{R: 84, G: 84, B: 252, A: 255},
		Color0E: color.RGBA{R: 252, G: 84, B: 252, A: 255},
		Color0F: color.RGBA{R: 0, G: 128, B: 0, A: 255},
	}
}
