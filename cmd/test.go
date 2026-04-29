package main

import (
	"context"
	"net/http"
	"os"

	go_anx_sdk "code.anexia.com/se/ks/go-anx-sdk"
	"code.anexia.com/se/ks/go-anx-sdk/config"
	"code.anexia.com/se/ks/go-anx-sdk/utils"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("API_KEY")

	httpClient := http.DefaultClient
	httpClient.Transport = utils.NewLoggingRoundTripper(httpClient.Transport)

	client := go_anx_sdk.NewClient(
		config.WithApiKey(apiKey),
		config.WithBaseURL("https://engine.anexia-it.com/"),
		config.WithHttpClient(httpClient),
	)

	err := client.V1().Vlans().Delete(ctx, "ff84f6df77ef408e843d1047668e04a5")
	if err != nil {
		panic(err)
	}
}
