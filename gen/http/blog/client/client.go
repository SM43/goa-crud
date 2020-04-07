// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog client HTTP transport
//
// Command:
// $ goa gen crud/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the blog service endpoint HTTP clients.
type Client struct {
	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Remove Doer is the HTTP client used to make requests to the remove endpoint.
	RemoveDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Oauth Doer is the HTTP client used to make requests to the oauth endpoint.
	OauthDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the blog service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateDoer:          doer,
		ListDoer:            doer,
		RemoveDoer:          doer,
		UpdateDoer:          doer,
		AddDoer:             doer,
		ShowDoer:            doer,
		OauthDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Create returns an endpoint that makes HTTP requests to the blog service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "create", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the blog service list
// server.
func (c *Client) List() goa.Endpoint {
	var (
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Remove returns an endpoint that makes HTTP requests to the blog service
// remove server.
func (c *Client) Remove() goa.Endpoint {
	var (
		decodeResponse = DecodeRemoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRemoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "remove", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the blog service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Add returns an endpoint that makes HTTP requests to the blog service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddRequest(c.encoder)
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the blog service show
// server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Oauth returns an endpoint that makes HTTP requests to the blog service oauth
// server.
func (c *Client) Oauth() goa.Endpoint {
	var (
		decodeResponse = DecodeOauthResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOauthRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OauthDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("blog", "oauth", err)
		}
		return decodeResponse(resp)
	}
}
