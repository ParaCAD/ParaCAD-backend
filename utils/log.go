package utils

import (
	"log"
	"strings"
)

const LINE_LENGTH = 40

func PrintLine(title string) {
	if len(title) == 0 {
		log.Println(strings.Repeat("-", LINE_LENGTH))
		return
	}
	log.Printf("%s %s", title, strings.Repeat("-", LINE_LENGTH-len(title)-1))
}
