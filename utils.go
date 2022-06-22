package main

import (
	"regexp"
	"strings"
)

func fileNameFromPath(path string) string {
	r := regexp.MustCompile(`(?mi)([^/]+$)`)
	return r.FindString(path)
}

func fileTypeFromPath(path string) string {
	r, _ := regexp.Compile(`(?mi)([^.]+$)`)
	return r.FindString(path)
}

func parentDirFromPath(path string) string {
	r := regexp.MustCompile(`(?mi)([^/]+$)`)
	return r.FindString(path)
}

func matchInArray(s string, a []string) bool {
	var matched bool
	for _, v := range a {
		matched = strings.Contains(s, v)
		if (matched) {
			break
		}
	}
	return matched
}