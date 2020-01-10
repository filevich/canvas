package canvas

import (
	"strings"
)

// Raw .
func Raw(obj string) string {
	return strings.Trim(obj, "\n ")
}

// Replace .
func Replace(this, that, here string) string {
	return strings.Replace(here, this, that, 1)
}
