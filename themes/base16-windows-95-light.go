package themes

import "image/color"

func init() {
	Base16["Windows 95 Light"] = Theme{
		Name: "Windows 95 Light",
		Author: "Fergus Collins (https://github.com/C-Fergus)",
		Color00: color.RGBA{R: 252, G: 252, B: 252, A: 255},
		Color01: color.RGBA{R: 224, G: 224, B: 224, A: 255},
		Color02: color.RGBA{R: 196, G: 196, B: 196, A: 255},
		Color03: color.RGBA{R: 168, G: 168, B: 168, A: 255},
		Color04: color.RGBA{R: 126, G: 126, B: 126, A: 255},
		Color05: color.RGBA{R: 84, G: 84, B: 84, A: 255},
		Color06: color.RGBA{R: 42, G: 42, B: 42, A: 255},
		Color07: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color08: color.RGBA{R: 168, G: 0, B: 0, A: 255},
		Color09: color.RGBA{R: 252, G: 252, B: 84, A: 255},
		Color0A: color.RGBA{R: 168, G: 84, B: 0, A: 255},
		Color0B: color.RGBA{R: 0, G: 168, B: 0, A: 255},
		Color0C: color.RGBA{R: 0, G: 168, B: 168, A: 255},
		Color0D: color.RGBA{R: 0, G: 0, B: 168, A: 255},
		Color0E: color.RGBA{R: 168, G: 0, B: 168, A: 255},
		Color0F: color.RGBA{R: 84, G: 252, B: 84, A: 255},
	}
}
