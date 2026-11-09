import "container/heap"

type Server struct {
	id   int
	free int
}

type MinHeap []Server

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].free < h[j].free }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Server))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type SegmentTree struct {
	size int
	tree []int
}

func NewSegmentTree(n int) *SegmentTree {
	size := 1
	for size < n {
		size *= 2
	}

	tree := make([]int, 2*size)

	for i := size; i < size+n; i++ {
		tree[i] = 1
	}

	for i := size - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}

	return &SegmentTree{size, tree}
}

func (st *SegmentTree) Update(pos, value int) {
	pos += st.size
	st.tree[pos] = value

	for pos > 1 {
		pos /= 2
		st.tree[pos] = st.tree[pos*2] + st.tree[pos*2+1]
	}
}

func (st *SegmentTree) FindFirst(node, left, right, start int) int {
	if right < start || st.tree[node] == 0 {
		return -1
	}

	if left == right {
		return left
	}

	mid := (left + right) / 2

	result := st.FindFirst(node*2, left, mid, start)
	if result != -1 {
		return result
	}

	return st.FindFirst(node*2+1, mid+1, right, start)
}

func busiestServers(k int, arrival []int, load []int) []int {
	count := make([]int, k)

	st := NewSegmentTree(k)

	busy := &MinHeap{}
	heap.Init(busy)

	for i := 0; i < len(arrival); i++ {

		// Release servers that have finished.
		for busy.Len() > 0 && (*busy)[0].free <= arrival[i] {
			server := heap.Pop(busy).(Server)
			st.Update(server.id, 1)
		}

		// Find the next available server.
		start := i % k

		serverID := st.FindFirst(1, 0, st.size-1, start)

		// Wrap around if necessary.
		if serverID == -1 || serverID >= k {
			serverID = st.FindFirst(1, 0, st.size-1, 0)
		}

		if serverID == -1 || serverID >= k {
			continue
		}

		// Mark the server as busy.
		st.Update(serverID, 0)

		// Increment its request count.
		count[serverID]++

		// Add it to the busy heap.
		heap.Push(busy, Server{
			id:   serverID,
			free: arrival[i] + load[i],
		})
	}

	// Find the maximum request count.
	maxCount := 0
	for _, c := range count {
		if c > maxCount {
			maxCount = c
		}
	}

	// Return all busiest servers.
	result := []int{}

	for i, c := range count {
		if c == maxCount {
			result = append(result, i)
		}
	}

	return result
}