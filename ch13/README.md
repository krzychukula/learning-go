# Chapter 13. Writing Tests

## The Basics of Testing

```
$ ch13 % go mod init adder
go: creating new go.mod: module adder
go: to add module requirements and sums:
        go mod tidy
$ ch13 % cd adder 
$ adder % go mod tidy
$ adder % go test
PASS
ok      adder/adder     0.005s
```

-> It's really good that `_test.go` live next to the files they test!

```go
func Test_addNumbersShouldReturn5For2and3(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
```

### Reporting Test Failures

`t.Error("incorrect result: expected 5, got", result)`
`t.Errorf("incorrect result: expected %d, got %d", 5, result)`
-> test still is running!

Use:
* `Fatal`
* `Fatalf`
-> to stop text execution

### Setting Up and Tearing Down

```go
func TestMain(m *testing.M) {
    // setup
    exitVal := m.Run()
    // teardown
    os.Exit(exitVal)
}
```

* only one `TestMain` for a package!

```go
t.Cleanup(func() {
    os.Remove(f.Name())
})
```

### Storing Sample Test Data

`./testdata/` directory

### Caching Test Results

`go test count=1` - to force running tests regardless of the cache

### Testing Your Public API

Use `packagename_test` for the package name of the test to force testing of the **Public API only**

### Use go-cmp to Compare Test Results

```go
if diff := cmp.Diff(expected, result); diff != "" {
    t.Error(diff)
}
```

## Table Tests

-> I don't think those loop tests should be promoted so much.

## Checking Your Code Coverage

`go test -v cover -coverprofile=c.out`

`go tool cover -html=c.out`

## Benchmarks

```go
var blackhole int

func BenchmarkFileLen1(b *testing.B) {
    // every benchmark needs to run b.N times
    for i := 0; i < b.N; i++ {
        result, err := FileLen("testdata/data.txt", 1)
        if err != nil {
            b.Fatal(err)
        }
        blackhole = result
    }
}
```

`go test -bench`

## Stubs in Go

You can embed an interface in a struct

```go
type Entities interface {
    GetUser(id string) (User, error)
    GetPets(userID string) ([]Pet, error)
    GetChildren(userID string) ([]Person, error)
    GetFriends(userID string) ([]Person, error)
    SaveUser(user User) error
}

type Logic struct {
    Entities Entities
}
```

> In short, a stub returns a canned value for a given input, whereas a mock validates that a set of calls happen in the expected order with the expected inputs.

## httptest

`httptest.NewServer`

## Integration Tests and Build Tags

```go
// +build integration
```

`go test -tags integration -v ./..`

or `-short` flag

## Finding Concurrency Problems with the Race Checker

`go test -race`

Binary with `-race` runs 10X slower.

