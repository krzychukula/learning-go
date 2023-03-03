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




