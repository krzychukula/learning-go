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


### Keep Your APIs Concurrency-Free

Do not expose **channels** or **Mutexes** (bottlenecks) in API. (don't export them)

### Goroutines, for Loops, and Varying Variables

Passing index or value from for loop to goroutine will lead to the same problem as in JS.

1. Shadow the varible

```go
for _, v := range a {
    v := v
    go func() {
        ch <- v * 2
    }()
}
```

2. Pass the value as parameter to the goroutine

```go
for _, v := range a {
    go func(val int) {
        ch <- val * 2
    }(v)
}
```

### Always Clean Up Your Goroutines

**goroutine leak** - runtime will give it time to run even if it does nothing.

-> But how do you clean them up???

### The Done Channel Pattern

```go
func searchData(s string, searchers []func(string) []string) []string {
    done := make(chan struct{})
    result := make(chan []string)
    for _, searcher := range searchers {
        go func(searcher func(string) []string) {
            select {
            case result <- searcher(s):
            case <-done:
            }
        }(searcher)
    }
    r := <-result
    close(done)
    return r
}
```
-> after one search finishes close(done) will close all other goroutines.

### Using a Cancel Funcsion to Terminate a Goroutine

```go
func countTo(max int) (<-chan int, func()) {
    done := make(chan struct{})
    ch := make(chan int)
    cancel := func() {
        close(done)
    }

    go func(searcher func(string) []string) {
        for i := 0; i < max; i++ {
            select {
            case <-done:
            case ch<-i:
            }
        }
        close(ch)
    }()
    return ch, cancel
}

func main() {
    ch, cancel := countTo(10)
    for i := range ch {
        if i > 5 {
            break
        }
        fmt.Println(i)
    }
    cancel() // this prevents goroutine leak
}
```

### When to Use Buffered and Unbuffered Channels

Buffered
- if you know how many goroutines you launched
- you want to limit the number of goroutines you will launch
- want to limit the amount of work that can queue up

Buffered channel - gathering data from known number of goroutines

```go
func processChannel(ch chan int) []int {
    const conc = 10
    results := make(chan int, conc)
    for i := 0; i < conc; i++ {
        go func() {
            v := <- ch
            results <- process(v)
        }()
    }
    var out []int
    for i := 0; i < conc; i++ {
        out = append(out, <-results)
    }
    return out
}
```

### Backpressure

> systems perform better overall when their comonents limit the amount of work they are willing to perform.

```go
type PressureGauge struct {
    ch chan struct{}
}

func New(limit int) *PressureGauge {
    ch := make(chan struct{}, limit)
    for i := 0; i < limit; i++ {
        ch <- struct{} // sends to channel initial number of events
    }
    return &PressureGauge{
        ch: ch
    }
}

func (pg *PressureGauge) Process(f func()) error {
    select {
    case <-pg.ch: // receives from channel
        f()
        pg.ch <- struct{}{} // sends to channel after it's done
        return nil
    default:
        return errors.New("no more capacity")
    }
}

func doThingThatShouldBeLimited() string {
    time.Sleep(2 * time.Second)
    return "done"
}

func main() {
    pg := New(10)
    http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
        err := pg.Process(func() {
            w.Write([]byte(doThingThatShouldBeLimited()))
        })
        if err != nil {
            w.WriteHeader(http.StatusTooManyRequests)
            w.Write([]byte("Too many requests"))
        }
    })
    http.ListenAndServe(":8080", nil)
}
```

### Turning Off a case in select

Set closed channel to `nil` to not waste time on it in `select`

```go
for {
    select {
    case <-done:
        return
    case v, ok := <-in:
        if !ok {
            in = nil // this case will never match again
            continue
        }
        fmt.Println(v)
    case v, ok := <-in2:
        if !ok {
            in2 = nil // this case will never match again
            continue
        }
        fmt.Println(v)
    }
}
```

### How to Time Out Code

-> 100 Go Mistakes #76: time.After and memory leaks
* time.After creates a channel that won't be garbage collected until it finishes.
* So don't use it in a loop when you would create a lot of those channels
Better use `context.WithTimeout`

```go
func timeLimit() (int, err) {
    var result int
    var err error
    done := make(chan struct{})
    go func() {
        result, err = doSomething()
        close(done)
    }()
    select {
    case <- done:
        return result, err
    case <- time.After(2 * time.Second):
        return 0, errors.New("work timed out")
    }
}
```


### Using WaitGroups

Waiting for one goroutine -> you can use done channel
Waiting for more `WaitGroup`

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go func() {
        defer wg.Done()
        doThing()
    }()
    go func() {
        defer wg.Done()
        doThing()
    }()
    go func() {
        defer wg.Done()
        doThing()
    }()
    wg.Wait()
}
```

```go
func processAndGather(in <-chan int, processor func(int) int, num int) []int {
    out := make(chan int, num)
    var wg sync.WaitGroup
    wg.Add(num)
    for i:=0; i<num; i++ {
        go func() {
            defer wg.Done()
            for v := range in {
                out <- processor(v)
            }
        }()
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    var result []int
    for v := range out {
        result = append(result, v)
    }
    return result
}
```

### Running Code Exactly Once

`sync.Once` `once.Do()`

```go
type SlowComplicatedParser interface {
    Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
    once.Do(func() {
        parser = initParser()
    })
    return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
    // setup and loading that is slow
}
```

### Putting Our Concurrent Tools Together

```go
func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
    ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
    defer cancel()

    p := processor{
        outA: make(chan AOut, 1),
        outB: make(chan BOut, 1),
        inC: make(chan CIn, 1),
        outC: make(chan COut, 1),
        errs: make(chan error, 2),
    }
    p.launch(ctx, data)
    inputC, err := p.waitForAB(ctx)
    if err != nil {
        return COut{}, err
    }
    p.inC <- inputC
    out, err := p.waitForC(ctx)
    return out, err
}

