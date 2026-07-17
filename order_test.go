// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/bluehive-health/bluehive-sdk-go"
	"github.com/bluehive-health/bluehive-sdk-go/internal/testutil"
	"github.com/bluehive-health/bluehive-sdk-go/option"
)

func TestOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.New(context.TODO(), githubcombluehivehealthbluehivesdkgo.OrderNewParams{
		OfObject: &githubcombluehivehealthbluehivesdkgo.OrderNewParamsBodyObject{
			PaymentMethod: "self-pay",
			Person: githubcombluehivehealthbluehivesdkgo.OrderNewParamsBodyObjectPerson{
				City:      "x",
				Dob:       "7321-69-10",
				Email:     "email",
				FirstName: "x",
				LastName:  "x",
				Phone:     "+)() 92))()1)",
				State:     "xx",
				Street:    "x",
				Zipcode:   "73216-0225",
				Country:   githubcombluehivehealthbluehivesdkgo.String("country"),
				County:    githubcombluehivehealthbluehivesdkgo.String("county"),
				Street2:   githubcombluehivehealthbluehivesdkgo.String("street2"),
			},
			ProviderID: "providerId",
			Services: []githubcombluehivehealthbluehivesdkgo.OrderNewParamsBodyObjectService{{
				ID:         "x",
				Quantity:   1,
				AutoAccept: githubcombluehivehealthbluehivesdkgo.Bool(true),
			}},
			ID:      githubcombluehivehealthbluehivesdkgo.String("_id"),
			BrandID: githubcombluehivehealthbluehivesdkgo.String("brandId"),
			BundleIDs: map[string]string{
				"foo": "string",
			},
			DueDate:     githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
			DueDates:    []time.Time{time.Now()},
			EmployeeID:  githubcombluehivehealthbluehivesdkgo.String("employeeId"),
			EmployeeIDs: []string{"string"},
			EmployerID:  githubcombluehivehealthbluehivesdkgo.String("employerId"),
			Metadata: map[string]any{
				"foo": "bar",
			},
			Priority:        "normal",
			ProviderCreated: githubcombluehivehealthbluehivesdkgo.Bool(true),
			ProvidersIDs: []githubcombluehivehealthbluehivesdkgo.OrderNewParamsBodyObjectProvidersID{{
				ProviderID: "x",
				ServiceID:  githubcombluehivehealthbluehivesdkgo.String("x"),
			}},
			Quantities: map[string]int64{
				"foo": 1,
			},
			ReCaptchaToken: githubcombluehivehealthbluehivesdkgo.String("reCaptchaToken"),
			ServicesIDs:    []string{"string"},
			TokenID:        githubcombluehivehealthbluehivesdkgo.String("tokenId"),
		},
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderGet(t *testing.T) {
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
	_, err := client.Orders.Get(context.TODO(), "orderId")
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.Update(
		context.TODO(),
		"orderId",
		githubcombluehivehealthbluehivesdkgo.OrderUpdateParams{
			ExpirationDate: githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
			Metadata: map[string]any{
				"foo": "bar",
			},
			Services: []githubcombluehivehealthbluehivesdkgo.OrderUpdateParamsService{{
				ServiceID:      "x",
				DueDate:        githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
				ExpirationDate: githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
				Results: map[string]any{
					"foo": "bar",
				},
				Status: "pending",
			}},
			Status: githubcombluehivehealthbluehivesdkgo.OrderUpdateParamsStatusOrderSent,
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

func TestOrderGetResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.GetResults(
		context.TODO(),
		"orderId",
		githubcombluehivehealthbluehivesdkgo.OrderGetResultsParams{
			Page:      githubcombluehivehealthbluehivesdkgo.Int(1),
			PageSize:  githubcombluehivehealthbluehivesdkgo.Int(1),
			ServiceID: githubcombluehivehealthbluehivesdkgo.String("serviceId"),
			Since:     githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
			Status:    githubcombluehivehealthbluehivesdkgo.String("status"),
			Until:     githubcombluehivehealthbluehivesdkgo.Time(time.Now()),
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

func TestOrderScheduleAppointmentWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.ScheduleAppointment(
		context.TODO(),
		"orderId",
		githubcombluehivehealthbluehivesdkgo.OrderScheduleAppointmentParams{
			Appointment: githubcombluehivehealthbluehivesdkgo.OrderScheduleAppointmentParamsAppointmentUnion{
				OfOrderScheduleAppointmentsAppointmentObject: &githubcombluehivehealthbluehivesdkgo.OrderScheduleAppointmentParamsAppointmentObject{
					Date:     "date",
					DateTime: time.Now(),
					Time:     "time",
					Notes:    githubcombluehivehealthbluehivesdkgo.String("notes"),
					Type:     "appointment",
				},
			},
			OrderAccessCode: "orderAccessCode",
			ProviderID:      githubcombluehivehealthbluehivesdkgo.String("providerId"),
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

func TestOrderSendForEmployeeWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.SendForEmployee(context.TODO(), githubcombluehivehealthbluehivesdkgo.OrderSendForEmployeeParams{
		EmployeeID: "employeeId",
		EmployerID: "employerId",
		ProvidersIDs: []githubcombluehivehealthbluehivesdkgo.OrderSendForEmployeeParamsProvidersID{{
			ProviderID: "providerId",
			ServiceID:  githubcombluehivehealthbluehivesdkgo.String("serviceId"),
		}},
		ServicesIDs:    []string{"string"},
		LoginToken:     "login-token",
		UserID:         "user-id",
		BrandID:        githubcombluehivehealthbluehivesdkgo.String("brandId"),
		DueDate:        githubcombluehivehealthbluehivesdkgo.String("dueDate"),
		DueDates:       []string{"string"},
		ExpirationDate: githubcombluehivehealthbluehivesdkgo.String("expirationDate"),
		Metadata: map[string]any{
			"foo": "bar",
		},
		Priority:        githubcombluehivehealthbluehivesdkgo.OrderSendForEmployeeParamsPriorityNormal,
		ProviderCreated: githubcombluehivehealthbluehivesdkgo.Bool(true),
		ProviderID:      githubcombluehivehealthbluehivesdkgo.String("providerId"),
		Quantities: map[string]int64{
			"foo": 1,
		},
	})
	if err != nil {
		var apierr *githubcombluehivehealthbluehivesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderUploadResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.UploadResults(
		context.TODO(),
		"orderId",
		githubcombluehivehealthbluehivesdkgo.OrderUploadResultsParams{
			CaptchaToken:    "x",
			OrderAccessCode: "x",
			ServiceID:       "x",
			Dob:             githubcombluehivehealthbluehivesdkgo.String("7321-69-10"),
			FileIDs:         []string{"x"},
			Files: []githubcombluehivehealthbluehivesdkgo.OrderUploadResultsParamsFile{{
				Base64: "x",
				Name:   "x",
				Type:   "x",
			}},
			LastName: githubcombluehivehealthbluehivesdkgo.String("x"),
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
