package grange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionForRangeInterfaceWithRangeTimeType(t *testing.T) {
	var range1 RangeTime = RangeTime{
		time.Now().Add(-10 * time.Hour),
		time.Now().Add(-5 * time.Hour),
	}
	range2 := RangeTime{
		time.Now().Add(-15 * time.Hour),
		time.Now().Add(-7 * time.Hour),
	}
	intersectionRange := Intersection[RangeTime](range1, range2)
	expectedIntersectionRange := RangeTime{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0].UnixMilli(), intersectionRange[0].UnixMilli())
	assert.Equal(t, expectedIntersectionRange[1].UnixMilli(), intersectionRange[1].UnixMilli())
}

func TestIntersectionForRangeInterfaceWithNumberType(t *testing.T) {
	range1 := RangeNumber[int]{
		0, 20,
	}
	range2 := RangeNumber[int]{
		-10, 10,
	}
	intersectionRange := Intersection[RangeNumber[int]](range1, range2)
	expectedIntersectionRange := RangeNumber[int]{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0], intersectionRange[0])
	assert.Equal(t, expectedIntersectionRange[1], intersectionRange[1])
}
