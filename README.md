# wechatdomaincheck

A Go package to check if a domain is blocked by WeChat or not.

## Installation

To install the package, use the following command:

```bash
go get github.com/lib-x/wechatdomaincheck
```

## Usage

Here's a simple example of how to use the wechatdomaincheck package:

```go
package main

import (
	"fmt"
	"github.com/lib-x/wechatdomaincheck"
)

func main() {
    domain := "https://example.com"
    result, err := wechatdomaincheck.Check(domain)
    if err != nil {
        fmt.Printf("Error checking domain: %v\n", err)
        return
    }

    fmt.Printf("Check result for %s:\n", domain)
    fmt.Printf("Data: %s\n", result.Data)
    fmt.Printf("ReCode: %d\n", result.ReCode)

    if result.ReCode == 0 {
        fmt.Println("The domain is not blocked by WeChat.")
    } else {
        fmt.Println("The domain might be blocked by WeChat.")
    }
}

```

## API

### `Check(urlPath string) (*CheckResponse, error)`

Checks if a given URL is blocked by WeChat.

- `urlPath`: The URL to check (string)
- Returns: A pointer to `CheckResponse` and an error (if any)

### `CheckResponse` struct

```go
type CheckResponse struct {
	Data   string `json:"data"`
	ReCode int    `json:"reCode"`
}
```

- `Data`: Additional information from the WeChat API (string)
- `ReCode`: Response code from the WeChat API (int)
  - 0: The domain is not blocked
  - Other values: The domain might be blocked

## Documentation

For detailed documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/lib-x/wechatdomaincheck) page.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This package uses WeChat's public API for domain checking. The accuracy and availability of results depend on WeChat's service. Use this package responsibly and in accordance with WeChat's terms of service.
