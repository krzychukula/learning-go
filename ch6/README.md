# Chapter 6. Pointers

## A Quick Pointer Primer

Pointer is a variable that holds the location in memory where a value is stored.

```go
var x int32 = 10
var y bool = true

pointerX := &x
pointerY := &y

var pointerZ *string
```

Zero value for a pointer is `nil`. (maps, slices, functions)

Reminder: You can shadow `nil`!!!! 

&&&&&&&&&&&&&&&&
The `&` is the *address* operator. Gets memory address of a value:

```go
x := "hello"
pointerToX := &x
```
&&&&&&&&&&&&&&&&



***************
The `*` is a *indirection* operator. Returns the pointed-to value of a pointer. 

`*` is dereferencing

Dereferencing a `nil` pointer will cause runtime `panic`

```go
x := 10
pointerToX := 10
fmt.Println(pointerToX) // prints memory address
fmt.Println(*pointerToX) // prints 10

z := 5 + *pointerToX
fmt.Println(z) // 15
```
****************


**Pointer type** - type that represents a pointer. Can be based on any other type.
Pointer type has `*` before the type name. `*int` or `*string`

New:
```go
var x = new(int)
fmt.Println(x == nil)//false
fmt.Println(*x) // 0 - int type zero value put referenced by x pointer
```

Structs:
x := &Foo{}

For primitive literals (numbers, booleans and strings) you need to create a variable to get a pointer.

```go
func stringp(s string) *string {
  return &s
}
```

## Don't Fear the Pointers

So using pointers is a lot like using references to objects in JS?

## Pointers Indicate Mutable Parameters

You can't set `nil` pointer passe to a function to anything inside this function.

```go
func failedUpdate(px *int) {
    x2 := 20
    px = &x2
}

func update(px *int){
    *px = 10
}

func main () {
    x := 10
    failedUpdate(&x)
    fmt.Println(x) // still 10
    
    update(x)
    fmt.Println(x)// updated to 20
}
```

## Pointers Are a Last Resort

```go
f := struct {
    Name string `json:"name"`
    Age int `json:"age"`
}{}
err := json.Unmarshall([]byte('{"name":"Bob", "age":30}'), &f)
```

## Pointer Passing Performance

Passing a pointer into a function is 1 nanosecond.

For 10 megabytes of data - passing it into a function takes 1 milisecond.

1 milisecond = 1000000 nanoseconds

### Returned valus

Until around 1 megabyte it's faster to return value than a pointer!

## Ther Zero Value Versus No Value

Do not use pointer `nil` values as an api of exposing `undefined`

Use `comma ok` (value type and a boolean) to represent unassigned value like in maps.

## The Difference Between Maps and Slices

Don't use maps as function input parameters.
Don't return maps from functions.

Use `struct` for it.

If you pass a slice to a function then it can be changed.
But `append` won't change the original slice even if it had capacity!!!!!

Because slice is implemented as a struct with 3 fields:
* int field for length
* int field for capacity
* pointer to a block of memory

When you pass a slice to function then Go makes a copy of all 3 fields!
* changing existing value - works - copy of a pointer to a block of memory works

But, changing *length* or *capacity* will change them only in the COPY!
It won't affect them in the original slice.

If you `append` values to an array with existing capacity:
* memory will be updated in the array
* BUT - it won't be visible in the caller as it's length won't be updated

Most functions shouldn't modify passed slices in any way.
If they do they need to document it.

## Slices as Buffers

```go

file, err := os.Open(fileName)
if err != nil {
    return err
}
defer file.Close()

data := make([]byte, 100)
for {
    count, err := file.Read(data)
    if err != nil {
        return err
    }
    if count == 0 {
        return nil
    }
    process(data[:count])
}
```

## Reducing the Garbage Collector's Workload

To store something on **stack**:
* You have to know how big it is at compile time
  * primitive values, arrays, structs
  * This is why size is part of the array type
  * pointer size is also known at compile time
* Pointer value to be on stack:
  * local variable with known size at compile time
  * pointer cannot be returned from the function
  * If pointer is passed to a function then compiler still has to make sure the previous true rules are met.
- If data can't be stored on stack then it **escapes** the stack

Each garbage collection takes less than 500 microseconds

mechanical sympathy
2011 Martin Thompson began applying it to software developement

Use pointers sparingly - it's less work for garbage collector

Slices of structs or primitive types have their data lined up sequentially in RAM!


