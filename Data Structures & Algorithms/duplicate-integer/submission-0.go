func hasDuplicate(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}

	m := make(map[int]int)

	for _,num := range nums {
		_, ok := m[num]
		if !ok{
			m[num]=num
		}	
	}

	if len(m) != len(nums){
		return true
	}

	return false
    
}
