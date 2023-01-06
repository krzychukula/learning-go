# Chapter 1

go install github.com/rakyll/hey@latest
hey https://go.dev/


go fmt -> for tools

go install golang.org/x/tools/cmd/goimports@latest
goimports -l -w .

## Linting and Vetting
 
- [ ] read Effective Go
- [ ] Code Review Comments page on Go's wiki

Lint is not 100% correct. 

Use golangci-lint, staticcheck or revive.

go vet ./..

golangci-lint run

## Choose Your Tools

Sharing on go playground generates a unique url, just remember to not post anything confidential.

## Makefiles

go mod init ch1
make

## Staying Up to Date

go get golang.org/dl/go.1.15.6
go1.15.6 download

go1.15.6 build // test your program in the new compiler version

Clean it up by finding it GOROOT

go1.15.6 env GOROOT // -> some path
rm -rf $(go1.15.6 env GOROOT)
go $(go env GOPATH)/bin/go1.15.6

You can also use:

* https://magefile.org/
* https://taskfile.dev/