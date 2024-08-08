package readbanner

import (
	"ascii-art-web-stylize/checkbanner"
	"errors"
	"log"
	"os"
	"strings"
)

var ErrBannerTampered = errors.New("hash value of the selected banner does not correspond with the expected value. corrupt file")

// Function ReadBanner() is responsible for reading the contents of a seleted banner file while ensuring the banner is correct.
func ReadBanner(bannerFile string) ([]string, error) {
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		log.Printf("Error reading banner file: %v", err)
		return nil, err
	}

	bannerFileHash := checkbanner.CheckBanner(bannerFileData)
	var expectedHash string

	switch bannerFile {
	case "banners/standard.txt":
		expectedHash = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	case "banners/shadow.txt":
		expectedHash = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	case "banners/thinkertoy.txt":
		expectedHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	default:
		return nil, nil
	}

	if bannerFileHash != expectedHash {
		return nil, ErrBannerTampered
	}

	splitBannerFileData := strings.ReplaceAll(string(bannerFileData), "\r\n", "\n")
	return strings.Split(splitBannerFileData, "\n"), nil
}
