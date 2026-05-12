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
    Consumer[Consumer Application]

    SdkClient[go_anx_sdk.Client]
    V1Client[V1 Client]

    LocationsClient[LocationsClient]
    VlansClient[VlansClient]
    OtherClients[Other Resource Clients...]

    Transport[internal.Transport]
    HttpClient[net/http.Client]
    AnxApi[Anexia API]

    ClientOption[config.ClientOption]
    LoggingRoundTripper[LoggingRoundTripper]

    Consumer --> SdkClient
    SdkClient --> V1Client
    V1Client --> LocationsClient
    V1Client --> VlansClient
    V1Client --> OtherClients

    LocationsClient --> Transport
    VlansClient --> Transport
    OtherClients --> Transport

    Transport --> HttpClient
    HttpClient --> AnxApi

    ClientOption --> SdkClient
    LoggingRoundTripper --> HttpClient

%% -----------------------
%% Styling / grouping
%% -----------------------

    classDef app fill:#E3F2FD,stroke:#1E88E5,stroke-width:2px;
    classDef sdk_public fill:#E8F5E9,stroke:#43A047,stroke-width:2px;
    classDef sdk_internal fill:#FFF3E0,stroke:#FB8C00,stroke-width:2px;
    classDef api fill:#F3E5F5,stroke:#8E24AA,stroke-width:2px;

    class Consumer app;
    class SdkClient,V1Client,LocationsClient,VlansClient,OtherClients,ClientOption,LoggingRoundTripper sdk_public;
    class Transport,HttpClient sdk_internal;
    class AnxApi api;
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

- WithApiKey(string) - required or api will return a 401
- WithBaseURL(string) - if omitted 'https://engine.anexia-it.com' will be used
- WithHttpClient(*http.Client) - if omitted the default `http.Client` will be used

## Error handling

TBD

## Testing

    go test ./...

## License

TBD