package piscine

import (
	"bufio"
	"os"
)

func Load(file string) map[rune][]string {
	font, error := os.Open(file)
	if error != nil {
		return nil
	}
	defer font.Close()

	scanner := bufio.NewScanner(font)
	table := make(map[rune][]string)
	lines := make([]string, 8)
	i := ' '
	for j := 0; scanner.Scan(); j++ {
		line := scanner.Text()
		if j == 0 {
			lines = append(lines, line)
		}
		if j != 0 && j%9 != 0 {
			lines = append(lines, line)
		}
		if j != 0 && j%9 == 0 {
			table[i] = lines
			lines = nil
			i++
		}
	}
	if len(lines) > 0 {
		table[i] = lines
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return table
}
