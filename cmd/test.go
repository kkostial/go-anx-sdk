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

	httpClient := &http.Client{
		Transport: utils.NewLoggingRoundTripper(http.DefaultTransport),
	}

	client := go_anx_sdk.NewClient(
		config.WithAPIKey(apiKey),
		config.WithBaseURL("https://engine.anexia-it.com/"),
		config.WithHTTPClient(httpClient),
	)

	locations, err := client.V1().Locations().List(ctx, v1.LocationListParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range locations.Data {
		fmt.Printf("%+v\n", l)
	}
}
