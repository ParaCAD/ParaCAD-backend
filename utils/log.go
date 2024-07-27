package utils

import (
	"log"
	"strings"
)

const lineLength = 40

func PrintLine(title string) {
	if len(title) == 0 {
		log.Println(strings.Repeat("-", lineLength))
		return
	}
	log.Printf("%s %s", title, strings.Repeat("-", lineLength-len(title)-1))
}
