package colors

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

const (
	selectGraphicRendition = "\x1B["

	FgBlack         = 30
	FgRed           = 31
	FgGreen         = 32
	FgYellow        = 33
	FgBlue          = 34
	FgMagenta       = 35
	FgCyan          = 36
	FgWhite         = 37
	FgReset         = 39
	FgBrightBlack   = 90
	FgBrightRed     = 91
	FgBrightGreen   = 92
	FgBrightYellow  = 93
	FgBrightBlue    = 94
	FgBrightMagenta = 95
	FgBrightCyan    = 96
	FgBrightWhite   = 97

	BgBlack         = 40
	BgRed           = 41
	BgGreen         = 42
	BgYellow        = 43
	BgBlue          = 44
	BgMagenta       = 45
	BgCyan          = 46
	BgWhite         = 47
	BgReset         = 49
	BgBrightBlack   = 100
	BgBrightRed     = 101
	BgBrightGreen   = 102
	BgBrightYellow  = 103
	BgBrightBlue    = 104
	BgBrightMagenta = 105
	BgBrightCyan    = 106
	BgBrightWhite   = 107

	ResetFGColorTag = selectGraphicRendition + "39m"
	ResetBGColorTag = selectGraphicRendition + "49m"
	ResetColorsTag  = selectGraphicRendition + "39;49m"
)

// RGB functions
func NewRandomFGColorRGB() string {
	color, _ := NewFGColorRGB(127+rand.IntN(129), 127+rand.IntN(129), 127+rand.IntN(129))
	return color
}

func NewRandomBGColorRGB() string {
	color, _ := NewBGColorRGB(127+rand.IntN(129), 127+rand.IntN(129), 127+rand.IntN(129))
	return color
}

func NewFGColorRGB(r, g, b int) (string, error) {
	if (r < 127 || r > 255) && (g < 127 || g > 255) && (b < 127 || b > 255) {
		return "", errors.New("invalid parameters: numbers must be between 127 and 255")
	}
	return fmt.Sprintf("%s38;2;%d;%d;%dm", selectGraphicRendition, r, g, b), nil
}

func NewBGColorRGB(r, g, b int) (string, error) {
	if (r < 127 || r > 255) && (g < 127 || g > 255) && (b < 127 || b > 255) {
		return "", errors.New("invalid parameters: numbers must be between 127 and 255")
	}
	return fmt.Sprintf("%s48;2;%d;%d;%dm", selectGraphicRendition, r, g, b), nil
}

func SprintfForegroundRGB(str string, r, g, b int) string {
	rgbColor, errRGB := NewFGColorRGB(r, g, b)
	if errRGB != nil {
		return str
	}

	return fmt.Sprintf("%s%s%s", rgbColor, str, ResetFGColorTag)
}

func SprintfBackgroundRGB(str string, r, g, b int) string {
	rgbColor, errRGB := NewBGColorRGB(r, g, b)
	if errRGB != nil {
		return str
	}

	return fmt.Sprintf("%s%s%s", rgbColor, str, ResetBGColorTag)
}

func SprintfRGB(str string, rFG, gFG, bFG, rBG, gBG, bBG int) string {
	rgbFGColor, errRGB := NewFGColorRGB(rFG, gFG, bFG)
	if errRGB != nil {
		return str
	}
	rgbBGColor, errRGB := NewBGColorRGB(rBG, gBG, bBG)
	if errRGB != nil {
		return str
	}

	return fmt.Sprintf("%s%s%s%s", rgbFGColor, rgbBGColor, str, ResetColorsTag)
}

// ANSI functions
func SprintfANSI(str string, fg, bg int) string {
	if ((fg < 30 || fg > 37) && fg != 39) && (fg < 90 || fg > 97) {
		return str
	}
	if ((bg < 40 || bg > 47) && bg != 49) && (bg < 100 || bg > 107) {
		return str
	}

	return fmt.Sprintf("%s%d;%dm%s%s", selectGraphicRendition, fg, bg, str, ResetColorsTag)
}

func NewFGColorANSI(color string) (string, error) {
	if color == "" {
		return "", errors.New("empty parameter")
	}

	switch strings.ToLower(color) {
	case "black":
		return selectGraphicRendition + strconv.Itoa(FgBlack) + "m", nil
	case "red":
		return selectGraphicRendition + strconv.Itoa(FgRed) + "m", nil
	case "green":
		return selectGraphicRendition + strconv.Itoa(FgGreen) + "m", nil
	case "yellow":
		return selectGraphicRendition + strconv.Itoa(FgYellow) + "m", nil
	case "blue":
		return selectGraphicRendition + strconv.Itoa(FgBlue) + "m", nil
	case "magenta":
		return selectGraphicRendition + strconv.Itoa(FgMagenta) + "m", nil
	case "cyan":
		return selectGraphicRendition + strconv.Itoa(FgCyan) + "m", nil
	case "white":
		return selectGraphicRendition + strconv.Itoa(FgWhite) + "m", nil
	case "bright black", "bright-black", "brightblack":
		return selectGraphicRendition + strconv.Itoa(FgBrightBlack) + "m", nil
	case "bright red", "bright-red", "brightred":
		return selectGraphicRendition + strconv.Itoa(FgBrightRed) + "m", nil
	case "bright green", "bright-green", "brightgreen":
		return selectGraphicRendition + strconv.Itoa(FgBrightGreen) + "m", nil
	case "bright yellow", "bright-yellow", "brightyellow":
		return selectGraphicRendition + strconv.Itoa(FgBrightYellow) + "m", nil
	case "bright blue", "bright-blue", "brightblue":
		return selectGraphicRendition + strconv.Itoa(FgBrightBlue) + "m", nil
	case "bright magenta", "bright-magenta", "brightmagenta":
		return selectGraphicRendition + strconv.Itoa(FgBrightMagenta) + "m", nil
	case "bright cyan", "bright-cyan", "brightcyan":
		return selectGraphicRendition + strconv.Itoa(FgBrightCyan) + "m", nil
	case "bright white", "bright-white", "brightwhite":
		return selectGraphicRendition + strconv.Itoa(FgBrightWhite) + "m", nil
	default:
		return "", errors.New("invalid parameter")
	}
}
