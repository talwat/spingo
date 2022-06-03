package main

import (
	"fmt"
	"time"

	"github.com/talwat/spingo"
)

type Color struct {
	Black  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Violet string
	Cyan   string
	White  string
}

var TextColor = Color{
	Black:  "\x1B[30m",
	Red:    "\x1B[31m",
	Green:  "\x1B[32m",
	Yellow: "\x1B[33m",
	Blue:   "\x1B[34m",
	Violet: "\x1B[35m",
	Cyan:   "\x1B[36m",
	White:  "\x1B[37m",
}

var TextReset = "\x1B[0m"

func main() {
	spinner := spingo.Spinner{
		Prefix: "Default configuration ",
	}

	displaySpin := func() {
		for i := 0; i <= 40; i++ {
			spinner.DisplaySpinner()
			time.Sleep(time.Millisecond * 50)
		}
		fmt.Print("\n")
	}

	displaySpin()

	spinner = spingo.Spinner{
		Prefix: "Fancy dots ",
		SpinnerArt: []string{
			"⠋",
			"⠙",
			"⠹",
			"⠸",
			"⠼",
			"⠴",
			"⠦",
			"⠧",
			"⠇",
			"⠏",
		},
	}

	displaySpin()

	spinner = spingo.Spinner{
		Prefix: "Traditional arc ",
		SpinnerArt: []string{
			"◜",
			"◠",
			"◝",
			"◞",
			"◡",
			"◟",
		},
	}

	displaySpin()

	spinner = spingo.Spinner{
		Prefix: "Multiple characters work too! ",
		SpinnerArt: []string{
			"( ●    )",
			"(  ●   )",
			"(   ●  )",
			"(    ● )",
			"(     ●)",
			"(    ● )",
			"(   ●  )",
			"(  ●   )",
			"( ●    )",
			"(●     )",
		},
	}

	displaySpin()

	spinner = spingo.Spinner{
		Prefix: "Even Emojis work! ",
		SpinnerArt: []string{
			"🌑 ",
			"🌒 ",
			"🌓 ",
			"🌔 ",
			"🌕 ",
			"🌖 ",
			"🌗 ",
			"🌘 ",
			"🌑 ",
		},
	}

	displaySpin()

	spinner = spingo.Spinner{
		Prefix: "Rainbow: ",
		SpinnerArt: []string{
			TextColor.Red + "0" + TextColor.Yellow + "0" + TextColor.Green + "0" + TextColor.Cyan + "0" + TextColor.Blue + "0" + TextColor.Violet + "0" + TextReset,
			TextColor.Violet + "0" + TextColor.Red + "0" + TextColor.Yellow + "0" + TextColor.Green + "0" + TextColor.Cyan + "0" + TextColor.Blue + "0" + TextReset,
			TextColor.Blue + "0" + TextColor.Violet + "0" + TextColor.Red + "0" + TextColor.Yellow + "0" + TextColor.Green + "0" + TextColor.Cyan + "0" + TextReset,
			TextColor.Cyan + "0" + TextColor.Blue + "0" + TextColor.Violet + "0" + TextColor.Red + "0" + TextColor.Yellow + "0" + TextColor.Green + "0" + TextReset,
			TextColor.Green + "0" + TextColor.Cyan + "0" + TextColor.Blue + "0" + TextColor.Violet + "0" + TextColor.Red + "0" + TextColor.Yellow + "0" + TextReset,
			TextColor.Yellow + "0" + TextColor.Green + "0" + TextColor.Cyan + "0" + TextColor.Blue + "0" + TextColor.Violet + "0" + TextColor.Red + "0" + TextReset,
		},
	}

	displaySpin()
}
