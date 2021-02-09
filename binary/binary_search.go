package binary

// Comparable comarable item
type Comparable interface {
	//Compare compare with item i ，return 1 if j>Comparable[i] , return -1 if j<IndexComparably[i] , return 0 if j<IndexComparably[i]
	Compare(j interface{}, i int) int
}

// Search binary search
func Search(v interface{}, a Comparable, start, end int, asc bool) int {
	mid := start + (end-start)/2
	if start > end {
		return -1
	}
	if asc {
		c := a.Compare(v, mid)
		if c < 0 {
			return Search(v, a, start, mid-1, asc)
		} else if c > 0 {
			return Search(v, a, mid+1, end, asc)
		} else {
			return mid
		}
	} else {
		c := a.Compare(v, mid)
		if c < 0 {
			return Search(v, a, mid+1, end, asc)
		} else if c > 0 {
			return Search(v, a, start, mid-1, asc)
		} else {
			return mid
		}
	}
}
