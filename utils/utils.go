package utils

import (
	"log"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}
