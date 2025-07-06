// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bluehive_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bluehive-go"
	"github.com/stainless-sdks/bluehive-go/internal/testutil"
	"github.com/stainless-sdks/bluehive-go/option"
)

func TestFaxListProviders(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bluehive.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fax.ListProviders(context.TODO())
	if err != nil {
		var apierr *bluehive.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxGetStatus(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bluehive.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fax.GetStatus(context.TODO(), "id")
	if err != nil {
		var apierr *bluehive.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxSendWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bluehive.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fax.Send(context.TODO(), bluehive.FaxSendParams{
		Document: bluehive.FaxSendParamsDocument{
			Content:     "content",
			ContentType: "application/pdf",
			Filename:    bluehive.String("filename"),
		},
		To:       "to",
		From:     bluehive.String("from"),
		Provider: bluehive.String("provider"),
		Subject:  bluehive.String("subject"),
	})
	if err != nil {
		var apierr *bluehive.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
