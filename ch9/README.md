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
