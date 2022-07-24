package main

func main() {

}

func majorityElement(nums []int) int {
	n := len(nums)

	majority := nums[0]
	majorityCount := 1

	m := make(map[int]int)
	m[nums[0]]++

	for i := 1; i < n; i++ {
		m[nums[i]]++
		if m[nums[i]] > majorityCount {
			majorityCount = m[nums[i]]
			majority = nums[i]
		}
	}

	return majority
}
