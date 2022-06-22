package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"strings"
)

func hashFile(path string) string {
	input := strings.NewReader(path)

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
			log.Fatal(err)
	}

	sum := hash.Sum(nil)

	return hex.EncodeToString(sum[:])
}