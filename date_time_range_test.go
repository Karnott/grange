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
