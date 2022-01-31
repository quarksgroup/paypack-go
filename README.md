# Go Paypack
The official [Paypack](https://payments.paypack.rw/) Go client library.

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference paypack-go in a Go program with `import`:

``` go
import (
    "github.com/quarksgroup/paypack-go"
)
```

Run any of the normal `go mod tidy`. The Go
toolchain will resolve and fetch the paypack-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```
go get -u github.com/quarksgroup/paypack-go@latest
```

## Below are a few simple examples:

### Merchant

### Events

### Authentication

### Usage

[api-docs]: http://payments.paypack.rw/api
