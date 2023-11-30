package common

import (
	"strings"
)

func ConvertKebabCase(input string) string {
	// Replace spaces with dash
	kebabCase := strings.Replace(input, " ", "-", -1)
	return strings.ToLower(kebabCase)
}
