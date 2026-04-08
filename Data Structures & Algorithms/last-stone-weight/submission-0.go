type Heap struct{
	MaxHeap []int
}
func lastStoneWeight(stones []int) int {
	heap := NewHeap()

	heap.heapify(stones)

	for len(heap.MaxHeap) > 2{
		x := heap.Pop()
		y := heap.Pop()

		i := x - y

		if i > 0 {
			heap.Push(i)
		}
	}

	res := heap.Pop()
	if res == -1 { return 0 }
	return res
}

func NewHeap() *Heap{
	return &Heap{
		MaxHeap: nil,
	}
}

func (h *Heap) heapify(stones []int) {
	h.MaxHeap = make([]int, len(stones)+1)
	copy(h.MaxHeap[1:], stones)

	curr := len(h.MaxHeap)/2

	for curr > 0{
		i := curr
		for 2*i < len(h.MaxHeap){
			if 2*i+1 < len(h.MaxHeap) && h.MaxHeap[2*i+1] > h.MaxHeap[2*i] && h.MaxHeap[2*i+1] > h.MaxHeap[i]{
				h.MaxHeap[i], h.MaxHeap[2*i+1] = h.MaxHeap[2*i+1], h.MaxHeap[i]
				i = 2*i+1
			}else if h.MaxHeap[2*i] > h.MaxHeap[i]{
				h.MaxHeap[i], h.MaxHeap[2*i] = h.MaxHeap[2*i], h.MaxHeap[i]
				i = 2*i
			}else{
				break
			}
		}
		curr--
	}

}

func (h *Heap) Pop() int {
	if len(h.MaxHeap) == 1 {
		return -1 // In Go, we can't return nil for int type. So, we return -1 to denote error.
	}
	if len(h.MaxHeap) == 2 {
		popVal := h.MaxHeap[len(h.MaxHeap)-1]
		h.MaxHeap = h.MaxHeap[:len(h.MaxHeap)-1]
		return popVal
	}

	res := h.MaxHeap[1]
	// Move last value to root
	h.MaxHeap[1] = h.MaxHeap[len(h.MaxHeap)-1]
	h.MaxHeap = h.MaxHeap[:len(h.MaxHeap)-1]
	i := 1
	// Percolate down
	for 2*i < len(h.MaxHeap) {
		if 2*i+1 < len(h.MaxHeap) && h.MaxHeap[2*i+1] > h.MaxHeap[2*i] && h.MaxHeap[i] < h.MaxHeap[2*i+1] {
			// Swap right child
			tmp := h.MaxHeap[i]
			h.MaxHeap[i] = h.MaxHeap[2*i+1]
			h.MaxHeap[2*i+1] = tmp
			i = 2*i + 1
		} else if h.MaxHeap[i] < h.MaxHeap[2*i] {
			// Swap left child
			tmp := h.MaxHeap[i]
			h.MaxHeap[i] = h.MaxHeap[2*i]
			h.MaxHeap[2*i] = tmp
			i = 2 * i
		} else {
			break
		}
	}
	return res
}

func (h *Heap) Push(val int) {
	h.MaxHeap = append(h.MaxHeap, val)
	i := len(h.MaxHeap) - 1

	// Percolate up
	for i > 1 && h.MaxHeap[i] > h.MaxHeap[i/2] {
		// Swap values
		tmp := h.MaxHeap[i]
		h.MaxHeap[i] = h.MaxHeap[i/2]
		h.MaxHeap[i/2] = tmp
		i = i / 2
	}
}

