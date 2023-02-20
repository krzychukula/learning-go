# Chapter 9. Modules, Packages, and Imports

## Repositories, Modules, and Packages

- repository - git
- module - the root of a Go library or application (stored in the repository)
- package - modules are built from packages. Organizing blocks of modules. 

module and repository should be 1-1
Put only one module in the repository.

module - has globally unique identifier (like Java)

`github.com/jonbodner/proteus`

## go.mod

`go mod init MODULE_PATH` - creates a new module -> creates a new `go.mod` file

MODULE_PATH -> globally unique name -> module identifier

```go.mod
module github.com/learning-go-book/money

go 1.15

require (
    github.com/learning-go-book/formatter v0.0.0
    github.com/shopspring/decimal v1.2.0
)
```

`replace` -> override location of a dependent module
`exclude` -> prevents some module version from being used
- weird? Why we would even need exclude?!


## Building Packages

### Imports and Exports

Exported from a package
private
_private

### Creating and Accessing a Package

```go
// /math/math.go
package math

func Double(a int) int {
    return a * a
}
```

`package clause` -> `package math`



```go
// /formatter/formatter.go

package print

import "fmt"

func Format(num int) string {
    return fmt.Sprintf("The number is %d", num)
}
```


```go
// /main.go

package main

import (
    "fmt"

    "github.com/learning-go-book/package_example/formatter" // imports a file
    "github.com/learning-go-book/package_example/math"
)

func main() {
    num := math.Double(2)
    output := print.Format(num) // formatter had "print" package!
    fmt.Println(output)
}
```

```
$ go run main.go
// The number is 4
```

Every go file in a directory should have the same package.
`package clause` is important - not directory name
But, try to match them, as it's hard to find without looking at code/documentation.

### Naming Packages

Don't use `util`. 
Better:
* extract.Names
* format.Names

Don't duplicate exceptions:
* sort.Sort
* context.Context

### How to Organize Your Module

If your module is small -> use single package for everything

`/cmd` - directory for binaries

`/pkg` - put most of code into it if you have lot's of configuration files

### Overriding a Package's Name

`math/rand` and `crypto/rand` use the same package name!

```go
import (
    crand "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/rand"
)

func seedRand() *rand.Rand {
    var b [8]byte
    _, err := crand.Read(b[:]) // crand
    if err != nil {
        panic("cannot seed with cryptographic random number generator")
    }
    // rand
    r := rand.New(
        rand.NewSource(
            int64(
                binary.LittleEndian.Uint64(
                    b[:]
                )
            )
        )
    )
    return r
}
```
You can import with `.`, but don't use it.

```go
import (
    . "crypto/rand"
    "math/rand"
)

// Now Read can be called without crand prefix
```

BEWARE: Go let's you shadow a package with a variable.

### Package Comments and godoc

godoc - convention (btw. conventions are usually not explicit as what Go is supposedly about)

* Put comment directly abot what you document
* // NameOfTheItem
* //
* // (use blank comments for multiple lines)
* // preformated by indenting lines
* //    some list
* //    item
* //    item
* //    item

`go doc PACKAGE_NAME` - for viewing documentation of a package

`go doc PACKAGE_NAME.IDENTIFIER_NAME`

`go doc fmt.Println` for documentation of an identifier

Document at least all exported identifiers

### The `internal` Package

`internal` package exported identifiers are accessible to:
* direct parent package of `internal`
* sibling packages of `internal`

### The `init` Function: Avoid if Possible

Don't use it. 
But it's still there for:
* database drivers
* image formats

If you use `init` to set up anything package level - them make it **immutable**

### Circular Dependencies

**No** circular dependencies - directly or indirectly

If you get a compiler error then deal with it.
Join packages if you can, or extract the shared thing.

### Gracefully Renaming and Reorganizing Your API

Try to not remove any public API (exported identifiers)

For types use `alias`

```go
type Foo struct {
    x int
    S string
}

func (f Foo) Hell() string {
    return "hello"
}

func (f Foo) goodbye() string {
    return "goodbye"
}

// alias

type Bar = Foo

// You can use Bar the same way as Foo!

func MakeBar() Bar {
    bar := Bar {
        x: 20,
        S: "Hello"
    }
    var f Foo = bar
    fmt.Println(f.Hello())
    return bar
}
```

You can alias even from another package - but you won't have access to unexported methods and fields.
- You can still create your own versions and call the other package under the hood

Can't have alternatme names:
* package-level variable
* field in a struct (exported name?)






