package ds

type LinkedList struct {
	Data 	[]Node
	Len 	int
	Max 	[2]float64
	Min 	[2]float64
}

type Node struct {
	Body float64
	Next *Node
}

func (ll *LinkedList) Fill(source []float64) {
	min := source[0]
	max := source[0]
	ll.Len = len(source)
	for i := 0; i < ll.Len; i++ {
		if min > source[i] {
			min = source[i]
		}
		if max < source[i] {
			max = source[i]
		}
		ll.Data = append(ll.Data, Node{Body: source[i]})
		if i != 0 {
			ll.Data[i-1].Next = &ll.Data[i]
		}
	}
}

