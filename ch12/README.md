# Chapter 12. The Context

## What Is the Context?

`ctx := context.Background()` or `context.TODO()`

```go
func Middleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctx := req.Context()
        // wrap context with stuff
        req.WithContext(ctx)
        handler.ServeHTTP(rw, req)
    })
}

func hander(rw http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    err := req.ParseForm()
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        rw.Write([]byte(err.Error()))
        return
    }

    data := req.FormValue("data")
    // pass context to the logic
    result, err := logic(ctx, data)
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        rw.Write([]byte(err.Error()))
        return
    }
    rw.Write([]byte(result))
}
```

And client
```go
type ServiceCaller struct {
    client *http.Client
}

func (sc ServiceClient) callAnotherService(ctx context.Context, data string) (string, error) {
    req, err := http.NewRequest(
                        http.MethodGet,
                        "http://exmaple.com?data=" + data,
                        nil)
    if err != nil {
        return "", err
    }

    req = req.WithContext(ctx)

    resp, err := sc.client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
    }
    // do stuff
    id, err := processResponse(resp.Body)
    return id, err
}
```

## Cancellation

```go
func slowServer() * httptest.Server {
    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(2 * time.Second)
        w.Write([]byte("Slow response"))
    }))
    return s
}

func fastServer() * httptest.Server {
    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Query().Get("error") == "true" {
            w.Write([]byte("error"))
            return
        }
        w.Write([]byte("Slow response"))
    }))
    return s
}
```

```go
//client.go

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        err := callServer(ctx, "slow", slowURL)
        if eff != nil {
            cancel()
        }
    }()
    go func() {
        defer wg.Done()
        err := callServer(ctx, "fast", fastURL+"?error="+errVal)
        if eff != nil {
            cancel()
        }
    }()
    wg.Wait()
    fmt.Println("done with both")
}
```

- always call `cancel`
- It's ok to call it multiple times

## Timers

* context.WithTimeout
* context.WithDeadline

`time, ok := ctx.Deadline()` time when it will cancel if set (ok?)

## Handling Context Cancellation in Your Own Code

