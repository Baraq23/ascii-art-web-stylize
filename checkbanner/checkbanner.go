package checkbanner

import (
	"crypto/sha256"
	"encoding/hex"
)

// Function CheckBanner() uses a sha256 encryption algorithm to determine the integrity of a file.
func CheckBanner(bannerFileData []byte) string {
	hasher := sha256.New()
	hasher.Write(bannerFileData)
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
