package heap

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func smallestt(h Interface) {
	for i := h.Len()/2 - 1; i >= 0; i-- {
		down(h, i)
	}
}

func Pop(h Interface) interface{} {
	h.Swap(0, h.Len()-1)
	down(h, 0)
	return h.Pop()
}

func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func down(h Interface, i int) {
	n := h.Len()

	for {
		smallest := i

		cl := smallest*2 + 1
		cr := smallest*2 + 2

		if cr < n && h.Less(cr, smallest) {
			smallest = cr
		}

		if cl < n && h.Less(cl, smallest) {
			smallest = cl
		}

		if smallest == i {
			return
		}
		h.Swap(smallest, i)
	}
}

func up(h Interface, i int) {
	for i > 0 && h.Less(i, (i-1)/2) {
		h.Swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}
