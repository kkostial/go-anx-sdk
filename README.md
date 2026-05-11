# Go Anexia API Client

A Go client for the Anexia Engine API.

This SDK provides a simple, versioned interface for interacting with Anexia resources such as locations and VLANs.

The project is still in early development.

## Installation

```bash
go get code.anexia.com/se/ks/go-anx-sdk
Usage
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	go_anx_sdk "code.anexia.com/se/ks/go-anx-sdk"
	"code.anexia.com/se/ks/go-anx-sdk/config"
	"code.anexia.com/se/ks/go-anx-sdk/utils"
	v1 "code.anexia.com/se/ks/go-anx-sdk/v1"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("API_KEY")

	client := go_anx_sdk.NewClient(
		config.WithApiKey(apiKey),
		config.WithBaseURL("https://engine.anexia-it.com/"),
		config.WithHttpClient(&http.Client{
			Transport: utils.NewLoggingRoundTripper(http.DefaultTransport),
		}),
	)

	locations, err := client.V1().Locations().List(ctx, v1.LocationListParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range locations.Data {
		fmt.Printf("%+v\n", l)
	}
}
```

## API

### Versioning

All endpoints are accessed via a versioned client:

```go
client.V1()
```

### Resources

Currently available:

- Locations
- VLANs

Example:

```go 
client.V1().Locations().List(ctx, params)
client.V1().Vlans().List(ctx, params)
```

## Configuration

The client is configured using functional options:

- WithApiKey(string)
- WithBaseURL(string)
- WithHttpClient(*http.Client)

## Error handling

TBD

## Testing

    go test ./...

## License

TBD