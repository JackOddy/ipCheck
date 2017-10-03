# A IP Address Check, written in Golang

A tool to return your external IP address. This tool with ping multiple APIs concurrently and periodically until it has two IP addresses that match.

## Using the Tool

Clone the repo into `$GOPATH/src` and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```
build:
```bash
  # inside the repo
  $ go build
```

start the server:
```bash
  #inside the repo
  $ ./ipCheck
```

## Test

Clone the repo and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```

run the tests
```bash
  # inside the repo
  $ go test 
```
Multiple test servers will be automatically spun up and torn down for the integration test.

*a note on testing: before any further features are added I would like to unit test the public API of this program with enough variance in data to cover all cases. Due to time restrictions this has not been possible*
