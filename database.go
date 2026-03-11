// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/bluehive-health/bluehive-sdk-go/internal/apijson"
	"github.com/bluehive-health/bluehive-sdk-go/internal/requestconfig"
	"github.com/bluehive-health/bluehive-sdk-go/option"
	"github.com/bluehive-health/bluehive-sdk-go/packages/respjson"
)

// DatabaseService contains methods and other services that help with interacting
// with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseService] method instead.
type DatabaseService struct {
	Options []option.RequestOption
}

// NewDatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatabaseService(opts ...option.RequestOption) (r DatabaseService) {
	r = DatabaseService{}
	r.Options = opts
	return
}

// Check MongoDB database connectivity and retrieve health statistics.
func (r *DatabaseService) CheckHealth(ctx context.Context, opts ...option.RequestOption) (res *DatabaseCheckHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/database/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DatabaseCheckHealthResponse struct {
	// Database health status
	//
	// Any of "ok", "error".
	Status DatabaseCheckHealthResponseStatus `json:"status" api:"required"`
	// Health check timestamp
	Timestamp string `json:"timestamp" api:"required"`
	// Database name (hidden in production)
	Database string `json:"database"`
	// Error message if status is error
	Error string `json:"error"`
	// Database statistics (not available in production)
	Stats DatabaseCheckHealthResponseStats `json:"stats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Database    respjson.Field
		Error       respjson.Field
		Stats       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatabaseCheckHealthResponse) RawJSON() string { return r.JSON.raw }
func (r *DatabaseCheckHealthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Database health status
type DatabaseCheckHealthResponseStatus string

const (
	DatabaseCheckHealthResponseStatusOk    DatabaseCheckHealthResponseStatus = "ok"
	DatabaseCheckHealthResponseStatusError DatabaseCheckHealthResponseStatus = "error"
)

// Database statistics (not available in production)
type DatabaseCheckHealthResponseStats struct {
	// Number of collections
	Collections float64 `json:"collections"`
	// Total data size in bytes
	DataSize float64 `json:"dataSize"`
	// Total number of documents
	Documents float64 `json:"documents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collections respjson.Field
		DataSize    respjson.Field
		Documents   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatabaseCheckHealthResponseStats) RawJSON() string { return r.JSON.raw }
func (r *DatabaseCheckHealthResponseStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
