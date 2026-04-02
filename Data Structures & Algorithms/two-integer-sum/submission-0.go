func twoSum(nums []int, target int) []int {
    mapNum := make(map[int]int)

	for i, num := range nums{
		sub := target - num
		if v, ok := mapNum[sub]; ok {
			return []int{v, i}
		}
		
		mapNum[num] = i		
	}

	return nil
}
