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

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)// 0 ...9 
}
```
can't use `var` only `:=` in `i := 0`

comparison:
1. before the loop body
2. after initialization
3. after loop reaches the end

### The Condition-Only for Statement

```go
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
```


### The Infinite for Statement

```go
i := 1
for {
    fmt.Println("infinity")
}
```

### break and contiune

```go
i := 1
for {
    if i > 1 {
        break
    }
}
```

```go

for i := 1; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")
        continue
    }
    if i%3 == 0 {
        fmt.Println("Fizz")
        continue
    }
    if i%5 == 0 {
        fmt.Println("Buzz")
        continue
    }
    fmt.Println(i)
   
}
```

### The for-range Statement

only for compound types

uniqueNames := map[string]bool{
		"Fred":  true,
		"Bob":   false,
		"Alice": true,
	}
	for k := range uniqueNames {
		// you can't depend on the order when iterating over a map
		fmt.Println(k)
	}

Formatting functions always log maps with keys sorted

```go
for i, r := range "abcąbć" {
    // only use value (rune) in strings!!!
    fmt.Println(i, r, string(r))
}
/*
    0 97 a
    1 98 b
    2 99 c
    3 261 ą
    5 98 b
    6 263 ć
*/
```

Value is a copy. range makes a copy of the value from your slice/map/string. 

### Lebeling Your for Statements


```go
outer:
    for i := 0; i < 10; i++ {
        for i, r := range "abcąbć" {
            if r == "b" {
                continue outer
            }
            fmt.Println(i, r, string(r))
        }
    }
### Choosing the Right for Statement

Prefer `for range` loop

## switch

expression switch statements vs type swich statements?
here only "expression switch statements"



## Blank Switches

```go
for _, word := range words {
    switch wordLen := len(word); {
    case wordLen < 5:
        fmt.Println(word, "is a short word")
    case wordLen > 10:
        fmt.Println(word, "is a long word")
    default:
        fmt.Println(word, "is exactly right len:", wordLen)
    }
}
/*
    a is a short word
    cow is a short word
    smile is exactly right len: 5
    gopher is exactly right len: 6
    octopus is exactly right len: 7
    anthropologist is a long word
*/
```

## Choosing between if and switch



## goto-Yes, goto

