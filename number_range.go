package grange

type RangeNumber[K Number] Range[K]

func (r RangeNumber[K]) Intersection(r1 RangeNumber[K]) *RangeNumber[K] {
	return intersection(r, r1)
}
