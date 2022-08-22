# Grange

Grange is range library to manipulate date and number range

# Install

```bash
# v1
go get github.com/karnott/grange

# v2
go get github.com/karnott/grange/v2
```

# [V2] How to use it (check v1 package for correct struct name)

```go
import github.com/karnott/grange

func main() {
  range1 := grange.NumberRange[int]{1, 10}
  range2 := grange.NumberRange[int]{5, 20}

  rangeIntersection := range.Intersection(range1, range2) 
  // rangeIntersection == [2]NumberRange[int]{5, 10}
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