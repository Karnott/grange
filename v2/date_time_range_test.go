package grange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInclusiveIntersectionForTimeRange(t *testing.T) {
	t.Run("expect intersection", func(t *testing.T) {
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour)},
			{Value: time.Now().Add(-5 * time.Hour)},
		}
		range2 := DateTimeRange{
			{Value: time.Now().Add(-15 * time.Hour)},
			{Value: time.Now().Add(-7 * time.Hour)},
		}
		intersectionRange := range1.Intersection(range2)
		expectedIntersectionRange := DateTimeRange{
			range1[0],
			range2[1],
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())

		// reverse test
		intersectionRange = range2.Intersection(range1)
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
	})

	t.Run("expect no intersection", func(t *testing.T) {
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour)},
			{Value: time.Now().Add(-5 * time.Hour)},
		}
		range2 := DateTimeRange{
			{Value: time.Now().Add(-15 * time.Hour)},
			{Value: range1[0].Value.Add(-time.Second)},
		}
		intersectionRange := range1.Intersection(range2)
		assert.Nil(t, intersectionRange)
	})

	t.Run("expect intersection with infinity value", func(t *testing.T) {
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour)},
			{Value: time.Now().Add(-5 * time.Hour)},
		}
		range2 := DateTimeRange{
			{},
			range1[1],
		}
		intersectionRange := range1.Intersection(range2)
		expectedIntersectionRange := DateTimeRange{
			range1[0],
			range1[1],
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())

		// reverse test
		intersectionRange = range1.Intersection(range2)
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
	})
}

func TestExclusiveIntersectionForTimeRange(t *testing.T) {
	t.Run("expect no intersection with range with same bound value with both exclusive", func(t *testing.T) {
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour)},
			{Value: time.Now().Add(-5 * time.Hour), IsExclusive: true},
		}
		range2 := DateTimeRange{
			range1[1],
			{Value: time.Now().Add(-3 * time.Hour)},
		}
		intersectionRange := range1.Intersection(range2)
		assert.Nil(t, intersectionRange)
	})

	t.Run("expect return IsExclusive to true for intersection with range with same end bound value but different IsExclusive", func(t *testing.T) {
		// expect keep IsExclusive value
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour)},
			{Value: time.Now().Add(-5 * time.Hour), IsExclusive: true},
		}
		range2 := DateTimeRange{
			{Value: time.Now().Add(-7 * time.Hour), IsExclusive: true},
			{Value: range1[1].Value},
		}
		expectedIntersectionRange := DateTimeRange{
			range2[0],
			range1[1],
		}
		intersectionRange := range1.Intersection(range2)
		if !assert.NotNil(t, intersectionRange) {
			return
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[1].IsExclusive)

		// reverse func call to ensure IsExclusive is set in both case
		intersectionRange = range2.Intersection(range1)
		if !assert.NotNil(t, intersectionRange) {
			return
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[1].IsExclusive)
	})

	t.Run("expect return exclusive range with range with same start/end bound value but different IsExclusive", func(t *testing.T) {
		// expect keep IsExclusive value
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour), IsExclusive: true},
			{Value: time.Now().Add(-5 * time.Hour)},
		}
		range2 := DateTimeRange{
			{Value: range1[0].Value},
			{Value: range1[1].Value, IsExclusive: true},
		}
		expectedIntersectionRange := DateTimeRange{
			range1[0],
			range2[1],
		}
		intersectionRange := range1.Intersection(range2)
		if !assert.NotNil(t, intersectionRange) {
			return
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[1].IsExclusive)

		// reverse func call to ensure IsExclusive is set in both case
		intersectionRange = range2.Intersection(range1)
		if !assert.NotNil(t, intersectionRange) {
			return
		}
		assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
		assert.True(t, expectedIntersectionRange[1].IsExclusive)
	})
}

func TestDifferenceForTimeRange(t *testing.T) {
	t.Run("expect difference between same ranges with one of the bounds is exclusive", func(t *testing.T) {
		now := time.Now()
		range1 := DateTimeRange{
			{Value: now.Add(-10 * time.Hour).Truncate(time.Minute)},
			{Value: now.Add(-5 * time.Hour).Truncate(time.Minute), IsExclusive: true},
		}
		range2 := DateTimeRange{
			{Value: now.Add(-10 * time.Hour).Truncate(time.Minute)},
			{Value: now.Add(-5 * time.Hour).Truncate(time.Minute)},
		}
		diffRange := range1.Difference(range2)
		assert.Equal(t, 1, len(diffRange))
		assert.True(t, diffRange[0][0].Value.Equal(range2[1].Value))
		assert.True(t, diffRange[0][1].Value.Equal(range2[1].Value))
	})

	t.Run("expect difference with second range includes in first range", func(t *testing.T) {
		// 1. --------------
		// 2.    ------
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour).Truncate(time.Minute)},
			{Value: time.Now().Add(-5 * time.Hour).Truncate(time.Minute)},
		}
		range2 := DateTimeRange{
			{Value: time.Now().Add(-7 * time.Hour).Truncate(time.Minute)},
			{Value: time.Now().Add(-6 * time.Hour).Truncate(time.Minute)},
		}
		diffRange := range1.Difference(range2)
		expectedDifferenceRange := []DateTimeRange{
			{range1[0], range2[0]},
			{range2[1], range1[1]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.True(t, diffRange[0][0].Value.Equal(expectedDifferenceRange[0][0].Value))
		assert.True(t, diffRange[0][1].Value.Equal(expectedDifferenceRange[0][1].Value))
		assert.True(t, diffRange[1][0].Value.Equal(expectedDifferenceRange[1][0].Value))
		assert.True(t, diffRange[1][1].Value.Equal(expectedDifferenceRange[1][1].Value))

		diffRange = range2.Difference(range1)
		assert.Equal(t, 0, len(diffRange))
	})

	t.Run("expect difference with 2 ranges in intersection", func(t *testing.T) {
		// 1. --------------
		// 2.            ------
		range1 := DateTimeRange{
			{Value: time.Now().Add(-10 * time.Hour).Truncate(time.Minute)},
			{Value: time.Now().Add(-6 * time.Hour).Truncate(time.Minute)},
		}
		range2 := DateTimeRange{
			{Value: time.Now().Add(-7 * time.Hour).Truncate(time.Minute)},
			{Value: time.Now().Add(-3 * time.Hour).Truncate(time.Minute)},
		}
		diffRange := range1.Difference(range2)
		expectedDifferenceRange := []DateTimeRange{
			{range1[0], range2[0]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.True(t, diffRange[0][0].Value.Equal(expectedDifferenceRange[0][0].Value))
		assert.True(t, diffRange[0][1].Value.Equal(expectedDifferenceRange[0][1].Value))

		diffRange = range2.Difference(range1)
		expectedDifferenceRange = []DateTimeRange{
			{range1[1], range2[1]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)
	})
}
