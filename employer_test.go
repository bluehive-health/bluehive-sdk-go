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

func TestEmployerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Employers.New(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployerNewParams{
		Address: githubcombluehivehealthbluehivesdkgo.EmployerNewParamsAddress{
			City:    "city",
			State:   "state",
			Street1: "street1",
			ZipCode: "zipCode",
			Country: githubcombluehivehealthbluehivesdkgo.String("country"),
			Street2: githubcombluehivehealthbluehivesdkgo.String("street2"),
		},
		Email: "dev@stainless.com",
		Name:  "name",
		Phones: []githubcombluehivehealthbluehivesdkgo.EmployerNewParamsPhone{{
			Number:  "number",
			Primary: githubcombluehivehealthbluehivesdkgo.Bool(true),
			Type:    githubcombluehivehealthbluehivesdkgo.String("type"),
		}},
		BillingAddress: githubcombluehivehealthbluehivesdkgo.EmployerNewParamsBillingAddress{
			City:    "city",
			State:   "state",
			Street1: "street1",
			ZipCode: "zipCode",
			Country: githubcombluehivehealthbluehivesdkgo.String("country"),
			Street2: githubcombluehivehealthbluehivesdkgo.String("street2"),
		},
		Checkr: githubcombluehivehealthbluehivesdkgo.EmployerNewParamsCheckr{
			ID:     "id",
			Status: githubcombluehivehealthbluehivesdkgo.String("status"),
		},
		Demo:            githubcombluehivehealthbluehivesdkgo.Bool(true),
		EmployeeConsent: githubcombluehivehealthbluehivesdkgo.Bool(true),
		Metadata:        map[string]interface{}{},
		OnsiteClinic:    githubcombluehivehealthbluehivesdkgo.Bool(true),
		Website:         githubcombluehivehealthbluehivesdkgo.String("website"),
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployerGet(t *testing.T) {
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
	_, err := client.Employers.Get(context.TODO(), "employerId")
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
