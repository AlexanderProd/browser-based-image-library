package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sum := hash.Sum(nil)

	return hex.EncodeToString(sum[:]), nil
}

func hashString(path string) string {
	input := strings.NewReader(path)

	hash := sha1.New()
	if _, err := io.Copy(hash, input); err != nil {
			log.Fatal(err)
	}

	sum := hash.Sum(nil)

	return hex.EncodeToString(sum[:])
}