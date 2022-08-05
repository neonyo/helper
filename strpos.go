package helper

import "strings"

func MbStrPos(haystack, needle string) int {
	index := strings.Index(haystack, needle)
	if index == -1 || index == 0 {
		return index
	}
	pos := 0
	total := 0
	reader := strings.NewReader(haystack)
	for {
		_, size, err := reader.ReadRune()
		if err != nil {
			return -1
		}
		total += size
		pos++
		// got it
		if total == index {
			return pos
		}
	}
}
