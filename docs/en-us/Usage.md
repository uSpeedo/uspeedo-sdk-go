## Usage

### New Client

sdk client for any uspeedo product openapi service should constructed by client config and credential config,

```go
import (
    "github.com/uspeedo/uspeedo-sdk-go/services/uhost"
)

client := uhost.NewClient(&cfg, &credential)
```

### Build Request

```go
req := client.NewCreateUHostInstanceRequest()

req.Password = uspeedo.String("xxx")
req.LoginMode = uspeedo.String("Password")
req.Region = uspeedo.String("cn-bj2")
req.Zone = uspeedo.String("cn-bj2-04")
req.ImageId = uspeedo.String("uimage-xxxx")
req.CPU = uspeedo.Int(1)
req.Memory = uspeedo.Int(1024)
```

#### Request Timeout

To set request timeout for current request, it will overwrite the client Timeout at current request only.

```go
req.WithTimeout(10 * time.Second)
```

#### Request Retring

To set request max retries > 0 to enable auto retry, it will overwrite the client MaxRetries at current request only.

```go
req.WithRetry(3)
```

### Send Request

```go
resp, err := client.CreateUHostInstance(req)
log.Printf("%#v", resp.UHostIds)
```

### Error Handling

#### General Usage

In general usage, you can serialize error to string for free, use ``err.Error()`` or

```go
fmt.Printf("got error: %s", err)
```

For business error with RetCode > 0 at the body of server response, you can check it easily, such as

```go
if uerr.IsCodeError(err) {
    fmt.Printf("%v %s", e.Code(), e.Message())
} else {
    fmt.Printf("%s", e)
}
```

#### Advanced Usage

In advanced usage, error could be infered uspeedo sdk client error / server error, such as

```go
import (
    uerr "github.com/uspeedo/uspeedo-sdk-go/uspeedo/error"
)

if err != nil {
    switch err.(type) {
    case uerr.(ClientError):
        fmt.Printf("client error: %s", err)
    case uerr.(ServerError):
        fmt.Printf("server error: %s", err)
    }
}
```

or enum by error name

```go
import (
    uerr "github.com/uspeedo/uspeedo-sdk-go/uspeedo/error"
)

if e, ok := err.(uerr.Error); err != nil && ok {
    switch e.Name() {
    case uerr.ErrRetCode:
        fmt.Printf("%v %s", e.Code(), e.Message())
    default:
        fmt.Printf("%s", e)
    }
}
```
