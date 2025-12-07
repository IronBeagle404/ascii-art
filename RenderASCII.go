// package ascii-art provides ASCII rendering utilities
//
// Available fonts :
//   - standard
//   - shadow
//   - thinkertoy
//
// Use asciiart.AvailableFonts to get the list programmatically
package ascii_art

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed fonts/*
var fontFiles embed.FS

var AvailableFonts = map[string]string{
	"standard":   "fonts/standard.txt",
	"shadow":     "fonts/shadow.txt",
	"thinkertoy": "fonts/thinkertoy.txt",
}

// RenderASCIIToString renders text using the specified font.
//
// The font must be one of the built-in fonts (see package documentation). Passing an invalid font name returns an error.
func RenderASCIIToString(input string, font string) (string, error) {
	input = strings.ReplaceAll(input, "\\n", "\n")

	path, ok := AvailableFonts[font]
	if !ok {
		return "", fmt.Errorf("font %v does not exist", font)
	}

	fontBytes, err := fontFiles.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("font file %v is unreadable", path)
	}

	bannerData := strings.Split(string(fontBytes), "\n\n")

	var final []string

	for _, r := range input {
		if r == '\n' {
			final = append(final, "")
		} else if int(r) >= 32 && int(r) <= 126 {
			final = append(final, bannerData[int(r)-32])
		}
	}

	var builder strings.Builder
	for i := 0; i < len(final); {
		if final[i] == "" {
			builder.WriteString("\n")
			i++
			continue
		}

		j := i
		for j < len(final) && final[j] != "" {
			j++
		}

		for line := range 8 {
			var result []string
			for k := i; k < j; k++ {
				lineParts := strings.Split(final[k], "\n")
				if line < len(lineParts) {
					result = append(result, lineParts[line])
				}
			}
			builder.WriteString(strings.Join(result, ""))
			if line != 7 {
				builder.WriteString("\n")
			} else {
				break
			}
		}

		i = j
	}

	return builder.String(), nil
}
