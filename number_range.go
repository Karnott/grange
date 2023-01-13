package grange

import "fmt"

type RangeNumber[K Number] Range[K]

func (r RangeNumber[K]) ToPostgresString() (string, string, string) {
	start := r[0]
	end := r[1]
	startBoundString := "["
	endBoundString := "]"
	if start.IsExclusive {
		startBoundString = "("
	}
	if end.IsExclusive {
		endBoundString = ")"
	}

	return fmt.Sprintf("%v", start.Value), fmt.Sprintf("%v", end.Value), startBoundString + endBoundString
}

func (r RangeNumber[K]) Intersection(r1 RangeNumber[K]) *RangeNumber[K] {
	return intersection(r, r1)
}

func (r RangeNumber[K]) IsEmpty() bool {
	return r[0].Value == r[1].Value && (r[0].IsExclusive || r[1].IsExclusive)
}

// if start and end bound have same value and one of the value is exclusive
// so the range is empty
func formatEmptyExclusiveRange[K Number](r RangeNumber[K]) *RangeNumber[K] {
	if r.IsEmpty() {
		return nil
	}
	return &r
}

func (r RangeNumber[K]) Difference(r1 RangeNumber[K]) []RangeNumber[K] {
	return diff(r, r1)
}
