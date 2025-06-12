// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttptrace

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"testing"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func ExampleNewClientTrace() {
	client := http.Client{
		Transport: otelhttp.NewTransport(
			http.DefaultTransport,
			otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
				return NewClientTrace(ctx)
			}),
		),
	}

	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func TestClient_NoHookContext(t *testing.T) {
	t.Run("Got100Continue", func(t *testing.T) {
		trace := NewClientTrace(context.Background())
		trace.Got100Continue()
	})
}
