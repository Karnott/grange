package grange

import "time"

type Integer interface {
	int | int64 | int32 | int16
}
type Float interface {
	float32 | float64
}
type Number interface {
	Integer | Float
}
type NumberTime interface {
	Number | time.Time
}
type RangeValue[K NumberTime] struct {
	Value       K
	IsExclusive bool
}
type Range[K NumberTime] [2]RangeValue[K]

type RangeTypes interface {
	NumberRange[int] | NumberRange[int16] | NumberRange[int32] | NumberRange[int64] | NumberRange[float32] | NumberRange[float64] | DateTimeRange
}

type RangeInterface[K RangeTypes] interface {
	Intersection(K) *K
	Difference(K) []K
	IsEmpty() bool
	ToPostgresString() (string, string, string) // return lowerBound, upperBound, inclusivity/exclusivity of bounds for postgres
}

func Intersection[K RangeTypes](t RangeInterface[K], t1 K) *K {
	return t.Intersection(t1)
}

func Difference[K RangeTypes](t RangeInterface[K], t1 K) []K {
	return t.Difference(t1)
}
