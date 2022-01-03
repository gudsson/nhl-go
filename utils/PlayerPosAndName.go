package utils

import (
	"strings"
)

func PlayerPosAndName(str string) []string {
	return strings.Split(str, " - ")
}