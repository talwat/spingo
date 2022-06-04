package spingo

import "testing"

func Test(t *testing.T) {
	spinner := Spinner{
		SpinnerArt: []string{"/", "|", "\\", "-"},
	}

	if spinner.ArtIndex != 0 {
		t.Errorf("Expected spinner.ArtIndex to be 0, got %d", spinner.ArtIndex)
	}

	spinner.DisplaySpinner()

	if spinner.ArtIndex != 1 {
		t.Errorf("Expected spinner.ArtIndex to be 1, got %d", spinner.ArtIndex)
	}

	for i := 0; i <= 100; i++ {
		spinner.DisplaySpinner()
	}

	spinner.End()
}
