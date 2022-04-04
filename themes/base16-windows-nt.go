package themes

import "image/color"

func init() {
	Base16["Windows NT"] = Theme{
		Name:    "Windows NT",
		Author:  "Fergus Collins (https://github.com/C-Fergus)",
		Color00: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color01: color.RGBA{R: 42, G: 42, B: 42, A: 255},
		Color02: color.RGBA{R: 85, G: 85, B: 85, A: 255},
		Color03: color.RGBA{R: 128, G: 128, B: 128, A: 255},
		Color04: color.RGBA{R: 161, G: 161, B: 161, A: 255},
		Color05: color.RGBA{R: 192, G: 192, B: 192, A: 255},
		Color06: color.RGBA{R: 224, G: 224, B: 224, A: 255},
		Color07: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Color08: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		Color09: color.RGBA{R: 128, G: 128, B: 0, A: 255},
		Color0A: color.RGBA{R: 255, G: 255, B: 0, A: 255},
		Color0B: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		Color0C: color.RGBA{R: 0, G: 255, B: 255, A: 255},
		Color0D: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		Color0E: color.RGBA{R: 255, G: 0, B: 255, A: 255},
		Color0F: color.RGBA{R: 0, G: 128, B: 0, A: 255},
	}
}
