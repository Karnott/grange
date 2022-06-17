package grange

import (
	"time"
)

type Range[K Number | time.Time] [2]K

type RangeTypes[K Number] interface {
	RangeNumber[K] | RangeTime
}

type RangeInterface[K RangeTypes[L], L Number] interface {
	Intersection(K) *K
}

func Intersection[K RangeTypes[L], L Number](t RangeInterface[K, L], t1 K) *K {
	return t.Intersection(t1)
}
