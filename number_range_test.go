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

func TestDifferenceForNumberRange(t *testing.T) {
	t.Run("expect difference between same ranges with one of the bounds is exclusive", func(t *testing.T) {
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: 0}, {Value: 20, IsExclusive: true},
		}
		diffRange := range1.Difference(range2)
		expectedDifferenceRange := []RangeNumber[int]{
			{range1[1], range1[1]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)
		assert.False(t, diffRange[0][0].IsExclusive)
		assert.False(t, diffRange[0][1].IsExclusive)

		diffRange = range2.Difference(range1)
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)
		assert.False(t, diffRange[0][0].IsExclusive)
		assert.False(t, diffRange[0][1].IsExclusive)
	})

	t.Run("expect difference with second range includes in first range", func(t *testing.T) {
		// 1. --------------
		// 2.    ------
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: 10}, {Value: 15},
		}
		diffRange := range1.Difference(range2)
		expectedDifferenceRange := []RangeNumber[int]{
			{range1[0], range2[0]},
			{range2[1], range1[1]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)
		assert.Equal(t, diffRange[1][0].Value, expectedDifferenceRange[1][0].Value)
		assert.Equal(t, diffRange[1][1].Value, expectedDifferenceRange[1][1].Value)

		diffRange = range2.Difference(range1)
		assert.Equal(t, 0, len(diffRange))
	})

	t.Run("expect difference with 2 ranges in intersection", func(t *testing.T) {
		// 1. --------------
		// 2.   		 ------
		range1 := RangeNumber[int]{
			{Value: 0}, {Value: 20},
		}
		range2 := RangeNumber[int]{
			{Value: 10}, {Value: 30},
		}
		diffRange := range1.Difference(range2)
		expectedDifferenceRange := []RangeNumber[int]{
			{range1[0], range2[0]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)

		diffRange = range2.Difference(range1)
		expectedDifferenceRange = []RangeNumber[int]{
			{range1[1], range2[1]},
		}
		assert.Equal(t, len(diffRange), len(expectedDifferenceRange))
		assert.Equal(t, diffRange[0][0].Value, expectedDifferenceRange[0][0].Value)
		assert.Equal(t, diffRange[0][1].Value, expectedDifferenceRange[0][1].Value)
	})

}
