package themes

import "image/color"

func init() {
	Base16["Windows NT Light"] = Theme{
		Name: "Windows NT Light",
		Author: "Fergus Collins (https://github.com/C-Fergus)",
		Color00: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Color01: color.RGBA{R: 234, G: 234, B: 234, A: 255},
		Color02: color.RGBA{R: 213, G: 213, B: 213, A: 255},
		Color03: color.RGBA{R: 192, G: 192, B: 192, A: 255},
		Color04: color.RGBA{R: 160, G: 160, B: 160, A: 255},
		Color05: color.RGBA{R: 128, G: 128, B: 128, A: 255},
		Color06: color.RGBA{R: 64, G: 64, B: 64, A: 255},
		Color07: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Color08: color.RGBA{R: 128, G: 0, B: 0, A: 255},
		Color09: color.RGBA{R: 255, G: 255, B: 0, A: 255},
		Color0A: color.RGBA{R: 128, G: 128, B: 0, A: 255},
		Color0B: color.RGBA{R: 0, G: 128, B: 0, A: 255},
		Color0C: color.RGBA{R: 0, G: 128, B: 128, A: 255},
		Color0D: color.RGBA{R: 0, G: 0, B: 128, A: 255},
		Color0E: color.RGBA{R: 128, G: 0, B: 128, A: 255},
		Color0F: color.RGBA{R: 0, G: 255, B: 0, A: 255},
	}
}
