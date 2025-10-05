// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/bluehive-health/bluehive-sdk-go/internal/apijson"
	"github.com/bluehive-health/bluehive-sdk-go/internal/requestconfig"
	"github.com/bluehive-health/bluehive-sdk-go/option"
	"github.com/bluehive-health/bluehive-sdk-go/packages/param"
)

// Hl7Service contains methods and other services that help with interacting with
// the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHl7Service] method instead.
type Hl7Service struct {
	Options []option.RequestOption
}

// NewHl7Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHl7Service(opts ...option.RequestOption) (r Hl7Service) {
	r = Hl7Service{}
	r.Options = opts
	return
}

// Process incoming HL7 messages from EHR systems. Accepts JSON with "message"
// field, raw text/plain HL7 content, or form-encoded data.
func (r *Hl7Service) Process(ctx context.Context, body Hl7ProcessParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/hl7/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send lab results or documents via HL7
func (r *Hl7Service) SendResults(ctx context.Context, body Hl7SendResultsParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/hl7/results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Hl7ProcessParams struct {
	// Form field (legacy support)
	F param.Opt[string] `json:"f,omitzero"`
	// Interface identifier (legacy support)
	Interface param.Opt[string] `json:"interface,omitzero"`
	// Login password (legacy support)
	LoginPasswd param.Opt[string] `json:"login_passwd,omitzero"`
	// Login user (legacy support)
	LoginUser param.Opt[string] `json:"login_user,omitzero"`
	// HL7 message content - the primary way to send HL7 data
	Message param.Opt[string] `json:"message,omitzero"`
	// Base64 encoded HL7 message (legacy support)
	MessageB64 param.Opt[string] `json:"message_b64,omitzero"`
	paramObj
}

func (r Hl7ProcessParams) MarshalJSON() (data []byte, err error) {
	type shadow Hl7ProcessParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Hl7ProcessParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Hl7SendResultsParams struct {
	// Employee ID to send results for
	EmployeeID string `json:"employeeId,required"`
	// File containing the results
	File Hl7SendResultsParamsFile `json:"file,omitzero,required"`
	paramObj
}

func (r Hl7SendResultsParams) MarshalJSON() (data []byte, err error) {
	type shadow Hl7SendResultsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Hl7SendResultsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File containing the results
//
// The properties Base64, Name, Type are required.
type Hl7SendResultsParamsFile struct {
	// Base64 encoded file content
	Base64 string `json:"base64,required"`
	// File name
	Name string `json:"name,required"`
	// MIME type of the file
	Type string `json:"type,required"`
	paramObj
}

func (r Hl7SendResultsParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow Hl7SendResultsParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Hl7SendResultsParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
