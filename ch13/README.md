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

