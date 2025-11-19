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

// EmployerServiceBundleService contains methods and other services that help with
// interacting with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmployerServiceBundleService] method instead.
type EmployerServiceBundleService struct {
	Options []option.RequestOption
}

// NewEmployerServiceBundleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmployerServiceBundleService(opts ...option.RequestOption) (r EmployerServiceBundleService) {
	r = EmployerServiceBundleService{}
	r.Options = opts
	return
}

// Create Service Bundle
func (r *EmployerServiceBundleService) New(ctx context.Context, employerID string, body EmployerServiceBundleNewParams, opts ...option.RequestOption) (res *EmployerServiceBundleNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if employerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s/service-bundles", employerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Service Bundle
func (r *EmployerServiceBundleService) Get(ctx context.Context, id string, query EmployerServiceBundleGetParams, opts ...option.RequestOption) (res *EmployerServiceBundleGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EmployerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s/service-bundles/%s", query.EmployerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Service Bundle
func (r *EmployerServiceBundleService) Update(ctx context.Context, id string, params EmployerServiceBundleUpdateParams, opts ...option.RequestOption) (res *EmployerServiceBundleUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EmployerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s/service-bundles/%s", params.EmployerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List Service Bundles
func (r *EmployerServiceBundleService) List(ctx context.Context, employerID string, opts ...option.RequestOption) (res *[]EmployerServiceBundleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if employerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s/service-bundles", employerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Service Bundle
func (r *EmployerServiceBundleService) Delete(ctx context.Context, id string, body EmployerServiceBundleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EmployerID == "" {
		err = errors.New("missing required employerId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/employers/%s/service-bundles/%s", body.EmployerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type EmployerServiceBundleNewResponse struct {
	ID         string   `json:"_id,required"`
	BundleName string   `json:"bundleName,required"`
	EmployerID string   `json:"employerId,required"`
	ServiceIDs []string `json:"serviceIds,required"`
	CreatedAt  string   `json:"createdAt"`
	CreatedBy  string   `json:"createdBy"`
	Limit      float64  `json:"limit"`
	Occurrence string   `json:"occurrence"`
	Recurring  bool     `json:"recurring"`
	Roles      []string `json:"roles,nullable"`
	StartDate  string   `json:"startDate"`
	UpdatedAt  string   `json:"updatedAt"`
	UpdatedBy  string   `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BundleName  respjson.Field
		EmployerID  respjson.Field
		ServiceIDs  respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Limit       respjson.Field
		Occurrence  respjson.Field
		Recurring   respjson.Field
		Roles       respjson.Field
		StartDate   respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerServiceBundleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerServiceBundleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleGetResponse struct {
	ID         string   `json:"_id,required"`
	BundleName string   `json:"bundleName,required"`
	EmployerID string   `json:"employerId,required"`
	ServiceIDs []string `json:"serviceIds,required"`
	CreatedAt  string   `json:"createdAt"`
	CreatedBy  string   `json:"createdBy"`
	Limit      float64  `json:"limit"`
	Occurrence string   `json:"occurrence"`
	Recurring  bool     `json:"recurring"`
	Roles      []string `json:"roles,nullable"`
	StartDate  string   `json:"startDate"`
	UpdatedAt  string   `json:"updatedAt"`
	UpdatedBy  string   `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BundleName  respjson.Field
		EmployerID  respjson.Field
		ServiceIDs  respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Limit       respjson.Field
		Occurrence  respjson.Field
		Recurring   respjson.Field
		Roles       respjson.Field
		StartDate   respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerServiceBundleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerServiceBundleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleUpdateResponse struct {
	ID         string   `json:"_id,required"`
	BundleName string   `json:"bundleName,required"`
	EmployerID string   `json:"employerId,required"`
	ServiceIDs []string `json:"serviceIds,required"`
	CreatedAt  string   `json:"createdAt"`
	CreatedBy  string   `json:"createdBy"`
	Limit      float64  `json:"limit"`
	Occurrence string   `json:"occurrence"`
	Recurring  bool     `json:"recurring"`
	Roles      []string `json:"roles,nullable"`
	StartDate  string   `json:"startDate"`
	UpdatedAt  string   `json:"updatedAt"`
	UpdatedBy  string   `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BundleName  respjson.Field
		EmployerID  respjson.Field
		ServiceIDs  respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Limit       respjson.Field
		Occurrence  respjson.Field
		Recurring   respjson.Field
		Roles       respjson.Field
		StartDate   respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerServiceBundleUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerServiceBundleUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleListResponse struct {
	ID         string   `json:"_id,required"`
	BundleName string   `json:"bundleName,required"`
	EmployerID string   `json:"employerId,required"`
	ServiceIDs []string `json:"serviceIds,required"`
	CreatedAt  string   `json:"createdAt"`
	CreatedBy  string   `json:"createdBy"`
	Limit      float64  `json:"limit"`
	Occurrence string   `json:"occurrence"`
	Recurring  bool     `json:"recurring"`
	Roles      []string `json:"roles,nullable"`
	StartDate  string   `json:"startDate"`
	UpdatedAt  string   `json:"updatedAt"`
	UpdatedBy  string   `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BundleName  respjson.Field
		EmployerID  respjson.Field
		ServiceIDs  respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Limit       respjson.Field
		Occurrence  respjson.Field
		Recurring   respjson.Field
		Roles       respjson.Field
		StartDate   respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployerServiceBundleListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployerServiceBundleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleNewParams struct {
	BundleName string             `json:"bundleName,required"`
	ServiceIDs []string           `json:"serviceIds,omitzero,required"`
	ID         param.Opt[string]  `json:"_id,omitzero"`
	Limit      param.Opt[float64] `json:"limit,omitzero"`
	Occurrence param.Opt[string]  `json:"occurrence,omitzero"`
	Recurring  param.Opt[bool]    `json:"recurring,omitzero"`
	StartDate  param.Opt[string]  `json:"startDate,omitzero"`
	Roles      []string           `json:"roles,omitzero"`
	paramObj
}

func (r EmployerServiceBundleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployerServiceBundleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerServiceBundleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleGetParams struct {
	EmployerID string `path:"employerId,required" json:"-"`
	paramObj
}

type EmployerServiceBundleUpdateParams struct {
	EmployerID string             `path:"employerId,required" json:"-"`
	BundleName string             `json:"bundleName,required"`
	ServiceIDs []string           `json:"serviceIds,omitzero,required"`
	ID         param.Opt[string]  `json:"_id,omitzero"`
	Limit      param.Opt[float64] `json:"limit,omitzero"`
	Occurrence param.Opt[string]  `json:"occurrence,omitzero"`
	Recurring  param.Opt[bool]    `json:"recurring,omitzero"`
	StartDate  param.Opt[string]  `json:"startDate,omitzero"`
	Roles      []string           `json:"roles,omitzero"`
	paramObj
}

func (r EmployerServiceBundleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployerServiceBundleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployerServiceBundleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerServiceBundleDeleteParams struct {
	EmployerID string `path:"employerId,required" json:"-"`
	paramObj
}
