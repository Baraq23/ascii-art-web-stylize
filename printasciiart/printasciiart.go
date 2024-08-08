package printasciiart

import (
	"fmt"
	"strings"
)

// Function PrintAsciiArt() is responsible for matching every character in the input string with its corresponding equivalent in a banner file.
func PrintAsciiArt(bannerFileSlice []string, inputString string, output *strings.Builder) error {
	switch inputString {
	case "":
		return nil
	}

	// Handle unprintable sequences
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}
	for _, sequence := range unprintableSequences {
		if strings.Contains(inputString, sequence) {
			return fmt.Errorf("input string contains an unprintable sequence")
		}
	}

	// Split inputString along its newlines
	splitString := strings.Split(inputString, "\r\n")

	// Handle characters absent in the range of 32 - 126 of the ascii manual
	for _, split := range splitString {
		for _, char := range split {
			if char < 32 || char > 126 {
				return fmt.Errorf("input string contains a character absent in the required range of the ascii manual")
			}
		}
	}

	// Compare characters with their equivalent banner representation
	for _, text := range splitString {
		if text == "" {
			output.WriteString("\n")
			continue
		}

		asciiHeight := 8

		for i := 0; i < asciiHeight; i++ {
			for _, char := range text {
				startingIndex := int(char-32)*9 + 1
				output.WriteString(bannerFileSlice[startingIndex+i])
			}
			output.WriteString("\n")
		}
	}
	return nil
}
