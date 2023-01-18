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

x := []int{1, 2, 3, 4}
// y := x[startIndexInclusive|0:endIndexExclusive|endOfSlice]
y := x[:2] // [1, 2]
d := x[1:3] // [2, 3]

#### Slices share storage sometimes

Beware of memory leaks!
Beware of changing underlying data when two slices share data!

##### append makes it even more broken

x := []int{1, 2, 3, 4}
y := x[:2] // [1, 2]
fmt.Println(cap(x), cap(y))// 4 4
y = append(y, 30)
fmt.Println("x:", x) // [1, 2, 30, 4] !!!!!!!
fmt.Println("y:", y) // [1, 2, 30]

##### Full slice expression
// y := x[startIndexInclusive|0:endIndexExclusive|endOfSlice:lastPositionInOriginalSliceCapacityAvailable|fullCapacity]
y := x[:2:2]
z := x[2:4:4]

** THIS API is broken and should NOT be used **

### You can convert array to slice by slicing it

slice := array[:]

### copy

numberOfElementsCopied = copy(destination, source)

num = copy(x[:3], x[1:])


## Strings and Runes and Bytes

Just don't use slicing on strings.
1. Memory leaks
2. It won't know anything about UTF-8 and will break any characters greater than 1 byte

var s string = "Hello, ę"
var bs []byte = []byte(s)
var rs []rune = []rune(s)

Use `string` or `unicode/utf8` if you need to do something with the strings.

## Maps
