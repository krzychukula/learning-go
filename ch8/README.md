# Chapter 8. Errors

## How to Handle Errors: The Basics

* returning a value of type `error` as the last return value of a function.

* Convention -> but strong.
* no error? Then return `nil`

```go
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("denominator is 0")
    }
    return numerator / denominator, numerator % denominator, nil
}
```

Error message:
* should not be capitalized
* should not return with punctuation

When returning `error` then other return values should return their zero values.

```go
func main() {
    numerator := 20
    denominator := 3
    remainder, mod, err := calcRemainderAndMod(numerator, denominator)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(remainder, mod)
}
```

error is a built-in interface with a single method:

```go
type error interface {
    Error() string
}
```

* Anything that implements `Error` method (error interface) has `error` type
* `nil` is a zero value for any interface.

* blablabla why not throw exceptions. Not convinced at all!
* blablabla all variables (errors) must be read or explicitly ignored. Not convinced at all!
If you said because designers of the language like it that way it would be more honest. 

* Shorter code MAKES code easier to understand and maintain!
* So I'm not drinking the cool-aid.
* This is compared to the worst exceptions.
* No mention of Either monad!

So much BS about why it's so great. Not sure if I should laugh or cry. 

## Use Strings for Simple Errors

```go
err := errors.New("some error")
fmt.Println(err) // this will call Error() method be default - and print the passed string

fmt.Errorf("%d isn't an even number", i)
// this allows formatting in error messages. 
```

## Sentinel Errors

Dave Cheney

* package variable with a name starting with `Err`
* exception `io.EOF`
* they should be readonly (unless you make a mistake)

`archive/zip`

```go
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
```

`crypto/rsa` has:
* `rsa.ErrMessageTooLong`

Another common: `context.Canceled`

Weird things here. So how do you create those sentinel errors anyway?

## Errors Are Values

```go
type Status int

const (
    InvalidLogin Status = iota + 1
    NotFound
)

type StatusErr struct {
    Status Status
    Message string
}

func (se StatusErr) Error() string() {
    return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
    err := login(uid, pwd)
    if err != nil {
        return nil, StatusErr{
            Status: InvalidLogin,
            Message: fmt.Sprintf("invalid credentials for user %s", uid)
        }
    }
    data, err := getData(file)
    if err != nil {
        return nil, StatusErr{
            Status: NotFound,
            Message: fmt.Sprintf("file %s not found", file)
        }
    }
    return data, nil
}
```

always return `error` type
Do not return uninitialised variable as custom error

```go
func GenerateError(flag bool) error {
    // don't do this
    var genErr StatusErr
    if flag {
        genErr = StatusErr{
            Status: NotFound
        }
    }
    return genErr
}

func main() {
    err := GenerateError(true)
    fmt.Println(err != nil) // true
    
    err = GenerateError(false)
    fmt.Println(err != nil) // true!!!!!!
}
```

err is never `nil` because interface values have to fields:
* interface (is not nil here)
* value (is nil)

How to fix it?

1. Return `nil` explicitly.

```go
func GenerateError(flag bool) error {
    if flag {
        return StatusErr{
            Status: NotFound
        }
    }
    return nil
}

func main() {
    err := GenerateError(true)
    fmt.Println(err != nil) // true
    
    err = GenerateError(false)
    fmt.Println(err != nil) // should be fixed
}
```

2. Or make sure that you use `error` type!

```go
func GenerateError(flag bool) error {
    var genErr error // not StatusErr
    if flag {
        genErr = StatusErr{
            Status: NotFound
        }
    }
    return genErr
}

func main() {
    err := GenerateError(true)
    fmt.Println(err != nil) // true
    
    err = GenerateError(false)
    fmt.Println(err != nil) // should be fixed
}
```

## Wrapping Errors

`fmt.Errorf` with `%w` will wrap an error. 

convention:
* make `%w` last thing in the error
* make error last parameter

`errors.Unwrap` can unwrap errors.

For custom errors:

```go
type StatusErr struct {
    Status Status
    Message string
    Err error
}
func (se StatusErr) Error() string {
    return se.Message
}
func (se StatusErr) Unwrap() error {
    return se.Err
}
```

If you just want a message from the error without wrapping it:
* Use `%v`

```go
err := internalFunction()
if err != nil {
    return fmt.Errorf("internal failure: %v", err)
}
```

