package grange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionForTimeRange(t *testing.T) {
	range1 := RangeTime{
		time.Now().Add(-10 * time.Hour),
		time.Now().Add(-5 * time.Hour),
	}
	range2 := RangeTime{
		time.Now().Add(-15 * time.Hour),
		time.Now().Add(-7 * time.Hour),
	}
	intersectionRange := range1.Intersection(range2)
	expectedIntersectionRange := RangeTime{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].UnixMilli(), intersectionRange[0].UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].UnixMilli(), intersectionRange[1].UnixMilli())

	intersectionRange = range2.Intersection(range1)
	assert.Equal(t, expectedIntersectionRange[0].UnixMilli(), intersectionRange[0].UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].UnixMilli(), intersectionRange[1].UnixMilli())

	range3 := RangeTime{
		time.Now().Add(-15 * time.Hour),
		range1[0].Add(-time.Second),
	}
	intersectionRange = range1.Intersection(range3)
	assert.Nil(t, intersectionRange)

	range4 := RangeTime{
		time.Time{},
		range1[1],
	}
	intersectionRange = range1.Intersection(range4)
	expectedIntersectionRange = RangeTime{
		range1[0],
		range1[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].UnixMilli(), intersectionRange[0].UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].UnixMilli(), intersectionRange[1].UnixMilli())
}
