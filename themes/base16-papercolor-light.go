package themes

import "image/color"

func init() {
	Base16["PaperColor Light"] = Theme{
		Name:    "PaperColor Light",
		Author:  "Jon Leopard (http://github.com/jonleopard) based on PaperColor Theme (https://github.com/NLKNguyen/papercolor-theme)",
		Color00: color.RGBA{R: 238, G: 238, B: 238, A: 255},
		Color01: color.RGBA{R: 175, G: 0, B: 0, A: 255},
		Color02: color.RGBA{R: 0, G: 135, B: 0, A: 255},
		Color03: color.RGBA{R: 95, G: 135, B: 0, A: 255},
		Color04: color.RGBA{R: 0, G: 135, B: 175, A: 255},
		Color05: color.RGBA{R: 68, G: 68, B: 68, A: 255},
		Color06: color.RGBA{R: 0, G: 95, B: 135, A: 255},
		Color07: color.RGBA{R: 135, G: 135, B: 135, A: 255},
		Color08: color.RGBA{R: 188, G: 188, B: 188, A: 255},
		Color09: color.RGBA{R: 215, G: 0, B: 0, A: 255},
		Color0A: color.RGBA{R: 215, G: 0, B: 135, A: 255},
		Color0B: color.RGBA{R: 135, G: 0, B: 175, A: 255},
		Color0C: color.RGBA{R: 215, G: 95, B: 0, A: 255},
		Color0D: color.RGBA{R: 215, G: 95, B: 0, A: 255},
		Color0E: color.RGBA{R: 0, G: 95, B: 175, A: 255},
		Color0F: color.RGBA{R: 0, G: 95, B: 135, A: 255},
	}
}
