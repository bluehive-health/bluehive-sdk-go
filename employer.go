// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

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
	Options        []option.RequestOption
	ServiceBundles EmployerServiceBundleService
}

// NewEmployerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmployerService(opts ...option.RequestOption) (r EmployerService) {
	r = EmployerService{}
	r.Options = opts
	r.ServiceBundles = NewEmployerServiceBundleService(opts...)
	return
}

// Create Employer
func (r *EmployerService) New(ctx context.Context, body EmployerNewParams, opts ...option.RequestOption) (res *EmployerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Employer
func (r *EmployerService) Get(ctx context.Context, employerID string, opts ...option.RequestOption) (res *EmployerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if employerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s", employerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Employers for Current User
func (r *EmployerService) List(ctx context.Context, query EmployerListParams, opts ...option.RequestOption) (res *[]EmployerListResponse, err error) {
	if !param.IsOmitted(query.LoginToken) {
		opts = append(opts, option.WithHeader("login-token", fmt.Sprintf("%v", query.LoginToken)))
	}
	if !param.IsOmitted(query.UserID) {
		opts = append(opts, option.WithHeader("user-id", fmt.Sprintf("%v", query.UserID)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/employers/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmployerNewResponse struct {
	ID              string           `json:"_id,required"`
	Address         map[string]any   `json:"address,required"`
	Email           string           `json:"email,required"`
	Name            string           `json:"name,required"`
	Phones          []map[string]any `json:"phones,required"`
	CreatedAt       string           `json:"createdAt"`
	CreatedBy       string           `json:"createdBy"`
	Demo            bool             `json:"demo"`
	EmployeeConsent bool             `json:"employeeConsent"`
	OnsiteClinic    bool             `json:"onsiteClinic"`
	Website         string           `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Address         respjson.Field
		Email           respjson.Field
		Name            respjson.Field
		Phones          respjson.Field
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

type EmployerGetResponse map[string]any

type EmployerListResponse map[string]any

type EmployerNewParams struct {
	Address         EmployerNewParamsAddress `json:"address,omitzero,required"`
	Email           string                   `json:"email,required" format:"email"`
	Name            string                   `json:"name,required"`
	Phones          []EmployerNewParamsPhone `json:"phones,omitzero,required"`
	Demo            param.Opt[bool]          `json:"demo,omitzero"`
	EmployeeConsent param.Opt[bool]          `json:"employeeConsent,omitzero"`
	OnsiteClinic    param.Opt[bool]          `json:"onsiteClinic,omitzero"`
	Website         param.Opt[string]        `json:"website,omitzero"`
	BillingAddress  map[string]any           `json:"billingAddress,omitzero"`
	Checkr          EmployerNewParamsCheckr  `json:"checkr,omitzero"`
	Metadata        map[string]any           `json:"metadata,omitzero"`
	paramObj
}

func (r EmployerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, State, Street1, ZipCode are required.
type EmployerNewParamsAddress struct {
	City    string            `json:"city,required"`
	State   string            `json:"state,required"`
	Street1 string            `json:"street1,required"`
	ZipCode string            `json:"zipCode,required"`
	Country param.Opt[string] `json:"country,omitzero"`
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
	Number  string            `json:"number,required"`
	Primary param.Opt[bool]   `json:"primary,omitzero"`
	Type    param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r EmployerNewParamsPhone) MarshalJSON() (data []byte, err error) {
	type shadow EmployerNewParamsPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerNewParamsPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type EmployerNewParamsCheckr struct {
	ID     string            `json:"id,required"`
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

type EmployerListParams struct {
	LoginToken string `header:"login-token,required" json:"-"`
	UserID     string `header:"user-id,required" json:"-"`
	paramObj
}
