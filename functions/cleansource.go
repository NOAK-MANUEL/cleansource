package functions

import (
	"cleanSource/globalvar"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func matchIgnore(name string, ignore []string) bool {
	for _, pattern := range ignore {
		switch {
		case strings.HasPrefix(pattern, "*") && strings.HasSuffix(pattern, "*") && len(pattern) > 1:
			// *foo* -> contains "foo"
			middle := pattern[1 : len(pattern)-1]
			if middle != "" && strings.Contains(name, middle) {
				return true
			}
		case strings.HasPrefix(pattern, "*"):
			// *.git -> ends with ".git"
			suffix := pattern[1:]

			if strings.HasPrefix(name, suffix) {
				return true
			}
			

		case strings.HasSuffix(pattern, "*"):
			// git* -> starts with "git"
			prefix := pattern[:len(pattern)-1]
			if strings.HasSuffix(name, prefix) {
				return true
			}
		default:
			// exact match, e.g. "gitignore"
			if name == pattern {
				return true
			}
		}
	}
	return false
}

func CleanSource(path string, ignore []string, list bool) int {
	entries, err := os.ReadDir(path)
	if err != nil {
		println("Couldn't read directory", err.Error())
		return -1
	}
	amountDeleted := 0
	for _, entry := range entries {
		if matchIgnore(entry.Name(), ignore) || slices.Contains(globalvar.Dependencies, filepath.Join(path, entry.Name())) {
			continue
		}
		full := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			amountDeleted += CleanSource(full, ignore, list)

			// Only remove the folder if it's now actually empty.
			remaining, err := os.ReadDir(full)
			if err != nil {
				continue // couldn't check, skip deleting it
			}
			if len(remaining) == 0 {
				if list {
					println(full)
				} else {
					if err := os.Remove(full); err == nil {
						amountDeleted++
					}
				}
			}
		} else {
			if list {
				println(full)
			} else {
				if err := os.Remove(full); err == nil {
					amountDeleted++
				}
			}
		}
	}
	return amountDeleted
}