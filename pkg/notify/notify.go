package notify

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

// Send displays a native system notification with title and message.
// If notification fails (e.g. headless environment), it returns an error without panicking.
func Send(title, message string) error {
	err := beeep.Notify(title, message, "")
	if err != nil {
		// Log or suppress without aborting execution
		return fmt.Errorf("notification error: %w", err)
	}
	return nil
}
