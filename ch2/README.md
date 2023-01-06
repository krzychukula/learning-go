# Chapter 2. Primitive Types and Declarations

## Built-in Types
variables have default zero values

100_000 is a valid number in Go

6.03e23

'a' is a rune literal

'\141' 8-bit ocal number

'\x61' 8-bi hexadecimal number 

0o777 is used for rwxrwxrwx 

"use double quotes"
`
another string
`

literals in Go are untyped

var flag bool // false
var isAwesome = true

Always use float64 

Floats are the same as in JavaScript

real() and imag() are built-in functions for complext numbers.
This reminds me of old PHL which used a lot of built-in functions.

Imaginary numbers should be ignored in Go. Use some other language if you need math

Strings are immutable
rune = int32
byte = uint8

Go doesn't allow automatic type promotion

```go
var x int = 10
var y float64 = 30.2
var z float64 = float64(x) + y
var d int = x + int(y)
fmt.Println(z, d)
// 40.2 40
```

Go doesn't have truthy values.

Within a function you can use `:=`

`var x = 10` is the same as `x := 10`

```
var x, y 10, "hello"
x, y := 10, "hello"
```

`:=` will allow you to assign values to existing variables.
IMHO. that's not a good thing.

When you want variable with zero value use format:
`var x int`

If you want to specify type then use `var x int`
`var x byte = 20`

Avoid declaring variables outside of functions. 

## Using const

Can only be values that compiler can figure out at compile time!
Const needs to be a literals or expressions based on literals.

No const at runtime!

## Typed and Untyped Constants

Mostly leave constants untyped for convenience. 

## Unused Variables

Compiles won't work if you have a variable that is never read inside of a function. 
* Won't complain about package-level variables.
* Won't complain about unused assignments.
* Unused constants are also fine.

## Naming Variables and Constants

variablesUseCamelCase

Go variable names are shorter depending on the scope of usage.
Within function it is a single letter.

Using `f` to represent float -> woot

Using shorter variables to keep the code shorter is maybe a valid argument, but not in Golang.
This language is built around not having a short code. 

At least package variables are usually named properly.

