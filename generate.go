package piscine

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[38;5;93m"
	Cyan   = "\033[36m"
	Pink   = "\033[38;5;201m" // Using 256 colors
	Orange = "\033[38;5;214m" // Using 256 colors
	Grey   = "\033[38;5;245m" // Using 256 colors
	Brown  = "\033[38;5;94m"  // Using 256 colors
	Beige  = "\033[38;5;230m" // Using 256 colors
)

var colorMap = map[string]string{
	"red":    Red,
	"green":  Green,
	"yellow": Yellow,
	"blue":   Blue,
	"purple": Purple,
	"cyan":   Cyan,
	"pink":   Pink,
	"orange": Orange,
	"grey":   Grey,
	"brown":  Brown,
	"beige":  Beige,
}

func Exist(c *string) bool {
	if _, exists := colorMap[strings.ToLower(*c)]; !exists {
		return false // Default to reset if color is not valid
	}
	return true
}

func PrintOutput(t map[rune][]string, str string, substring string, color string) {
	colorCode := colorMap[strings.ToLower(color)]
	in := strings.ReplaceAll(str, "\\n", "\n")

	words := strings.Split(in, "\n")
	for _, word := range words {
		Recursion(word, t, substring, colorCode)
	}
}

func Recursion(word string, t2 map[rune][]string, substring string, colorCode string) {
	if word == "" {
		fmt.Println()
		return
	}

	// Prepare colored lines
	coloredLines := make([]string, 8)

	for j := 0; j < 8; j++ {
		coloredLines[j] = ""
	}

	if substring == "" {
		// Color the whole word if no substring is specified
		for _, ch := range word {
			if ch == ' ' {
				for j := 0; j < 8; j++ {
					coloredLines[j] += "        " // Add space for ASCII art representation
				}
				continue
			}
			if lines, ok := t2[ch]; ok {
				for j := 0; j < 8; j++ {
					coloredLines[j] += colorCode + lines[j] + Reset
				}
			} else {
				for j := 0; j < 8; j++ {
					coloredLines[j] += "        " // Placeholder for missing characters
				}
			}
		}
	} else {
		// Iterate through the word to find the substring
		i := 0
		for i < len(word) {
			ch := rune(word[i])

			// Handle spaces
			if ch == ' ' {
				for j := 0; j < 8; j++ {
					coloredLines[j] += "        " // Add space for ASCII art representation
				}
				i++ // Move to the next character
				continue
			}

			// Check for substring match
			if strings.HasPrefix(word[i:], substring) {
				// Color the corresponding ASCII art for the substring
				for _, ch := range substring {
					if lines, ok := t2[ch]; ok {
						for j := 0; j < 8; j++ {
							coloredLines[j] += colorCode + lines[j] + Reset
						}
					} else {
						for j := 0; j < 8; j++ {
							coloredLines[j] += "        " // Placeholder for missing characters
						}
					}
				}
				i += len(substring) // Move the index forward by the length of the substring
			} else {
				// Otherwise, color individual characters
				if lines, ok := t2[ch]; ok {
					for j := 0; j < 8; j++ {
						coloredLines[j] += lines[j]
					}
				} else {
					for j := 0; j < 8; j++ {
						coloredLines[j] += "        " // Placeholder for missing characters
					}
				}
				i++ // Move to the next character
			}
		}
	}

	// Print all lines for the current word
	for _, line := range coloredLines {
		fmt.Println(line)
	}
}
