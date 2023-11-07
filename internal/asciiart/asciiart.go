package asciiart

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"asciiartweb/internal/asciiartfs"
)

const fileLen = 855

// AsciiArt generates ASCII art based on the provided banner and font string.
func AsciiArt(banner string, fontStr string) (string, error) {
	// Read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(fontStr, "\r", "\n"), "\n")
	arr := make([]string, 0, fileLen)

	readFile, err := os.Open("../../internal/asciiart/fonts/" + banner + ".txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("file not found")
		}
		return "", errors.New("failed to open file: " + err.Error())
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		return "", errors.New("file is corrupted")
	}

	larg := len(argsArr)
	if larg >= 2 && argsArr[larg-1] == "" && argsArr[larg-2] != "" {
		argsArr = argsArr[:larg-1]
	}

	return asciiartfs.PrintBanners(argsArr, arr), nil
}
