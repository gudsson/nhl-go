package utils

import (
	"regexp"
)

func PeriodClock(str string) []string {
	re := regexp.MustCompile(`(\d{1,2}:\d{2})`)
	// str := "0:0020:00"
	split := re.FindAllString(str, 2)
	// fmt.Println(strings.SplitAfterN(str, ":", 2))
	times := []string{}
	for i := range split {
		times = append(times, split[i])
	}

	return times

	// fmt.Println(times) // ["Have", "a", "great", "day!"]
}