// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bluehive-health/bluehive-sdk-go/internal/apijson"
	"github.com/bluehive-health/bluehive-sdk-go/internal/requestconfig"
	"github.com/bluehive-health/bluehive-sdk-go/option"
	"github.com/bluehive-health/bluehive-sdk-go/packages/param"
	"github.com/bluehive-health/bluehive-sdk-go/packages/respjson"
)

// EmployerService contains methods and other services that help with interacting
// with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmployerService] method instead.
type EmployerService struct {
	Options []option.RequestOption
}

// NewEmployerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmployerService(opts ...option.RequestOption) (r EmployerService) {
	r = EmployerService{}
	r.Options = opts
	return
}

// Create a new employer in the system.
func (r *EmployerService) New(ctx context.Context, body EmployerNewParams, opts ...option.RequestOption) (res *EmployerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/employers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an employer by their unique ID.
func (r *EmployerService) Get(ctx context.Context, employerID string, opts ...option.RequestOption) (res *EmployerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if employerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s", employerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmployerNewResponse struct {
	// Unique identifier for the employer
	ID      string                     `json:"_id,required"`
	Address EmployerNewResponseAddress `json:"address,required"`
	Email   string                     `json:"email,required"`
	// The name of the employer
	Name            string                            `json:"name,required"`
	Phones          []EmployerNewResponsePhone        `json:"phones,required"`
	BillingAddress  EmployerNewResponseBillingAddress `json:"billingAddress"`
	Checkr          EmployerNewResponseCheckr         `json:"checkr"`
	CreatedAt       time.Time                         `json:"createdAt" format:"date-time"`
	CreatedBy       string                            `json:"createdBy"`
	Demo            bool                              `json:"demo"`
	EmployeeConsent bool                              `json:"employeeConsent"`
	OnsiteClinic    bool                              `json:"onsiteClinic"`
	Website         string                            `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Address         respjson.Field
		Email           respjson.Field
		Name            respjson.Field
		Phones          respjson.Field
		BillingAddress  respjson.Field
		Checkr          respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		Demo            respjson.Field
		EmployeeConsent respjson.Field
		OnsiteClinic    respjson.Field
		Website         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerNewResponseAddress struct {
	City    string `json:"city,required"`
	State   string `json:"state,required"`
	Street1 string `json:"street1,required"`
	ZipCode string `json:"zipCode,required"`
	Country string `json:"country"`
	Street2 string `json:"street2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		ZipCode     respjson.Field
		Country     respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerNewResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployerNewResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerNewResponsePhone struct {
	Number  string `json:"number,required"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Primary     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerNewResponsePhone) RawJSON() string { return r.JSON.raw }
func (r *EmployerNewResponsePhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerNewResponseBillingAddress struct {
	City    string `json:"city,required"`
	State   string `json:"state,required"`
	Street1 string `json:"street1,required"`
	ZipCode string `json:"zipCode,required"`
	Country string `json:"country"`
	Street2 string `json:"street2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		ZipCode     respjson.Field
		Country     respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerNewResponseBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployerNewResponseBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerNewResponseCheckr struct {
	ID     string `json:"id,required"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerNewResponseCheckr) RawJSON() string { return r.JSON.raw }
func (r *EmployerNewResponseCheckr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerGetResponse struct {
	// Unique identifier for the employer
	ID      string                     `json:"_id,required"`
	Address EmployerGetResponseAddress `json:"address,required"`
	Email   string                     `json:"email,required"`
	// The name of the employer
	Name            string                            `json:"name,required"`
	Phones          []EmployerGetResponsePhone        `json:"phones,required"`
	BillingAddress  EmployerGetResponseBillingAddress `json:"billingAddress"`
	Checkr          EmployerGetResponseCheckr         `json:"checkr"`
	CreatedAt       time.Time                         `json:"createdAt" format:"date-time"`
	CreatedBy       string                            `json:"createdBy"`
	Demo            bool                              `json:"demo"`
	EmployeeConsent bool                              `json:"employeeConsent"`
	OnsiteClinic    bool                              `json:"onsiteClinic"`
	Website         string                            `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Address         respjson.Field
		Email           respjson.Field
		Name            respjson.Field
		Phones          respjson.Field
		BillingAddress  respjson.Field
		Checkr          respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		Demo            respjson.Field
		EmployeeConsent respjson.Field
		OnsiteClinic    respjson.Field
		Website         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerGetResponseAddress struct {
	City    string `json:"city,required"`
	State   string `json:"state,required"`
	Street1 string `json:"street1,required"`
	Country string `json:"country"`
	Street2 string `json:"street2"`
	ZipCode string `json:"zipCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		Country     respjson.Field
		Street2     respjson.Field
		ZipCode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerGetResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployerGetResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerGetResponsePhone struct {
	Number  string `json:"number,required"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Primary     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerGetResponsePhone) RawJSON() string { return r.JSON.raw }
func (r *EmployerGetResponsePhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerGetResponseBillingAddress struct {
	City    string `json:"city,required"`
	State   string `json:"state,required"`
	Street1 string `json:"street1,required"`
	Country string `json:"country"`
	Street2 string `json:"street2"`
	ZipCode string `json:"zipCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		Country     respjson.Field
		Street2     respjson.Field
		ZipCode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerGetResponseBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployerGetResponseBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerGetResponseCheckr struct {
	ID     string `json:"id,required"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerGetResponseCheckr) RawJSON() string { return r.JSON.raw }
func (r *EmployerGetResponseCheckr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerNewParams struct {
	// Primary address of the employer
	Address EmployerNewParamsAddress `json:"address,omitzero,required"`
	// Email address for the employer administrator
	Email string `json:"email,required" format:"email"`
	// The name of the employer
	Name string `json:"name,required"`
	// Phone numbers for the employer
	Phones []EmployerNewParamsPhone `json:"phones,omitzero,required"`
	// Whether this is a demo employer
	Demo param.Opt[bool] `json:"demo,omitzero"`
	// Whether employee consent is required
	EmployeeConsent param.Opt[bool] `json:"employeeConsent,omitzero"`
	// Whether the employer has an onsite clinic
	OnsiteClinic param.Opt[bool] `json:"onsiteClinic,omitzero"`
	// Website URL for the employer
	Website param.Opt[string] `json:"website,omitzero"`
	// Billing address of the employer (optional)
	BillingAddress EmployerNewParamsBillingAddress `json:"billingAddress,omitzero"`
	// Checkr information (excluding sensitive token)
	Checkr EmployerNewParamsCheckr `json:"checkr,omitzero"`
	// Additional metadata for the employer
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r EmployerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary address of the employer
//
// The properties City, State, Street1, ZipCode are required.
type EmployerNewParamsAddress struct {
	// City
	City string `json:"city,required"`
	// State or province
	State string `json:"state,required"`
	// Primary street address
	Street1 string `json:"street1,required"`
	// ZIP/postal code
	ZipCode string `json:"zipCode,required"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Secondary street address
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r EmployerNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Number is required.
type EmployerNewParamsPhone struct {
	// Phone number
	Number string `json:"number,required"`
	// Is this the primary phone number
	Primary param.Opt[bool] `json:"primary,omitzero"`
	// Phone type (e.g., office, mobile)
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r EmployerNewParamsPhone) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParamsPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParamsPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing address of the employer (optional)
//
// The properties City, State, Street1, ZipCode are required.
type EmployerNewParamsBillingAddress struct {
	// City
	City string `json:"city,required"`
	// State or province
	State string `json:"state,required"`
	// Primary street address
	Street1 string `json:"street1,required"`
	// ZIP/postal code
	ZipCode string `json:"zipCode,required"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Secondary street address
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r EmployerNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParamsBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParamsBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Checkr information (excluding sensitive token)
//
// The property ID is required.
type EmployerNewParamsCheckr struct {
	// Checkr Account ID
	ID string `json:"id,required"`
	// Checkr Account Status
	Status param.Opt[string] `json:"status,omitzero"`
	paramObj
}

func (r EmployerNewParamsCheckr) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParamsCheckr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParamsCheckr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
