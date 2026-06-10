// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/bluehive-health/bluehive-sdk-go/internal/apijson"
	"github.com/bluehive-health/bluehive-sdk-go/internal/apiquery"
	"github.com/bluehive-health/bluehive-sdk-go/internal/requestconfig"
	"github.com/bluehive-health/bluehive-sdk-go/option"
	"github.com/bluehive-health/bluehive-sdk-go/packages/param"
	"github.com/bluehive-health/bluehive-sdk-go/packages/respjson"
)

// OrderService contains methods and other services that help with interacting with
// the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options []option.RequestOption
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r OrderService) {
	r = OrderService{}
	r.Options = opts
	return
}

// Create orders for consumers (self-pay or employer-sponsored), employers, or bulk
// orders. Consolidates functionality from legacy Order.createOrder and
// Order.SendOrder methods.
func (r *OrderService) New(ctx context.Context, body OrderNewParams, opts ...option.RequestOption) (res *OrderNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details for a specific order
func (r *OrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *OrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update order details and associated order items. Allows updating order status,
// metadata, and modifying order item services.
func (r *OrderService) Update(ctx context.Context, orderID string, body OrderUpdateParams, opts ...option.RequestOption) (res *OrderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve results for an order. Supports filtering by serviceId, status, date
// window, and pagination.
func (r *OrderService) GetResults(ctx context.Context, orderID string, query OrderGetResultsParams, opts ...option.RequestOption) (res *OrderGetResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orders/%s/results", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Schedule an appointment or walk-in for an existing order. Sends HL7 SIU^S12
// message for appointment booking.
func (r *OrderService) ScheduleAppointment(ctx context.Context, orderID string, body OrderScheduleAppointmentParams, opts ...option.RequestOption) (res *OrderScheduleAppointmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orders/%s/schedule-appointment", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send an order for a specific employee. Requires API key, login token, and user
// ID. This endpoint specifically handles employer-to-employee order sending.
func (r *OrderService) SendForEmployee(ctx context.Context, params OrderSendForEmployeeParams, opts ...option.RequestOption) (res *OrderSendForEmployeeResponseUnion, err error) {
	if !param.IsOmitted(params.LoginToken) {
		opts = append(opts, option.WithHeader("login-token", fmt.Sprintf("%v", params.LoginToken)))
	}
	if !param.IsOmitted(params.UserID) {
		opts = append(opts, option.WithHeader("user-id", fmt.Sprintf("%v", params.UserID)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/orders/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Upload test results for a specific order item. Supports both existing fileIds
// and base64 encoded files. Requires order access code and employee verification.
func (r *OrderService) UploadResults(ctx context.Context, orderID string, body OrderUploadResultsParams, opts ...option.RequestOption) (res *OrderUploadResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orders/%s/upload-results", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// OrderNewResponseUnion contains all possible properties and values from
// [OrderNewResponseObject], [OrderNewResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OrderNewResponseUnion struct {
	// This field is from variant [OrderNewResponseObject].
	OrderID string `json:"orderId"`
	// This field is from variant [OrderNewResponseObject].
	OrderNumber string `json:"orderNumber"`
	Success     bool   `json:"success"`
	// This field is from variant [OrderNewResponseObject].
	HostedInvoiceURL string `json:"hostedInvoiceUrl"`
	Message          string `json:"message"`
	PartialSuccess   bool   `json:"partialSuccess"`
	// This field is from variant [OrderNewResponseObject].
	SelfPay bool `json:"selfPay"`
	// This field is a union of [[]OrderNewResponseObjectUnavailableService],
	// [[]OrderNewResponseObject2UnavailableService]
	UnavailableServices OrderNewResponseUnionUnavailableServices `json:"unavailableServices"`
	// This field is from variant [OrderNewResponseObject2].
	OrderResults []OrderNewResponseObject2OrderResult `json:"orderResults"`
	// This field is from variant [OrderNewResponseObject2].
	Status string `json:"status"`
	JSON   struct {
		OrderID             respjson.Field
		OrderNumber         respjson.Field
		Success             respjson.Field
		HostedInvoiceURL    respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		SelfPay             respjson.Field
		UnavailableServices respjson.Field
		OrderResults        respjson.Field
		Status              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u OrderNewResponseUnion) AsOrderNewResponseObject() (v OrderNewResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderNewResponseUnion) AsOrderNewResponseObject2() (v OrderNewResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OrderNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OrderNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OrderNewResponseUnionUnavailableServices is an implicit subunion of
// [OrderNewResponseUnion]. OrderNewResponseUnionUnavailableServices provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OrderNewResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOrderNewResponseObjectUnavailableServices
// OfOrderNewResponseObject2UnavailableServices]
type OrderNewResponseUnionUnavailableServices struct {
	// This field will be present if the value is a
	// [[]OrderNewResponseObjectUnavailableService] instead of an object.
	OfOrderNewResponseObjectUnavailableServices []OrderNewResponseObjectUnavailableService `json:",inline"`
	// This field will be present if the value is a
	// [[]OrderNewResponseObject2UnavailableService] instead of an object.
	OfOrderNewResponseObject2UnavailableServices []OrderNewResponseObject2UnavailableService `json:",inline"`
	JSON                                         struct {
		OfOrderNewResponseObjectUnavailableServices  respjson.Field
		OfOrderNewResponseObject2UnavailableServices respjson.Field
		raw                                          string
	} `json:"-"`
}

func (r *OrderNewResponseUnionUnavailableServices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObject struct {
	OrderID     string `json:"orderId" api:"required"`
	OrderNumber string `json:"orderNumber" api:"required"`
	// Any of true.
	Success             bool                                       `json:"success" api:"required"`
	HostedInvoiceURL    string                                     `json:"hostedInvoiceUrl" format:"uri"`
	Message             string                                     `json:"message"`
	PartialSuccess      bool                                       `json:"partialSuccess"`
	SelfPay             bool                                       `json:"selfPay"`
	UnavailableServices []OrderNewResponseObjectUnavailableService `json:"unavailableServices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID             respjson.Field
		OrderNumber         respjson.Field
		Success             respjson.Field
		HostedInvoiceURL    respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		SelfPay             respjson.Field
		UnavailableServices respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderNewResponseObject) RawJSON() string { return r.JSON.raw }
func (r *OrderNewResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObjectUnavailableService struct {
	Reason      string `json:"reason" api:"required"`
	ServiceID   string `json:"serviceId" api:"required"`
	ServiceName string `json:"serviceName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ServiceID   respjson.Field
		ServiceName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderNewResponseObjectUnavailableService) RawJSON() string { return r.JSON.raw }
func (r *OrderNewResponseObjectUnavailableService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObject2 struct {
	OrderResults []OrderNewResponseObject2OrderResult `json:"orderResults" api:"required"`
	// Any of "split".
	Status string `json:"status" api:"required"`
	// Any of true.
	Success             bool                                        `json:"success" api:"required"`
	Message             string                                      `json:"message"`
	PartialSuccess      bool                                        `json:"partialSuccess"`
	UnavailableServices []OrderNewResponseObject2UnavailableService `json:"unavailableServices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderResults        respjson.Field
		Status              respjson.Field
		Success             respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		UnavailableServices respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderNewResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *OrderNewResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObject2OrderResult struct {
	OrderID     string `json:"orderId" api:"required"`
	OrderNumber string `json:"orderNumber" api:"required"`
	ProviderID  string `json:"providerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID     respjson.Field
		OrderNumber respjson.Field
		ProviderID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderNewResponseObject2OrderResult) RawJSON() string { return r.JSON.raw }
func (r *OrderNewResponseObject2OrderResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObject2UnavailableService struct {
	Reason      string `json:"reason" api:"required"`
	ServiceID   string `json:"serviceId" api:"required"`
	ServiceName string `json:"serviceName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ServiceID   respjson.Field
		ServiceName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderNewResponseObject2UnavailableService) RawJSON() string { return r.JSON.raw }
func (r *OrderNewResponseObject2UnavailableService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderGetResponse struct {
	OrderID     string `json:"orderId"`
	OrderNumber string `json:"orderNumber"`
	Status      string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID     respjson.Field
		OrderNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderUpdateResponse struct {
	Message     string `json:"message" api:"required"`
	OrderID     string `json:"orderId" api:"required"`
	OrderNumber string `json:"orderNumber" api:"required"`
	// Any of true.
	Success       bool     `json:"success" api:"required"`
	UpdatedFields []string `json:"updatedFields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message       respjson.Field
		OrderID       respjson.Field
		OrderNumber   respjson.Field
		Success       respjson.Field
		UpdatedFields respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderGetResultsResponse struct {
	Meta     OrderGetResultsResponseMeta      `json:"meta" api:"required"`
	Services []OrderGetResultsResponseService `json:"services" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		Services    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderGetResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderGetResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderGetResultsResponseMeta struct {
	OrderID       string  `json:"orderId" api:"required"`
	Page          float64 `json:"page" api:"required"`
	PageSize      float64 `json:"pageSize" api:"required"`
	Returned      float64 `json:"returned" api:"required"`
	TotalServices float64 `json:"totalServices" api:"required"`
	EmployeeID    string  `json:"employeeId"`
	OrderNumber   string  `json:"orderNumber"`
	ProviderID    string  `json:"providerId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID       respjson.Field
		Page          respjson.Field
		PageSize      respjson.Field
		Returned      respjson.Field
		TotalServices respjson.Field
		EmployeeID    respjson.Field
		OrderNumber   respjson.Field
		ProviderID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderGetResultsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OrderGetResultsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderGetResultsResponseService struct {
	ServiceID         string    `json:"serviceId" api:"required"`
	Status            string    `json:"status" api:"required"`
	AltTxt            string    `json:"altTxt"`
	CompletedDatetime time.Time `json:"completed_datetime" format:"date-time"`
	Contacts          []string  `json:"contacts"`
	DrawnDatetime     time.Time `json:"drawn_datetime" format:"date-time"`
	FileIDs           []string  `json:"fileIds"`
	Message           string    `json:"message"`
	Result            string    `json:"result"`
	ResultsPosted     time.Time `json:"resultsPosted" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServiceID         respjson.Field
		Status            respjson.Field
		AltTxt            respjson.Field
		CompletedDatetime respjson.Field
		Contacts          respjson.Field
		DrawnDatetime     respjson.Field
		FileIDs           respjson.Field
		Message           respjson.Field
		Result            respjson.Field
		ResultsPosted     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderGetResultsResponseService) RawJSON() string { return r.JSON.raw }
func (r *OrderGetResultsResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderScheduleAppointmentResponse struct {
	Message string `json:"message" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderScheduleAppointmentResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderScheduleAppointmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OrderSendForEmployeeResponseUnion contains all possible properties and values
// from [OrderSendForEmployeeResponseObject],
// [OrderSendForEmployeeResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OrderSendForEmployeeResponseUnion struct {
	// This field is from variant [OrderSendForEmployeeResponseObject].
	OrderID string `json:"orderId"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	OrderNumber    string `json:"orderNumber"`
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	PartialSuccess bool   `json:"partialSuccess"`
	// This field is a union of
	// [[]OrderSendForEmployeeResponseObjectUnavailableService],
	// [[]OrderSendForEmployeeResponseObject2UnavailableService]
	UnavailableServices OrderSendForEmployeeResponseUnionUnavailableServices `json:"unavailableServices"`
	// This field is from variant [OrderSendForEmployeeResponseObject2].
	OrderResults []OrderSendForEmployeeResponseObject2OrderResult `json:"orderResults"`
	// This field is from variant [OrderSendForEmployeeResponseObject2].
	Status string `json:"status"`
	JSON   struct {
		OrderID             respjson.Field
		OrderNumber         respjson.Field
		Success             respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		UnavailableServices respjson.Field
		OrderResults        respjson.Field
		Status              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u OrderSendForEmployeeResponseUnion) AsOrderSendForEmployeeResponseObject() (v OrderSendForEmployeeResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderSendForEmployeeResponseUnion) AsOrderSendForEmployeeResponseObject2() (v OrderSendForEmployeeResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OrderSendForEmployeeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OrderSendForEmployeeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OrderSendForEmployeeResponseUnionUnavailableServices is an implicit subunion of
// [OrderSendForEmployeeResponseUnion].
// OrderSendForEmployeeResponseUnionUnavailableServices provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OrderSendForEmployeeResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOrderSendForEmployeeResponseObjectUnavailableServices
// OfOrderSendForEmployeeResponseObject2UnavailableServices]
type OrderSendForEmployeeResponseUnionUnavailableServices struct {
	// This field will be present if the value is a
	// [[]OrderSendForEmployeeResponseObjectUnavailableService] instead of an object.
	OfOrderSendForEmployeeResponseObjectUnavailableServices []OrderSendForEmployeeResponseObjectUnavailableService `json:",inline"`
	// This field will be present if the value is a
	// [[]OrderSendForEmployeeResponseObject2UnavailableService] instead of an object.
	OfOrderSendForEmployeeResponseObject2UnavailableServices []OrderSendForEmployeeResponseObject2UnavailableService `json:",inline"`
	JSON                                                     struct {
		OfOrderSendForEmployeeResponseObjectUnavailableServices  respjson.Field
		OfOrderSendForEmployeeResponseObject2UnavailableServices respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (r *OrderSendForEmployeeResponseUnionUnavailableServices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObject struct {
	OrderID     string `json:"orderId" api:"required"`
	OrderNumber string `json:"orderNumber" api:"required"`
	// Any of true.
	Success bool   `json:"success" api:"required"`
	Message string `json:"message"`
	// True when some services were unavailable but order was still created
	PartialSuccess bool `json:"partialSuccess"`
	// Services that could not be included in the order
	UnavailableServices []OrderSendForEmployeeResponseObjectUnavailableService `json:"unavailableServices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID             respjson.Field
		OrderNumber         respjson.Field
		Success             respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		UnavailableServices respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderSendForEmployeeResponseObject) RawJSON() string { return r.JSON.raw }
func (r *OrderSendForEmployeeResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObjectUnavailableService struct {
	// Why the service was unavailable
	Reason      string `json:"reason" api:"required"`
	ServiceID   string `json:"serviceId" api:"required"`
	ServiceName string `json:"serviceName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ServiceID   respjson.Field
		ServiceName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderSendForEmployeeResponseObjectUnavailableService) RawJSON() string { return r.JSON.raw }
func (r *OrderSendForEmployeeResponseObjectUnavailableService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObject2 struct {
	OrderResults []OrderSendForEmployeeResponseObject2OrderResult `json:"orderResults" api:"required"`
	// Any of "split".
	Status string `json:"status" api:"required"`
	// Any of true.
	Success bool   `json:"success" api:"required"`
	Message string `json:"message"`
	// True when some services were unavailable but orders were still created
	PartialSuccess bool `json:"partialSuccess"`
	// Services that could not be included in any order
	UnavailableServices []OrderSendForEmployeeResponseObject2UnavailableService `json:"unavailableServices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderResults        respjson.Field
		Status              respjson.Field
		Success             respjson.Field
		Message             respjson.Field
		PartialSuccess      respjson.Field
		UnavailableServices respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderSendForEmployeeResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *OrderSendForEmployeeResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObject2OrderResult struct {
	OrderID     string `json:"orderId" api:"required"`
	OrderNumber string `json:"orderNumber" api:"required"`
	ProviderID  string `json:"providerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID     respjson.Field
		OrderNumber respjson.Field
		ProviderID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderSendForEmployeeResponseObject2OrderResult) RawJSON() string { return r.JSON.raw }
func (r *OrderSendForEmployeeResponseObject2OrderResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObject2UnavailableService struct {
	// Why the service was unavailable
	Reason      string `json:"reason" api:"required"`
	ServiceID   string `json:"serviceId" api:"required"`
	ServiceName string `json:"serviceName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ServiceID   respjson.Field
		ServiceName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderSendForEmployeeResponseObject2UnavailableService) RawJSON() string { return r.JSON.raw }
func (r *OrderSendForEmployeeResponseObject2UnavailableService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderUploadResultsResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderUploadResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderUploadResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *OrderNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfOrderNewsBodyObject2 *OrderNewParamsBodyObject2 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfOrderNewsBodyObject3 *OrderNewParamsBodyObject3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfOrderNewsBodyObject4 *OrderNewParamsBodyObject4 `json:",inline"`

	paramObj
}

func (u OrderNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfOrderNewsBodyObject2, u.OfOrderNewsBodyObject3, u.OfOrderNewsBodyObject4)
}
func (r *OrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PaymentMethod, Person, ProviderID, Services are required.
type OrderNewParamsBodyObject struct {
	// Any of "self-pay", "employer-sponsored".
	PaymentMethod   string                            `json:"paymentMethod,omitzero" api:"required"`
	Person          OrderNewParamsBodyObjectPerson    `json:"person,omitzero" api:"required"`
	ProviderID      string                            `json:"providerId" api:"required"`
	Services        []OrderNewParamsBodyObjectService `json:"services,omitzero" api:"required"`
	ID              param.Opt[string]                 `json:"_id,omitzero"`
	BrandID         param.Opt[string]                 `json:"brandId,omitzero"`
	DueDate         param.Opt[time.Time]              `json:"dueDate,omitzero" format:"date-time"`
	EmployeeID      param.Opt[string]                 `json:"employeeId,omitzero"`
	EmployerID      param.Opt[string]                 `json:"employerId,omitzero"`
	ProviderCreated param.Opt[bool]                   `json:"providerCreated,omitzero"`
	ReCaptchaToken  param.Opt[string]                 `json:"reCaptchaToken,omitzero"`
	TokenID         param.Opt[string]                 `json:"tokenId,omitzero"`
	DueDates        []time.Time                       `json:"dueDates,omitzero" format:"date-time"`
	EmployeeIDs     []string                          `json:"employeeIds,omitzero"`
	// Optional arbitrary metadata (<=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
	// Order priority level
	//
	// Any of "normal", "high".
	Priority     string                                `json:"priority,omitzero"`
	ProvidersIDs []OrderNewParamsBodyObjectProvidersID `json:"providersIds,omitzero"`
	Quantities   map[string]int64                      `json:"quantities,omitzero"`
	ServicesIDs  []string                              `json:"servicesIds,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject](
		"paymentMethod", "self-pay", "employer-sponsored",
	)
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject](
		"priority", "normal", "high",
	)
}

// The properties City, Dob, Email, FirstName, LastName, Phone, State, Street,
// Zipcode are required.
type OrderNewParamsBodyObjectPerson struct {
	City string `json:"city" api:"required"`
	// Date of birth in YYYY-MM-DD format
	Dob       string `json:"dob" api:"required"`
	Email     string `json:"email" api:"required"`
	FirstName string `json:"firstName" api:"required"`
	LastName  string `json:"lastName" api:"required"`
	Phone     string `json:"phone" api:"required"`
	State     string `json:"state" api:"required"`
	Street    string `json:"street" api:"required"`
	// US ZIP code in 12345 or 12345-6789 format
	Zipcode string            `json:"zipcode" api:"required"`
	Country param.Opt[string] `json:"country,omitzero"`
	County  param.Opt[string] `json:"county,omitzero"`
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObjectPerson) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObjectPerson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObjectPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Quantity are required.
type OrderNewParamsBodyObjectService struct {
	ID         string          `json:"_id" api:"required"`
	Quantity   int64           `json:"quantity" api:"required"`
	AutoAccept param.Opt[bool] `json:"autoAccept,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObjectService) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObjectService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObjectService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProviderID is required.
type OrderNewParamsBodyObjectProvidersID struct {
	ProviderID string            `json:"providerId" api:"required"`
	ServiceID  param.Opt[string] `json:"serviceId,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObjectProvidersID) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObjectProvidersID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObjectProvidersID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EmployeeID, EmployerID, Services are required.
type OrderNewParamsBodyObject2 struct {
	EmployeeID      string                             `json:"employeeId" api:"required"`
	EmployerID      string                             `json:"employerId" api:"required"`
	Services        []OrderNewParamsBodyObject2Service `json:"services,omitzero" api:"required"`
	ID              param.Opt[string]                  `json:"_id,omitzero"`
	BrandID         param.Opt[string]                  `json:"brandId,omitzero"`
	DueDate         param.Opt[time.Time]               `json:"dueDate,omitzero" format:"date-time"`
	ProviderCreated param.Opt[bool]                    `json:"providerCreated,omitzero"`
	ProviderID      param.Opt[string]                  `json:"providerId,omitzero"`
	ReCaptchaToken  param.Opt[string]                  `json:"reCaptchaToken,omitzero"`
	TokenID         param.Opt[string]                  `json:"tokenId,omitzero"`
	DueDates        []time.Time                        `json:"dueDates,omitzero" format:"date-time"`
	EmployeeIDs     []string                           `json:"employeeIds,omitzero"`
	// Optional arbitrary metadata (<=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
	// Any of "self-pay", "employer-sponsored".
	PaymentMethod string                          `json:"paymentMethod,omitzero"`
	Person        OrderNewParamsBodyObject2Person `json:"person,omitzero"`
	// Order priority level
	//
	// Any of "normal", "high".
	Priority     string                                 `json:"priority,omitzero"`
	ProvidersIDs []OrderNewParamsBodyObject2ProvidersID `json:"providersIds,omitzero"`
	Quantities   map[string]int64                       `json:"quantities,omitzero"`
	ServicesIDs  []string                               `json:"servicesIds,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject2) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject2](
		"paymentMethod", "self-pay", "employer-sponsored",
	)
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject2](
		"priority", "normal", "high",
	)
}

// The properties ID, Quantity are required.
type OrderNewParamsBodyObject2Service struct {
	ID         string          `json:"_id" api:"required"`
	Quantity   int64           `json:"quantity" api:"required"`
	AutoAccept param.Opt[bool] `json:"autoAccept,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject2Service) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject2Service
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject2Service) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, Dob, Email, FirstName, LastName, Phone, State, Street,
// Zipcode are required.
type OrderNewParamsBodyObject2Person struct {
	City string `json:"city" api:"required"`
	// Date of birth in YYYY-MM-DD format
	Dob       string `json:"dob" api:"required"`
	Email     string `json:"email" api:"required"`
	FirstName string `json:"firstName" api:"required"`
	LastName  string `json:"lastName" api:"required"`
	Phone     string `json:"phone" api:"required"`
	State     string `json:"state" api:"required"`
	Street    string `json:"street" api:"required"`
	// US ZIP code in 12345 or 12345-6789 format
	Zipcode string            `json:"zipcode" api:"required"`
	Country param.Opt[string] `json:"country,omitzero"`
	County  param.Opt[string] `json:"county,omitzero"`
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject2Person) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject2Person
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject2Person) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProviderID is required.
type OrderNewParamsBodyObject2ProvidersID struct {
	ProviderID string            `json:"providerId" api:"required"`
	ServiceID  param.Opt[string] `json:"serviceId,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject2ProvidersID) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject2ProvidersID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject2ProvidersID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EmployeeID, EmployerID, ProvidersIDs, ServicesIDs are required.
type OrderNewParamsBodyObject3 struct {
	EmployeeID      string                                 `json:"employeeId" api:"required"`
	EmployerID      string                                 `json:"employerId" api:"required"`
	ProvidersIDs    []OrderNewParamsBodyObject3ProvidersID `json:"providersIds,omitzero" api:"required"`
	ServicesIDs     []string                               `json:"servicesIds,omitzero" api:"required"`
	ID              param.Opt[string]                      `json:"_id,omitzero"`
	BrandID         param.Opt[string]                      `json:"brandId,omitzero"`
	DueDate         param.Opt[time.Time]                   `json:"dueDate,omitzero" format:"date-time"`
	ProviderCreated param.Opt[bool]                        `json:"providerCreated,omitzero"`
	ProviderID      param.Opt[string]                      `json:"providerId,omitzero"`
	ReCaptchaToken  param.Opt[string]                      `json:"reCaptchaToken,omitzero"`
	TokenID         param.Opt[string]                      `json:"tokenId,omitzero"`
	DueDates        []time.Time                            `json:"dueDates,omitzero" format:"date-time"`
	EmployeeIDs     []string                               `json:"employeeIds,omitzero"`
	// Optional arbitrary metadata (<=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
	// Any of "self-pay", "employer-sponsored".
	PaymentMethod string                          `json:"paymentMethod,omitzero"`
	Person        OrderNewParamsBodyObject3Person `json:"person,omitzero"`
	// Order priority level
	//
	// Any of "normal", "high".
	Priority   string                             `json:"priority,omitzero"`
	Quantities map[string]int64                   `json:"quantities,omitzero"`
	Services   []OrderNewParamsBodyObject3Service `json:"services,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject3) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject3](
		"paymentMethod", "self-pay", "employer-sponsored",
	)
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject3](
		"priority", "normal", "high",
	)
}

// The property ProviderID is required.
type OrderNewParamsBodyObject3ProvidersID struct {
	ProviderID string            `json:"providerId" api:"required"`
	ServiceID  param.Opt[string] `json:"serviceId,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject3ProvidersID) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject3ProvidersID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject3ProvidersID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, Dob, Email, FirstName, LastName, Phone, State, Street,
// Zipcode are required.
type OrderNewParamsBodyObject3Person struct {
	City string `json:"city" api:"required"`
	// Date of birth in YYYY-MM-DD format
	Dob       string `json:"dob" api:"required"`
	Email     string `json:"email" api:"required"`
	FirstName string `json:"firstName" api:"required"`
	LastName  string `json:"lastName" api:"required"`
	Phone     string `json:"phone" api:"required"`
	State     string `json:"state" api:"required"`
	Street    string `json:"street" api:"required"`
	// US ZIP code in 12345 or 12345-6789 format
	Zipcode string            `json:"zipcode" api:"required"`
	Country param.Opt[string] `json:"country,omitzero"`
	County  param.Opt[string] `json:"county,omitzero"`
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject3Person) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject3Person
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject3Person) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Quantity are required.
type OrderNewParamsBodyObject3Service struct {
	ID         string          `json:"_id" api:"required"`
	Quantity   int64           `json:"quantity" api:"required"`
	AutoAccept param.Opt[bool] `json:"autoAccept,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject3Service) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject3Service
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject3Service) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EmployeeIDs, EmployerID, ProvidersIDs, ServicesIDs are required.
type OrderNewParamsBodyObject4 struct {
	EmployeeIDs     []string                               `json:"employeeIds,omitzero" api:"required"`
	EmployerID      string                                 `json:"employerId" api:"required"`
	ProvidersIDs    []OrderNewParamsBodyObject4ProvidersID `json:"providersIds,omitzero" api:"required"`
	ServicesIDs     []string                               `json:"servicesIds,omitzero" api:"required"`
	ID              param.Opt[string]                      `json:"_id,omitzero"`
	BrandID         param.Opt[string]                      `json:"brandId,omitzero"`
	DueDate         param.Opt[time.Time]                   `json:"dueDate,omitzero" format:"date-time"`
	EmployeeID      param.Opt[string]                      `json:"employeeId,omitzero"`
	ProviderCreated param.Opt[bool]                        `json:"providerCreated,omitzero"`
	ProviderID      param.Opt[string]                      `json:"providerId,omitzero"`
	ReCaptchaToken  param.Opt[string]                      `json:"reCaptchaToken,omitzero"`
	TokenID         param.Opt[string]                      `json:"tokenId,omitzero"`
	DueDates        []time.Time                            `json:"dueDates,omitzero" format:"date-time"`
	// Optional arbitrary metadata (<=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
	// Any of "self-pay", "employer-sponsored".
	PaymentMethod string                          `json:"paymentMethod,omitzero"`
	Person        OrderNewParamsBodyObject4Person `json:"person,omitzero"`
	// Order priority level
	//
	// Any of "normal", "high".
	Priority   string                             `json:"priority,omitzero"`
	Quantities map[string]int64                   `json:"quantities,omitzero"`
	Services   []OrderNewParamsBodyObject4Service `json:"services,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject4) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject4](
		"paymentMethod", "self-pay", "employer-sponsored",
	)
	apijson.RegisterFieldValidator[OrderNewParamsBodyObject4](
		"priority", "normal", "high",
	)
}

// The property ProviderID is required.
type OrderNewParamsBodyObject4ProvidersID struct {
	ProviderID string            `json:"providerId" api:"required"`
	ServiceID  param.Opt[string] `json:"serviceId,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject4ProvidersID) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject4ProvidersID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject4ProvidersID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, Dob, Email, FirstName, LastName, Phone, State, Street,
// Zipcode are required.
type OrderNewParamsBodyObject4Person struct {
	City string `json:"city" api:"required"`
	// Date of birth in YYYY-MM-DD format
	Dob       string `json:"dob" api:"required"`
	Email     string `json:"email" api:"required"`
	FirstName string `json:"firstName" api:"required"`
	LastName  string `json:"lastName" api:"required"`
	Phone     string `json:"phone" api:"required"`
	State     string `json:"state" api:"required"`
	Street    string `json:"street" api:"required"`
	// US ZIP code in 12345 or 12345-6789 format
	Zipcode string            `json:"zipcode" api:"required"`
	Country param.Opt[string] `json:"country,omitzero"`
	County  param.Opt[string] `json:"county,omitzero"`
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject4Person) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject4Person
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject4Person) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Quantity are required.
type OrderNewParamsBodyObject4Service struct {
	ID         string          `json:"_id" api:"required"`
	Quantity   int64           `json:"quantity" api:"required"`
	AutoAccept param.Opt[bool] `json:"autoAccept,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyObject4Service) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyObject4Service
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyObject4Service) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderUpdateParams struct {
	// Order expiration date (ISO 8601 format). Set to null to remove the expiration
	// date.
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	// Arbitrary metadata to update on the order (non-indexed passthrough, <=10KB when
	// JSON stringified)
	Metadata map[string]any             `json:"metadata,omitzero"`
	Services []OrderUpdateParamsService `json:"services,omitzero"`
	// Any of "order_sent", "order_accepted", "order_refused", "employee_confirmed",
	// "order_fulfilled", "order_complete".
	Status OrderUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r OrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ServiceID is required.
type OrderUpdateParamsService struct {
	ServiceID string               `json:"serviceId" api:"required"`
	DueDate   param.Opt[time.Time] `json:"dueDate,omitzero" format:"date-time"`
	// Service-level expiration date
	ExpirationDate param.Opt[time.Time] `json:"expirationDate,omitzero" format:"date-time"`
	Results        map[string]any       `json:"results,omitzero"`
	// Any of "pending", "in_progress", "completed", "cancelled", "rejected".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r OrderUpdateParamsService) MarshalJSON() (data []byte, err error) {
	type shadow OrderUpdateParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUpdateParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderUpdateParamsService](
		"status", "pending", "in_progress", "completed", "cancelled", "rejected",
	)
}

type OrderUpdateParamsStatus string

const (
	OrderUpdateParamsStatusOrderSent         OrderUpdateParamsStatus = "order_sent"
	OrderUpdateParamsStatusOrderAccepted     OrderUpdateParamsStatus = "order_accepted"
	OrderUpdateParamsStatusOrderRefused      OrderUpdateParamsStatus = "order_refused"
	OrderUpdateParamsStatusEmployeeConfirmed OrderUpdateParamsStatus = "employee_confirmed"
	OrderUpdateParamsStatusOrderFulfilled    OrderUpdateParamsStatus = "order_fulfilled"
	OrderUpdateParamsStatusOrderComplete     OrderUpdateParamsStatus = "order_complete"
)

type OrderGetResultsParams struct {
	Page      param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64]     `query:"pageSize,omitzero" json:"-"`
	ServiceID param.Opt[string]    `query:"serviceId,omitzero" json:"-"`
	Since     param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	Status    param.Opt[string]    `query:"status,omitzero" json:"-"`
	Until     param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OrderGetResultsParams]'s query parameters as `url.Values`.
func (r OrderGetResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrderScheduleAppointmentParams struct {
	Appointment OrderScheduleAppointmentParamsAppointmentUnion `json:"appointment,omitzero" api:"required"`
	// Order access code for authorization
	OrderAccessCode string `json:"orderAccessCode" api:"required"`
	// Provider ID for authorization
	ProviderID param.Opt[string] `json:"providerId,omitzero"`
	paramObj
}

func (r OrderScheduleAppointmentParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderScheduleAppointmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderScheduleAppointmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OrderScheduleAppointmentParamsAppointmentUnion struct {
	OfOrderScheduleAppointmentsAppointmentObject  *OrderScheduleAppointmentParamsAppointmentObject  `json:",omitzero,inline"`
	OfOrderScheduleAppointmentsAppointmentObject2 *OrderScheduleAppointmentParamsAppointmentObject2 `json:",omitzero,inline"`
	paramUnion
}

func (u OrderScheduleAppointmentParamsAppointmentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOrderScheduleAppointmentsAppointmentObject, u.OfOrderScheduleAppointmentsAppointmentObject2)
}
func (u *OrderScheduleAppointmentParamsAppointmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OrderScheduleAppointmentParamsAppointmentUnion) asAny() any {
	if !param.IsOmitted(u.OfOrderScheduleAppointmentsAppointmentObject) {
		return u.OfOrderScheduleAppointmentsAppointmentObject
	} else if !param.IsOmitted(u.OfOrderScheduleAppointmentsAppointmentObject2) {
		return u.OfOrderScheduleAppointmentsAppointmentObject2
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetDate() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Date)
	} else if vt := u.OfOrderScheduleAppointmentsAppointmentObject2; vt != nil && vt.Date.Valid() {
		return &vt.Date.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetTime() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Time)
	} else if vt := u.OfOrderScheduleAppointmentsAppointmentObject2; vt != nil && vt.Time.Valid() {
		return &vt.Time.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetNotes() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil && vt.Notes.Valid() {
		return &vt.Notes.Value
	} else if vt := u.OfOrderScheduleAppointmentsAppointmentObject2; vt != nil && vt.Notes.Valid() {
		return &vt.Notes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetType() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOrderScheduleAppointmentsAppointmentObject2; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's DateTime property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetDateTime() *time.Time {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return &vt.DateTime
	} else if vt := u.OfOrderScheduleAppointmentsAppointmentObject2; vt != nil && vt.DateTime.Valid() {
		return &vt.DateTime.Value
	}
	return nil
}

// The properties Date, DateTime, Time are required.
type OrderScheduleAppointmentParamsAppointmentObject struct {
	// Required for appointment type
	Date string `json:"date" api:"required"`
	// Required for appointment type
	DateTime time.Time `json:"dateTime" api:"required" format:"date-time"`
	// Required for appointment type
	Time string `json:"time" api:"required"`
	// Optional for walkin type
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Any of "appointment".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r OrderScheduleAppointmentParamsAppointmentObject) MarshalJSON() (data []byte, err error) {
	type shadow OrderScheduleAppointmentParamsAppointmentObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderScheduleAppointmentParamsAppointmentObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderScheduleAppointmentParamsAppointmentObject](
		"type", "appointment",
	)
}

type OrderScheduleAppointmentParamsAppointmentObject2 struct {
	// Required for appointment type
	Date param.Opt[string] `json:"date,omitzero"`
	// Required for appointment type
	DateTime param.Opt[time.Time] `json:"dateTime,omitzero" format:"date-time"`
	// Optional for walkin type
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Required for appointment type
	Time param.Opt[string] `json:"time,omitzero"`
	// Any of "walkin".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r OrderScheduleAppointmentParamsAppointmentObject2) MarshalJSON() (data []byte, err error) {
	type shadow OrderScheduleAppointmentParamsAppointmentObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderScheduleAppointmentParamsAppointmentObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OrderScheduleAppointmentParamsAppointmentObject2](
		"type", "walkin",
	)
}

type OrderSendForEmployeeParams struct {
	// Employee ID to send order to
	EmployeeID string `json:"employeeId" api:"required"`
	// Employer ID sending the order
	EmployerID string `json:"employerId" api:"required"`
	// Array mapping each service (by index) to a provider; serviceId optional
	ProvidersIDs []OrderSendForEmployeeParamsProvidersID `json:"providersIds,omitzero" api:"required"`
	// Array of service IDs to include in the order
	ServicesIDs []string `json:"servicesIds,omitzero" api:"required"`
	LoginToken  string   `header:"login-token" api:"required" json:"-"`
	UserID      string   `header:"user-id" api:"required" json:"-"`
	// Brand ID for branded orders
	BrandID param.Opt[string] `json:"brandId,omitzero"`
	// Due date for the order (date or date-time ISO string)
	DueDate param.Opt[string] `json:"dueDate,omitzero"`
	// Expiration date for the order (date or date-time ISO string)
	ExpirationDate param.Opt[string] `json:"expirationDate,omitzero"`
	// Whether this order is being created by a provider (affects permission checking)
	ProviderCreated param.Opt[bool] `json:"providerCreated,omitzero"`
	// Single provider ID (shortcut when all services map to one provider)
	ProviderID param.Opt[string] `json:"providerId,omitzero"`
	// Array of due dates per service
	DueDates []string `json:"dueDates,omitzero"`
	// Optional arbitrary metadata to store on the order (non-indexed passthrough,
	// <=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
	// Order priority level
	//
	// Any of "normal", "high".
	Priority OrderSendForEmployeeParamsPriority `json:"priority,omitzero"`
	// Service ID to quantity mapping
	Quantities map[string]int64 `json:"quantities,omitzero"`
	paramObj
}

func (r OrderSendForEmployeeParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderSendForEmployeeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderSendForEmployeeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProviderID is required.
type OrderSendForEmployeeParamsProvidersID struct {
	ProviderID string            `json:"providerId" api:"required"`
	ServiceID  param.Opt[string] `json:"serviceId,omitzero"`
	paramObj
}

func (r OrderSendForEmployeeParamsProvidersID) MarshalJSON() (data []byte, err error) {
	type shadow OrderSendForEmployeeParamsProvidersID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderSendForEmployeeParamsProvidersID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order priority level
type OrderSendForEmployeeParamsPriority string

const (
	OrderSendForEmployeeParamsPriorityNormal OrderSendForEmployeeParamsPriority = "normal"
	OrderSendForEmployeeParamsPriorityHigh   OrderSendForEmployeeParamsPriority = "high"
)

type OrderUploadResultsParams struct {
	CaptchaToken    string `json:"captchaToken" api:"required"`
	OrderAccessCode string `json:"orderAccessCode" api:"required"`
	ServiceID       string `json:"serviceId" api:"required"`
	// Date of birth in YYYY-MM-DD format
	Dob      param.Opt[string]              `json:"dob,omitzero"`
	LastName param.Opt[string]              `json:"lastName,omitzero"`
	FileIDs  []string                       `json:"fileIds,omitzero"`
	Files    []OrderUploadResultsParamsFile `json:"files,omitzero"`
	paramObj
}

func (r OrderUploadResultsParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderUploadResultsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUploadResultsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Base64, Name, Type are required.
type OrderUploadResultsParamsFile struct {
	Base64 string `json:"base64" api:"required"`
	Name   string `json:"name" api:"required"`
	Type   string `json:"type" api:"required"`
	paramObj
}

func (r OrderUploadResultsParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow OrderUploadResultsParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUploadResultsParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
