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








