// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bluehive

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/bluehive-go/internal/apijson"
	"github.com/stainless-sdks/bluehive-go/internal/requestconfig"
	"github.com/stainless-sdks/bluehive-go/option"
	"github.com/stainless-sdks/bluehive-go/packages/param"
	"github.com/stainless-sdks/bluehive-go/packages/respjson"
)

// FaxService contains methods and other services that help with interacting with
// the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFaxService] method instead.
type FaxService struct {
	Options []option.RequestOption
}

// NewFaxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFaxService(opts ...option.RequestOption) (r FaxService) {
	r = FaxService{}
	r.Options = opts
	return
}

// Get a list of available fax providers and their configuration status.
func (r *FaxService) ListProviders(ctx context.Context, opts ...option.RequestOption) (res *FaxListProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/fax/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the current status and details of a fax by its ID.
func (r *FaxService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *FaxGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/fax/status/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Send a fax document to a specified number using the configured fax provider.
func (r *FaxService) Send(ctx context.Context, body FaxSendParams, opts ...option.RequestOption) (res *FaxSendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/fax/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FaxListProvidersResponse struct {
	Providers []FaxListProvidersResponseProvider `json:"providers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Providers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxListProvidersResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxListProvidersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxListProvidersResponseProvider struct {
	// Whether the provider is properly configured
	Configured bool `json:"configured,required"`
	// Whether this is the default provider
	IsDefault bool `json:"isDefault,required"`
	// Provider name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Configured  respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxListProvidersResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *FaxListProvidersResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxGetStatusResponse struct {
	// Fax identifier
	ID string `json:"id,required"`
	// ISO timestamp when fax was created
	CreatedAt string `json:"createdAt,required"`
	// Sender fax number
	From string `json:"from,required"`
	// Provider used to send the fax
	Provider string `json:"provider,required"`
	// Current fax status
	//
	// Any of "queued", "dialing", "sending", "delivered", "failed", "cancelled",
	// "retrying".
	Status FaxGetStatusResponseStatus `json:"status,required"`
	// Recipient fax number
	To string `json:"to,required"`
	// ISO timestamp when status was last updated
	UpdatedAt string `json:"updatedAt,required"`
	// Cost of the fax
	Cost float64 `json:"cost"`
	// ISO timestamp when fax was delivered
	DeliveredAt string `json:"deliveredAt"`
	// Call duration in seconds
	Duration float64 `json:"duration"`
	// Error message if fax failed
	ErrorMessage string `json:"errorMessage"`
	// Number of pages in the fax
	PageCount float64 `json:"pageCount"`
	// Provider-specific additional data
	ProviderData any `json:"providerData"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		From         respjson.Field
		Provider     respjson.Field
		Status       respjson.Field
		To           respjson.Field
		UpdatedAt    respjson.Field
		Cost         respjson.Field
		DeliveredAt  respjson.Field
		Duration     respjson.Field
		ErrorMessage respjson.Field
		PageCount    respjson.Field
		ProviderData respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current fax status
type FaxGetStatusResponseStatus string

const (
	FaxGetStatusResponseStatusQueued    FaxGetStatusResponseStatus = "queued"
	FaxGetStatusResponseStatusDialing   FaxGetStatusResponseStatus = "dialing"
	FaxGetStatusResponseStatusSending   FaxGetStatusResponseStatus = "sending"
	FaxGetStatusResponseStatusDelivered FaxGetStatusResponseStatus = "delivered"
	FaxGetStatusResponseStatusFailed    FaxGetStatusResponseStatus = "failed"
	FaxGetStatusResponseStatusCancelled FaxGetStatusResponseStatus = "cancelled"
	FaxGetStatusResponseStatusRetrying  FaxGetStatusResponseStatus = "retrying"
)

type FaxSendResponse struct {
	// Unique fax identifier
	ID string `json:"id,required"`
	// ISO timestamp when fax was created
	CreatedAt string `json:"createdAt,required"`
	// Sender fax number
	From string `json:"from,required"`
	// Provider used to send the fax
	Provider string `json:"provider,required"`
	// Current fax status
	//
	// Any of "queued", "dialing", "sending", "delivered", "failed", "cancelled",
	// "retrying".
	Status FaxSendResponseStatus `json:"status,required"`
	// Recipient fax number
	To string `json:"to,required"`
	// Estimated delivery time (ISO timestamp)
	EstimatedDelivery string `json:"estimatedDelivery"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		From              respjson.Field
		Provider          respjson.Field
		Status            respjson.Field
		To                respjson.Field
		EstimatedDelivery respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current fax status
type FaxSendResponseStatus string

const (
	FaxSendResponseStatusQueued    FaxSendResponseStatus = "queued"
	FaxSendResponseStatusDialing   FaxSendResponseStatus = "dialing"
	FaxSendResponseStatusSending   FaxSendResponseStatus = "sending"
	FaxSendResponseStatusDelivered FaxSendResponseStatus = "delivered"
	FaxSendResponseStatusFailed    FaxSendResponseStatus = "failed"
	FaxSendResponseStatusCancelled FaxSendResponseStatus = "cancelled"
	FaxSendResponseStatusRetrying  FaxSendResponseStatus = "retrying"
)

type FaxSendParams struct {
	Document FaxSendParamsDocument `json:"document,omitzero,required"`
	// Recipient fax number (E.164 format preferred)
	To string `json:"to,required"`
	// Sender fax number (optional, uses default if not provided)
	From param.Opt[string] `json:"from,omitzero"`
	// Optional provider override (uses default if not specified)
	Provider param.Opt[string] `json:"provider,omitzero"`
	// Subject line for the fax
	Subject param.Opt[string] `json:"subject,omitzero"`
	paramObj
}

func (r FaxSendParams) MarshalJSON() (data []byte, err error) {
	type shadow FaxSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, ContentType are required.
type FaxSendParamsDocument struct {
	// Base64 encoded document content
	Content string `json:"content,required"`
	// MIME type of the document
	//
	// Any of "application/pdf", "image/tiff", "image/tif", "image/jpeg", "image/jpg",
	// "image/png", "text/plain".
	ContentType string `json:"contentType,omitzero,required"`
	// Optional filename for the document
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r FaxSendParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow FaxSendParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxSendParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FaxSendParamsDocument](
		"contentType", "application/pdf", "image/tiff", "image/tif", "image/jpeg", "image/jpg", "image/png", "text/plain",
	)
}
