package algs

func twoSum(lst []int, target int) []int {
	indexes := make(map[int]int)
	for k, v := range lst {
		if idx, ok := indexes[target-v]; ok {
			return []int{idx, k}
		}
		indexes[v] = k
	}
	return nil
}
