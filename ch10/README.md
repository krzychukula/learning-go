# Chapter 10. Concurrency in Go

Concurrency - breaking up single processes into independent components and specifying how those components safely share data.

Communicating Sequential Processes (CSP)

## When to Use Concurrency

-> It's WAY harder than advertised. Especially if people are moving into Mutexes (bottlenecks)

Amdahl's Law - cost of synchronization!!! Bottlenecks again

> Use concurrency when you want to combine data from multiple operations that can operate independently.

There is a cost of concurrency -> This is why it's mostly used for I/O!
Write a sequential program first.

## Goroutines

process - run by OS
thread - run by OS

Go doesn't use OS threads. Creates its own.

You can invoke function with `go`, but any return argument are ignored.

* NICE: you can invoke ANY function as a goroutine. No async/await (blue/red functions) needed. 
* BUT: because Go is explicit it still needs a wrapper because of all the admin work you have to do.
* So: JavaScript wins?! 

```go
func process(val int) {
    // do something
}

func runThingConcurrently(in <-chan int, out chan<- int) {
    go func() {
        for val := range in {
            result := process(val)
            out <- result
        }
    }()
}
```

-> It seems a lot of good practices are skipped here.
-> If I remember correctly you always need to think about exit of a coroutine.
-> Maybe it's later in the chapter?

## Channels

`ch := make(chan int)`

Like maps, channels are reference types.

zero value is nil.

### Reading, Writing, and Buffering

-> From 100 Go Mistakes - don't use buffering
-> And changing from no-buffering to buffering can expose/introduce bugs to your existing code

`a := <-ch` // reads from ch

`ch <- b` // write to ch

Arrow on the left - reads
Arror on the right - writes

Each value can be read only once (unless it's a weird broadcast)

Type of channel when passing to a function:
* `(ch <-chan int)` Function only reads from the channel
* `(ch chan<- int)` Function only writes to a channel

If you need buffer:
`ch := make(chan int, 10)`

Unbuffered channel has `len` and `cap` of `0`.

### for-range and Channels

```go
for v := range ch {
    fmt.Println(v)
}
```

### Closing a Channel

`close(ch)` - built-in function. Of course.

Just like with maps you can check if channel is closed when getting values:

`v, ok := <- ch`

ok: true - channel is open

Only goroutine that writes to a channel should `close` it. 

-> I won't be fooled. Golang in practice depends a lot on a shared mutable state as any other language.

### How Channels Behave

`nil` channel on read or write will hang forever and will panic on close. 
-> and the funniest thing is that it's useful in practice.

Because channels API is too simplistic and they panic so often you need `sync.WaitGroup` to make them practical.

## `select`

```go
select {
    case v := <- ch:
        fmt.Println(v)
    case v := <- ch2:
        fmt.Println(v)
    case ch3 <- x:
        fmt.Println("wrote", x)
    case <- ch4:
        fmt.Println("got value from ch4, but ignored it")
}
```

* If multiple cases can run - select picks at RANDOM. 

Deadlock:

```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        v := 1
        ch1 <- v
        v2 := <-ch2
        fmt.Println(v, v2)
        // writes to ch1 and blocks
    }()

    v := 2
    ch2 <- v // writes to ch2 and blocks
    v2 := <-ch1
    fmt.Println(v, v2)
    // fatal error: deadlock
}
```

Fix:

```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        v := 1
        ch1 <- v
        v2 := <-ch2
        fmt.Println(v, v2)
    }()

    v := 2
    var v2 int
    select {
        case ch2 <- v:
        case v2 = <-ch1:
    }
    fmt.Println(v, v2)
    // 2 1
}
```

**for-select loop**

```go
for {
    select {
        case <- done:
            return
        case v = <-ch:
            fmt.Println(v)
    }
}
```

When exactly is called `default` in the `select`???

```go
select {
case v = <-ch:
    fmt.Println("read from ch:", v)
default:
    fmt.Println("no value written to ch") // how many times is this called?
}
```

// how many times is this called?
The `default` will be called every time there is nothing!
So don't use it in *for-select* as it will be invoked all the time.

## Concurrency Practices and Patterns


