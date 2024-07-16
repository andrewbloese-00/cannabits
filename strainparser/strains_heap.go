package strainparser

type StrainHeapNode struct {
	Score       int
	StrainEntry *StrainEntry
}

type StrainsHeap struct {
	Items []*StrainHeapNode
	Size  int
}

func NewStrainsHeap() *StrainsHeap {
	return &StrainsHeap{Items: make([]*StrainHeapNode, 0), Size: 0}
}

func heap_parent(i int) int {
	return (i - 1) / 2
}
func heap_left(i int) int {
	return 2*i + 1
}
func heap_right(i int) int {
	return 2*i + 2
}

func (h *StrainsHeap) swap(a, b int) {
	h.Items[a], h.Items[b] = h.Items[b], h.Items[a]
}

func (h *StrainsHeap) heapifyUp(idx int) {
	for h.Items[heap_parent(idx)].Score < h.Items[idx].Score {
		p := heap_parent(idx)
		h.swap(p, idx)
		idx = p
	}
}

func (h *StrainsHeap) Insert(entry *StrainEntry, score int) {
	node := &StrainHeapNode{Score: score, StrainEntry: entry}
	h.Items = append(h.Items, node)
	h.Size += 1
	h.heapifyUp(h.Size - 1)
}

func (h *StrainsHeap) heapifyDown(idx int) {
	greatest := idx
	left, right := heap_left(idx), heap_right(idx)

	if left < h.Size && h.Items[left].Score > h.Items[greatest].Score {
		greatest = left
	}

	if right < h.Size && h.Items[right].Score > h.Items[greatest].Score {
		greatest = right
	}

	if greatest != idx {
		h.swap(idx, greatest)
		h.heapifyDown(greatest)
	}

}

func (h *StrainsHeap) Extract() *StrainEntry {
	if h.Size == 0 {
		return nil
	}

	root := h.Items[0]
	h.Items[0] = h.Items[h.Size-1]
	h.Items = h.Items[:h.Size-1]
	h.Size--
	h.heapifyDown(0)
	return root.StrainEntry

}
