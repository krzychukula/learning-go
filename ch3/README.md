# Chapter 3. Composite Types

## Arrays—Too Rigid to Use Directly

Just don't use them

var x = [...]int{1, 2, 3} // size of the array based on the contents

## Slices

[] makes a slice

var x = []int{1, 2, 3}

vax x []int // zero value of slice is nil

> nil is an identifier that represents the lack of a value
-> Not `null`... Is this more like `undefined`?

Slices are not comparable with `==` or `!=` you can only compare them to nil `x == nil`

`reflect.DeepEqual` if you need to.

### len

len(nil) -> 0

### append

var x []int // nil slice
x = append(x, 10) // you can append to nil slice

x = append(x, 20, 30)

y := []int{4, 5, 6}
x = append(x, y...)

Is x passed as a copy of slice value or a copy of a slice pointer? 
Is the book incorrect here saying that copy of the slice lands in the append? 

### cap - capacity

Go 1.14
1. double the size of the slice if capacity is less than 1_024
2. After that grow the slice by 25%

### make

x := make([]int, 5)// len 5, cap 5
x := make([]int, 5, 10)// len 5, cap 10

### Declaring Your Slice

Most of the time use nil slice, it may stay that way and you will avoid allocations

`var data []int // nil`

For JSON use empty slice:
`var x = []int{}`

If you know capacity but not values then:
`x := make([]int, 0, 20)`

Append always increases len, so beware of:

```go
x := make([]int, 10, 20)
x = append(x, 1, 2, 3)
```
This will give you 10 `0` values at the biginning of the slice!!!

### Slicing slices

