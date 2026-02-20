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

func TestEmployeeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Employees.New(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployeeNewParams{
		Email:         "dev@stainless.com",
		FirstName:     "x",
		LastName:      "x",
		ActiveAccount: githubcombluehivehealthbluehivesdkgo.EmployeeNewParamsActiveAccountActive,
		Address: githubcombluehivehealthbluehivesdkgo.EmployeeNewParamsAddress{
			City:       "x",
			PostalCode: "x",
			State:      "x",
			Street1:    "x",
			Country:    githubcombluehivehealthbluehivesdkgo.String("country"),
			County:     githubcombluehivehealthbluehivesdkgo.String("county"),
			Street2:    githubcombluehivehealthbluehivesdkgo.String("street2"),
		},
		Blurb:       githubcombluehivehealthbluehivesdkgo.String("blurb"),
		Departments: []string{"string"},
		Dob:         githubcombluehivehealthbluehivesdkgo.String("7321-69-10"),
		EmployerID:  githubcombluehivehealthbluehivesdkgo.String("employer_id"),
		ExtendedFields: []githubcombluehivehealthbluehivesdkgo.EmployeeNewParamsExtendedField{{
			Name:  "x",
			Value: "x",
		}},
		Phone: []githubcombluehivehealthbluehivesdkgo.EmployeeNewParamsPhone{{
			Number: "x",
			Type:   "Cell",
		}},
		Title: githubcombluehivehealthbluehivesdkgo.String("title"),
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeGet(t *testing.T) {
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
	_, err := client.Employees.Get(context.TODO(), "employeeId")
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Employees.Update(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployeeUpdateParams{
		ID:            "x",
		ActiveAccount: githubcombluehivehealthbluehivesdkgo.EmployeeUpdateParamsActiveAccountActive,
		Address: githubcombluehivehealthbluehivesdkgo.EmployeeUpdateParamsAddress{
			City:       "x",
			PostalCode: "x",
			State:      "x",
			Street1:    "x",
			Country:    githubcombluehivehealthbluehivesdkgo.String("country"),
			County:     githubcombluehivehealthbluehivesdkgo.String("county"),
			Street2:    githubcombluehivehealthbluehivesdkgo.String("street2"),
		},
		Blurb:       githubcombluehivehealthbluehivesdkgo.String("blurb"),
		Departments: []string{"string"},
		Dob:         githubcombluehivehealthbluehivesdkgo.String("7321-69-10"),
		Email:       githubcombluehivehealthbluehivesdkgo.String("dev@stainless.com"),
		EmployerID:  githubcombluehivehealthbluehivesdkgo.String("employer_id"),
		ExtendedFields: []githubcombluehivehealthbluehivesdkgo.EmployeeUpdateParamsExtendedField{{
			Name:  "x",
			Value: "x",
		}},
		FirstName: githubcombluehivehealthbluehivesdkgo.String("x"),
		LastName:  githubcombluehivehealthbluehivesdkgo.String("x"),
		Phone: []githubcombluehivehealthbluehivesdkgo.EmployeeUpdateParamsPhone{{
			Number: "x",
			Type:   "Cell",
		}},
		Title: githubcombluehivehealthbluehivesdkgo.String("title"),
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Employees.List(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployeeListParams{
		EmployerID: "employerId",
		Limit:      githubcombluehivehealthbluehivesdkgo.String("269125115713"),
		Offset:     githubcombluehivehealthbluehivesdkgo.String("269125115713"),
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeDelete(t *testing.T) {
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
	_, err := client.Employees.Delete(context.TODO(), "employeeId")
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeLinkUserWithOptionalParams(t *testing.T) {
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
	_, err := client.Employees.LinkUser(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployeeLinkUserParams{
		EmployeeID: "x",
		UserID:     "x",
		Role:       []string{"string"},
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmployeeUnlinkUser(t *testing.T) {
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
	_, err := client.Employees.UnlinkUser(context.TODO(), githubcombluehivehealthbluehivesdkgo.EmployeeUnlinkUserParams{
		EmployeeID: "employeeId",
		UserID:     "userId",
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
