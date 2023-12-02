package shared

import (
	"io/fs"
	"strings"
)

func ReadLinesFromFile(fileSystem fs.FS, filePath string) ([]string, error) {
	// Read the entire file using fs.ReadFile
	fileContent, err := fs.ReadFile(fileSystem, filePath)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and split into lines
	lines := splitLines(string(fileContent))

	return lines, nil
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}
