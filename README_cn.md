[English](README.md) | 简体中文

<h1 align="center">uspeedo Go SDK</h1>

[![GitHub (pre-)release](https://img.shields.io/github/release/uspeedo/uspeedo-sdk-go/all.svg)](https://github.com/uspeedo/uspeedo-sdk-go/releases)
[![GoDoc](https://godoc.org/github.com/uspeedo/uspeedo-sdk-go?status.svg)](https://godoc.org/github.com/uspeedo/uspeedo-sdk-go)
[![GitHub](https://img.shields.io/github/license/uspeedo/uspeedo-sdk-go.svg)](http://www.apache.org/licenses/LICENSE-2.0)

- [Website](https://uspeedo.com/)

## 安装

### 环境要求

- Go 1.10+

### 使用 go get 安装

```bash
go get github.com/uspeedo/uspeedo-sdk-go
```

**Note** 如果遇到网络不稳定，可以使用代理服务器来加速下载，例如使用 GOPROXY 加速：

```go
export GOPROXY = https: //goproxy.cn
```

再次执行安装命令即可。

### 使用 go mod 安装

在任意代码中添加

```go
import _ "github.com/uspeedo/uspeedo-sdk-go"
```

之后在项目根目录执行：

```bash
go mod init
go mod tidy
```

### 使用 dep 安装

```bash
dep ensure -add github.com/uspeedo/uspeedo-sdk-go
```

## 初次使用

目前，Go SDK 使用 PublicKey/PrivateKey 作为唯一的鉴权方式，该公私钥可以从以下途径获取：

- [AccessKey](https://console.uspeedo.com/dashboard)

下面提供一个简单的示例：

```go
package main

import (
	"fmt"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/services/usms"
)

func main() {
	cfg := uspeedo.NewConfig()

	// replace the public/private key by your own
	credential := auth.NewCredential()
	credential.PrivateKey = "my_private_key"
	credential.PublicKey = "my_public_key"

	usmsClient := usms.NewClient(&cfg, &credential)

	req := usmsClient.NewSendBatchUSMSMessageRequest()
	req.ChargeType = uspeedo.String("Dynamic")
	req.TaskContent = []usms.SendBatchInfo{
		{
			TemplateId: "...",
			SenderId:   "",
			Target: []usms.SendBatchTarget{
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
	resp, err := usmsClient.SendBatchUSMSMessage(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", resp)
	}
}
```

## 反馈 & 贡献

- [Issue](https://github.com/uspeedo/uspeedo-sdk-go/issues)
- [Pull Request](https://github.com/uspeedo/uspeedo-sdk-go/pulls)
