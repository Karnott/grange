package grange

// return [start, end] intersection range between 2 ranges
// return nil if no intersection
func intersection[K Number, L RangeNumber[K]](r L, r1 L) *RangeNumber[K] {
	// if one bound with some value IsExclusive
	// 		=> all the other bound with the same value must be exclusive
	//			 to ensure to return intersection range with exclusive to `true` for this value
	for i, valueRange1 := range r {
		for i2, valueRange2 := range r1 {
			if valueRange2.Value == valueRange2.Value && valueRange1.IsExclusive != valueRange2.IsExclusive {
				r[i].IsExclusive = true
				r1[i2].IsExclusive = true
			}
		}
	}
	start1 := r[0]
	end1 := r[1]
	start2 := r1[0]
	end2 := r1[1]

	if end2.Value < start1.Value || end1.Value < start2.Value {
		return nil
	}

	if start1.Value < start2.Value {
		if end1.Value > end2.Value {
			return formatEmptyExclusiveRange(
				RangeNumber[K]{
					start2,
					end2,
				},
			)
		}
		if end1.Value <= end2.Value {
			return formatEmptyExclusiveRange(
				RangeNumber[K]{
					start2,
					end1,
				},
			)
		}
	}
	if start2.Value <= start1.Value {
		if end2.Value > end1.Value {
			return formatEmptyExclusiveRange(
				RangeNumber[K]{
					start1,
					end1,
				},
			)
		}

		if end2.Value <= end1.Value {
			return formatEmptyExclusiveRange(
				RangeNumber[K]{
					start1,
					end2,
				},
			)
		}
	}

	return nil
}
