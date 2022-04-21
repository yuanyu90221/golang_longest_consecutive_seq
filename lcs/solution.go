package lcs

func longestConsecutive(nums []int) int {
	visitedMap := make(map[int]bool)
	uniqueValues := make([]int, 0)
	for _, num := range nums {
		if _, ok := visitedMap[num]; !ok {
			uniqueValues = append(uniqueValues, num)
			visitedMap[num] = true
		}
	}
	if len(uniqueValues) == 1 {
		return 1
	}
	max_len := 0
	for _, num := range uniqueValues {
		if _, ok := visitedMap[num-1]; !ok {
			currentNum := num
			currentAccum := 1
			for _, existed := visitedMap[currentNum+1]; existed; _, existed = visitedMap[currentNum+1] {
				currentNum += 1
				currentAccum += 1
			}
			if currentAccum > max_len {
				max_len = currentAccum
			}
		}
	}
	return max_len
}
