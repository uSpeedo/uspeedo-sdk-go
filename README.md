English | [简体中文](README_cn.md)

<h1 align="center">uspeedo Go SDK</h1>

[![GitHub release](https://img.shields.io/github/release/uspeedo/uspeedo-sdk-go/all.svg)](https://github.com/uspeedo/uspeedo-sdk-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/uspeedo/uspeedo-sdk-go)](https://goreportcard.com/report/github.com/uspeedo/uspeedo-sdk-go)
[![GoDoc](https://godoc.org/github.com/uspeedo/uspeedo-sdk-go?status.svg)](https://godoc.org/github.com/uspeedo/uspeedo-sdk-go)
[![GitHub](https://img.shields.io/github/license/uspeedo/uspeedo-sdk-go.svg)](http://www.apache.org/licenses/LICENSE-2.0)

- [Website](https://uspeedo.com/)

## Installation

### Requirements

- Go 1.10+

### Use `go get`

```bash
go get github.com/uspeedo/uspeedo-sdk-go
```

**Note** if meet network problem, you can use go proxy to speed up the downloaded, eg: use GOPROXY environment variable

```go
export GOPROXY=https://goproxy.io
```

Replay the command to retry installation.

### Use `go mod`

Add the following snippet to any code.

```go
import _ "github.com/uspeedo/uspeedo-sdk-go"
```

And execute this commands：

```bash
go mod init
go mod tidy
```

### Use `dep`

```bash
dep ensure -add github.com/uspeedo/uspeedo-sdk-go
```

## First Using

Currently, Go SDK use `PublicKey/PrivateKey` as authentication method, the key can be found from：

- [AccessKey](https://console.uspeedo.com/dashboard)

Here is a simple example：

```go
package main

import (
	"fmt"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/services/asms"
)

func main() {
	cfg := uspeedo.NewConfig()

	// replace the public/private key by your own
	credential := auth.NewCredential()
	credential.PrivateKey = "my_private_key"
	credential.PublicKey = "my_public_key"

	asmsClient := asms.NewClient(&cfg, &credential)

	req := asmsClient.NewSendBatchUSMSMessageRequest()
	req.ChargeType = uspeedo.String("Dynamic")
	req.TaskContent = []asms.SendBatchInfo{
		{
			TemplateId: "...",
			SenderId:   "",
			Target: []asms.SendBatchTarget{
				{
					Phone: "130xxxx1321",
				},
				{
					Phone: "130xxxx1321",
				},
			},
		},
	}

	// send request
	resp, err := asmsClient.SendBatchUSMSMessage(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", resp)
	}
}
```

## Feedback & Contribution

- [Issue](https://github.com/uspeedo/uspeedo-sdk-go/issues)
- [Pull Request](https://github.com/uspeedo/uspeedo-sdk-go/pulls)
