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


## Working with Modules

### Importing Third-Party Code

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/learning-go-book/formatter"
    "github.com/shopspring/decimal"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Need two parameters: amound and percent")
        os.Exit(1)
    }
    amount, err := decimal.NewFromString(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    percent, err := decimal.NewFromString(os.Args[2])
    if err != nil {
        log.Fatal(err)
    }
    percent = percent.Div(decimal.NewFromInt(100))
    total := amount.Add(amount.Mul(percent)).Round(2)
    fmt.Println(formatter.Space(80, os.Args[1], os.Args[2], total.StringFixed(2)))
}
```

`go build` will download packages and put them in `go.mod` file.

*******
**But when I tried it I got an error**
```
ch9 % go build
main.go:8:5: no required module provides package github.com/learning-go-book/formatter; to add it:
        go get github.com/learning-go-book/formatter
main.go:9:5: no required module provides package github.com/shopspring/decimal; to add it:
        go get github.com/shopspring/decimal
```
****

Something to try:
`go get ./...`

It works! :) 


```go
// go.mod
module github.com/krzychukula/learning-go/ch9

go 1.20

```

Will create a `go.sum` as well.

```
$ ./money 99.99 7.25
99.99  7.25  107.24
```

### Working with Versions

`go list -m -versions github.com/learning-go-book/simpletax`

Go uses SemVer

### Minimal Version Selection

Minimum version that satisfies all requirements.

If new version is not forward compatible, then you need to ask library authors to fix it. 

### Updating to Compatible Versions

`-u=patch`:

`go get -u=patch github.com/learning-go-book/simpletax`

Update to the newest `-u`

`go get -u github.com/learning-go-book/simpletax`

### Updating to Incompatible Versions

SemVer - breaking change:
* The major version of the module must be incremented
* For all major versions besides 0 and 1, the path to the module must end with `vN` where N is the major version.

Path changes because incompatible versions are DIFFERENT packages

`"github.com/learning-go-book/simpletax/v2"`

To remove old versions form `go.mod` and `go.sum` files use: 

`go mod tidy`


### Vendoring

* just copy the exact thing

Enable -> `go mod vendor`
Will create `/vendor` directory

After that you need to run `go mod vendor` after **any** change to dependencies.

-> So vendoring is like a flag you set for a module.

### pkg.go.dev

-> documentation of open-source packages

