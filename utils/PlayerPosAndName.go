package utils

import (
	"strings"
)

func PlayerPosAndName(str string) []string {
	playerSlice := strings.Split(str, " - ")
	playerSlice[0] = playerSlice[0][0:1]
	return playerSlice
}