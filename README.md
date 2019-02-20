# go-mod-tidy

### 墙内专供

墙内使用 `go mod tidy` 时，经常遇到 `golang.org/x/xxx` 之类墙内访问不到的包，在 `go.mod` 中可以使用 
`replace` 来替换包地址，但手动修改太麻烦，因此做了一个自动工具

### 安装

```go
go get -v -u github.com/hqpko/go-mod-tidy
```

> 请检查是否设置了 `PATH=$PATH:$GOPATH/bin`


### 使用

在需要更新 `go.mod` 的项目中，使用 `go-mod-tidy` 命令，会自动添加 `replace`

### 现有的 replace 包


```go
    replaceMap = map[string]string{
		"golang.org/x/tools":          "github.com/golang/tools",
		"golang.org/x/sys":            "github.com/golang/sys",
		"golang.org/x/sync":           "github.com/golang/sync",
		"golang.org/x/oauth2":         "github.com/golang/oauth2",
		"golang.org/x/net":            "github.com/golang/net",
		"golang.org/x/lint":           "github.com/golang/lint",
		"golang.org/x/text":           "github.com/golang/text",
		"google.golang.org/genproto":  "github.com/google/go-genproto",
		"google.golang.org/grpc":      "github.com/grpc/grpc-go",
		"google.golang.org/appengine": "github.com/golang/appengine",
		"cloud.google.com/go":         "github.com/googleapis/google-cloud-go",
		"google.golang.org/api":       "github.com/googleapis/google-api-go-client",
	}
```