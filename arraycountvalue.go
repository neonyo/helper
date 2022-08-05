package helper

func ArrayCountValues(s []interface{}) map[interface{}]uint {
	r := make(map[interface{}]uint)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
		} else {
			r[v] = 1
		}
	}
	return r
}

func ArrayUniqueValue(s []string) (unique []string) {
	r := make(map[string]uint)
	unique = make([]string, 0)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
			unique = append(unique, v)
		} else {
			r[v] = 1
		}
	}
	return
}
