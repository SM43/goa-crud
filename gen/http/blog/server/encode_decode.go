// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog HTTP server encoders and decoders
//
// Command:
// $ goa gen crud/design

package server

import (
	"context"
	blog "crud/gen/blog"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the blog
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*blog.Blog)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the blog create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateBlog(&body)

		return payload, nil
	}
}

// EncodeListResponse returns an encoder for responses returned by the blog
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*blog.Storedblog)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the blog
// remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the blog remove
// endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint32
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewRemovePayload(id)

		return payload, nil
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the blog
// update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the blog update
// endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint32

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdatePayload(&body, id)

		return payload, nil
	}
}

// EncodeAddResponse returns an encoder for responses returned by the blog add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*blog.NewComment)
		enc := encoder(ctx, w)
		body := NewAddResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the blog add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			id uint32

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddNewComment(&body, id)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the blog
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*blog.Blog)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the blog show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ShowRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateShowRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint32

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowBlog(&body, id)

		return payload, nil
	}
}

// EncodeOauthResponse returns an encoder for responses returned by the blog
// oauth endpoint.
func EncodeOauthResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// unmarshalCommentsRequestBodyToBlogComments builds a value of type
// *blog.Comments from a value of type *CommentsRequestBody.
func unmarshalCommentsRequestBodyToBlogComments(v *CommentsRequestBody) *blog.Comments {
	if v == nil {
		return nil
	}
	res := &blog.Comments{
		ID:       v.ID,
		Comments: v.Comments,
	}

	return res
}

// marshalBlogCommentsToCommentsResponseBody builds a value of type
// *CommentsResponseBody from a value of type *blog.Comments.
func marshalBlogCommentsToCommentsResponseBody(v *blog.Comments) *CommentsResponseBody {
	if v == nil {
		return nil
	}
	res := &CommentsResponseBody{
		ID:       v.ID,
		Comments: v.Comments,
	}

	return res
}

// marshalBlogStoredblogToStoredblogResponse builds a value of type
// *StoredblogResponse from a value of type *blog.Storedblog.
func marshalBlogStoredblogToStoredblogResponse(v *blog.Storedblog) *StoredblogResponse {
	res := &StoredblogResponse{
		ID:   v.ID,
		Name: v.Name,
	}
	if v.Comments != nil {
		res.Comments = make([]*CommentsResponse, len(v.Comments))
		for i, val := range v.Comments {
			res.Comments[i] = marshalBlogCommentsToCommentsResponse(val)
		}
	}

	return res
}

// marshalBlogCommentsToCommentsResponse builds a value of type
// *CommentsResponse from a value of type *blog.Comments.
func marshalBlogCommentsToCommentsResponse(v *blog.Comments) *CommentsResponse {
	if v == nil {
		return nil
	}
	res := &CommentsResponse{
		ID:       v.ID,
		Comments: v.Comments,
	}

	return res
}
