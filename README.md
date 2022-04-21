# lognest_consecutive_seq

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

## 題目解析

要找最大連續整數長度 並且需要在 O(n)

可以透過 hashTable 紀錄有出現過的元素

連續的意思是兩個值之間只差 1

每次只要找出某一個值 檢查那個值 - 1 不存在 hash 內就知道他最小

然後透過 最小值 開始 +1 是否有存在 開始紀錄次數

這樣只要跑 N 次即可找到最大連續長度

## 實作

```golang
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
```
