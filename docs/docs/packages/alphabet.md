---
layout: page
title:  "Alphabet"
nav_order: 2
heading_anchors: true
parent: Packages
---

## Alphabet

All biological sequences have a finite alphabet of valid characters.
In Go this translates into the following interface:

```go
type Interface interface {
    // Must be stringable
    fmt.Stringer

    // Contains checks that the given elements are in the Alphabet
    Contains(...string) []bool

    // Length is the number of letters in the Alphabet
    Length() int
}
```

What the above means is that any valid biological alphabet can:

1. `fmt.Stringer`: produce a printable version of the alphabet (in order) as a string
2. `Contains(...string) []bool`: validate whether a given entry is part of the alphabet
3. `Length() int`: specify how many valid characters there are in the alphabet

Not that there does not exist a `Letters() []Letter`, if you need to get at the individual letters use the `String() string` method defined by `fmt.Stringer`.

### hashmap

This version of alphabet uses Go's internal hashmap to provide constant time lookup (`Contains(...string) string`) of potentially valid characters.
`Length()` is the same as `len(...)` on the underlying map.
`String()` is the alphabetized characters (done by sorting after iterating over all map keys)

#### Performance Optimization

**Fun Go fact**: the empty struct `struct{}` does not require any memory to store which is precisely why it is used as the values for our hashmap. We get constant lookup by using the hashmap, but require no additional memory to store values that we will never use.
