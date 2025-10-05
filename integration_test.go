// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bluehive-health/bluehive-sdk-go"
	"github.com/bluehive-health/bluehive-sdk-go/internal/testutil"
	"github.com/bluehive-health/bluehive-sdk-go/option"
)

func TestIntegrationList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcombluehivehealthbluehivesdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Integrations.List(context.TODO(), githubcombluehivehealthbluehivesdkgo.IntegrationListParams{
		XBrandID: "x",
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntegrationCheckActive(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcombluehivehealthbluehivesdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Integrations.CheckActive(
		context.TODO(),
		"name",
		githubcombluehivehealthbluehivesdkgo.IntegrationCheckActiveParams{
			XBrandID: "x",
		},
	)
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
