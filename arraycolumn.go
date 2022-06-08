package helper

func ArrayColumn(input map[int]map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

func ArrayColumnKey(input map[int]map[string]interface{}, indexKey string) map[interface{}]map[string]interface{} {
	columns := make(map[interface{}]map[string]interface{}, len(input))
	for _, val := range input {
		if _, ok := val[indexKey]; ok {
			columns[indexKey] = val
		}
	}
	return columns
}
