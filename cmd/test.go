package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	go_anx_sdk "github.com/kkostial/go-anx-sdk"
	"github.com/kkostial/go-anx-sdk/config"
	"github.com/kkostial/go-anx-sdk/internal/utils/ptr"
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

	cluster, err := client.V1().Clusters().Update(ctx, "b06c68f4a2154fe58d11ae12bed7039f", v1.ClusterUpdateRequest{
		KubeConfig: ptr.To("hallo!"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cluster)
}
