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

### Merchant Profile
```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/quarksgroup/paypack-go"
)

func main() {
    client := paypack.NewDefault()

    merchant, err := client.Profile(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(merchant.Name)
}
```

### Cashin (Deposit)

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/quarksgroup/paypack-go"
)

func main() {
    client := paypack.NewDefault()

    tx, err := client.Cashin(context.Background(), &paypack.TransactionRequest{
        Amount: 100,
        Number: "0789898989",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(tx.Ref)
}
```

### List Events
```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/quarksgroup/paypack-go"
)

func main() {
    client := paypack.NewDefault()

    events, err := client.ListEvents(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(events.Transactions)
}
```

### Authentication

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/quarksgroup/paypack-go"
)

func main() {
    client := paypack.NewDefault()
    clientID := os.Getenv("CLIENT_ID")
    clientSecret := os.Getenv("CLIENT_SECRET")

    token, err := client.Login(context.Background(), clientID, clientSecret)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(token.Access)
}
```

### Usage

Access full documentation [here](https://docs.paypack.rw/sdk/go)
