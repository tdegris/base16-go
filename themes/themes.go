package themes

import "image/color"

// Theme is a base16 theme.
type Theme struct {
	Name    string
	Author  string
	Color00 color.RGBA
	Color01 color.RGBA
	Color02 color.RGBA
	Color03 color.RGBA
	Color04 color.RGBA
	Color05 color.RGBA
	Color06 color.RGBA
	Color07 color.RGBA
	Color08 color.RGBA
	Color09 color.RGBA
	Color0A color.RGBA
	Color0B color.RGBA
	Color0C color.RGBA
	Color0D color.RGBA
	Color0E color.RGBA
	Color0F color.RGBA
}

// Base maps theme name to theme data.
var Base16 = make(map[string]Theme)
