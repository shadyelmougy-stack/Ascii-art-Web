package main

import (
	"fmt"
	"os"
	"strings"
)

// validateInput rejects any character outside the printable ASCII range (32–126).
func validateInput(content string) error {
	for _, b := range content {
		if b < ' ' || b > '~' {
			return fmt.Errorf("please use only printable ASCII characters (space to ~)")
		}
	}
	return nil
}

// buildBannerMap parses the banner file into a map from byte → [8]lines.
func buildBannerMap(content string) (map[byte][8]string, error) {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	if len(lines) < 855 {
		return nil, fmt.Errorf("the banner file is invalid or incomplete")
	}
	charMap := map[byte][8]string{}
	for ch := ' '; ch <= '~'; ch++ {
		offset := int(ch-' ')*9 + 1
		rows := [8]string{}
		copy(rows[:], lines[offset:offset+8])
		charMap[byte(ch)] = rows
	}
	return charMap, nil
}

// slashNHandling splits the input on the literal two-character sequence \n.
func slashNHandling(content string) []string {
	return strings.Split(content, `\n`)
}

// readFile reads and returns the entire content of the named file.
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// renderArt builds the ASCII-art result as a single string, instead of printing it.
func renderArt(charMap map[byte][8]string, lines []string) string {
	var result strings.Builder

	li := []string{}
	if len(lines) == 0 {
		return ""
	}
	for i, line := range lines {
		if line != "" {
			li = lines
			break
		} else if i == len(lines)-1 && line == "" {
			li = lines[0 : len(lines)-1]
		}
	}

	var sb strings.Builder
	for _, line := range li {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for row := range 8 {
			sb.Reset()
			for _, ch := range line {
				sb.WriteString(charMap[byte(ch)][row])
			}
			result.WriteString(sb.String())
			result.WriteString("\n")
		}
	}
	return result.String()
}
