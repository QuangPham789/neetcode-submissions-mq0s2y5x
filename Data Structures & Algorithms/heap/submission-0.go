type MinHeap struct {
	Heap []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Heap: []int{0},
	}
}

func (mh *MinHeap) Push(val int) {
	mh.Heap = append(mh.Heap, val)

	i := len(mh.Heap) - 1
	for i > 1 && mh.Heap[i] < mh.Heap[i/2]{
		mh.Heap[i], mh.Heap[i/2] = mh.Heap[i/2], mh.Heap[i]
		 i /= 2
	}
}

func (mh *MinHeap) Pop() int {
	if len(mh.Heap) == 1 {
		return -1
	}

	if len(mh.Heap) == 2{
		res := mh.Heap[1]
    mh.Heap = mh.Heap[:1] // Đưa mảng về chỉ còn phần tử [0]
    return res
	}

	res := mh.Heap[1]
	mh.Heap[1] = mh.Heap[len(mh.Heap) - 1] 
	mh.Heap = mh.Heap[:len(mh.Heap) - 1] 
	i := 1

	for 2*i < len(mh.Heap){
		if 2*i+1 < len(mh.Heap) && mh.Heap[2*i] > mh.Heap[2*i+1] && mh.Heap[i] > mh.Heap[2*i+1]{
			mh.Heap[i], mh.Heap[2*i+1] = mh.Heap[2*i+1], mh.Heap[i]
			i = 2*i+1
		}else if mh.Heap[2*i] < mh.Heap[i]{
			mh.Heap[i], mh.Heap[2*i] = mh.Heap[2*i], mh.Heap[i]
			i = 2*i
		}else{
			break
		}
	}

	return res
}

func (mh *MinHeap) Top() int {
	if len(mh.Heap) > 1 {
		return mh.Heap[1]
	}
	return -1
}

func (mh *MinHeap) Heapify(nums []int) {
	mh.Heap = make([]int, len(nums)+1)
    copy(mh.Heap[1:], nums)
    
    cur := (len(mh.Heap) - 1) / 2

	for cur > 0 {
		i := cur
		for 2*i < len(mh.Heap){
			if 2*i+1 < len(mh.Heap) && mh.Heap[2*i] > mh.Heap[2*i+1] && mh.Heap[i] > mh.Heap[2*i+1]{
				mh.Heap[i], mh.Heap[2*i+1] = mh.Heap[2*i+1], mh.Heap[i]
				i = 2*i+1
			}else if mh.Heap[2*i] < mh.Heap[i]{
				mh.Heap[i], mh.Heap[2*i] = mh.Heap[2*i], mh.Heap[i]
				i = 2*i
			}else{
				break
			}
		}
		cur--
	}
}
