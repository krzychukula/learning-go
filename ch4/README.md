# Chapter 4: Blocks, Shadows, and Control Structures

## Blocks

file block (imports)
each {} creates a block

### Shadowing Variables

You can shadow with `:=`

You can shadow packages like `fmt`

### Detecting Shadowed Variables

go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

// REALLY???!!!
true := 10
fmt.Println(true) // true

## if

```go
if n == 0 {

} else if n > 5 {

} else {

}
```

```go
if n := rand.Intn(10); n == 0 {
    fmt.Println(n)
} else if n > 5 {
    fmt.Println(n) // n is accessible here
} else {
    fmt.Println(n) // n is accessible here as well
}
// no "n" here
```

## for, Four Ways

### The Complete for Statement

### The Condition-Only for Statement

### The Infinite for Statement

### break and contiune

### The for-range Statement 

### Lebeling Your for Statements

### Choosing the Right for Statement

## switch

## Blank Switches

## Choosing between if and switch

## goto-Yes, goto

