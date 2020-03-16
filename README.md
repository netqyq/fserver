# fserver
fserver is a file server.

## Installation
Download the binary release. or install by go command
```go
go install github.com/netqyq/fserver
```

## Usage
start a file server
```bash
$ fserver -dir /tmp/www -addr 0.0.0.0:5000
```
help message
```bash
# fserver -h
Usage of fserver:
  -addr string
        server listen address (default "0.0.0.0:8000")
  -dir string
        www root directory (default "/tmp/www")
```
## License
MIT