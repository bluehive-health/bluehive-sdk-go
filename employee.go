// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombluehivehealthbluehivesdkgo

import (
	"context"
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

// EmployeeService contains methods and other services that help with interacting
// with the bluehive API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmployeeService] method instead.
type EmployeeService struct {
	Options []option.RequestOption
}

// NewEmployeeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmployeeService(opts ...option.RequestOption) (r EmployeeService) {
	r = EmployeeService{}
	r.Options = opts
	return
}

// Create a new employee in the system.
func (r *EmployeeService) New(ctx context.Context, body EmployeeNewParams, opts ...option.RequestOption) (res *EmployeeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employees"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve an employee by their unique ID.
func (r *EmployeeService) Get(ctx context.Context, employeeID string, opts ...option.RequestOption) (res *EmployeeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if employeeID == "" {
		err = errors.New("missing required employeeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/employees/%s", employeeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing employee in the system.
func (r *EmployeeService) Update(ctx context.Context, body EmployeeUpdateParams, opts ...option.RequestOption) (res *EmployeeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employees"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List all employees for a given employer with pagination.
func (r *EmployeeService) List(ctx context.Context, query EmployeeListParams, opts ...option.RequestOption) (res *EmployeeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employees"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an employee from the system. Cannot delete employees with existing
// orders.
func (r *EmployeeService) Delete(ctx context.Context, employeeID string, opts ...option.RequestOption) (res *EmployeeDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if employeeID == "" {
		err = errors.New("missing required employeeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/employees/%s", employeeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Link an employee to a user account with specified roles
func (r *EmployeeService) LinkUser(ctx context.Context, body EmployeeLinkUserParams, opts ...option.RequestOption) (res *EmployeeLinkUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employees/link-user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove the link between an employee and a user account
func (r *EmployeeService) UnlinkUser(ctx context.Context, body EmployeeUnlinkUserParams, opts ...option.RequestOption) (res *EmployeeUnlinkUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/employees/unlink-user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Employee created successfully
type EmployeeNewResponse struct {
	// ID of the created employee
	EmployeeID string `json:"employeeId" api:"required"`
	Message    string `json:"message" api:"required"`
	Success    bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmployeeID  respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee found successfully
type EmployeeGetResponse struct {
	// Employee details
	Employee EmployeeGetResponseEmployee `json:"employee" api:"required"`
	Message  string                      `json:"message" api:"required"`
	Success  bool                        `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Employee    respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee details
type EmployeeGetResponseEmployee struct {
	// Unique identifier
	ID string `json:"_id" api:"required"`
	// Email address
	Email string `json:"email" api:"required"`
	// ID of associated employer
	EmployerID string `json:"employer_id" api:"required"`
	// First name
	FirstName string `json:"firstName" api:"required"`
	// Last name
	LastName string `json:"lastName" api:"required"`
	// Account status
	//
	// Any of "Active", "Inactive".
	ActiveAccount string `json:"activeAccount"`
	// Employee address
	Address EmployeeGetResponseEmployeeAddress `json:"address"`
	// Brief description or bio
	Blurb string `json:"blurb"`
	// Creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// ID of user who created the employee
	CreatedBy string `json:"createdBy"`
	// List of department names
	Departments []string `json:"departments"`
	// Date of birth
	Dob string `json:"dob"`
	// Additional custom fields
	ExtendedFields []EmployeeGetResponseEmployeeExtendedField `json:"extendedFields"`
	// Contact phone numbers
	Phone []EmployeeGetResponseEmployeePhone `json:"phone"`
	// Job title
	Title string `json:"title"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// ID of user who last updated the employee
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Email          respjson.Field
		EmployerID     respjson.Field
		FirstName      respjson.Field
		LastName       respjson.Field
		ActiveAccount  respjson.Field
		Address        respjson.Field
		Blurb          respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Departments    respjson.Field
		Dob            respjson.Field
		ExtendedFields respjson.Field
		Phone          respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		UpdatedBy      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeGetResponseEmployee) RawJSON() string { return r.JSON.raw }
func (r *EmployeeGetResponseEmployee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee address
type EmployeeGetResponseEmployeeAddress struct {
	// City
	City string `json:"city" api:"required"`
	// Postal code
	PostalCode string `json:"postalCode" api:"required"`
	// State
	State string `json:"state" api:"required"`
	// Street address line 1
	Street1 string `json:"street1" api:"required"`
	// Country
	Country string `json:"country"`
	// County
	County string `json:"county"`
	// Street address line 2
	Street2 string `json:"street2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		Country     respjson.Field
		County      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeGetResponseEmployeeAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployeeGetResponseEmployeeAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeGetResponseEmployeeExtendedField struct {
	// Field name
	Name string `json:"name" api:"required"`
	// Field value
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeGetResponseEmployeeExtendedField) RawJSON() string { return r.JSON.raw }
func (r *EmployeeGetResponseEmployeeExtendedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeGetResponseEmployeePhone struct {
	// Phone number
	Number string `json:"number" api:"required"`
	// Type of phone number
	//
	// Any of "Cell", "Home", "Work", "Other".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeGetResponseEmployeePhone) RawJSON() string { return r.JSON.raw }
func (r *EmployeeGetResponseEmployeePhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee updated successfully
type EmployeeUpdateResponse struct {
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
func (r EmployeeUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employees retrieved successfully
type EmployeeListResponse struct {
	// List of employees
	Employees []EmployeeListResponseEmployee `json:"employees" api:"required"`
	Message   string                         `json:"message" api:"required"`
	Success   bool                           `json:"success" api:"required"`
	// Total number of employees returned
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Employees   respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee details
type EmployeeListResponseEmployee struct {
	// Unique identifier
	ID string `json:"_id" api:"required"`
	// Email address
	Email string `json:"email" api:"required"`
	// ID of associated employer
	EmployerID string `json:"employer_id" api:"required"`
	// First name
	FirstName string `json:"firstName" api:"required"`
	// Last name
	LastName string `json:"lastName" api:"required"`
	// Account status
	//
	// Any of "Active", "Inactive".
	ActiveAccount string `json:"activeAccount"`
	// Employee address
	Address EmployeeListResponseEmployeeAddress `json:"address"`
	// Brief description or bio
	Blurb string `json:"blurb"`
	// Creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// ID of user who created the employee
	CreatedBy string `json:"createdBy"`
	// List of department names
	Departments []string `json:"departments"`
	// Date of birth
	Dob string `json:"dob"`
	// Additional custom fields
	ExtendedFields []EmployeeListResponseEmployeeExtendedField `json:"extendedFields"`
	// Contact phone numbers
	Phone []EmployeeListResponseEmployeePhone `json:"phone"`
	// Job title
	Title string `json:"title"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// ID of user who last updated the employee
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Email          respjson.Field
		EmployerID     respjson.Field
		FirstName      respjson.Field
		LastName       respjson.Field
		ActiveAccount  respjson.Field
		Address        respjson.Field
		Blurb          respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Departments    respjson.Field
		Dob            respjson.Field
		ExtendedFields respjson.Field
		Phone          respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		UpdatedBy      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeListResponseEmployee) RawJSON() string { return r.JSON.raw }
func (r *EmployeeListResponseEmployee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee address
type EmployeeListResponseEmployeeAddress struct {
	// City
	City string `json:"city" api:"required"`
	// Postal code
	PostalCode string `json:"postalCode" api:"required"`
	// State
	State string `json:"state" api:"required"`
	// Street address line 1
	Street1 string `json:"street1" api:"required"`
	// Country
	Country string `json:"country"`
	// County
	County string `json:"county"`
	// Street address line 2
	Street2 string `json:"street2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street1     respjson.Field
		Country     respjson.Field
		County      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeListResponseEmployeeAddress) RawJSON() string { return r.JSON.raw }
func (r *EmployeeListResponseEmployeeAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeListResponseEmployeeExtendedField struct {
	// Field name
	Name string `json:"name" api:"required"`
	// Field value
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeListResponseEmployeeExtendedField) RawJSON() string { return r.JSON.raw }
func (r *EmployeeListResponseEmployeeExtendedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeListResponseEmployeePhone struct {
	// Phone number
	Number string `json:"number" api:"required"`
	// Type of phone number
	//
	// Any of "Cell", "Home", "Work", "Other".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeListResponseEmployeePhone) RawJSON() string { return r.JSON.raw }
func (r *EmployeeListResponseEmployeePhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee deleted successfully
type EmployeeDeleteResponse struct {
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
func (r EmployeeDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee linked successfully
type EmployeeLinkUserResponse struct {
	// ID of the created link
	LinkID  string `json:"linkId" api:"required"`
	Message string `json:"message" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkID      respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmployeeLinkUserResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeLinkUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee unlinked successfully
type EmployeeUnlinkUserResponse struct {
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
func (r EmployeeUnlinkUserResponse) RawJSON() string { return r.JSON.raw }
func (r *EmployeeUnlinkUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeNewParams struct {
	Email      string            `json:"email" api:"required" format:"email"`
	FirstName  string            `json:"firstName" api:"required"`
	LastName   string            `json:"lastName" api:"required"`
	Blurb      param.Opt[string] `json:"blurb,omitzero"`
	Dob        param.Opt[string] `json:"dob,omitzero"`
	EmployerID param.Opt[string] `json:"employer_id,omitzero"`
	Title      param.Opt[string] `json:"title,omitzero"`
	// Any of "Active", "Inactive".
	ActiveAccount  EmployeeNewParamsActiveAccount   `json:"activeAccount,omitzero"`
	Address        EmployeeNewParamsAddress         `json:"address,omitzero"`
	Departments    []string                         `json:"departments,omitzero"`
	ExtendedFields []EmployeeNewParamsExtendedField `json:"extendedFields,omitzero"`
	Phone          []EmployeeNewParamsPhone         `json:"phone,omitzero"`
	paramObj
}

func (r EmployeeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeNewParamsActiveAccount string

const (
	EmployeeNewParamsActiveAccountActive   EmployeeNewParamsActiveAccount = "Active"
	EmployeeNewParamsActiveAccountInactive EmployeeNewParamsActiveAccount = "Inactive"
)

// The properties City, PostalCode, State, Street1 are required.
type EmployeeNewParamsAddress struct {
	City       string            `json:"city" api:"required"`
	PostalCode string            `json:"postalCode" api:"required"`
	State      string            `json:"state" api:"required"`
	Street1    string            `json:"street1" api:"required"`
	Country    param.Opt[string] `json:"country,omitzero"`
	County     param.Opt[string] `json:"county,omitzero"`
	Street2    param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r EmployeeNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type EmployeeNewParamsExtendedField struct {
	Name  string `json:"name" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r EmployeeNewParamsExtendedField) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeNewParamsExtendedField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeNewParamsExtendedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Number, Type are required.
type EmployeeNewParamsPhone struct {
	Number string `json:"number" api:"required"`
	// Any of "Cell", "Home", "Work", "Other".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r EmployeeNewParamsPhone) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeNewParamsPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeNewParamsPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmployeeNewParamsPhone](
		"type", "Cell", "Home", "Work", "Other",
	)
}

type EmployeeUpdateParams struct {
	ID         string            `json:"_id" api:"required"`
	Blurb      param.Opt[string] `json:"blurb,omitzero"`
	Dob        param.Opt[string] `json:"dob,omitzero"`
	Email      param.Opt[string] `json:"email,omitzero" format:"email"`
	EmployerID param.Opt[string] `json:"employer_id,omitzero"`
	FirstName  param.Opt[string] `json:"firstName,omitzero"`
	LastName   param.Opt[string] `json:"lastName,omitzero"`
	Title      param.Opt[string] `json:"title,omitzero"`
	// Any of "Active", "Inactive".
	ActiveAccount  EmployeeUpdateParamsActiveAccount   `json:"activeAccount,omitzero"`
	Address        EmployeeUpdateParamsAddress         `json:"address,omitzero"`
	Departments    []string                            `json:"departments,omitzero"`
	ExtendedFields []EmployeeUpdateParamsExtendedField `json:"extendedFields,omitzero"`
	Phone          []EmployeeUpdateParamsPhone         `json:"phone,omitzero"`
	paramObj
}

func (r EmployeeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeUpdateParamsActiveAccount string

const (
	EmployeeUpdateParamsActiveAccountActive   EmployeeUpdateParamsActiveAccount = "Active"
	EmployeeUpdateParamsActiveAccountInactive EmployeeUpdateParamsActiveAccount = "Inactive"
)

// The properties City, PostalCode, State, Street1 are required.
type EmployeeUpdateParamsAddress struct {
	City       string            `json:"city" api:"required"`
	PostalCode string            `json:"postalCode" api:"required"`
	State      string            `json:"state" api:"required"`
	Street1    string            `json:"street1" api:"required"`
	Country    param.Opt[string] `json:"country,omitzero"`
	County     param.Opt[string] `json:"county,omitzero"`
	Street2    param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r EmployeeUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeUpdateParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeUpdateParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type EmployeeUpdateParamsExtendedField struct {
	Name  string `json:"name" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r EmployeeUpdateParamsExtendedField) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeUpdateParamsExtendedField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeUpdateParamsExtendedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Number, Type are required.
type EmployeeUpdateParamsPhone struct {
	Number string `json:"number" api:"required"`
	// Any of "Cell", "Home", "Work", "Other".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r EmployeeUpdateParamsPhone) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeUpdateParamsPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeUpdateParamsPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmployeeUpdateParamsPhone](
		"type", "Cell", "Home", "Work", "Other",
	)
}

type EmployeeListParams struct {
	// ID of the employer to list employees for
	EmployerID string `query:"employerId" api:"required" json:"-"`
	// Maximum number of employees to return (default: 50)
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Number of employees to skip (default: 0)
	Offset param.Opt[string] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmployeeListParams]'s query parameters as `url.Values`.
func (r EmployeeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmployeeLinkUserParams struct {
	EmployeeID string   `json:"employeeId" api:"required"`
	UserID     string   `json:"userId" api:"required"`
	Role       []string `json:"role,omitzero"`
	paramObj
}

func (r EmployeeLinkUserParams) MarshalJSON() (data []byte, err error) {
	type shadow EmployeeLinkUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmployeeLinkUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmployeeUnlinkUserParams struct {
	// ID of the employee to unlink
	EmployeeID string `query:"employeeId" api:"required" json:"-"`
	// ID of the user to unlink from
	UserID string `query:"userId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [EmployeeUnlinkUserParams]'s query parameters as
// `url.Values`.
func (r EmployeeUnlinkUserParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
