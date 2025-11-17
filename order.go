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
	return
}

// Retrieve details for a specific order
func (r *OrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *OrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update order details and associated order items. Allows updating order status,
// metadata, and modifying order item services.
func (r *OrderService) Update(ctx context.Context, orderID string, body OrderUpdateParams, opts ...option.RequestOption) (res *OrderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve results for an order. Supports filtering by serviceId, status, date
// window, and pagination.
func (r *OrderService) GetResults(ctx context.Context, orderID string, query OrderGetResultsParams, opts ...option.RequestOption) (res *OrderGetResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s/results", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Schedule an appointment or walk-in for an existing order. Sends HL7 SIU^S12
// message for appointment booking.
func (r *OrderService) ScheduleAppointment(ctx context.Context, orderID string, body OrderScheduleAppointmentParams, opts ...option.RequestOption) (res *OrderScheduleAppointmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s/schedule-appointment", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send an order for a specific employee. Requires API key, login token, and user
// ID. This endpoint specifically handles employer-to-employee order sending.
func (r *OrderService) SendForEmployee(ctx context.Context, params OrderSendForEmployeeParams, opts ...option.RequestOption) (res *OrderSendForEmployeeResponseUnion, err error) {
	if !param.IsOmitted(params.LoginToken) {
		opts = append(opts, option.WithHeader("login-token", fmt.Sprintf("%s", params.LoginToken)))
	}
	if !param.IsOmitted(params.UserID) {
		opts = append(opts, option.WithHeader("user-id", fmt.Sprintf("%s", params.UserID)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/orders/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the status of an existing order
func (r *OrderService) UpdateStatus(ctx context.Context, orderID string, body OrderUpdateStatusParams, opts ...option.RequestOption) (res *OrderUpdateStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s/status", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Upload test results for a specific order item. Supports both existing fileIds
// and base64 encoded files. Requires order access code and employee verification.
func (r *OrderService) UploadResults(ctx context.Context, orderID string, body OrderUploadResultsParams, opts ...option.RequestOption) (res *OrderUploadResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v1/orders/%s/upload-results", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// OrderNewResponseUnion contains all possible properties and values from
// [OrderNewResponseObject], [OrderNewResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OrderNewResponseUnion struct {
	// This field is from variant [OrderNewResponseObject].
	OrderID string `json:"orderId"`
	// This field is from variant [OrderNewResponseObject].
	OrderNumber string `json:"orderNumber"`
	// This field is from variant [OrderNewResponseObject].
	Success bool `json:"success"`
	// This field is from variant [OrderNewResponseObject].
	HostedInvoiceURL string `json:"hostedInvoiceUrl"`
	// This field is from variant [OrderNewResponseObject].
	Message string `json:"message"`
	// This field is from variant [OrderNewResponseObject].
	PartialSuccess bool `json:"partialSuccess"`
	// This field is from variant [OrderNewResponseObject].
	SelfPay bool `json:"selfPay"`
	// This field is from variant [OrderNewResponseObject].
	UnavailableServices []OrderNewResponseObjectUnavailableService `json:"unavailableServices"`
	// This field is from variant [OrderNewResponseObject].
	OrderResults []OrderNewResponseObjectOrderResult `json:"orderResults"`
	// This field is from variant [OrderNewResponseObject].
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

func (u OrderNewResponseUnion) AsVariant2() (v OrderNewResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OrderNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OrderNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewResponseObject struct {
	OrderID     string `json:"orderId,required"`
	OrderNumber string `json:"orderNumber,required"`
	// Any of true.
	Success             bool                                       `json:"success,required"`
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
	Reason      string `json:"reason,required"`
	ServiceID   string `json:"serviceId,required"`
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
	Message     string `json:"message,required"`
	OrderID     string `json:"orderId,required"`
	OrderNumber string `json:"orderNumber,required"`
	// Any of true.
	Success       bool     `json:"success,required"`
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
	Meta     OrderGetResultsResponseMeta      `json:"meta,required"`
	Services []OrderGetResultsResponseService `json:"services,required"`
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
	OrderID       string  `json:"orderId,required"`
	Page          float64 `json:"page,required"`
	PageSize      float64 `json:"pageSize,required"`
	Returned      float64 `json:"returned,required"`
	TotalServices float64 `json:"totalServices,required"`
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
	ServiceID         string    `json:"serviceId,required"`
	Status            string    `json:"status,required"`
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
	Message string `json:"message,required"`
	Success bool   `json:"success,required"`
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
// from [OrderSendForEmployeeResponseObject], [OrderSendForEmployeeResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OrderSendForEmployeeResponseUnion struct {
	// This field is from variant [OrderSendForEmployeeResponseObject].
	OrderID string `json:"orderId"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	OrderNumber string `json:"orderNumber"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	Success bool `json:"success"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	Message string `json:"message"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	PartialSuccess bool `json:"partialSuccess"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	UnavailableServices []OrderSendForEmployeeResponseObjectUnavailableService `json:"unavailableServices"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
	OrderResults []OrderSendForEmployeeResponseObjectOrderResult `json:"orderResults"`
	// This field is from variant [OrderSendForEmployeeResponseObject].
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

func (u OrderSendForEmployeeResponseUnion) AsVariant2() (v OrderSendForEmployeeResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OrderSendForEmployeeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OrderSendForEmployeeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderSendForEmployeeResponseObject struct {
	OrderID     string `json:"orderId,required"`
	OrderNumber string `json:"orderNumber,required"`
	// Any of true.
	Success bool   `json:"success,required"`
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
	Reason      string `json:"reason,required"`
	ServiceID   string `json:"serviceId,required"`
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

type OrderUpdateStatusResponse struct {
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
func (r OrderUpdateStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderUpdateStatusResponse) UnmarshalJSON(data []byte) error {
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
	OfOrderNewsBodyObject *OrderNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfVariant2 *OrderNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfVariant3 *OrderNewParamsBodyObject `json:",inline"`

	paramObj
}

func (u OrderNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfOrderNewsBodyObject, u.OfVariant2, u.OfVariant3)
}
func (r *OrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PaymentMethod, Person, ProviderID, Services are required.
type OrderNewParamsBodyObject struct {
	// Any of "self-pay", "employer-sponsored".
	PaymentMethod   string                            `json:"paymentMethod,omitzero,required"`
	Person          OrderNewParamsBodyObjectPerson    `json:"person,omitzero,required"`
	ProviderID      string                            `json:"providerId,required"`
	Services        []OrderNewParamsBodyObjectService `json:"services,omitzero,required"`
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
	Metadata     map[string]any                        `json:"metadata,omitzero"`
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
}

// The properties City, Dob, Email, FirstName, LastName, Phone, State, Street,
// Zipcode are required.
type OrderNewParamsBodyObjectPerson struct {
	City string `json:"city,required"`
	// Date of birth in YYYY-MM-DD format
	Dob       string `json:"dob,required"`
	Email     string `json:"email,required"`
	FirstName string `json:"firstName,required"`
	LastName  string `json:"lastName,required"`
	Phone     string `json:"phone,required"`
	State     string `json:"state,required"`
	Street    string `json:"street,required"`
	// US ZIP code in 12345 or 12345-6789 format
	Zipcode string            `json:"zipcode,required"`
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
	ID         string          `json:"_id,required"`
	Quantity   int64           `json:"quantity,required"`
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
	ProviderID string            `json:"providerId,required"`
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

type OrderUpdateParams struct {
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
	ServiceID string               `json:"serviceId,required"`
	DueDate   param.Opt[time.Time] `json:"dueDate,omitzero" format:"date-time"`
	Results   map[string]any       `json:"results,omitzero"`
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
	Appointment OrderScheduleAppointmentParamsAppointmentUnion `json:"appointment,omitzero,required"`
	// Order access code for authorization
	OrderAccessCode param.Opt[string] `json:"orderAccessCode,omitzero"`
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
	OfOrderScheduleAppointmentsAppointmentObject *OrderScheduleAppointmentParamsAppointmentObject `json:",omitzero,inline"`
	OfVariant2                                   *OrderScheduleAppointmentParamsAppointmentObject `json:",omitzero,inline"`
	paramUnion
}

func (u OrderScheduleAppointmentParamsAppointmentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOrderScheduleAppointmentsAppointmentObject, u.OfVariant2)
}
func (u *OrderScheduleAppointmentParamsAppointmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OrderScheduleAppointmentParamsAppointmentUnion) asAny() any {
	if !param.IsOmitted(u.OfOrderScheduleAppointmentsAppointmentObject) {
		return u.OfOrderScheduleAppointmentsAppointmentObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetDate() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Date)
	} else if vt := u.OfVariant2; vt != nil && vt.Date.Valid() {
		return &vt.Date.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetTime() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Time)
	} else if vt := u.OfVariant2; vt != nil && vt.Time.Valid() {
		return &vt.Time.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetNotes() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil && vt.Notes.Valid() {
		return &vt.Notes.Value
	} else if vt := u.OfVariant2; vt != nil && vt.Notes.Valid() {
		return &vt.Notes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetType() *string {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's DateTime property, if present.
func (u OrderScheduleAppointmentParamsAppointmentUnion) GetDateTime() *time.Time {
	if vt := u.OfOrderScheduleAppointmentsAppointmentObject; vt != nil {
		return &vt.DateTime
	} else if vt := u.OfVariant2; vt != nil && vt.DateTime.Valid() {
		return &vt.DateTime.Value
	}
	return nil
}

// The properties Date, DateTime, Time are required.
type OrderScheduleAppointmentParamsAppointmentObject struct {
	// Required for appointment type
	Date string `json:"date,required"`
	// Required for appointment type
	DateTime time.Time `json:"dateTime,required" format:"date-time"`
	// Required for appointment type
	Time string `json:"time,required"`
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

type OrderSendForEmployeeParams struct {
	// Employee ID to send order to
	EmployeeID string `json:"employeeId,required"`
	// Employer ID sending the order
	EmployerID string `json:"employerId,required"`
	// Array mapping each service (by index) to a provider; serviceId optional
	ProvidersIDs []OrderSendForEmployeeParamsProvidersID `json:"providersIds,omitzero,required"`
	// Array of service IDs to include in the order
	ServicesIDs []string `json:"servicesIds,omitzero,required"`
	LoginToken  string   `header:"login-token,required" json:"-"`
	UserID      string   `header:"user-id,required" json:"-"`
	// Brand ID for branded orders
	BrandID param.Opt[string] `json:"brandId,omitzero"`
	// Due date for the order (date or date-time ISO string)
	DueDate param.Opt[string] `json:"dueDate,omitzero"`
	// Whether this order is being created by a provider (affects permission checking)
	ProviderCreated param.Opt[bool] `json:"providerCreated,omitzero"`
	// Single provider ID (shortcut when all services map to one provider)
	ProviderID param.Opt[string] `json:"providerId,omitzero"`
	// Array of due dates per service
	DueDates []string `json:"dueDates,omitzero"`
	// Optional arbitrary metadata to store on the order (non-indexed passthrough,
	// <=10KB when JSON stringified)
	Metadata map[string]any `json:"metadata,omitzero"`
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
	ProviderID string            `json:"providerId,required"`
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

type OrderUpdateStatusParams struct {
	// Any of "order_sent", "order_accepted", "order_refused", "employee_confirmed",
	// "order_fulfilled", "order_complete".
	Status  OrderUpdateStatusParamsStatus `json:"status,omitzero,required"`
	Message param.Opt[string]             `json:"message,omitzero"`
	paramObj
}

func (r OrderUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderUpdateStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderUpdateStatusParamsStatus string

const (
	OrderUpdateStatusParamsStatusOrderSent         OrderUpdateStatusParamsStatus = "order_sent"
	OrderUpdateStatusParamsStatusOrderAccepted     OrderUpdateStatusParamsStatus = "order_accepted"
	OrderUpdateStatusParamsStatusOrderRefused      OrderUpdateStatusParamsStatus = "order_refused"
	OrderUpdateStatusParamsStatusEmployeeConfirmed OrderUpdateStatusParamsStatus = "employee_confirmed"
	OrderUpdateStatusParamsStatusOrderFulfilled    OrderUpdateStatusParamsStatus = "order_fulfilled"
	OrderUpdateStatusParamsStatusOrderComplete     OrderUpdateStatusParamsStatus = "order_complete"
)

type OrderUploadResultsParams struct {
	CaptchaToken    string `json:"captchaToken,required"`
	OrderAccessCode string `json:"orderAccessCode,required"`
	ServiceID       string `json:"serviceId,required"`
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
	Base64 string `json:"base64,required"`
	Name   string `json:"name,required"`
	Type   string `json:"type,required"`
	paramObj
}

func (r OrderUploadResultsParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow OrderUploadResultsParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUploadResultsParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
