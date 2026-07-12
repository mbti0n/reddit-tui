package colors

import (
	"github.com/charmbracelet/lipgloss"
)

type Color int

const (
	Accent = iota
	Link
	Negative
	Text
	Subtext
)

type Palette struct {
	Accent   string `toml:"accent"`
	Link     string `toml:"link"`
	Negative string `toml:"negative"`
	Text     string `toml:"text"`
	Subtext  string `toml:"subtext"`
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
	default:
		return p.Text
	}
}

var ColorTheme = Palette{
	Accent:   "#ff713e",
	Link:     "#6688E4",
	Negative: "#6B5DFB",
	Text:     "#F5F5F5",
	Subtext:  "#D0D0D0",
}

func AdaptiveColor(color Color) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Light: ColorTheme.ToHex(color),
		Dark:  ColorTheme.ToHex(color),
	}
}
