# Chapter 11. The Standard Library

## io and Friends

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

-> This is one of the reasons I think that Golang is a good fit if your level of work is at the level of **bytes**.

```go
func countLetter(r io.Reader) (map[string]int, error) {
    buf := make([]byte, 2048)
    out := map[string]int{}
    for {
        n, err := r.Read(buf)
        for _, b := range buf[:n] {
            if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
                out[string(b)]++
            }
        }
        if err == io.EOF {
            return out, nil
        }
        if err != nil {
            return nil, err
        }
    }
}
```

-> The argument of saving garbage collector work with this passing of buffers/sliced
-> But, doesn't `range` already do that *under the hood* for you?
-> This doesn't make sense.

-> But, I like the idea of a Reader interface as an input to the function.

normally you check error first
But, reader is an exception to this rule, and you check error **LAST**
-> WOOT

```go
s := "aaaaa bbbbbb cccccc dddddd eeeeee ffffff"

sr := strings.NewReader(s)
counts, err := countLetters(sr)
if err != nil {
    return err
}
fmt.Println(counts)
```

You can decorate readers.

```go
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
    r, err := os.Open(fileName)
    if err != nil {
        return nil, nil, err
    }
    gr, err := gzip.NewReader(r)
    if err != nil {
        // we probably should close the file here
        // r.Close()
        return nil, nil, err
    }
    return gr, func() {
        gr.Close()
        r.Close()
    }, nil
}
// if gzip returns an error do we properly close the file here? 

r, closer, err := buildGZipReader("my_data.txt.gz")
if err != nil {
    return err
}
defer closer()
counts, err := countLetters(r)
if err != nil {
    return err
}
fmt.Printl(counts)
```

* io.Copy
* io.MultiReader
* io.LimitReader
* io.MultiWriter

```go
// I don't like this one. In practice it's easy to miss.
type Closer interface {
    Close() error
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
    // whence shouldn't be an int as it accepts only:
    // io.SeekStart
    // io.SeekCurrent
    // io.SeekEnd
}
```

```go
f, err := os.Open(filename)
if err != nil {
    return nil, err
}
defer f.Close()
// use f
```

`ioutil.ReadAll`
`ioutil.ReadFile`
`ioutil.WriteFile` 
-> those look quite nice :) 

How to create a fake closer:

```go
type nopCloser struct {
    io.Reader
}

func (nopCloser) Close error { return nil }

func NopCloser(r io.Reader) io.ReadCloser {
    return nopCloser{r}
}
```

## time

* `time.Duration` represented by `int64`
* `time.Time`

time.Duration (smallest is 1 nanosecond)

when comparing times use `Equals` method instead of `==` because of time-zones. 

`t.Format("January 2, 2006 at 3:04:05PM MST")`

## encoding/json

marshalling: from Go data types -> JSON (or other encoding)
unmarshalling: JSON -> Go

### Use Struct Tags to Add Metadata

```go
type Item struct {
    ID string `json:"id"`
    Name string `json:"name"`
}
```

`json:-` to skip a field
`json:"name",omitempty`

I can agree that annotation in Java are overused. 

### Unmarshalling and Marshalling

```go
var i Item

err := json.Unmarshall([]byte(data), &i)
if err != nil {
    return err
}
```
One of the reasons is Go lack of generics

+ that it gives you control over memory allocation (NOT CONVINCED!) 

`out, err := json.Marshall(i)`

Why is there no passing of slice of bytes here? (because it's much nicer API)

Marshal and Unmarshall use Reflection

### JSON, Readers, and Writers

`encoding/json`

```go
type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

toFile := Person {
    Name: "Fred",
    Age: 40,
}

// streaming(?) reading and writing

tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
if err != nil {
    panic(err)
}
defer os.Remove(tmpFile.Name())

err = json.NewEncoder(tmpFile).Encode(toFile)
if err != nil {
    panic(err)
}
err = tmpFile.Close()
if err != nil {
    panic(err)
}

// now read

tmpFile2, err := os.Open(tmpFile.Name())
if err != nil {
    panic(err)
}

var fromFile Person
err = json.NewDecoder(tmpFile2).Decode(&fromFile)
if err != nil {
    panic(err)
}
err = tmpFile2.Close()
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", fromFile)

```

### Encoding and Decoding JSON Streams

-> This is weird. It's not really a valid JSON in this example. 

```go
dec := json.NewDecoder(strings.NewReader(data))
for dec.More() {
    err := dec.Decode(&t)
    if err != nil {
        panic(err)
    }
    // process t here
}
```


```go
var b bytes.Buffer
enc := json.NewEncoder(&b)
for _, input := range allInputs {
    t := process(input)
    err = enc.Encode(t)
    if err != nil {
        panic(err)
    }
}
out := b.String()
```

### Custom JSON Parsing

```go
type RFC822ZTime struct {
    time.Time
}

func (rc RFC822ZTime) MarshallJSON() ([]byte, error) {
    out := rt.Time.Format(time.RFC822Z)
    return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshallJSON(b []byte) error {
    if string(b) == "null" {
        return nil
    }
    t, err := time.Parse(`"` + time.RFC822Z + `"`, string(b))
    if err != nil {
        return err
    }
    *rt = RFC822ZTime{t}
    return nil
}
```

## `net/http`

### The Client

```go
client := &http.Client{
    Timeout: 30 * time.Second
}

req, err := http.NewRequestWithcontext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
if err != nil {
    panic(err)
}

req.Header.Add("X-My-Client", "Learning Go")
res, err := client.Do(req)
if err != nil {
    panic(err)
}

defer res.Body.Close() // important!

if res.StatusCode != http.StatusOK {
    panic(fmt.Sprintf("unexpected status: got %v", res.Status))
}

fmt.Println(res.Header.Get("Content-Type"))

var data struct {
    UserID int `json:"userID"`
    ID     int `json:"id"`
    Title  string `json:"title"`
    Completed bool `json:"completed"`
}

err = json.NewDecoder(res.Body).Decode(&data)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", data)
```

Always instantiate a Client - because you need to explicitly set the **timeout**.

## The Server

* http.Server
* http.Handler

```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}

type ResponseWriter interface {
    Header() http.Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

Always use this order:
1. Header()
2. WriteHeader()
3. Write()

```go
type HelloHandler struct{}

func (hh HelloHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!\n"))
}

s := http.Server{
    Addr: ":8080",
    ReadTimeout: 30 * time.Second,
    WriteTimeout: 90 * time.Second,
    IdleTimeout: 120 * time.Second,
    Handler: HelloHandler{},
}
err := s.ListenAndServe()
if err != nil {
    if err != http.ErrServerClosed {
        panic(err)
    }
}
```

reqest router -> `*http.ServeMux` `http.NewServeMux`

```go
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!\n"))
})
```

You can nest them:

```go
person := http.NewServeMux()
person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("greetings!\n"))
})

dog := http.NewServeMux()
dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("good putty!\n"))
})

mux := http.NewServeMux()
mux.Handle("/person", http.StripPrefix("/person", person))
mux.Handle("/dog", http.StripPrefix("/dog", dog))
```

### Middleware

```go
func RequestTimer(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        h.ServeHTTP(w, r)
        end := time.Now()
        log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
    })
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
    return func (h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.Header.Get("X-Secret-Password") != password {
                w.WriteHeader(http.StatusUnauthorized)
                w.Write(securityMsg)
                return
            }

            h.ServeHTTP(w, r)
        })
    }
}

terribleSecurity := TerribleSecurityProvider("GOPHER")

mux.Handle("/hello", terribleSecurity(RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello\n"))
}))))

alice - for chaining of middleware




```





