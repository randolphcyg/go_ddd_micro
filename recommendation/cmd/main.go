package main

import (
	"log/slog"
	"net/http"

	"go_ddd_micro/recommendation/internal/recommendation"
	"go_ddd_micro/recommendation/internal/transport"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	// 可重试的 http 客户端
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	partnerAdaptor, err := recommendation.NewPartnerShipAdaptor(
		c.StandardClient(),
		"http://localhost:3031",
	)
	if err != nil {
		slog.Warn("failed to create a partnerAdaptor: ", err)
		panic(err)
	}

	svc, err := recommendation.NewService(partnerAdaptor)
	if err != nil {
		slog.Warn("failed to create a service: ", err)
		panic(err)
	}

	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		slog.Warn("failed to create a handler: ", err)
		panic(err)
	}

	m := transport.NewMux(*handler)

	if err := http.ListenAndServe(":4040", m); err != nil {
		slog.Warn("server errored: ", err)
		panic(err)
	}
}
