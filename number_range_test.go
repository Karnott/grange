package grange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInclusiveIntersectionForNumberRange(t *testing.T) {
	t.Run("expect intersection", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: -10}, {Value: 10},
		}
		intersectionRange := range1.Intersection(range2)
		expectedIntersectionRange := RangeNumber[int]{
			range1[0],
			range2[1],
		}
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)

		intersectionRange = range2.Intersection(range1)
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
	})

	t.Run("expect intersection", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: -20},
			{Value: range1[0].Value - 1},
		}
		intersectionRange := range1.Intersection(range2)
		assert.Nil(t, intersectionRange)
		// reverse test
		intersectionRange = range2.Intersection(range1)
		assert.Nil(t, intersectionRange)
	})
}

func TestExclusiveIntersectionForNumberRange(t *testing.T) {
	t.Run("expect intersection with exclusive end bound", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: -10}, {Value: 10, IsExclusive: true},
		}
		intersectionRange := range1.Intersection(range2)
		expectedIntersectionRange := RangeNumber[int]{
			range1[0],
			range2[1],
		}
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.False(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
		assert.True(t, expectedIntersectionRange[1].IsExclusive)

		intersectionRange = range2.Intersection(range1)
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.False(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
		assert.True(t, expectedIntersectionRange[1].IsExclusive)
	})

	t.Run("expect intersection with exclusive start bound", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0, IsExclusive: true}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: -10}, {Value: 10},
		}
		intersectionRange := range1.Intersection(range2)
		expectedIntersectionRange := RangeNumber[int]{
			range1[0],
			range2[1],
		}
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
		assert.False(t, expectedIntersectionRange[1].IsExclusive)

		intersectionRange = range2.Intersection(range1)
		assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
		assert.True(t, expectedIntersectionRange[0].IsExclusive)
		assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
		assert.False(t, expectedIntersectionRange[1].IsExclusive)
	})

	t.Run("expect no intersection because of exclusive intersection", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0, IsExclusive: true}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: -10}, {Value: 0},
		}
		intersectionRange := range1.Intersection(range2)
		assert.Nil(t, intersectionRange)

		intersectionRange = range2.Intersection(range1)
		assert.Nil(t, intersectionRange)
	})
}
