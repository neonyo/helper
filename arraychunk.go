package helper

import (
	"errors"
	"math"
)

func ArrayChunk(s []interface{}, size int) (arr [][]interface{}, err error) {
	if size < 1 {
		return nil, errors.New("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		arr = append(arr, s[i*size:end])
		i++
	}
	return
}
