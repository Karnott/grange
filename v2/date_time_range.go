package grange

import (
	"math"
	"time"
)

type DateTimeRange Range[time.Time]

var (
	leftInfinityTimestamp  = int64(math.MinInt64)
	rightInfinityTimestamp = int64(math.MaxInt64)
	leftInfinityDate       = time.UnixMilli(leftInfinityTimestamp)
	rightInfinityDate      = time.UnixMilli(rightInfinityTimestamp)
)

func (r DateTimeRange) ToPostgresString() (string, string, string) {
	fromDateString := "-infinity"
	toDateString := "infinity"
	fromDate := r[0]
	toDate := r[1]
	if !fromDate.Value.IsZero() {
		fromDateString = fromDate.Value.Format(time.RFC3339)
	}
	if !toDate.Value.IsZero() {
		toDateString = toDate.Value.Format(time.RFC3339)
	}
	fromDateBoundString := "["
	toDateBoundString := "]"
	if fromDate.IsExclusive {
		fromDateBoundString = "("
	}
	if toDate.IsExclusive {
		toDateBoundString = ")"
	}

	return fromDateString, toDateString, fromDateBoundString + toDateBoundString
}

// date.IsZero() is considered as `-infinty` for fromDate & `infinity` for toDate
// infinity var are defined as leftInfinityDate and rightInfinityDate
func (r DateTimeRange) Intersection(r1 DateTimeRange) *DateTimeRange {
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

func (r DateTimeRange) IsEmpty() bool {
	return r[0].Value.Equal(r[1].Value) && (r[0].IsExclusive || r[1].IsExclusive)
}

func (dateRange DateTimeRange) formatIsZero() DateTimeRange {
	start := dateRange[0]
	end := dateRange[1]
	if start.Value.IsZero() {
		start.Value = leftInfinityDate
	}
	if end.Value.IsZero() {
		end.Value = rightInfinityDate
	}
	return DateTimeRange{
		start,
		end,
	}
}

func (dateRange DateTimeRange) formatToNumberRange() NumberRange[int64] {
	dateRangeFormatted := dateRange.formatIsZero()
	start := dateRangeFormatted[0]
	end := dateRangeFormatted[1]
	return NumberRange[int64]{
		{Value: start.Value.UnixMilli(), IsExclusive: start.IsExclusive},
		{Value: end.Value.UnixMilli(), IsExclusive: end.IsExclusive},
	}
}

func formatNumberRangeToDateRange(numberRange NumberRange[int64]) DateTimeRange {
	start := numberRange[0]
	end := numberRange[1]
	var dateTimeRange DateTimeRange
	dateTimeRange[0].IsExclusive = start.IsExclusive
	dateTimeRange[1].IsExclusive = end.IsExclusive
	if start.Value != leftInfinityTimestamp {
		dateTimeRange[0].Value = time.UnixMilli(start.Value)
	}
	if end.Value != rightInfinityTimestamp {
		dateTimeRange[1].Value = time.UnixMilli(end.Value)
	}
	return dateTimeRange
}

func (r DateTimeRange) Difference(r1 DateTimeRange) []DateTimeRange {
	var differenceResult = diff(
		r.formatToNumberRange(),
		r1.formatToNumberRange(),
	)
	if len(differenceResult) == 0 {
		return nil
	}
	diffDateTimeRange := make([]DateTimeRange, 0)
	for _, rangeNumber := range differenceResult {
		diffDateTimeRange = append(diffDateTimeRange, formatNumberRangeToDateRange(rangeNumber))
	}
	return diffDateTimeRange
}
