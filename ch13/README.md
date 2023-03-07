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



