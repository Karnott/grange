package grange

func diff[K Number, L NumberRange[K]](r L, r1 L) []NumberRange[K] {
	// return substract dateRange2 into dateRange1
	// if no intersection, the dateRange1 is returneed
	start1 := r[0]
	end1 := r[1]
	start2 := r1[0]
	end2 := r1[1]

	if end1.Value < start2.Value || end2.Value < start1.Value {
		return []NumberRange[K]{
			{start1, end1},
		}
	}

	ranges := make([]NumberRange[K], 0)
	if start1.Value < start2.Value && start2.Value < end1.Value {
		ranges = append(ranges, NumberRange[K]{start1, start2})
	}

	if start1.Value == start2.Value && (start2.IsExclusive != start1.IsExclusive) {
		if !start1.IsExclusive {
			ranges = append(ranges, NumberRange[K]{start1, start1})
		} else {
			ranges = append(ranges, NumberRange[K]{start2, start2})

		}
	}

	if end1.Value == end2.Value && (end2.IsExclusive != end1.IsExclusive) {
		if !end1.IsExclusive {
			ranges = append(ranges, NumberRange[K]{end1, end1})
		} else {
			ranges = append(ranges, NumberRange[K]{end2, end2})

		}
	}

	if end2.Value < end1.Value && end2.Value > start1.Value && end2.Value != end1.Value {
		ranges = append(ranges, NumberRange[K]{end2, end1})
	}

	return ranges
}
