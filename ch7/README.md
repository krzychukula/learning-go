# Chapter 7. Types, Methods, and Interfaces

## Types in Go

```go
type Person struct {
    FirstName string
    LastName string
    Age int
}

type Score int
type Converter func(string)Score
type TeamScores map[string]Score

```

## Methods

no overloading
methods only for types you control (in the same package as type)

```go
type Person struct {
    FirstName string
    LastName string
    Age int
}
func (p Person) String() string {
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

p := Person {
    FirstName: "Fred",
    LastName: "Fredson",
    Age: 52,
}
output := p.String()
```

### Pointer Receivers and Value Receivers

1. If your method modifies the receiver - then it must be a pointer `*`
2. If your method needs to handle `nil` receiver - then it must be a pointer `*` reveiver
3. If your method doesn't modify the reveiver - you can use a value receiver.

If type has any pointer receivers then you use pointer for all receivers (common practice).

**From 100 Go Mistakes:**
> A receiver MUST be a pointer:
> * If method needs to mutate the reveiver. 
> * If the receiver is a slice and you need to append to it. 
> * If the receiver contains a field that cannot be copied (a type part of the `sync` package).
> 
> A receiver SHOULD be a pointer:
> * If the receiver is a large object (but benchmark to check).
> 
> A receiver MUST be a value:
> * If we have to enforce a receiver's immutability.
> * If the receiver is a `map`, `function`, or `channel` (or you will get compilation error).
> 
> A receiver SHOULD be a value:
> * If the receiver is a slice that doesn't have to be mutated.
> * If the receiver is naturally a value type without mutable fields. Examples: small `array` or `struct` such as `time.Time`.
> * If the receiver is a basic type such as `int`, `float64`, or `string`.


```go
type Counter struct {
    total int
    lastUpdated time.Time
}

func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

var c Counter
fmt.Println(c.String())

c.Increment()
fmt.Println(c.String())
```

Go converts `c.Increment()` to `(&c).Increment()` for you for value types.

Type `*Counter` method set:
* pointer receiver methods
* value receiver methods

Type `Counter` method set;
* ONLY value reveiver methods

### Code Your Methods for `nil` instances

Go will try to invoke methods on `nil`.
1. Will `panic` on value reveiver
2. But, pointer reveiver methods can support it (but needs to be handle well in code)

```go
type IntTree struct {
    val int
    left, rigth *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
    if it == nil {
        return &IntTree{val: val}
    }
    if val < it.val {
        it.left = it.left.Insert(val)
    } else if val > it.val {
        it.rigth = it.right.Insert(val)
    }
    return it
}

func (it *IntTree) Contains(val int) bool {
    switch {
        case it == nil:
            return false
        case val < it.val:
            return it.left.Contains(val)
        case val > it.val:
            return it.right.Contains(val)
        default:
            return true
    }
}

func main() {
    var it *IntTree
    it = it.Insert(5)
    it = it.Insert(3)
    it = it.Insert(10)
    it = it.Insert(2)
    fmt.Println(it.Contains(2)) // true
    fmt.Println(it.Contains(12)) // false
}

```

Methods as functions can't change the copy of the pointer they received.
You can't get `nil` and make the original pointer non-nil.
If your method can'd do anything useful with `nil` then return error!

### Methods Are Functions Too

```go
type Adder struct {
    start int
}

func (a Adder) AddTo(val int) int {
    return a.start + val
}

myAdder := Adder{start:10}
fmt.Println(myAdder.AddTo(5)) // 15

f1 := myAdder.AddTo
// method value
fmt.Println(f1(10)) // 20

f2 := Adder.AddTo
// Method expression
fmt.Println(f2(myAdder, 15)) // 25
```

### Functions Versus Methods

Function:
* When your logic only depends on the input parameters

Method:
* Logic depends on values configured at startup
* Logic depends on vales that change while the program is running
* => Those values should be stored in a struct and this logic should be put into methods

### Type Declarations Aren't Inheritance

Go doesn't have typical inheritance.

Types in go are nominal types. You can't assing one to the other without conversion.

```go
type HighScore Score
type Employee Person

// assigning untyped constants is valid
var i int = 300
var s Score = 100
var hs HighScore = 200

hs = s // error
s = i // error

s = Score(i)
hs = HighScore(s)
```
Operators from built-in types should still work.
So like `s + s`?

### Types Are Executable Documentation

### `iota` Is for Enumeration—Sometimes

```go
type MailCategory int

const (
    Uncategorized MailCategory = iota // 0
    Personal // 1
    Spam // 2
    Social // 3
    Advertisements // 4
)
```

`iota` is only for internal constants. Don't use it to repserent anything from elsewhere.

If someone adds a value in the middle it will be renumbered. Potentially breaking things elsewhere. 

`iota` is super **FRAGILE**.
Anyone can add another value of your type. `iota` is **NOT** an enum. 
 IMHO. better to not use it at all.

```go

type BitField int

const (
    Field1 BitField = 1 << iota // 1
    Field2 // 2
    Field3 // 4
    Field4 // 8
)
```

## Use Embedding for Composition

```go
type Empleoyee struct {
    Name string
    ID string
}

func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
    // NO name assigned to type Employee here
    // embedded field
    Employee
    Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
    // some logic
}

m := Manager{
    Employee: Employee{
        Name: "Bob Bobson",
        ID: "123"
    }
    Reports: []Employee{}
}
fmt.Println(m.ID)//123
fmt.Println(m.Description())// Bob Bobson (123)

```

> You can embed any type within a struct, not just another struct.
> This promotes the methods on the embedded type to the containing struct.

Outer struct can **shadow** methods and fields from the inner type.

```go
type Inner struct {
    X int
}

type Outer struct {
    Inner
    X int
}

o := Outer {
    Inner: Inner{
        X: 10
    },
    X: 200
}

o.X // 200
o.Inner.X // 10
```

## Embedding is Not Inheritance

Methods of the embedded type won't call methods of the containing type (even if they share the same name).

```go

type Inner struct {
    A int
}

func (i Inner) IntPrinter(val int) string {
    return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
    return i.IntPrinter(i.A * 2)
}

type Outer struct {
    Inner
    S string
}

func (o Outer) IntPrinter(val int) string {
    return fmt.Sprintf("Outer: %d", val)
}

func main() {
    o := Outer{
        Inner: Inner{
            A: 10,
        },
        S: "Hello",
    }

    fmt.Println(o.Double()) // Inner: 20
}
```

> The methods on an embedded field do count toward the *method set* of the containg struct. 
> This means they can make the containing struct implement na interface.

## A Quick Lesson on Interfaces





