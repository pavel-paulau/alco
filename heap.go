package main

type heapOrder int

const (
	maxOrder heapOrder = iota
	minOrder
)

type heap struct {
	items      []int64
	comparator func(int64, int64) bool
}

func max(i, j int64) bool { return i < j }

func min(i, j int64) bool { return i > j }

func newHeap(order heapOrder) *heap {
	var cmp func(int64, int64) bool

	if order == maxOrder {
		cmp = max
	} else {
		cmp = min
	}

	return &heap{
		items:      []int64{0},
		comparator: cmp,
	}
}

func (h *heap) push(item int64) {
	h.items = append(h.items, item)
	h.swim(h.size())
}

func (h *heap) pop() int64 {
	if h.size() < 1 {
		return 0
	}
	item := h.items[1]

	h.exch(1, h.size())
	h.items = h.items[0:h.size()]
	h.sink(1)

	return item
}

func (h *heap) head() int64 {
	if h.size() < 1 {
		return 0
	}
	return h.items[1]
}

func (h *heap) size() int {
	return len(h.items)-1
}

func (h *heap) less(i, j int) bool {
	return h.comparator(h.items[i], h.items[j])
}

func (h *heap) exch(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *heap) swim(k int) {
	for k > 1 && h.less(k/2, k) {
		h.exch(k/2, k)
		k = k / 2
	}
}

func (h *heap) sink(k int) {
	for 2*k <= h.size() {
		var j = 2 * k
		if j < h.size() && h.less(j, j+1) {
			j++
		}
		if !h.less(k, j) {
			break
		}
		h.exch(k, j)
		k = j
	}
}

func medianMaintenance(array []int64) int64 {
	hLeft := newHeap(maxOrder)
	hRight := newHeap(minOrder)

	var medians int64
	for _, v := range array {
		if hLeft.size() == 0 {
			hLeft.push(v)
		} else if hLeft.size() > hRight.size() {
			if v < hLeft.head() {
				nv := hLeft.pop()
				hLeft.push(v)
				hRight.push(nv)
			} else {
				hRight.push(v)
			}
		} else {
			if v > hRight.head() {
				nv := hRight.pop()
				hLeft.push(nv)
				hRight.push(v)
			} else {
				hLeft.push(v)
			}
		}
		medians += hLeft.head()
	}
	return medians
}
