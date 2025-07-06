// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bluehive-health/bluehive-sdk-go/internal/apijson"
	"github.com/bluehive-health/bluehive-sdk-go/internal/apiquery"
	"github.com/bluehive-health/bluehive-sdk-go/internal/requestconfig"
	"github.com/bluehive-health/bluehive-sdk-go/option"
	"github.com/bluehive-health/bluehive-sdk-go/packages/param"
	"github.com/bluehive-health/bluehive-sdk-go/packages/respjson"
)

// ProviderService contains methods and other services that help with interacting
// with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r ProviderService) {
	r = ProviderService{}
	r.Options = opts
	return
}

// Search for healthcare providers by NPI number, name, or location proximity.
func (r *ProviderService) Lookup(ctx context.Context, query ProviderLookupParams, opts ...option.RequestOption) (res *ProviderLookupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/providers/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProviderLookupResponse struct {
	// Number of providers found
	Count float64 `json:"count,required"`
	// List of matching providers
	Providers []ProviderLookupResponseProvider `json:"providers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Providers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *ProviderLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderLookupResponseProvider struct {
	// Primary address line
	Address1 string `json:"address_1,required"`
	// Secondary address line (suite, unit, etc.)
	Address2 string `json:"address_2,required"`
	// City
	City string `json:"city,required"`
	// Country code
	Country string `json:"country,required"`
	// Distance in miles from the provided ZIP code
	Distance float64 `json:"distance,required"`
	// Fax number
	FaxNumber string `json:"fax_number,required"`
	// Provider first name
	Firstname string `json:"firstname,required"`
	// Provider last name or organization name
	Lastname string `json:"lastname,required"`
	// National Provider Identifier (NPI) number
	Npi string `json:"npi,required"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code,required"`
	// State or province code
	StateProvince string `json:"state_province,required"`
	// Work phone number
	WorkPhone string `json:"work_phone,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1      respjson.Field
		Address2      respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Distance      respjson.Field
		FaxNumber     respjson.Field
		Firstname     respjson.Field
		Lastname      respjson.Field
		Npi           respjson.Field
		PostalCode    respjson.Field
		StateProvince respjson.Field
		WorkPhone     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderLookupResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *ProviderLookupResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderLookupParams struct {
	// Provider first name
	Firstname param.Opt[string] `query:"firstname,omitzero" json:"-"`
	// Provider last name
	Lastname param.Opt[string] `query:"lastname,omitzero" json:"-"`
	// Provider NPI number
	Npi param.Opt[string] `query:"npi,omitzero" json:"-"`
	// ZIP code to filter results by proximity
	Zipcode param.Opt[string] `query:"zipcode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProviderLookupParams]'s query parameters as `url.Values`.
func (r ProviderLookupParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
