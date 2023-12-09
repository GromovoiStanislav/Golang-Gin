## Gin simple example

### Setup:

```
go mod init gin-example
go get -u github.com/gin-gonic/gin
```

### Start the server:

```
go run main.go
```

### Start the test:

```bash
go test -v
#or
go test -v
```

### Start the test benchmark:

```bash
go test -benchmem -bench=BenchmarkRequestGin
#or
go test -benchmem -bench RequestGin -run=^$
```
