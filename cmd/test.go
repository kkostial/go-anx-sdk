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

	cluster, err := client.V1().Clusters().Get(ctx, "6f2e578fee7741528ea3d94ff156141b")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cluster)
}
