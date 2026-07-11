package colors

import "github.com/charmbracelet/lipgloss"

type Color int

const (
	Accent = iota
	Link
	Negative
	Text
	Subtext
	White
)

type Palette struct {
	Accent   string
	Link     string
	Negative string
	Text     string
	Subtext  string
	White    string
}

func (p Palette) ToHex(color Color) string {
	switch color {
	case Accent:
		return p.Accent
	case Link:
		return p.Link
	case Negative:
		return p.Negative
	case Text:
		return p.Text
	case Subtext:
		return p.Subtext
	case White:
		return p.White
	default:
		return p.Text
	}
}

// Reddit Theme Dark
var Dark = Palette{
	Accent:   "#ff713e",
	Link:     "#6688E4",
	Negative: "#6B5DFB",
	Text:     "#F5F5F5",
	Subtext:  "#D0D0D0",
	White:    "#ffffff",
}

// Reddit Theme Light
var Light = Palette{
	Accent:   "#fb4300",
	Link:     "#6688E4",
	Negative: "#6B5DFB",
	Text:     "#4c4f69",
	Subtext:  "#5c5f77",
	White:    "#ffffff",
}

func AdaptiveColors(light, dark Color) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Light: Light.ToHex(light),
		Dark:  Dark.ToHex(dark),
	}
}

func AdaptiveColor(color Color) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Light: Light.ToHex(color),
		Dark:  Dark.ToHex(color),
	}
}
