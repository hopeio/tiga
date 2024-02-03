package queue

// 就是heap,length可省略
type TinyQueue struct {
	data []Item
}

type Item interface {
	Less(Item) bool
}

func New(data []Item) *TinyQueue {
	q := &TinyQueue{}
	q.data = data
	n := len(data)
	if n > 0 {
		i := n >> 1
		for ; i >= 0; i-- {
			q.down(i)
		}
	}
	return q
}

func (q *TinyQueue) Push(item Item) {
	q.data = append(q.data, item)
	q.up(len(q.data) - 1)
}
func (q *TinyQueue) Pop() Item {
	n := len(q.data)
	if n == 0 {
		return nil
	}
	top := q.data[0]
	if n > 0 {
		q.data[0] = q.data[n]
		q.down(0)
	}
	q.data = q.data[:n-1]
	return top
}
func (q *TinyQueue) Peek() Item {
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}
func (q *TinyQueue) Len() int {
	return len(q.data)
}
func (q *TinyQueue) down(pos int) {
	data := q.data
	halfLength := len(q.data) >> 1
	item := data[pos]
	for pos < halfLength {
		left := (pos << 1) + 1
		right := left + 1
		best := data[left]
		if right < len(q.data) && data[right].Less(best) {
			left = right
			best = data[right]
		}
		if !best.Less(item) {
			break
		}
		data[pos] = best
		pos = left
	}
	data[pos] = item
}

func (q *TinyQueue) up(pos int) {
	data := q.data
	item := data[pos]
	for pos > 0 {
		parent := (pos - 1) >> 1
		current := data[parent]
		if !item.Less(current) {
			break
		}
		data[pos] = current
		pos = parent
	}
	data[pos] = item
}
