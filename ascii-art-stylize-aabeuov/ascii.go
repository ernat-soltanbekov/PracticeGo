package main

import (
	"fmt"
	"os"
	"strings"
)

const CharHeight = 8

// ReadBannerFile reads a banner file and returns a map of character representations
func ReadBannerFile(bannerName string) (map[rune][]string, error) {
	filename := fmt.Sprintf("banners/%s.txt", bannerName)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read banner file %s: %v", filename, err)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	// Remove last empty line if present
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// Map to store character representations
	characters := make(map[rune][]string)

	// Skip leading empty line if present
	lineIndex := 0
	if len(lines) > 0 && lines[0] == "" {
		lineIndex = 1
	}

	// ASCII printable characters start at 32 (space) and end at 126 (~)
	charIndex := 32

	// Banner format: 8 lines per character, separated by blank lines
	for charIndex <= 126 && lineIndex < len(lines) {
		charLines := make([]string, 0, CharHeight)

		// Read 8 lines for this character
		for i := 0; i < CharHeight && lineIndex < len(lines); i++ {
			charLines = append(charLines, lines[lineIndex])
			lineIndex++
		}

		if len(charLines) == CharHeight {
			characters[rune(charIndex)] = charLines
		}

		// Skip blank line separator if present
		if lineIndex < len(lines) && lines[lineIndex] == "" {
			lineIndex++
		}

		charIndex++
	}

	return characters, nil
}

// GetCharacterWidth returns the width of characters in the banner
func GetCharacterWidth(characters map[rune][]string) int {
	for _, charLines := range characters {
		if len(charLines) > 0 && len(charLines[0]) > 0 {
			return len(charLines[0])
		}
	}
	return 0
}

// RenderASCII renders the input string as ASCII art
func RenderASCII(input, bannerName string) (string, error) {
	// Handle empty input
	if input == "" {
		return "", nil
	}

	characters, err := ReadBannerFile(bannerName)
	if err != nil {
		return "", err
	}

	charWidth := GetCharacterWidth(characters)
	if charWidth == 0 {
		return "", fmt.Errorf("invalid banner file: no characters found")
	}

	lines := strings.Split(input, "\n")
	var result strings.Builder

	for i, line := range lines {
		// Handle empty lines - just output a single newline
		if line == "" {
			// Only add blank line if not the first element
			if i > 0 {
				result.WriteRune('\n')
			}
			continue
		}

		// For each non-empty line in the input, build 8 lines of ASCII art
		for row := 0; row < CharHeight; row++ {
			for _, ch := range line {
				if charLines, exists := characters[ch]; exists {
					if row < len(charLines) {
						result.WriteString(charLines[row])
					} else {
						result.WriteString(strings.Repeat(" ", charWidth))
					}
				} else {
					// Character not found in banner, write spaces
					result.WriteString(strings.Repeat(" ", charWidth))
				}
			}
			result.WriteRune('\n')
		}
	}

	return result.String(), nil
}
