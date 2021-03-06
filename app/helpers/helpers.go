// Package helpers implements simple functions used across the application.
package helpers

import "os"

// FileExists returns true if the specified path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
