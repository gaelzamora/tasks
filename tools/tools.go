package tools

import "strings"

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")

	return desc
}