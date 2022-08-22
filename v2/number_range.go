package grange

import "fmt"

type NumberRange[K Number] Range[K]

func (r NumberRange[K]) ToPostgresString() (string, string, string) {
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

func (r NumberRange[K]) Intersection(r1 NumberRange[K]) *NumberRange[K] {
	return intersection(r, r1)
}

func (r NumberRange[K]) IsEmpty() bool {
	return r[0].Value == r[1].Value && (r[0].IsExclusive || r[1].IsExclusive)
}

// if start and end bound have same value and one of the value is exclusive
// so the range is empty
func formatEmptyExclusiveRange[K Number](r NumberRange[K]) *NumberRange[K] {
	if r.IsEmpty() {
		return nil
	}
	return &r
}
