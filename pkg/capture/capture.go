package capture

import (
	"errors"
	"image"
)

// ErrCanceled is returned when the user aborts screen selection (e.g., pressed ESC).
var ErrCanceled = errors.New("capture canceled by user")

// Snip captures an interactive user-selected region from the screen.
func Snip() (image.Image, error) {
	return snipScreen()
}
