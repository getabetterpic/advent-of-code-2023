package shared

import (
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	// Use os.DirFS to create an fs.FS from the current directory
	fileSystem := os.DirFS(".")
	// Read the entire file using fs.ReadFile
	fileContent, err := fs.ReadFile(fileSystem, filePath)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and split into lines
	lines := splitLines(string(fileContent))

	return lines, nil
}

func StringToInt(s string) int {
	// Convert the string to an int
	// Use strconv.Atoi to convert the string to an int
	integer, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return integer
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}
