package piscine

import "strings"

func GetFontFile(fontType string) string {
	switch strings.ToLower(fontType) {
	case "standard":
		return "standard.txt"
	case "shadow":
		return "shadow.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "standard2":
		return "standard2.txt"
	case "cross":
		return "cross.txt"
	default:
		return "standard.txt"
	}
}

func FontType(s string) bool {
	return s == "cross" || s == "standard" || s == "shadow" || s == "thinkertoy" || s == "standard2"
}
