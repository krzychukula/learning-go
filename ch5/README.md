# Chapterr 5. Functions

## Declaring and Calling Functions

use struct to emulate named and optional parameters

```go
type MyFuncOpts struct {
    FirstName string
    LastName string
    Age int
}

func MyFunc(opts MyFuncOpts) error {
    //
}

func main() {
    MyFunc(MyFuncOpts {
        LastName: "Patel",
        Age: 50
    })

    MyFunc(MyFuncOpts {
        FirstName: "Joe",
        LastName: "Smith",
    })
}
```

### Variadic Intup Parameters and Slices

```go
func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))
    for _, v := range vals {
        out = append(out, base+v)
    }
    return out
}

addTo(3, 1, 2, 3)
addTo(3, []int{1, 2, 3}...)
```

### Multiple Return Values

```go
func divAndRemainder(numerator int, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return numerator / denominator, numerator % denominator, nil
}
func main() {
    result, remainder, err := divAndRemainder(5, 2)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(result, remainder)
}
```

multiple return values are not a tuple or slice. Each one has to be assigned to something separately.

### Ignoring Returned Values

_

### Named Return Values

```go
func divAndRemainder(numerator int, denominator int) (result int, reminder int, err error) {
    if denominator == 0 {
        err := errors.New("cannot divide by zero")
        return result, reminder, err
    }
    result, reminder := numerator / denominator, numerator % denominator, nil
    return result, reminder, err
}
```

Use _ to leave some without name. 
```go
func divAndRemainder(numerator int, denominator int) (_ int, reminder int, _ error) {
    
}
```

```go
func divAndRemainder(numerator int, denominator int) (result int, reminder int, err error) {
    result, reminder := 100, 200 // those will be ignored
    if denominator == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return numerator / denominator, numerator % denominator, nil
}
```

### Blank returns 

100 Go mistakes -> use them with named return values

```go
func divAndRemainder(numerator int, denominator int) (result int, reminder int, err error) {
    // remember that zero values of Named Return Values 
    // will be returned from the function
    // make sure they make sense!
    // result int 0
    //  reminder int 0
    // err error nil
    if denominator == 0 {
        err := errors.New("cannot divide by zero")
        return
    }
    result, reminder := numerator / denominator, numerator % denominator, nil
    return
}
```

## Functions Are Values

calculator.go

### Function Type Declarations

```go
var opMap = map[string]func(int, int) int{}

// with named

type opFuncType func(int, int) int

var opMap = map[string]opFuncType {}

```

### Anonymous Functions

[[anon.go]]

## Closures

Like in JS.

## defer



## Go Is Call by Value