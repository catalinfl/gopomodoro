package utils

import (
	"github.com/gookit/color"
)

func SetLoadingBarColor(s string) color.Color {
	switch s {
	case "red":
		return color.BgRed
	case "green":
		return color.BgGreen
	case "yellow":
		return color.BgYellow
	case "blue":
		return color.BgBlue
	case "magenta":
		return color.BgMagenta
	case "cyan":
		return color.BgCyan
	case "white":
		return color.BgWhite
	default:
		return color.BgCyan
	}
}
