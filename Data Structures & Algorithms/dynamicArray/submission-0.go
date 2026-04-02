
type DynamicArray struct {
	length   int
	capacity int
	values   []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    if capacity <= 0 {
        capacity = 1
    }
	return &DynamicArray{
		length:   0,
		capacity: capacity,
		values:   make([]int, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.values[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.values[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.length == da.capacity {
		da.resize()
	}
	da.values[da.length] = n
	da.length++
	
}

func (da *DynamicArray) Popback() int {
	value := da.values[da.length-1]
	da.length -= 1
	return value

}

func (da *DynamicArray) resize() {
	da.capacity *= 2
	newValues := make([]int, da.capacity)
	copy(newValues, da.values)
	da.values = newValues
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
