package grange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionForNumberRange(t *testing.T) {
	range1 := RangeNumber[int]{
		0, 20,
	}
	range2 := RangeNumber[int]{
		-10, 10,
	}
	intersectionRange := range1.Intersection(range2)
	expectedIntersectionRange := RangeNumber[int]{
		range1[0],
		range2[1],
	}
	assert.Equal(t, expectedIntersectionRange[0], intersectionRange[0])
	assert.Equal(t, expectedIntersectionRange[1], intersectionRange[1])

	intersectionRange = range2.Intersection(range1)
	assert.Equal(t, expectedIntersectionRange[0], intersectionRange[0])
	assert.Equal(t, expectedIntersectionRange[1], intersectionRange[1])

	range3 := RangeNumber[int]{
		-20,
		range1[0] - 1,
	}
	intersectionRange = range1.Intersection(range3)
	assert.Nil(t, intersectionRange)

}
