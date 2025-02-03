package device

import (
	"fmt"
	"strings"
)

func labelMapping() map[string]string {
	result := map[string]string{
		// X1 carbon
		"x1c":       "3DPrinter-X1-Carbon",
		"x1 carbon": "3DPrinter-X1-Carbon",

		// X1
		"x1": "3DPrinter-X1",

		// P1P
		"p1p": "C11",

		// P1S
		"p1s": "C12",

		// X1E
		"x1e": "C13",

		// A1 mini
		"a1 mini": "N1",
		"a1m":     "N1",

		// A1
		"a1": "N2S",
	}

	return result
}

func FromLabel(label string) (string, error) {
	l := strings.ToLower(label)
	mapping := labelMapping()

	if value, ok := mapping[l]; ok {
		return value, nil
	}

	return "", fmt.Errorf(`Unknown label "%s".`, label)
}
