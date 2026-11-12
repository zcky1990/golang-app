// Package helpers mirrors Rails' app/helpers directory and holds small
// view/presentation helpers shared across controllers and templates.
package helpers

import (
	"fmt"
	"time"
)

// CurrentYear returns the current year, useful in layout footers.
func CurrentYear() int {
	return time.Now().Year()
}

// Presence returns fallback when value is empty, otherwise value.
func Presence(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

// Titleize builds a page title like "Home | MyApp".
func Titleize(page, app string) string {
	if page == "" {
		return app
	}
	return fmt.Sprintf("%s | %s", page, app)
}
