func replaceElements(arr []int) []int {
   maxSoFar := -1
   for i:= len(arr) - 1; i >= 0; i--{
    temp := arr[i]
    arr[i] = maxSoFar
    if maxSoFar < temp{
        maxSoFar = temp
    }
   }
   return arr
}
