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

// IntegrationService contains methods and other services that help with
// interacting with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationService] method instead.
type IntegrationService struct {
	Options []option.RequestOption
}

// NewIntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntegrationService(opts ...option.RequestOption) (r IntegrationService) {
	r = IntegrationService{}
	r.Options = opts
	return
}

// Returns the current brand integrations object keyed by integration name (empty
// object if none). Brand resolved via x-brand-id header.
func (r *IntegrationService) List(ctx context.Context, query IntegrationListParams, opts ...option.RequestOption) (res *IntegrationListResponse, err error) {
	if !param.IsOmitted(query.XBrandID) {
		opts = append(opts, option.WithHeader("x-brand-id", fmt.Sprintf("%v", query.XBrandID)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns true if the named integration is active for the given brand (brand
// resolved via x-brand-id header).
func (r *IntegrationService) CheckActive(ctx context.Context, name string, query IntegrationCheckActiveParams, opts ...option.RequestOption) (res *IntegrationCheckActiveResponse, err error) {
	if !param.IsOmitted(query.XBrandID) {
		opts = append(opts, option.WithHeader("x-brand-id", fmt.Sprintf("%v", query.XBrandID)))
	}
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("v1/integrations/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IntegrationListResponse struct {
	Integrations map[string]IntegrationListResponseIntegration `json:"integrations,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Integrations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationListResponseIntegration struct {
	Active      bool           `json:"active,required"`
	DisplayName string         `json:"displayName,required"`
	Config      map[string]any `json:"config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		DisplayName respjson.Field
		Config      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationCheckActiveResponse struct {
	Active bool `json:"active,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationCheckActiveResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationCheckActiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationListParams struct {
	XBrandID string `header:"x-brand-id,required" json:"-"`
	paramObj
}

type IntegrationCheckActiveParams struct {
	XBrandID string `header:"x-brand-id,required" json:"-"`
	paramObj
}
