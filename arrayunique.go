package helper

func arrayUniqueString(arr []string) []string {
	size := len(arr)
	result := make([]string, 0, size)
	temp := map[string]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

func arrayUniqueInt(arr []int) []int {
	size := len(arr)
	result := make([]int, 0, size)
	temp := map[int]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

func arrayUnique(arr []interface{}) []interface{} {
	size := len(arr)
	result := make([]interface{}, 0, size)
	temp := map[interface{}]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
