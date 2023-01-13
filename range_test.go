package grange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionForRangeInterfaceWithRangeTimeType(t *testing.T) {
	var range1 RangeTime = RangeTime{
		{Value: time.Now().Add(-10 * time.Hour)},
		{Value: time.Now().Add(-5 * time.Hour)},
	}
	range2 := RangeTime{
		{Value: time.Now().Add(-15 * time.Hour)},
		{Value: time.Now().Add(-7 * time.Hour)},
	}
	intersectionRange := Intersection[RangeTime](range1, range2)
	expectedIntersectionRange := RangeTime{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
}

func TestIntersectionForRangeInterfaceWithNumberType(t *testing.T) {
	range1 := RangeNumber[int]{
		{Value: 0}, {Value: 20},
	}
	range2 := RangeNumber[int]{
		{Value: -10}, {Value: 10},
	}
	intersectionRange := Intersection[RangeNumber[int]](range1, range2)
	expectedIntersectionRange := RangeNumber[int]{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
	assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
}

func TestDifferenceForRangeInterfaceWithNumberType(t *testing.T) {
	range1 := RangeNumber[int]{
		{Value: 0}, {Value: 20},
	}
	range2 := RangeNumber[int]{
		{Value: -10}, {Value: 10},
	}
	differenceRange := Difference[RangeNumber[int]](range1, range2)
	expectedDifferenceRange := []RangeNumber[int]{
		{range2[1],
			range1[1]},
	}
	assert.Equal(t, expectedDifferenceRange[0][0].Value, differenceRange[0][0].Value)
	assert.Equal(t, expectedDifferenceRange[0][1].Value, differenceRange[0][1].Value)
}
