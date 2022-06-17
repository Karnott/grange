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

func (dateRange RangeTime) formatIsZero() RangeTime {
	start := dateRange[0]
	end := dateRange[1]
	if start.IsZero() {
		start = leftInfinityDate
	}
	if end.IsZero() {
		end = rightInfinityDate
	}
	return RangeTime{
		start,
		end,
	}
}

func (dateRange RangeTime) formatToNumberRange() RangeNumber[int64] {
	dateRangeFormatted := dateRange.formatIsZero()
	return RangeNumber[int64]{
		dateRangeFormatted[0].UnixMilli(),
		dateRangeFormatted[1].UnixMilli(),
	}
}

func formatNumberRangeToDateRange(numberRange RangeNumber[int64]) RangeTime {
	start := numberRange[0]
	end := numberRange[1]
	var timeRange RangeTime
	if start != leftInfinityTimestamp {
		timeRange[0] = time.UnixMilli(start)
	}
	if end != rightInfinityTimestamp {
		timeRange[1] = time.UnixMilli(end)
	}
	return timeRange
}