type processor struct {
    outA chan AOut
    outB chan BOut
    outC chan COut
    inC chan CIn
    errs chan error
}

func (p *processor) launch(ctx context.Context, data Input) {
    go func() {
        aOut, err := getResultA(ctx, data.A)
        if err != nil {
            p.errs <- err
            return
        }
        p.outA <- aOut
    }()
    go func() {
        bOut, err := getResultB(ctx, data.B)
        if err != nil {
            p.errs <- err
            return
        }
        p.outB <- bOut
    }()
    go func() {
        select {
        case <-ctx.Done():
            return
        case inputC := <-p.inC:
            cOut, err := getResult(ctx, inputC)
            if err != nil {
                p.errs <- err
                return
            }
            p.outC <- cOut
        }
    }()
    // -> Don't we need to wait for those goroutines somewhere outside? 
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
    var intupC CIn
    count := 0
    for count < 2 {
        select {
        case a := <-p.outA:
            inputC.A = a
            count++
        case b := <-p.outB:
            inputC.B = b
            count++
        case err := <-p.errs:
            return CIn{}, err
        case <-ctx.Done():
            return CIn{}, ctx.Err()
        }
    }
    return inputC, nil
}

// This code is so bad...

func (p *processor) waitForAB(ctx context.Context) (COut, error) {
    select {
    case outC := <-p.outC:
        return outC, nil
    case err := <-p.errs:
        return COut{}, err
    case <-ctx.Done():
        return COut{}, ctx.Err()
    }
}

```

Try implement it in another language and see how hard it is. 
This is super hard already!
Also. Using something like promises in JS would make this a lot simpler. The only complication is cancellation.

## When to Use Mutexes Instead of Channels

```go

type BottleneckScoreboardManager struct {
    bottleneck sync.RWMutex
    scoreboard map[strin]int
}

func New() *BottleneckScoreboardManager {
    return &BottleneckScoreboardManager{
        scoreboard: map[string]int{}
    }
}

func (bsm *BottleneckScoreboardManager) Update(name string, val int) {
    bsm.bottleneck.Lock()
    defer bsm.bottleneck.Unlock()
    bsm.scoreboard[name] = val
}

func (bsm *BottleneckScoreboardManager) Read(name string) (int, bool) {
    bsm.bottleneck.RLock()
    defer bsm.bottleneck.RUnlock()
    val, ok = bsm.scoreboard[name]
    return val, ok
}
```

> Concurrency in Go by Katherine Cox-Buday
> * If you are coordinating goroutines or tracking a value as it is transformed by a series of goroutines, use **channels**.
> * If you are sharing access to a field in a struct, use **mutexes**.
> * If you discover a critical performance issue when using channels, and you cannot find any other way to fix the issue, modify your code to use a **mutex**.

You can send funcions over a channel. Nice!

Locks in Go are NOT reentrant:
If you call `Lock()` twice then it will deadlock!

Do not copy:
* sync.WaitGroup
* sync.Once
* mutex

`sync.Map`
* shared map: keys inserted once and read many times
* goroutines do not access each others keys

## Atomics—You Probably Don't Need These

`sync/atomic`

Not for normal people to use modern CPUs atomic variable operations:
* swap
* load
* store
* compare
* swap (CAS)
This is for singler register values. 

## Learn More

Concurrency in Go -> Katherine Cox-Buday

