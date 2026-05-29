package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	go_anx_sdk "github.com/kkostial/go-anx-sdk"
	"github.com/kkostial/go-anx-sdk/config"
	"github.com/kkostial/go-anx-sdk/paging"
	"github.com/kkostial/go-anx-sdk/utils"
	v1 "github.com/kkostial/go-anx-sdk/v1"
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

	clusters := paging.Paginate(ctx, client.V1().DevClusters().ListPageFetcher(v1.ClusterListParams{}))
	for cluster, err := range clusters {
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", cluster)
	}
}
