package helpers

import (
	"os"
	"strings"
)

func ParseLines(path string) []string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}
