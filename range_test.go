package grange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionForRangeInterfaceWithDateTimeRangeType(t *testing.T) {
	var range1 DateTimeRange = DateTimeRange{
		{Value: time.Now().Add(-10 * time.Hour)},
		{Value: time.Now().Add(-5 * time.Hour)},
	}
	range2 := DateTimeRange{
		{Value: time.Now().Add(-15 * time.Hour)},
		{Value: time.Now().Add(-7 * time.Hour)},
	}
	intersectionRange := Intersection[DateTimeRange](range1, range2)
	expectedIntersectionRange := DateTimeRange{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].Value.UnixMilli(), intersectionRange[0].Value.UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].Value.UnixMilli(), intersectionRange[1].Value.UnixMilli())
}

func TestIntersectionForRangeInterfaceWithNumberType(t *testing.T) {
	range1 := NumberRange[int]{
		{Value: 0}, {Value: 20},
	}
	range2 := NumberRange[int]{
		{Value: -10}, {Value: 10},
	}
	intersectionRange := Intersection[NumberRange[int]](range1, range2)
	expectedIntersectionRange := NumberRange[int]{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].Value, intersectionRange[0].Value)
	assert.Equal(t, expectedIntersectionRange[1].Value, intersectionRange[1].Value)
}
