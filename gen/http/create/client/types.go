// Code generated by goa v3.1.1, DO NOT EDIT.
//
// create HTTP client types
//
// Command:
// $ goa gen crud/design

package client

import (
	create "crud/gen/create"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "create" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ID of a person
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// UpdateRequestBody is the type of the "create" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Details of blog to be updated
	Name string `form:"name" json:"name" xml:"name"`
	// Comments to be updated
	Comments []string `form:"comments" json:"comments" xml:"comments"`
}

// CreateResponseBody is the type of the "create" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID of a person
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// ListResponseBody is the type of the "create" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredblogResponse

// StoredblogResponse is used to define fields on response body types.
type StoredblogResponse struct {
	// ID is the unique id of the blog.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "create" service.
func NewCreateRequestBody(p *create.Blog) *CreateRequestBody {
	body := &CreateRequestBody{
		ID:   p.ID,
		Name: p.Name,
	}
	if p.Comments != nil {
		body.Comments = make([]string, len(p.Comments))
		for i, val := range p.Comments {
			body.Comments[i] = val
		}
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "create" service.
func NewUpdateRequestBody(p *create.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name: p.Name,
	}
	if p.Comments != nil {
		body.Comments = make([]string, len(p.Comments))
		for i, val := range p.Comments {
			body.Comments[i] = val
		}
	}
	return body
}

// NewCreateBlogCreated builds a "create" service "create" endpoint result from
// a HTTP "Created" response.
func NewCreateBlogCreated(body *CreateResponseBody) *create.Blog {
	v := &create.Blog{
		ID:   body.ID,
		Name: body.Name,
	}
	if body.Comments != nil {
		v.Comments = make([]string, len(body.Comments))
		for i, val := range body.Comments {
			v.Comments[i] = val
		}
	}

	return v
}

// NewListStoredblogOK builds a "create" service "list" endpoint result from a
// HTTP "OK" response.
func NewListStoredblogOK(body []*StoredblogResponse) []*create.Storedblog {
	v := make([]*create.Storedblog, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredblogResponseToCreateStoredblog(val)
	}
	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if len(body.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.comments", body.Comments, len(body.Comments), 100, false))
	}
	return
}

// ValidateStoredblogResponse runs the validations defined on StoredblogResponse
func ValidateStoredblogResponse(body *StoredblogResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if len(body.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.comments", body.Comments, len(body.Comments), 100, false))
	}
	return
}
