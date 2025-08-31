package ascii_art

import (
	_ "embed"
	"strings"
)

//go:embed ressources.txt
var banner string

func RenderASCIIToString(input string) string {
	input = strings.ReplaceAll(input, "\\n", "\n")

	bannerData := strings.Split(banner, "\n\n")

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

		for line := 0; line < 8; line++ {
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

	return builder.String()
}
