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
	RangeNumber[int] | RangeNumber[int16] | RangeNumber[int32] | RangeNumber[int64] | RangeNumber[float32] | RangeNumber[float64] | RangeTime
}

type RangeInterface[K RangeTypes] interface {
	Intersection(K) *K
	IsEmpty() bool
}

func Intersection[K RangeTypes](t RangeInterface[K], t1 K) *K {
	return t.Intersection(t1)
}
