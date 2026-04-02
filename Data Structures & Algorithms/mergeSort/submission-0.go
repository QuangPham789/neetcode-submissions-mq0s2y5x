// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    if len(pairs) <= 1 {
        return pairs
    }

    mid := len(pairs) / 2
    left := mergeSort(pairs[:mid])
    right := mergeSort(pairs[mid:])

    return merge(left, right)
}

func merge(left, right []Pair) []Pair{
    result := []Pair{}
    i, j := 0, 0
    for i < len(left) && j < len(right){
        if left[i].Key <= right[j].Key{
           result = append(result, left[i])
           i++
        }else {
            result = append(result, right[j])
            j++
        }
    }

    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}
