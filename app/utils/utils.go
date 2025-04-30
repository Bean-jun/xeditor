package utils

import (
	"log"
	"regexp"
)

func RefindNumber(strs string) string {
	patten := regexp.MustCompile(`cpython-(\d+\.\d+\.\d+)-windows-x86_64-none`)
	match := patten.FindStringSubmatch(strs)
	log.Println(match)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
