func guessNumber(n int) int {
    low := 1
    high := n

    for low <= high {
        mid := low + (high-low)/2
        res := guess(mid)

        if res > 0 {
            low = mid + 1
        } else if res < 0 {
            high = mid - 1
        } else {
            return mid
        }
    }

    return -1
}