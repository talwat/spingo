// An extremely basic spinner library
package spingo

import (
	"fmt"
)

// Spinner Struct.
type Spinner struct {
	// Index of the current spinner art, it's fine to not set this when creating a new spinner.
	ArtIndex int

	// Whether the spinner has started spinning or not.
	AlreadyStarted bool

	// Prefix to display before the spinner.
	Prefix string

	// Suffix to display after the spinner.
	Suffix string

	// Spinner art to display,
	//   eg. ["/", "|", "\\", "-"]
	SpinnerArt []string
}

// Displays the spinners current state.
// Make sure to call End() when you are done with the spinner.
// It is not possible to print something in the middle of a spinner spinning or else there will be multiple spinners.
// For example, do not do this:
//   spinner.DisplaySpinner()
//   fmt.Println("Do not do this")
//   spinner.DisplaySpinner()
func (spinner *Spinner) DisplaySpinner() {
	// If there is no spinner art, set it to a default
	if len(spinner.SpinnerArt) == 0 {
		spinner.SpinnerArt = []string{"/", "|", "\\", "-"}
	}

	// If the spinner has not started, erase the line
	if !spinner.AlreadyStarted {
		fmt.Print("\x1b[2K")
		spinner.AlreadyStarted = true
	}

	fmt.Print("\r" + spinner.Prefix + spinner.SpinnerArt[spinner.ArtIndex] + spinner.Suffix)

	spinner.ArtIndex += 1

	// If the index is greater than the length of the spinner art, reset it
	if spinner.ArtIndex >= len(spinner.SpinnerArt) {
		spinner.ArtIndex = 0
	}
}

// Prints a new line to prevent issues with multiple spinners consecutively.
func (spinner Spinner) End() {
	fmt.Print("\n")
}
