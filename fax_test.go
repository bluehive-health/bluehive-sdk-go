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

func TestFaxListProviders(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Fax.ListProviders(context.TODO())
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Fax.GetStatus(context.TODO(), "id")
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxSendWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Fax.Send(context.TODO(), githubcombluehivehealthbluehivesdkgo.FaxSendParams{
		Document: githubcombluehivehealthbluehivesdkgo.FaxSendParamsDocument{
			Content:     "content",
			ContentType: "application/pdf",
			Filename:    githubcombluehivehealthbluehivesdkgo.String("filename"),
		},
		To:       "to",
		From:     githubcombluehivehealthbluehivesdkgo.String("from"),
		Provider: githubcombluehivehealthbluehivesdkgo.String("provider"),
		Subject:  githubcombluehivehealthbluehivesdkgo.String("subject"),
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
