package grange

// return [start, end] intersection range between 2 ranges
// return nil if no intersection
func intersection[K Number, L RangeNumber[K]](r L, r1 L) *RangeNumber[K] {
	start1 := r[0]
	end1 := r[1]
	start2 := r1[0]
	end2 := r1[1]

	if end2 < start1 || end1 < start2 {
		return nil
	}

	if start1 < start2 {
		if end1 > end2 {
			return &RangeNumber[K]{
				start2,
				end2,
			}
		}
		if end1 <= end2 {
			return &RangeNumber[K]{
				start2,
				end1,
			}
		}
	}
	if start2 <= start1 {
		if end2 > end1 {
			return &RangeNumber[K]{
				start1,
				end1,
			}
		}

		if end2 <= end1 {
			return &RangeNumber[K]{
				start1,
				end2,
			}
		}
	}

	return nil
}
