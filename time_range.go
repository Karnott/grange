package grange

import (
	"math"
	"time"
)

type RangeTime Range[time.Time]

var (
	leftInfinityTimestamp  = int64(math.MinInt64)
	rightInfinityTimestamp = int64(math.MaxInt64)
	leftInfinityDate       = time.UnixMilli(leftInfinityTimestamp)
	rightInfinityDate      = time.UnixMilli(rightInfinityTimestamp)
)

// date.IsZero() is considered as `-infinty` for fromDate & `infinity` for toDate
// infinity var are defined as leftInfinityDate and rightInfinityDate
func (r RangeTime) Intersection(r1 RangeTime) *RangeTime {
	var intersectionResult = intersection(
		r.formatToNumberRange(),
		r1.formatToNumberRange(),
	)
	if intersectionResult == nil {
		return nil
	}
	intersectionRange := formatNumberRangeToDateRange(*intersectionResult)
	return &intersectionRange
}

func (r RangeTime) IsEmpty() bool {
	return r[0].Value.Equal(r[1].Value) && (r[0].IsExclusive || r[1].IsExclusive)
}

func (dateRange RangeTime) formatIsZero() RangeTime {
	start := dateRange[0]
	end := dateRange[1]
	if start.Value.IsZero() {
		start.Value = leftInfinityDate
	}
	if end.Value.IsZero() {
		end.Value = rightInfinityDate
	}
	return RangeTime{
		start,
		end,
	}
}

func (dateRange RangeTime) formatToNumberRange() RangeNumber[int64] {
	dateRangeFormatted := dateRange.formatIsZero()
	start := dateRangeFormatted[0]
	end := dateRangeFormatted[1]
	return RangeNumber[int64]{
		{Value: start.Value.UnixMilli(), IsExclusive: start.IsExclusive},
		{Value: end.Value.UnixMilli(), IsExclusive: end.IsExclusive},
	}
}

func formatNumberRangeToDateRange(numberRange RangeNumber[int64]) RangeTime {
	start := numberRange[0]
	end := numberRange[1]
	var timeRange RangeTime
	timeRange[0].IsExclusive = start.IsExclusive
	timeRange[1].IsExclusive = end.IsExclusive
	if start.Value != leftInfinityTimestamp {
		timeRange[0].Value = time.UnixMilli(start.Value)
	}
	if end.Value != rightInfinityTimestamp {
		timeRange[1].Value = time.UnixMilli(end.Value)
	}
	return timeRange
}
