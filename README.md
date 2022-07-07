# Grange

Grange is range library to manipulate date and number range

# Install

```bash
go get github.com/karnott/grange
```

# How to use it

```go
import github.com/karnott/grange

func main() {
  range1 := grange.RangeNumber[int]{1, 10}
  range2 := grange.RangeNumber[int]{5, 20}

  rangeIntersection := range.Intersection(range1, range2) 
  // rangeIntersection == [2]RangeNumber[int]{5, 10}
}
```
# Compatibility

You need to use Golang 1.18 or newer version

# TODO

- [x] range intersection
- [ ] range union
- [ ] range difference
- [ ] range contain
- [ ] range before
- [ ] range after