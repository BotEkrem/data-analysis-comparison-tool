package utils

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
	"strings"
)

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var colorCode = map[Color]string{
	Black:   "\u001b[30m,\u001b[40m",
	Red:     "\u001b[31m,\u001b[41m",
	Green:   "\u001b[32m,\u001b[42m",
	Yellow:  "\u001b[33m,\u001b[43m",
	Blue:    "\u001b[34m,\u001b[44m",
	Magenta: "\u001b[35m,\u001b[45m",
	Cyan:    "\u001b[36m,\u001b[46m",
	White:   "\u001b[37m,\u001b[47m",
}

func (c Color) Code() string {
	return colorCode[c]
}

func ColoredLog(background bool, text string, color Color) {
	var colorName string

	if background {
		colorName = strings.Split(colorCode[color], ",")[1]
	} else {
		colorName = strings.Split(colorCode[color], ",")[0]
	}

	fmt.Println(colorName, text, "\u001b[0m")
}

func ClearConsole() {
	screen.Clear()
	screen.MoveTopLeft()
}

func Selection(choices []string, title string, hoverColor Color, choiceColor Color, titleColor Color) string {
	ClearConsole()
	var currentIndex = 0

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		ClearConsole()

		ColoredLog(true, title, titleColor)

		for i := 0; i < len(choices); i++ {
			if i == currentIndex {
				ColoredLog(true, choices[i], hoverColor)
			} else if choices[i] == "Finish" {
				ColoredLog(true, choices[i], Green)
			} else {
				ColoredLog(false, choices[i], choiceColor)
			}
		}

		_, key, err := keyboard.GetKey()

		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowDown {
			if (currentIndex + 1) != len(choices) {
				currentIndex++
			}
		}

		if key == keyboard.KeyArrowUp {
			if (currentIndex - 1) >= 0 {
				currentIndex--
			}
		}

		if key == keyboard.KeyEnter {
			break
		}
	}

	ClearConsole()
	println(strings.Split(colorCode[Green], ",")[1], "You selected", choices[currentIndex], "(Press any key to continue)", "\u001b[0m")

	return choices[currentIndex]
}
