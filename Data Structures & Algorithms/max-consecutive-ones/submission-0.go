func findMaxConsecutiveOnes(nums []int) int {
	count := 0
    max := 0
    
    for _, num := range nums{
        if num == 1{
            count++
            if max < count{
                max = count
            }
        }else{
            count = 0
        }
    } 

    return max
}
