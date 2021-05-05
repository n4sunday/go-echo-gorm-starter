<p align="center">
<img src="https://cdn.labstack.com/images/echo-logo.svg" height="60" />
</p>


#### Quick Start
`Installation`
```bash
$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/**v4**
```

Create `server.go`

```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```
Start server
```bash
$ go run server.go
```