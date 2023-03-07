# Chapter 15. Welcome to the Future: Generics in Go

## Generics Reduce Repetitive Code and Increase Type Safety

## Introducing Generics in Go

## Generic Functions Abstract Algorithms

```go
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
    r := make([]T2, len(s))
    for i, v := range s {
        r[i] = f(v)
    }
    return r
}

words := []string{"ast", "ast", "ost", "stst"}

lengths := Map(words, func(s string) int {
    return len(s)
})

```

## Generics and Interfaces

## Use Type Terms to Specify Operators

```go
type BuiltInOrdered interface {
    string | int | int8 | int16 | int32 | int64 | float32 | float64 |
        uint | uint8 | uint16 | uint32 | uint64 | uintptr
}
```

remember to use `~int` 

## Type Inference and Generics

## Type Elements Limit Constants

## Combining Generic Functions with Generic Data Structures

## Things That Are Left Out

## Idiomatic Go and Generics

Use `any` instead of `interface{}`

## Further Features Unlocked

---

-> overall I don't feel as informed about Generics as I would hope, but that's understandable given when it was released. 



