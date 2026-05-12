# Go Anexia API Client

A Go client for the Anexia Engine API.

This SDK provides a simple, versioned interface for interacting with Anexia resources such as locations and VLANs.

The project is still in early development.

## Installation

```bash
go get code.anexia.com/se/ks/go-anx-sdk
```

## Usage

The following shows how to use the api client.

```go
func main() {
	ctx := context.Background()

	apiKey := os.Getenv("API_KEY")

	client := go_anx_sdk.NewClient(
		config.WithAPIKey(apiKey),
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
// entry point to the v1 api endpoints
v1Client := client.V1()

// v1 locations api endpoints
locationV1Client := client.V1().Locations()
```

### Structure

The following diagram explains the structure of the api client and how it is used end to end.

```mermaid
flowchart TD
    A[Consumer Application]

    A --> B[go_anx_sdk.Client]

    B --> C[V1 Client]

    C --> D[LocationsClient]
    C --> E[VlansClient]
    C --> F[Other Resource Clients...]

    D --> G[internal.Transport]
    E --> G
    F --> G

    G --> H[net/http.Client]

    H --> I[Anexia API]

    J[config.ClientOption]
    J --> B

    K[LoggingRoundTripper]
    K --> H
```

The following diagram explains how a request flows through the different architectural layers. 

```mermaid
sequenceDiagram
    participant App
    participant Client
    participant V1
    participant Locations
    participant Transport
    participant HTTP
    participant API

    App->>Client: NewClient(opts...)
    App->>V1: client.V1()
    App->>Locations: Locations().List(ctx, params)

    Locations->>Transport: Get(ctx, endpoint, response, params)

    Transport->>Transport: buildRequestURL()
    Transport->>Transport: newRequest()

    Transport->>HTTP: client.Do(req)

    HTTP->>API: HTTPS Request
    API-->>HTTP: JSON Response

    HTTP-->>Transport: *http.Response

    Transport->>Transport: Decode JSON
    Transport-->>Locations: typed response

    Locations-->>App: PagedResponse[Location]
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