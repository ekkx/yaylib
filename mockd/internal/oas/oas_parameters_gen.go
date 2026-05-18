// Code generated; DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// AcceptGroupJoinRequestParams is parameters of acceptGroupJoinRequest operation.
type AcceptGroupJoinRequestParams struct {
	ID     int64
	UserId int64
}

func unpackAcceptGroupJoinRequestParams(packed middleware.Parameters) (params AcceptGroupJoinRequestParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(int64)
	}
	return params
}

func decodeAcceptGroupJoinRequestParams(args [2]string, argsEscaped bool, r *http.Request) (params AcceptGroupJoinRequestParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// AcceptGroupTransferParams is parameters of acceptGroupTransfer operation.
type AcceptGroupTransferParams struct {
	ID int64
}

func unpackAcceptGroupTransferParams(packed middleware.Parameters) (params AcceptGroupTransferParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeAcceptGroupTransferParams(args [1]string, argsEscaped bool, r *http.Request) (params AcceptGroupTransferParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// AddThreadMemberParams is parameters of addThreadMember operation.
type AddThreadMemberParams struct {
	ThreadID int64
	ID       int64
}

func unpackAddThreadMemberParams(packed middleware.Parameters) (params AddThreadMemberParams) {
	{
		key := middleware.ParameterKey{
			Name: "thread_id",
			In:   "path",
		}
		params.ThreadID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeAddThreadMemberParams(args [2]string, argsEscaped bool, r *http.Request) (params AddThreadMemberParams, _ error) {
	// Decode path: thread_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "thread_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ThreadID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// AgreePolicyParams is parameters of agreePolicy operation.
type AgreePolicyParams struct {
	Type string
}

func unpackAgreePolicyParams(packed middleware.Parameters) (params AgreePolicyParams) {
	{
		key := middleware.ParameterKey{
			Name: "type",
			In:   "path",
		}
		params.Type = packed[key].(string)
	}
	return params
}

func decodeAgreePolicyParams(args [1]string, argsEscaped bool, r *http.Request) (params AgreePolicyParams, _ error) {
	// Decode path: type.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "type",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Type = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "type",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// BanGroupUserParams is parameters of banGroupUser operation.
type BanGroupUserParams struct {
	ID     int64
	UserId int64
}

func unpackBanGroupUserParams(packed middleware.Parameters) (params BanGroupUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(int64)
	}
	return params
}

func decodeBanGroupUserParams(args [2]string, argsEscaped bool, r *http.Request) (params BanGroupUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// BlockUserParams is parameters of blockUser operation.
type BlockUserParams struct {
	ID int64
}

func unpackBlockUserParams(packed middleware.Parameters) (params BlockUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeBlockUserParams(args [1]string, argsEscaped bool, r *http.Request) (params BlockUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// BulkInviteToCallParams is parameters of bulkInviteToCall operation.
type BulkInviteToCallParams struct {
	CallID  int64
	GroupID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackBulkInviteToCallParams(packed middleware.Parameters) (params BulkInviteToCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeBulkInviteToCallParams(args [1]string, argsEscaped bool, r *http.Request) (params BulkInviteToCallParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// BumpCallParams is parameters of bumpCall operation.
type BumpCallParams struct {
	CallID           int64
	ParticipantLimit OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackBumpCallParams(packed middleware.Parameters) (params BumpCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "participant_limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ParticipantLimit = v.(OptNilInt32)
		}
	}
	return params
}

func decodeBumpCallParams(args [1]string, argsEscaped bool, r *http.Request) (params BumpCallParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: participant_limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "participant_limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotParticipantLimitVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotParticipantLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ParticipantLimit.SetTo(paramsDotParticipantLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "participant_limit",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// CancelGroupTransferParams is parameters of cancelGroupTransfer operation.
type CancelGroupTransferParams struct {
	ID int64
}

func unpackCancelGroupTransferParams(packed middleware.Parameters) (params CancelGroupTransferParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeCancelGroupTransferParams(args [1]string, argsEscaped bool, r *http.Request) (params CancelGroupTransferParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CreateBookmarkParams is parameters of createBookmark operation.
type CreateBookmarkParams struct {
	UserID int64
	ID     int64
}

func unpackCreateBookmarkParams(packed middleware.Parameters) (params CreateBookmarkParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeCreateBookmarkParams(args [2]string, argsEscaped bool, r *http.Request) (params CreateBookmarkParams, _ error) {
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CreatePostParams is parameters of createPost operation.
type CreatePostParams struct {
	XJwt OptString `json:",omitempty,omitzero"`
}

func unpackCreatePostParams(packed middleware.Parameters) (params CreatePostParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Jwt",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.XJwt = v.(OptString)
		}
	}
	return params
}

func decodeCreatePostParams(args [0]string, argsEscaped bool, r *http.Request) (params CreatePostParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Jwt.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotXJwtVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotXJwtVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.XJwt.SetTo(paramsDotXJwtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Jwt",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// CreateThreadPostParams is parameters of createThreadPost operation.
type CreateThreadPostParams struct {
	XJwt OptString `json:",omitempty,omitzero"`
	ID   int64
}

func unpackCreateThreadPostParams(packed middleware.Parameters) (params CreateThreadPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Jwt",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.XJwt = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeCreateThreadPostParams(args [1]string, argsEscaped bool, r *http.Request) (params CreateThreadPostParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Jwt.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotXJwtVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotXJwtVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.XJwt.SetTo(paramsDotXJwtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Jwt",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeclineGroupJoinRequestParams is parameters of declineGroupJoinRequest operation.
type DeclineGroupJoinRequestParams struct {
	ID     int64
	UserId int64
}

func unpackDeclineGroupJoinRequestParams(packed middleware.Parameters) (params DeclineGroupJoinRequestParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(int64)
	}
	return params
}

func decodeDeclineGroupJoinRequestParams(args [2]string, argsEscaped bool, r *http.Request) (params DeclineGroupJoinRequestParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteBookmarkParams is parameters of deleteBookmark operation.
type DeleteBookmarkParams struct {
	UserID int64
	ID     int64
}

func unpackDeleteBookmarkParams(packed middleware.Parameters) (params DeleteBookmarkParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeDeleteBookmarkParams(args [2]string, argsEscaped bool, r *http.Request) (params DeleteBookmarkParams, _ error) {
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteChatMessageParams is parameters of deleteChatMessage operation.
type DeleteChatMessageParams struct {
	RoomID    int64
	MessageID int64
}

func unpackDeleteChatMessageParams(packed middleware.Parameters) (params DeleteChatMessageParams) {
	{
		key := middleware.ParameterKey{
			Name: "room_id",
			In:   "path",
		}
		params.RoomID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "message_id",
			In:   "path",
		}
		params.MessageID = packed[key].(int64)
	}
	return params
}

func decodeDeleteChatMessageParams(args [2]string, argsEscaped bool, r *http.Request) (params DeleteChatMessageParams, _ error) {
	// Decode path: room_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "room_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.RoomID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "room_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: message_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "message_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.MessageID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "message_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteFootprintParams is parameters of deleteFootprint operation.
type DeleteFootprintParams struct {
	UserID      int64
	FootprintID int64
}

func unpackDeleteFootprintParams(packed middleware.Parameters) (params DeleteFootprintParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "footprint_id",
			In:   "path",
		}
		params.FootprintID = packed[key].(int64)
	}
	return params
}

func decodeDeleteFootprintParams(args [2]string, argsEscaped bool, r *http.Request) (params DeleteFootprintParams, _ error) {
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: footprint_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "footprint_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.FootprintID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "footprint_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteMuteKeywordParams is parameters of deleteMuteKeyword operation.
type DeleteMuteKeywordParams struct {
	Ids []int64 `json:",omitempty"`
}

func unpackDeleteMuteKeywordParams(packed middleware.Parameters) (params DeleteMuteKeywordParams) {
	{
		key := middleware.ParameterKey{
			Name: "ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.([]int64)
		}
	}
	return params
}

func decodeDeleteMuteKeywordParams(args [0]string, argsEscaped bool, r *http.Request) (params DeleteMuteKeywordParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Ids = append(params.Ids, paramsDotIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteMyReviewsParams is parameters of deleteMyReviews operation.
type DeleteMyReviewsParams struct {
	ReviewIds []int64 `json:",omitempty"`
}

func unpackDeleteMyReviewsParams(packed middleware.Parameters) (params DeleteMyReviewsParams) {
	{
		key := middleware.ParameterKey{
			Name: "review_ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ReviewIds = v.([]int64)
		}
	}
	return params
}

func decodeDeleteMyReviewsParams(args [0]string, argsEscaped bool, r *http.Request) (params DeleteMyReviewsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: review_ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "review_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotReviewIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotReviewIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.ReviewIds = append(params.ReviewIds, paramsDotReviewIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "review_ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteThreadParams is parameters of deleteThread operation.
type DeleteThreadParams struct {
	ID int64
}

func unpackDeleteThreadParams(packed middleware.Parameters) (params DeleteThreadParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeDeleteThreadParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteThreadParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeputizeGroupUsersParams is parameters of deputizeGroupUsers operation.
type DeputizeGroupUsersParams struct {
	ID int64
}

func unpackDeputizeGroupUsersParams(packed middleware.Parameters) (params DeputizeGroupUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeDeputizeGroupUsersParams(args [1]string, argsEscaped bool, r *http.Request) (params DeputizeGroupUsersParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeputizeGroupUsersMassParams is parameters of deputizeGroupUsersMass operation.
type DeputizeGroupUsersMassParams struct {
	GroupID int64
}

func unpackDeputizeGroupUsersMassParams(packed middleware.Parameters) (params DeputizeGroupUsersMassParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	return params
}

func decodeDeputizeGroupUsersMassParams(args [1]string, argsEscaped bool, r *http.Request) (params DeputizeGroupUsersMassParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// FireGroupUserParams is parameters of fireGroupUser operation.
type FireGroupUserParams struct {
	GroupID int64
	UserID  int64
}

func unpackFireGroupUserParams(packed middleware.Parameters) (params FireGroupUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeFireGroupUserParams(args [2]string, argsEscaped bool, r *http.Request) (params FireGroupUserParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: user_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// FollowUserParams is parameters of followUser operation.
type FollowUserParams struct {
	ID int64
}

func unpackFollowUserParams(packed middleware.Parameters) (params FollowUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeFollowUserParams(args [1]string, argsEscaped bool, r *http.Request) (params FollowUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// FollowUsersParams is parameters of followUsers operation.
type FollowUsersParams struct {
	UserIds []int64 `json:",omitempty"`
}

func unpackFollowUsersParams(packed middleware.Parameters) (params FollowUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserIds = v.([]int64)
		}
	}
	return params
}

func decodeFollowUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params FollowUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: user_ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotUserIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotUserIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.UserIds = append(params.UserIds, paramsDotUserIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetActiveCallPostParams is parameters of getActiveCallPost operation.
type GetActiveCallPostParams struct {
	UserID OptInt64 `json:",omitempty,omitzero"`
}

func unpackGetActiveCallPostParams(packed middleware.Parameters) (params GetActiveCallPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserID = v.(OptInt64)
		}
	}
	return params
}

func decodeGetActiveCallPostParams(args [0]string, argsEscaped bool, r *http.Request) (params GetActiveCallPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: user_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID.SetTo(paramsDotUserIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetActiveFollowingsParams is parameters of getActiveFollowings operation.
type GetActiveFollowingsParams struct {
	OnlyOnline     OptBool     `json:",omitempty,omitzero"`
	FromLoggedinAt OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetActiveFollowingsParams(packed middleware.Parameters) (params GetActiveFollowingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "only_online",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OnlyOnline = v.(OptBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_loggedin_at",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromLoggedinAt = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetActiveFollowingsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetActiveFollowingsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: only_online.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "only_online",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOnlyOnlineVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotOnlyOnlineVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OnlyOnline.SetTo(paramsDotOnlyOnlineVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "only_online",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_loggedin_at.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_loggedin_at",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromLoggedinAtVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromLoggedinAtVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromLoggedinAt.SetTo(paramsDotFromLoggedinAtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_loggedin_at",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetAgoraRtmTokenParams is parameters of getAgoraRtmToken operation.
type GetAgoraRtmTokenParams struct {
	CallID int64
}

func unpackGetAgoraRtmTokenParams(packed middleware.Parameters) (params GetAgoraRtmTokenParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	return params
}

func decodeGetAgoraRtmTokenParams(args [1]string, argsEscaped bool, r *http.Request) (params GetAgoraRtmTokenParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetAppConfigParams is parameters of getAppConfig operation.
type GetAppConfigParams struct {
	App string
}

func unpackGetAppConfigParams(packed middleware.Parameters) (params GetAppConfigParams) {
	{
		key := middleware.ParameterKey{
			Name: "app",
			In:   "path",
		}
		params.App = packed[key].(string)
	}
	return params
}

func decodeGetAppConfigParams(args [1]string, argsEscaped bool, r *http.Request) (params GetAppConfigParams, _ error) {
	// Decode path: app.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "app",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.App = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "app",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetBannedWordsParams is parameters of getBannedWords operation.
type GetBannedWordsParams struct {
	CountryApiValue string
}

func unpackGetBannedWordsParams(packed middleware.Parameters) (params GetBannedWordsParams) {
	{
		key := middleware.ParameterKey{
			Name: "countryApiValue",
			In:   "path",
		}
		params.CountryApiValue = packed[key].(string)
	}
	return params
}

func decodeGetBannedWordsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetBannedWordsParams, _ error) {
	// Decode path: countryApiValue.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "countryApiValue",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CountryApiValue = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "countryApiValue",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetBlockedUsersParams is parameters of getBlockedUsers operation.
type GetBlockedUsersParams struct {
	FromID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetBlockedUsersParams(packed middleware.Parameters) (params GetBlockedUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetBlockedUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params GetBlockedUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetBookmarkedPostsParams is parameters of getBookmarkedPosts operation.
type GetBookmarkedPostsParams struct {
	UserID int64
	From   OptNilString `json:",omitempty,omitzero"`
}

func unpackGetBookmarkedPostsParams(packed middleware.Parameters) (params GetBookmarkedPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetBookmarkedPostsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetBookmarkedPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetBucketPresignedUrlsParams is parameters of getBucketPresignedUrls operation.
type GetBucketPresignedUrlsParams struct {
	FileNames []string `json:",omitempty"`
}

func unpackGetBucketPresignedUrlsParams(packed middleware.Parameters) (params GetBucketPresignedUrlsParams) {
	{
		key := middleware.ParameterKey{
			Name: "file_names[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FileNames = v.([]string)
		}
	}
	return params
}

func decodeGetBucketPresignedUrlsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetBucketPresignedUrlsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: file_names[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "file_names[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotFileNamesVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotFileNamesVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.FileNames = append(params.FileNames, paramsDotFileNamesVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "file_names[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetCallFollowersTimelineParams is parameters of getCallFollowersTimeline operation.
type GetCallFollowersTimelineParams struct {
	FromTimestamp          OptNilInt64  `json:",omitempty,omitzero"`
	Number                 OptNilInt32  `json:",omitempty,omitzero"`
	CategoryID             OptNilInt64  `json:",omitempty,omitzero"`
	CallType               OptNilString `json:",omitempty,omitzero"`
	IncludeCircleCall      OptNilBool   `json:",omitempty,omitzero"`
	ExcludeRecentGomimushi OptNilBool   `json:",omitempty,omitzero"`
}

func unpackGetCallFollowersTimelineParams(packed middleware.Parameters) (params GetCallFollowersTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "category_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CategoryID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "call_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CallType = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "include_circle_call",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.IncludeCircleCall = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "exclude_recent_gomimushi",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ExcludeRecentGomimushi = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetCallFollowersTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetCallFollowersTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: category_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCategoryIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotCategoryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CategoryID.SetTo(paramsDotCategoryIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "category_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: call_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "call_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCallTypeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCallTypeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CallType.SetTo(paramsDotCallTypeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: include_circle_call.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "include_circle_call",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIncludeCircleCallVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotIncludeCircleCallVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IncludeCircleCall.SetTo(paramsDotIncludeCircleCallVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "include_circle_call",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: exclude_recent_gomimushi.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "exclude_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotExcludeRecentGomimushiVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotExcludeRecentGomimushiVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ExcludeRecentGomimushi.SetTo(paramsDotExcludeRecentGomimushiVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "exclude_recent_gomimushi",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetCallGiftHistoryParams is parameters of getCallGiftHistory operation.
type GetCallGiftHistoryParams struct {
	CallID int64
	From   OptNilString `json:",omitempty,omitzero"`
	Number OptNilInt32  `json:",omitempty,omitzero"`
}

func unpackGetCallGiftHistoryParams(packed middleware.Parameters) (params GetCallGiftHistoryParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetCallGiftHistoryParams(args [1]string, argsEscaped bool, r *http.Request) (params GetCallGiftHistoryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetCallTimelineParams is parameters of getCallTimeline operation.
type GetCallTimelineParams struct {
	GroupID                OptNilInt64  `json:",omitempty,omitzero"`
	FromTimestamp          OptNilInt64  `json:",omitempty,omitzero"`
	Number                 OptNilInt32  `json:",omitempty,omitzero"`
	CategoryID             OptNilInt64  `json:",omitempty,omitzero"`
	CallType               OptNilString `json:",omitempty,omitzero"`
	IncludeCircleCall      OptNilBool   `json:",omitempty,omitzero"`
	CrossGeneration        OptNilBool   `json:",omitempty,omitzero"`
	ExcludeRecentGomimushi OptNilBool   `json:",omitempty,omitzero"`
}

func unpackGetCallTimelineParams(packed middleware.Parameters) (params GetCallTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "category_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CategoryID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "call_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CallType = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "include_circle_call",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.IncludeCircleCall = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "cross_generation",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CrossGeneration = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "exclude_recent_gomimushi",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ExcludeRecentGomimushi = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetCallTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetCallTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: category_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCategoryIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotCategoryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CategoryID.SetTo(paramsDotCategoryIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "category_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: call_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "call_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCallTypeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCallTypeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CallType.SetTo(paramsDotCallTypeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: include_circle_call.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "include_circle_call",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIncludeCircleCallVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotIncludeCircleCallVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IncludeCircleCall.SetTo(paramsDotIncludeCircleCallVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "include_circle_call",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: cross_generation.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "cross_generation",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCrossGenerationVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotCrossGenerationVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CrossGeneration.SetTo(paramsDotCrossGenerationVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cross_generation",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: exclude_recent_gomimushi.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "exclude_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotExcludeRecentGomimushiVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotExcludeRecentGomimushiVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ExcludeRecentGomimushi.SetTo(paramsDotExcludeRecentGomimushiVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "exclude_recent_gomimushi",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatMessagesParams is parameters of getChatMessages operation.
type GetChatMessagesParams struct {
	ID            int64
	Number        OptNilInt32 `json:",omitempty,omitzero"`
	FromMessageID OptNilInt64 `json:",omitempty,omitzero"`
	ToMessageID   OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetChatMessagesParams(packed middleware.Parameters) (params GetChatMessagesParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_message_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromMessageID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to_message_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ToMessageID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetChatMessagesParams(args [1]string, argsEscaped bool, r *http.Request) (params GetChatMessagesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_message_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_message_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromMessageIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromMessageIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromMessageID.SetTo(paramsDotFromMessageIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_message_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: to_message_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to_message_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToMessageIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotToMessageIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ToMessageID.SetTo(paramsDotToMessageIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to_message_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatRequestsParams is parameters of getChatRequests operation.
type GetChatRequestsParams struct {
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetChatRequestsParams(packed middleware.Parameters) (params GetChatRequestsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetChatRequestsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetChatRequestsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatRoomParams is parameters of getChatRoom operation.
type GetChatRoomParams struct {
	ID int64
}

func unpackGetChatRoomParams(packed middleware.Parameters) (params GetChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params GetChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatUnreadStatusParams is parameters of getChatUnreadStatus operation.
type GetChatUnreadStatusParams struct {
	FromTime OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetChatUnreadStatusParams(packed middleware.Parameters) (params GetChatUnreadStatusParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_time",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTime = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetChatUnreadStatusParams(args [0]string, argsEscaped bool, r *http.Request) (params GetChatUnreadStatusParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimeVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTime.SetTo(paramsDotFromTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_time",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatableFollowingsParams is parameters of getChatableFollowings operation.
type GetChatableFollowingsParams struct {
	FromFollowID  OptNilInt64  `json:",omitempty,omitzero"`
	FromTimestamp OptNilInt64  `json:",omitempty,omitzero"`
	OrderBy       OptNilString `json:",omitempty,omitzero"`
}

func unpackGetChatableFollowingsParams(packed middleware.Parameters) (params GetChatableFollowingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_follow_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromFollowID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order_by",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OrderBy = v.(OptNilString)
		}
	}
	return params
}

func decodeGetChatableFollowingsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetChatableFollowingsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_follow_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_follow_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromFollowIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromFollowIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromFollowID.SetTo(paramsDotFromFollowIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_follow_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: order_by.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderByVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderByVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderBy.SetTo(paramsDotOrderByVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order_by",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetConferenceCallParams is parameters of getConferenceCall operation.
type GetConferenceCallParams struct {
	CallID int64
}

func unpackGetConferenceCallParams(packed middleware.Parameters) (params GetConferenceCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	return params
}

func decodeGetConferenceCallParams(args [1]string, argsEscaped bool, r *http.Request) (params GetConferenceCallParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetConversationParams is parameters of getConversation operation.
type GetConversationParams struct {
	ID         int64
	GroupID    OptNilInt64 `json:",omitempty,omitzero"`
	ThreadID   OptNilInt64 `json:",omitempty,omitzero"`
	Number     OptNilInt32 `json:",omitempty,omitzero"`
	FromPostID OptNilInt64 `json:",omitempty,omitzero"`
	Reverse    OptBool     `json:",omitempty,omitzero"`
}

func unpackGetConversationParams(packed middleware.Parameters) (params GetConversationParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "thread_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ThreadID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "reverse",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Reverse = v.(OptBool)
		}
	}
	return params
}

func decodeGetConversationParams(args [1]string, argsEscaped bool, r *http.Request) (params GetConversationParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: thread_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "thread_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotThreadIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotThreadIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ThreadID.SetTo(paramsDotThreadIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: reverse.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "reverse",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotReverseVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotReverseVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Reverse.SetTo(paramsDotReverseVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "reverse",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetFollowRequestsParams is parameters of getFollowRequests operation.
type GetFollowRequestsParams struct {
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetFollowRequestsParams(packed middleware.Parameters) (params GetFollowRequestsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetFollowRequestsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetFollowRequestsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetFollowingTimelineParams is parameters of getFollowingTimeline operation.
type GetFollowingTimelineParams struct {
	From                  OptNilString `json:",omitempty,omitzero"`
	FromPostID            OptNilInt64  `json:",omitempty,omitzero"`
	OnlyRoot              OptNilBool   `json:",omitempty,omitzero"`
	OrderBy               OptString    `json:",omitempty,omitzero"`
	Number                OptNilInt32  `json:",omitempty,omitzero"`
	Mxn                   OptNilInt32  `json:",omitempty,omitzero"`
	ReduceSelfie          OptNilBool   `json:",omitempty,omitzero"`
	CustomGenerationRange OptNilBool   `json:",omitempty,omitzero"`
}

func unpackGetFollowingTimelineParams(packed middleware.Parameters) (params GetFollowingTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "only_root",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OnlyRoot = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order_by",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OrderBy = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "mxn",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Mxn = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "reduce_selfie",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ReduceSelfie = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "custom_generation_range",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CustomGenerationRange = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetFollowingTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetFollowingTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: only_root.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "only_root",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOnlyRootVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotOnlyRootVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OnlyRoot.SetTo(paramsDotOnlyRootVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "only_root",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: order_by.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderByVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderByVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderBy.SetTo(paramsDotOrderByVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order_by",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: mxn.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "mxn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotMxnVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotMxnVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Mxn.SetTo(paramsDotMxnVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "mxn",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: reduce_selfie.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "reduce_selfie",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotReduceSelfieVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotReduceSelfieVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ReduceSelfie.SetTo(paramsDotReduceSelfieVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "reduce_selfie",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: custom_generation_range.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "custom_generation_range",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCustomGenerationRangeVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotCustomGenerationRangeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CustomGenerationRange.SetTo(paramsDotCustomGenerationRangeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "custom_generation_range",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetFollowingsBornTodayParams is parameters of getFollowingsBornToday operation.
type GetFollowingsBornTodayParams struct {
	Birthdate OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetFollowingsBornTodayParams(packed middleware.Parameters) (params GetFollowingsBornTodayParams) {
	{
		key := middleware.ParameterKey{
			Name: "birthdate",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Birthdate = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetFollowingsBornTodayParams(args [0]string, argsEscaped bool, r *http.Request) (params GetFollowingsBornTodayParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: birthdate.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "birthdate",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBirthdateVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotBirthdateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Birthdate.SetTo(paramsDotBirthdateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "birthdate",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetFootprintsParams is parameters of getFootprints operation.
type GetFootprintsParams struct {
	Mode   OptNilString `json:",omitempty,omitzero"`
	Number OptInt32     `json:",omitempty,omitzero"`
	From   OptNilString `json:",omitempty,omitzero"`
}

func unpackGetFootprintsParams(packed middleware.Parameters) (params GetFootprintsParams) {
	{
		key := middleware.ParameterKey{
			Name: "mode",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Mode = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetFootprintsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetFootprintsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: mode.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "mode",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotModeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotModeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Mode.SetTo(paramsDotModeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "mode",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetFreshUserParams is parameters of getFreshUser operation.
type GetFreshUserParams struct {
	ID int64
}

func unpackGetFreshUserParams(packed middleware.Parameters) (params GetFreshUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetFreshUserParams(args [1]string, argsEscaped bool, r *http.Request) (params GetFreshUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupParams is parameters of getGroup operation.
type GetGroupParams struct {
	ID int64
}

func unpackGetGroupParams(packed middleware.Parameters) (params GetGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupBanListParams is parameters of getGroupBanList operation.
type GetGroupBanListParams struct {
	ID   int64
	Page OptInt32 `json:",omitempty,omitzero"`
}

func unpackGetGroupBanListParams(packed middleware.Parameters) (params GetGroupBanListParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt32)
		}
	}
	return params
}

func decodeGetGroupBanListParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupBanListParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupGiftHistoryParams is parameters of getGroupGiftHistory operation.
type GetGroupGiftHistoryParams struct {
	GroupID int64
	Number  OptNilInt32  `json:",omitempty,omitzero"`
	From    OptNilString `json:",omitempty,omitzero"`
}

func unpackGetGroupGiftHistoryParams(packed middleware.Parameters) (params GetGroupGiftHistoryParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetGroupGiftHistoryParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupGiftHistoryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupGiftTransactionsParams is parameters of getGroupGiftTransactions operation.
type GetGroupGiftTransactionsParams struct {
	GroupID int64
	Number  OptNilInt32  `json:",omitempty,omitzero"`
	From    OptNilString `json:",omitempty,omitzero"`
}

func unpackGetGroupGiftTransactionsParams(packed middleware.Parameters) (params GetGroupGiftTransactionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetGroupGiftTransactionsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupGiftTransactionsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupHighlightsParams is parameters of getGroupHighlights operation.
type GetGroupHighlightsParams struct {
	GroupID int64
	From    OptNilInt64 `json:",omitempty,omitzero"`
	Number  OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetGroupHighlightsParams(packed middleware.Parameters) (params GetGroupHighlightsParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetGroupHighlightsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupHighlightsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupMemberParams is parameters of getGroupMember operation.
type GetGroupMemberParams struct {
	ID     int64
	UserId int64
}

func unpackGetGroupMemberParams(packed middleware.Parameters) (params GetGroupMemberParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(int64)
	}
	return params
}

func decodeGetGroupMemberParams(args [2]string, argsEscaped bool, r *http.Request) (params GetGroupMemberParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupMembersParams is parameters of getGroupMembers operation.
type GetGroupMembersParams struct {
	ID     int64
	Number OptNilInt32 `json:",omitempty,omitzero"`
	FromID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetGroupMembersParams(packed middleware.Parameters) (params GetGroupMembersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetGroupMembersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupMembersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupNotificationSettingsParams is parameters of getGroupNotificationSettings operation.
type GetGroupNotificationSettingsParams struct {
	ID int64
}

func unpackGetGroupNotificationSettingsParams(packed middleware.Parameters) (params GetGroupNotificationSettingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetGroupNotificationSettingsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetGroupNotificationSettingsParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupReceivedGiftSendersParams is parameters of getGroupReceivedGiftSenders operation.
type GetGroupReceivedGiftSendersParams struct {
	GroupID   int64
	GiftID    int64
	UserID    OptInt64    `json:",omitempty,omitzero"`
	Timestamp OptInt64    `json:",omitempty,omitzero"`
	Number    OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetGroupReceivedGiftSendersParams(packed middleware.Parameters) (params GetGroupReceivedGiftSendersParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "gift_id",
			In:   "path",
		}
		params.GiftID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Timestamp = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetGroupReceivedGiftSendersParams(args [2]string, argsEscaped bool, r *http.Request) (params GetGroupReceivedGiftSendersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: gift_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "gift_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GiftID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "gift_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: user_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID.SetTo(paramsDotUserIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Timestamp.SetTo(paramsDotTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupTimelineParams is parameters of getGroupTimeline operation.
type GetGroupTimelineParams struct {
	GroupID    OptInt64    `json:",omitempty,omitzero"`
	FromPostID OptNilInt64 `json:",omitempty,omitzero"`
	Reverse    OptNilBool  `json:",omitempty,omitzero"`
	PostType   OptPostType `json:",omitempty,omitzero"`
	Number     OptNilInt32 `json:",omitempty,omitzero"`
	OnlyRoot   OptNilBool  `json:",omitempty,omitzero"`
}

func unpackGetGroupTimelineParams(packed middleware.Parameters) (params GetGroupTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "reverse",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Reverse = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "post_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PostType = v.(OptPostType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "only_root",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OnlyRoot = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetGroupTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetGroupTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: reverse.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "reverse",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotReverseVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotReverseVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Reverse.SetTo(paramsDotReverseVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "reverse",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: post_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPostTypeVal PostType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPostTypeVal = PostType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.PostType.SetTo(paramsDotPostTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PostType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: only_root.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "only_root",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOnlyRootVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotOnlyRootVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OnlyRoot.SetTo(paramsDotOnlyRootVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "only_root",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetGroupUnreadStatusParams is parameters of getGroupUnreadStatus operation.
type GetGroupUnreadStatusParams struct {
	FromTime OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetGroupUnreadStatusParams(packed middleware.Parameters) (params GetGroupUnreadStatusParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_time",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTime = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetGroupUnreadStatusParams(args [0]string, argsEscaped bool, r *http.Request) (params GetGroupUnreadStatusParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimeVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTime.SetTo(paramsDotFromTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_time",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetHimaUsersParams is parameters of getHimaUsers operation.
type GetHimaUsersParams struct {
	FromHimaID OptNilInt64 `json:",omitempty,omitzero"`
	Number     OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetHimaUsersParams(packed middleware.Parameters) (params GetHimaUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_hima_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromHimaID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetHimaUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params GetHimaUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_hima_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_hima_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromHimaIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromHimaIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromHimaID.SetTo(paramsDotFromHimaIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_hima_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetInvitableCallUsersParams is parameters of getInvitableCallUsers operation.
type GetInvitableCallUsersParams struct {
	CallID        int64
	FromTimestamp OptNilInt64  `json:",omitempty,omitzero"`
	UserNickname  OptNilString `json:",omitempty,omitzero"`
}

func unpackGetInvitableCallUsersParams(packed middleware.Parameters) (params GetInvitableCallUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user[nickname]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserNickname = v.(OptNilString)
		}
	}
	return params
}

func decodeGetInvitableCallUsersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetInvitableCallUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: user[nickname].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserNicknameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserNicknameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserNickname.SetTo(paramsDotUserNicknameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user[nickname]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetInvitableGroupUsersParams is parameters of getInvitableGroupUsers operation.
type GetInvitableGroupUsersParams struct {
	GroupID       int64
	FromTimestamp OptNilInt64  `json:",omitempty,omitzero"`
	UserNickname  OptNilString `json:",omitempty,omitzero"`
}

func unpackGetInvitableGroupUsersParams(packed middleware.Parameters) (params GetInvitableGroupUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user[nickname]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserNickname = v.(OptNilString)
		}
	}
	return params
}

func decodeGetInvitableGroupUsersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetInvitableGroupUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: user[nickname].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserNicknameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserNicknameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserNickname.SetTo(paramsDotUserNicknameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user[nickname]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetJoinedGroupStatusesParams is parameters of getJoinedGroupStatuses operation.
type GetJoinedGroupStatusesParams struct {
	Ids []int64 `json:",omitempty"`
}

func unpackGetJoinedGroupStatusesParams(packed middleware.Parameters) (params GetJoinedGroupStatusesParams) {
	{
		key := middleware.ParameterKey{
			Name: "ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.([]int64)
		}
	}
	return params
}

func decodeGetJoinedGroupStatusesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetJoinedGroupStatusesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Ids = append(params.Ids, paramsDotIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetJoinedThreadStatusesParams is parameters of getJoinedThreadStatuses operation.
type GetJoinedThreadStatusesParams struct {
	Ids []int64 `json:",omitempty"`
}

func unpackGetJoinedThreadStatusesParams(packed middleware.Parameters) (params GetJoinedThreadStatusesParams) {
	{
		key := middleware.ParameterKey{
			Name: "ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.([]int64)
		}
	}
	return params
}

func decodeGetJoinedThreadStatusesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetJoinedThreadStatusesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Ids = append(params.Ids, paramsDotIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetMainChatRoomsParams is parameters of getMainChatRooms operation.
type GetMainChatRoomsParams struct {
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetMainChatRoomsParams(packed middleware.Parameters) (params GetMainChatRoomsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetMainChatRoomsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetMainChatRoomsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetMyPostsParams is parameters of getMyPosts operation.
type GetMyPostsParams struct {
	FromPostID       OptNilInt64 `json:",omitempty,omitzero"`
	Number           OptInt32    `json:",omitempty,omitzero"`
	IncludeGroupPost OptNilBool  `json:",omitempty,omitzero"`
}

func unpackGetMyPostsParams(packed middleware.Parameters) (params GetMyPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "include_group_post",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.IncludeGroupPost = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetMyPostsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetMyPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: include_group_post.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "include_group_post",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIncludeGroupPostVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotIncludeGroupPostVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IncludeGroupPost.SetTo(paramsDotIncludeGroupPostVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "include_group_post",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetMyReviewsParams is parameters of getMyReviews operation.
type GetMyReviewsParams struct {
	FromID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetMyReviewsParams(packed middleware.Parameters) (params GetMyReviewsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetMyReviewsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetMyReviewsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPhoneStatusParams is parameters of getPhoneStatus operation.
type GetPhoneStatusParams struct {
	OpponentID int64
}

func unpackGetPhoneStatusParams(packed middleware.Parameters) (params GetPhoneStatusParams) {
	{
		key := middleware.ParameterKey{
			Name: "opponent_id",
			In:   "path",
		}
		params.OpponentID = packed[key].(int64)
	}
	return params
}

func decodeGetPhoneStatusParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPhoneStatusParams, _ error) {
	// Decode path: opponent_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "opponent_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.OpponentID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "opponent_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPopularWordsParams is parameters of getPopularWords operation.
type GetPopularWordsParams struct {
	CountryApiValue string
	App             string
}

func unpackGetPopularWordsParams(packed middleware.Parameters) (params GetPopularWordsParams) {
	{
		key := middleware.ParameterKey{
			Name: "countryApiValue",
			In:   "path",
		}
		params.CountryApiValue = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "app",
			In:   "path",
		}
		params.App = packed[key].(string)
	}
	return params
}

func decodeGetPopularWordsParams(args [2]string, argsEscaped bool, r *http.Request) (params GetPopularWordsParams, _ error) {
	// Decode path: countryApiValue.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "countryApiValue",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CountryApiValue = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "countryApiValue",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: app.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "app",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.App = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "app",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostParams is parameters of getPost operation.
type GetPostParams struct {
	ID           int64
	CacheControl OptNilString `json:",omitempty,omitzero"`
}

func unpackGetPostParams(packed middleware.Parameters) (params GetPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "Cache-Control",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.CacheControl = v.(OptNilString)
		}
	}
	return params
}

func decodeGetPostParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPostParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode header: Cache-Control.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Cache-Control",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCacheControlVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCacheControlVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CacheControl.SetTo(paramsDotCacheControlVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Cache-Control",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostGiftTransactionsParams is parameters of getPostGiftTransactions operation.
type GetPostGiftTransactionsParams struct {
	PostID int64
	Number OptNilInt32  `json:",omitempty,omitzero"`
	From   OptNilString `json:",omitempty,omitzero"`
}

func unpackGetPostGiftTransactionsParams(packed middleware.Parameters) (params GetPostGiftTransactionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "post_id",
			In:   "path",
		}
		params.PostID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetPostGiftTransactionsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPostGiftTransactionsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: post_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "post_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PostID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostLikersParams is parameters of getPostLikers operation.
type GetPostLikersParams struct {
	ID     int64
	FromID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetPostLikersParams(packed middleware.Parameters) (params GetPostLikersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetPostLikersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPostLikersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostRepostsParams is parameters of getPostReposts operation.
type GetPostRepostsParams struct {
	ID         int64
	FromPostID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetPostRepostsParams(packed middleware.Parameters) (params GetPostRepostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetPostRepostsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPostRepostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostUrlMetadataParams is parameters of getPostUrlMetadata operation.
type GetPostUrlMetadataParams struct {
	URL OptString `json:",omitempty,omitzero"`
}

func unpackGetPostUrlMetadataParams(packed middleware.Parameters) (params GetPostUrlMetadataParams) {
	{
		key := middleware.ParameterKey{
			Name: "url",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.URL = v.(OptString)
		}
	}
	return params
}

func decodeGetPostUrlMetadataParams(args [0]string, argsEscaped bool, r *http.Request) (params GetPostUrlMetadataParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: url.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "url",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotURLVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotURLVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.URL.SetTo(paramsDotURLVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "url",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostsByIdsParams is parameters of getPostsByIds operation.
type GetPostsByIdsParams struct {
	PostIds []int64 `json:",omitempty"`
}

func unpackGetPostsByIdsParams(packed middleware.Parameters) (params GetPostsByIdsParams) {
	{
		key := middleware.ParameterKey{
			Name: "post_ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PostIds = v.([]int64)
		}
	}
	return params
}

func decodeGetPostsByIdsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetPostsByIdsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: post_ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "post_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotPostIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotPostIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.PostIds = append(params.PostIds, paramsDotPostIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPostsByTagParams is parameters of getPostsByTag operation.
type GetPostsByTagParams struct {
	Tag        string
	FromPostID OptNilInt64 `json:",omitempty,omitzero"`
	Number     OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetPostsByTagParams(packed middleware.Parameters) (params GetPostsByTagParams) {
	{
		key := middleware.ParameterKey{
			Name: "tag",
			In:   "path",
		}
		params.Tag = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetPostsByTagParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPostsByTagParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: tag.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "tag",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Tag = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tag",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetReceivedGiftSendersParams is parameters of getReceivedGiftSenders operation.
type GetReceivedGiftSendersParams struct {
	GiftID int64
	Number OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetReceivedGiftSendersParams(packed middleware.Parameters) (params GetReceivedGiftSendersParams) {
	{
		key := middleware.ParameterKey{
			Name: "gift_id",
			In:   "path",
		}
		params.GiftID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetReceivedGiftSendersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetReceivedGiftSendersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: gift_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "gift_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GiftID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "gift_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetReceivedGiftTransactionParams is parameters of getReceivedGiftTransaction operation.
type GetReceivedGiftTransactionParams struct {
	GiftTransactionUUID string
}

func unpackGetReceivedGiftTransactionParams(packed middleware.Parameters) (params GetReceivedGiftTransactionParams) {
	{
		key := middleware.ParameterKey{
			Name: "gift_transaction_uuid",
			In:   "path",
		}
		params.GiftTransactionUUID = packed[key].(string)
	}
	return params
}

func decodeGetReceivedGiftTransactionParams(args [1]string, argsEscaped bool, r *http.Request) (params GetReceivedGiftTransactionParams, _ error) {
	// Decode path: gift_transaction_uuid.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "gift_transaction_uuid",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.GiftTransactionUUID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "gift_transaction_uuid",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetRecentEngagementPostsParams is parameters of getRecentEngagementPosts operation.
type GetRecentEngagementPostsParams struct {
	Number OptInt32 `json:",omitempty,omitzero"`
}

func unpackGetRecentEngagementPostsParams(packed middleware.Parameters) (params GetRecentEngagementPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	return params
}

func decodeGetRecentEngagementPostsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetRecentEngagementPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRecommendedFollowUsersParams is parameters of getRecommendedFollowUsers operation.
type GetRecommendedFollowUsersParams struct {
	ID     int64
	Number OptNilInt32 `json:",omitempty,omitzero"`
	Page   OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetRecommendedFollowUsersParams(packed middleware.Parameters) (params GetRecommendedFollowUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetRecommendedFollowUsersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetRecommendedFollowUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRecommendedTimelineParams is parameters of getRecommendedTimeline operation.
type GetRecommendedTimelineParams struct {
	ExperimentNum OptInt32 `json:",omitempty,omitzero"`
	VariantNum    OptInt32 `json:",omitempty,omitzero"`
	Number        OptInt32 `json:",omitempty,omitzero"`
}

func unpackGetRecommendedTimelineParams(packed middleware.Parameters) (params GetRecommendedTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "experiment_num",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ExperimentNum = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "variant_num",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.VariantNum = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	return params
}

func decodeGetRecommendedTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetRecommendedTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: experiment_num.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "experiment_num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotExperimentNumVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotExperimentNumVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ExperimentNum.SetTo(paramsDotExperimentNumVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "experiment_num",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: variant_num.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "variant_num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotVariantNumVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotVariantNumVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.VariantNum.SetTo(paramsDotVariantNumVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "variant_num",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRelatableGroupsParams is parameters of getRelatableGroups operation.
type GetRelatableGroupsParams struct {
	ID      int64
	Keyword OptNilString `json:",omitempty,omitzero"`
	From    OptNilString `json:",omitempty,omitzero"`
}

func unpackGetRelatableGroupsParams(packed middleware.Parameters) (params GetRelatableGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetRelatableGroupsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetRelatableGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRelatedGroupsParams is parameters of getRelatedGroups operation.
type GetRelatedGroupsParams struct {
	ID      int64
	Keyword OptNilString `json:",omitempty,omitzero"`
	From    OptNilString `json:",omitempty,omitzero"`
}

func unpackGetRelatedGroupsParams(packed middleware.Parameters) (params GetRelatedGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetRelatedGroupsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetRelatedGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRootPostsParams is parameters of getRootPosts operation.
type GetRootPostsParams struct {
	Ids []int64 `json:",omitempty"`
}

func unpackGetRootPostsParams(packed middleware.Parameters) (params GetRootPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.([]int64)
		}
	}
	return params
}

func decodeGetRootPostsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetRootPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Ids = append(params.Ids, paramsDotIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetThreadParams is parameters of getThread operation.
type GetThreadParams struct {
	ID int64
}

func unpackGetThreadParams(packed middleware.Parameters) (params GetThreadParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetThreadParams(args [1]string, argsEscaped bool, r *http.Request) (params GetThreadParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetThreadPostsParams is parameters of getThreadPosts operation.
type GetThreadPostsParams struct {
	ID       int64
	PostType OptPostType `json:",omitempty,omitzero"`
	Number   OptNilInt32 `json:",omitempty,omitzero"`
	From     OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetThreadPostsParams(packed middleware.Parameters) (params GetThreadPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "post_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PostType = v.(OptPostType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetThreadPostsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetThreadPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: post_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPostTypeVal PostType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPostTypeVal = PostType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.PostType.SetTo(paramsDotPostTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PostType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetTimelineParams is parameters of getTimeline operation.
type GetTimelineParams struct {
	NoreplyMode             NoreplyMode
	OrderBy                 OptString    `json:",omitempty,omitzero"`
	ExperimentOlderAgeRules OptNilBool   `json:",omitempty,omitzero"`
	From                    OptNilString `json:",omitempty,omitzero"`
	FromPostID              OptNilInt64  `json:",omitempty,omitzero"`
	Number                  OptNilInt32  `json:",omitempty,omitzero"`
	Mxn                     OptNilInt32  `json:",omitempty,omitzero"`
	En                      OptNilInt32  `json:",omitempty,omitzero"`
	Vn                      OptNilInt32  `json:",omitempty,omitzero"`
	ReduceSelfie            OptNilBool   `json:",omitempty,omitzero"`
	CustomGenerationRange   OptNilBool   `json:",omitempty,omitzero"`
}

func unpackGetTimelineParams(packed middleware.Parameters) (params GetTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "noreply_mode",
			In:   "path",
		}
		params.NoreplyMode = packed[key].(NoreplyMode)
	}
	{
		key := middleware.ParameterKey{
			Name: "order_by",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OrderBy = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "experiment_older_age_rules",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ExperimentOlderAgeRules = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "mxn",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Mxn = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "en",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.En = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "vn",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Vn = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "reduce_selfie",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ReduceSelfie = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "custom_generation_range",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CustomGenerationRange = v.(OptNilBool)
		}
	}
	return params
}

func decodeGetTimelineParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: noreply_mode.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "noreply_mode",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.NoreplyMode = NoreplyMode(c)
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.NoreplyMode.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "noreply_mode",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: order_by.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderByVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderByVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderBy.SetTo(paramsDotOrderByVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order_by",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: experiment_older_age_rules.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "experiment_older_age_rules",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotExperimentOlderAgeRulesVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotExperimentOlderAgeRulesVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ExperimentOlderAgeRules.SetTo(paramsDotExperimentOlderAgeRulesVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "experiment_older_age_rules",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: mxn.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "mxn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotMxnVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotMxnVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Mxn.SetTo(paramsDotMxnVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "mxn",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: en.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "en",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEnVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotEnVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.En.SetTo(paramsDotEnVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "en",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: vn.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "vn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotVnVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotVnVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Vn.SetTo(paramsDotVnVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "vn",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: reduce_selfie.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "reduce_selfie",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotReduceSelfieVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotReduceSelfieVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ReduceSelfie.SetTo(paramsDotReduceSelfieVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "reduce_selfie",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: custom_generation_range.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "custom_generation_range",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCustomGenerationRangeVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotCustomGenerationRangeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CustomGenerationRange.SetTo(paramsDotCustomGenerationRangeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "custom_generation_range",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUpdatedChatRoomsParams is parameters of getUpdatedChatRooms operation.
type GetUpdatedChatRoomsParams struct {
	FromTime OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetUpdatedChatRoomsParams(packed middleware.Parameters) (params GetUpdatedChatRoomsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_time",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTime = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetUpdatedChatRoomsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUpdatedChatRoomsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimeVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTime.SetTo(paramsDotFromTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_time",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserParams is parameters of getUser operation.
type GetUserParams struct {
	ID int64
}

func unpackGetUserParams(packed middleware.Parameters) (params GetUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetUserParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserActivitiesParams is parameters of getUserActivities operation.
type GetUserActivitiesParams struct {
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetUserActivitiesParams(packed middleware.Parameters) (params GetUserActivitiesParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetUserActivitiesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUserActivitiesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserActivitiesV1Params is parameters of getUserActivitiesV1 operation.
type GetUserActivitiesV1Params struct {
	Important     OptBool     `json:",omitempty,omitzero"`
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
	Number        OptInt32    `json:",omitempty,omitzero"`
}

func unpackGetUserActivitiesV1Params(packed middleware.Parameters) (params GetUserActivitiesV1Params) {
	{
		key := middleware.ParameterKey{
			Name: "important",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Important = v.(OptBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	return params
}

func decodeGetUserActivitiesV1Params(args [0]string, argsEscaped bool, r *http.Request) (params GetUserActivitiesV1Params, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: important.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "important",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotImportantVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotImportantVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Important.SetTo(paramsDotImportantVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "important",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserByQrParams is parameters of getUserByQr operation.
type GetUserByQrParams struct {
	Qr string
}

func unpackGetUserByQrParams(packed middleware.Parameters) (params GetUserByQrParams) {
	{
		key := middleware.ParameterKey{
			Name: "qr",
			In:   "path",
		}
		params.Qr = packed[key].(string)
	}
	return params
}

func decodeGetUserByQrParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserByQrParams, _ error) {
	// Decode path: qr.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "qr",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Qr = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "qr",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserFollowersParams is parameters of getUserFollowers operation.
type GetUserFollowersParams struct {
	ID           int64
	From         OptNilInt64  `json:",omitempty,omitzero"`
	FollowedByMe OptNilBool   `json:",omitempty,omitzero"`
	UserNickname OptNilString `json:",omitempty,omitzero"`
}

func unpackGetUserFollowersParams(packed middleware.Parameters) (params GetUserFollowersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "followed_by_me",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FollowedByMe = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user[nickname]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserNickname = v.(OptNilString)
		}
	}
	return params
}

func decodeGetUserFollowersParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserFollowersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: followed_by_me.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "followed_by_me",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFollowedByMeVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotFollowedByMeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FollowedByMe.SetTo(paramsDotFollowedByMeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "followed_by_me",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: user[nickname].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserNicknameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserNicknameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserNickname.SetTo(paramsDotUserNicknameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user[nickname]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserFollowingsParams is parameters of getUserFollowings operation.
type GetUserFollowingsParams struct {
	ID            int64
	From          OptNilInt64  `json:",omitempty,omitzero"`
	FromTimestamp OptNilInt64  `json:",omitempty,omitzero"`
	OrderBy       OptNilString `json:",omitempty,omitzero"`
	UserNickname  OptNilString `json:",omitempty,omitzero"`
}

func unpackGetUserFollowingsParams(packed middleware.Parameters) (params GetUserFollowingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order_by",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OrderBy = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user[nickname]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserNickname = v.(OptNilString)
		}
	}
	return params
}

func decodeGetUserFollowingsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserFollowingsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: order_by.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderByVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderByVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderBy.SetTo(paramsDotOrderByVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order_by",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: user[nickname].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserNicknameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserNicknameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserNickname.SetTo(paramsDotUserNicknameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user[nickname]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserGiftTransactionsParams is parameters of getUserGiftTransactions operation.
type GetUserGiftTransactionsParams struct {
	UserID int64
	Number OptNilInt32  `json:",omitempty,omitzero"`
	From   OptNilString `json:",omitempty,omitzero"`
}

func unpackGetUserGiftTransactionsParams(packed middleware.Parameters) (params GetUserGiftTransactionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeGetUserGiftTransactionsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserGiftTransactionsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserGroupListParams is parameters of getUserGroupList operation.
type GetUserGroupListParams struct {
	Page   OptInt32 `json:",omitempty,omitzero"`
	UserID OptInt64 `json:",omitempty,omitzero"`
}

func unpackGetUserGroupListParams(packed middleware.Parameters) (params GetUserGroupListParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserID = v.(OptInt64)
		}
	}
	return params
}

func decodeGetUserGroupListParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUserGroupListParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: user_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID.SetTo(paramsDotUserIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserInfoParams is parameters of getUserInfo operation.
type GetUserInfoParams struct {
	ID int64
}

func unpackGetUserInfoParams(packed middleware.Parameters) (params GetUserInfoParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetUserInfoParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserInfoParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserPresignedUrlParams is parameters of getUserPresignedUrl operation.
type GetUserPresignedUrlParams struct {
	VideoFileName OptString `json:",omitempty,omitzero"`
}

func unpackGetUserPresignedUrlParams(packed middleware.Parameters) (params GetUserPresignedUrlParams) {
	{
		key := middleware.ParameterKey{
			Name: "video_file_name",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.VideoFileName = v.(OptString)
		}
	}
	return params
}

func decodeGetUserPresignedUrlParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUserPresignedUrlParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: video_file_name.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "video_file_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotVideoFileNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotVideoFileNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.VideoFileName.SetTo(paramsDotVideoFileNameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "video_file_name",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserReviewsParams is parameters of getUserReviews operation.
type GetUserReviewsParams struct {
	ID     int64
	FromID OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackGetUserReviewsParams(packed middleware.Parameters) (params GetUserReviewsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeGetUserReviewsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserReviewsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserTimelineParams is parameters of getUserTimeline operation.
type GetUserTimelineParams struct {
	UserID     OptInt64    `json:",omitempty,omitzero"`
	FromPostID OptNilInt64 `json:",omitempty,omitzero"`
	PostType   OptPostType `json:",omitempty,omitzero"`
	Number     OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackGetUserTimelineParams(packed middleware.Parameters) (params GetUserTimelineParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "post_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PostType = v.(OptPostType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeGetUserTimelineParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUserTimelineParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: user_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID.SetTo(paramsDotUserIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: post_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPostTypeVal PostType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPostTypeVal = PostType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.PostType.SetTo(paramsDotPostTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PostType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetUsersByIdsParams is parameters of getUsersByIds operation.
type GetUsersByIdsParams struct {
	XJwt    OptString `json:",omitempty,omitzero"`
	UserIds []int64   `json:",omitempty"`
}

func unpackGetUsersByIdsParams(packed middleware.Parameters) (params GetUsersByIdsParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Jwt",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.XJwt = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "user_ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserIds = v.([]int64)
		}
	}
	return params
}

func decodeGetUsersByIdsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetUsersByIdsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Jwt.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotXJwtVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotXJwtVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.XJwt.SetTo(paramsDotXJwtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Jwt",
			In:   "header",
			Err:  err,
		}
	}
	// Decode query: user_ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotUserIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotUserIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.UserIds = append(params.UserIds, paramsDotUserIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// InviteToChatRoomParams is parameters of inviteToChatRoom operation.
type InviteToChatRoomParams struct {
	ID int64
}

func unpackInviteToChatRoomParams(packed middleware.Parameters) (params InviteToChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeInviteToChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params InviteToChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// InviteToConferenceCallParams is parameters of inviteToConferenceCall operation.
type InviteToConferenceCallParams struct {
	CallID int64
}

func unpackInviteToConferenceCallParams(packed middleware.Parameters) (params InviteToConferenceCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	return params
}

func decodeInviteToConferenceCallParams(args [1]string, argsEscaped bool, r *http.Request) (params InviteToConferenceCallParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// InviteToGroupParams is parameters of inviteToGroup operation.
type InviteToGroupParams struct {
	ID int64
}

func unpackInviteToGroupParams(packed middleware.Parameters) (params InviteToGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeInviteToGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params InviteToGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// JoinGroupParams is parameters of joinGroup operation.
type JoinGroupParams struct {
	ID int64
}

func unpackJoinGroupParams(packed middleware.Parameters) (params JoinGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeJoinGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params JoinGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// KickFromChatRoomParams is parameters of kickFromChatRoom operation.
type KickFromChatRoomParams struct {
	ID int64
}

func unpackKickFromChatRoomParams(packed middleware.Parameters) (params KickFromChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeKickFromChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params KickFromChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// KickFromConferenceCallParams is parameters of kickFromConferenceCall operation.
type KickFromConferenceCallParams struct {
	CallID int64
}

func unpackKickFromConferenceCallParams(packed middleware.Parameters) (params KickFromConferenceCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	return params
}

func decodeKickFromConferenceCallParams(args [1]string, argsEscaped bool, r *http.Request) (params KickFromConferenceCallParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// LeaveGroupParams is parameters of leaveGroup operation.
type LeaveGroupParams struct {
	ID int64
}

func unpackLeaveGroupParams(packed middleware.Parameters) (params LeaveGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeLeaveGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params LeaveGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListGameAppsParams is parameters of listGameApps operation.
type ListGameAppsParams struct {
	Number OptInt32    `json:",omitempty,omitzero"`
	FromID OptNilInt64 `json:",omitempty,omitzero"`
	Ids    []int64     `json:",omitempty"`
}

func unpackListGameAppsParams(packed middleware.Parameters) (params ListGameAppsParams) {
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.([]int64)
		}
	}
	return params
}

func decodeListGameAppsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGameAppsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromID.SetTo(paramsDotFromIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Ids = append(params.Ids, paramsDotIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGameWalkthroughsParams is parameters of listGameWalkthroughs operation.
type ListGameWalkthroughsParams struct {
	AppID int64
}

func unpackListGameWalkthroughsParams(packed middleware.Parameters) (params ListGameWalkthroughsParams) {
	{
		key := middleware.ParameterKey{
			Name: "app_id",
			In:   "path",
		}
		params.AppID = packed[key].(int64)
	}
	return params
}

func decodeListGameWalkthroughsParams(args [1]string, argsEscaped bool, r *http.Request) (params ListGameWalkthroughsParams, _ error) {
	// Decode path: app_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "app_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.AppID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "app_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListGenresParams is parameters of listGenres operation.
type ListGenresParams struct {
	Number OptInt32     `json:",omitempty,omitzero"`
	From   OptNilString `json:",omitempty,omitzero"`
}

func unpackListGenresParams(packed middleware.Parameters) (params ListGenresParams) {
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeListGenresParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGenresParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGiftsParams is parameters of listGifts operation.
type ListGiftsParams struct {
	Currency OptNilString `json:",omitempty,omitzero"`
	Number   OptNilInt32  `json:",omitempty,omitzero"`
	From     OptNilString `json:",omitempty,omitzero"`
}

func unpackListGiftsParams(packed middleware.Parameters) (params ListGiftsParams) {
	{
		key := middleware.ParameterKey{
			Name: "currency",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Currency = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	return params
}

func decodeListGiftsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGiftsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: currency.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "currency",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCurrencyVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCurrencyVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Currency.SetTo(paramsDotCurrencyVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "currency",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGroupCategoriesParams is parameters of listGroupCategories operation.
type ListGroupCategoriesParams struct {
	Page   OptInt32    `json:",omitempty,omitzero"`
	Number OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackListGroupCategoriesParams(packed middleware.Parameters) (params ListGroupCategoriesParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeListGroupCategoriesParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGroupCategoriesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGroupsParams is parameters of listGroups operation.
type ListGroupsParams struct {
	GroupCategoryID OptNilInt64  `json:",omitempty,omitzero"`
	Keyword         OptNilString `json:",omitempty,omitzero"`
	FromTimestamp   OptNilInt64  `json:",omitempty,omitzero"`
	SubCategoryID   OptNilInt64  `json:",omitempty,omitzero"`
}

func unpackListGroupsParams(packed middleware.Parameters) (params ListGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_category_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupCategoryID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sub_category_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SubCategoryID = v.(OptNilInt64)
		}
	}
	return params
}

func decodeListGroupsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: group_category_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupCategoryIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupCategoryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupCategoryID.SetTo(paramsDotGroupCategoryIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_category_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sub_category_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sub_category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSubCategoryIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotSubCategoryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SubCategoryID.SetTo(paramsDotSubCategoryIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sub_category_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListHiddenChatsParams is parameters of listHiddenChats operation.
type ListHiddenChatsParams struct {
	Number        OptInt32    `json:",omitempty,omitzero"`
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackListHiddenChatsParams(packed middleware.Parameters) (params ListHiddenChatsParams) {
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeListHiddenChatsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListHiddenChatsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListHiddenUsersParams is parameters of listHiddenUsers operation.
type ListHiddenUsersParams struct {
	From   OptNilString `json:",omitempty,omitzero"`
	Number OptNilInt32  `json:",omitempty,omitzero"`
}

func unpackListHiddenUsersParams(packed middleware.Parameters) (params ListHiddenUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeListHiddenUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params ListHiddenUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListMutedGroupUsersParams is parameters of listMutedGroupUsers operation.
type ListMutedGroupUsersParams struct {
	ID      int64
	Keyword OptNilString `json:",omitempty,omitzero"`
	Cursor  OptString    `json:",omitempty,omitzero"`
	Size    OptInt32     `json:",omitempty,omitzero"`
}

func unpackListMutedGroupUsersParams(packed middleware.Parameters) (params ListMutedGroupUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "cursor",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Cursor = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "size",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Size = v.(OptInt32)
		}
	}
	return params
}

func decodeListMutedGroupUsersParams(args [1]string, argsEscaped bool, r *http.Request) (params ListMutedGroupUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: cursor.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "cursor",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCursorVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCursorVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Cursor.SetTo(paramsDotCursorVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cursor",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: size.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSizeVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotSizeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Size.SetTo(paramsDotSizeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "size",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListMyGroupsParams is parameters of listMyGroups operation.
type ListMyGroupsParams struct {
	FromTimestamp OptNilInt64 `json:",omitempty,omitzero"`
}

func unpackListMyGroupsParams(packed middleware.Parameters) (params ListMyGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	return params
}

func decodeListMyGroupsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListMyGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListReceivedGiftsV1Params is parameters of listReceivedGiftsV1 operation.
type ListReceivedGiftsV1Params struct {
	From   OptNilString `json:",omitempty,omitzero"`
	Number OptNilInt32  `json:",omitempty,omitzero"`
}

func unpackListReceivedGiftsV1Params(packed middleware.Parameters) (params ListReceivedGiftsV1Params) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeListReceivedGiftsV1Params(args [0]string, argsEscaped bool, r *http.Request) (params ListReceivedGiftsV1Params, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListThreadsParams is parameters of listThreads operation.
type ListThreadsParams struct {
	GroupID    OptInt64     `json:",omitempty,omitzero"`
	From       OptNilString `json:",omitempty,omitzero"`
	JoinStatus OptNilString `json:",omitempty,omitzero"`
}

func unpackListThreadsParams(packed middleware.Parameters) (params ListThreadsParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "join_status",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.JoinStatus = v.(OptNilString)
		}
	}
	return params
}

func decodeListThreadsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListThreadsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: join_status.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "join_status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotJoinStatusVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotJoinStatusVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.JoinStatus.SetTo(paramsDotJoinStatusVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "join_status",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// MovePostToSpecificThreadParams is parameters of movePostToSpecificThread operation.
type MovePostToSpecificThreadParams struct {
	ID       int64
	ThreadID int64
}

func unpackMovePostToSpecificThreadParams(packed middleware.Parameters) (params MovePostToSpecificThreadParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "thread_id",
			In:   "path",
		}
		params.ThreadID = packed[key].(int64)
	}
	return params
}

func decodeMovePostToSpecificThreadParams(args [2]string, argsEscaped bool, r *http.Request) (params MovePostToSpecificThreadParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: thread_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "thread_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ThreadID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// MovePostToThreadParams is parameters of movePostToThread operation.
type MovePostToThreadParams struct {
	ID int64
}

func unpackMovePostToThreadParams(packed middleware.Parameters) (params MovePostToThreadParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeMovePostToThreadParams(args [1]string, argsEscaped bool, r *http.Request) (params MovePostToThreadParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// MuteGroupUserParams is parameters of muteGroupUser operation.
type MuteGroupUserParams struct {
	ID     int64
	UserID int64
}

func unpackMuteGroupUserParams(packed middleware.Parameters) (params MuteGroupUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeMuteGroupUserParams(args [2]string, argsEscaped bool, r *http.Request) (params MuteGroupUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: user_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PinChatRoomParams is parameters of pinChatRoom operation.
type PinChatRoomParams struct {
	ID int64
}

func unpackPinChatRoomParams(packed middleware.Parameters) (params PinChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodePinChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params PinChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PinGroupHighlightPostParams is parameters of pinGroupHighlightPost operation.
type PinGroupHighlightPostParams struct {
	GroupID int64
	PostID  int64
}

func unpackPinGroupHighlightPostParams(packed middleware.Parameters) (params PinGroupHighlightPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "post_id",
			In:   "path",
		}
		params.PostID = packed[key].(int64)
	}
	return params
}

func decodePinGroupHighlightPostParams(args [2]string, argsEscaped bool, r *http.Request) (params PinGroupHighlightPostParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: post_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "post_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PostID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadChatAttachmentsParams is parameters of readChatAttachments operation.
type ReadChatAttachmentsParams struct {
	ID int64
}

func unpackReadChatAttachmentsParams(packed middleware.Parameters) (params ReadChatAttachmentsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeReadChatAttachmentsParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadChatAttachmentsParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadChatMessageParams is parameters of readChatMessage operation.
type ReadChatMessageParams struct {
	ID        int64
	MessageID int64
}

func unpackReadChatMessageParams(packed middleware.Parameters) (params ReadChatMessageParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "message_id",
			In:   "path",
		}
		params.MessageID = packed[key].(int64)
	}
	return params
}

func decodeReadChatMessageParams(args [2]string, argsEscaped bool, r *http.Request) (params ReadChatMessageParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: message_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "message_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.MessageID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "message_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadChatVideosParams is parameters of readChatVideos operation.
type ReadChatVideosParams struct {
	ID int64
}

func unpackReadChatVideosParams(packed middleware.Parameters) (params ReadChatVideosParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeReadChatVideosParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadChatVideosParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveChatRoomBackgroundParams is parameters of removeChatRoomBackground operation.
type RemoveChatRoomBackgroundParams struct {
	ID int64
}

func unpackRemoveChatRoomBackgroundParams(packed middleware.Parameters) (params RemoveChatRoomBackgroundParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRemoveChatRoomBackgroundParams(args [1]string, argsEscaped bool, r *http.Request) (params RemoveChatRoomBackgroundParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveGroupCoverParams is parameters of removeGroupCover operation.
type RemoveGroupCoverParams struct {
	ID int64
}

func unpackRemoveGroupCoverParams(packed middleware.Parameters) (params RemoveGroupCoverParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRemoveGroupCoverParams(args [1]string, argsEscaped bool, r *http.Request) (params RemoveGroupCoverParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveGroupDeputiesParams is parameters of removeGroupDeputies operation.
type RemoveGroupDeputiesParams struct {
	ID int64
}

func unpackRemoveGroupDeputiesParams(packed middleware.Parameters) (params RemoveGroupDeputiesParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRemoveGroupDeputiesParams(args [1]string, argsEscaped bool, r *http.Request) (params RemoveGroupDeputiesParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveGroupIconParams is parameters of removeGroupIcon operation.
type RemoveGroupIconParams struct {
	ID int64
}

func unpackRemoveGroupIconParams(packed middleware.Parameters) (params RemoveGroupIconParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRemoveGroupIconParams(args [1]string, argsEscaped bool, r *http.Request) (params RemoveGroupIconParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveRelatedGroupsParams is parameters of removeRelatedGroups operation.
type RemoveRelatedGroupsParams struct {
	ID             int64
	RelatedGroupID []int64 `json:",omitempty"`
}

func unpackRemoveRelatedGroupsParams(packed middleware.Parameters) (params RemoveRelatedGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "related_group_id[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.RelatedGroupID = v.([]int64)
		}
	}
	return params
}

func decodeRemoveRelatedGroupsParams(args [1]string, argsEscaped bool, r *http.Request) (params RemoveRelatedGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: related_group_id[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "related_group_id[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotRelatedGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotRelatedGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.RelatedGroupID = append(params.RelatedGroupID, paramsDotRelatedGroupIDVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "related_group_id[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RemoveThreadMemberParams is parameters of removeThreadMember operation.
type RemoveThreadMemberParams struct {
	ThreadID int64
	ID       int64
}

func unpackRemoveThreadMemberParams(packed middleware.Parameters) (params RemoveThreadMemberParams) {
	{
		key := middleware.ParameterKey{
			Name: "thread_id",
			In:   "path",
		}
		params.ThreadID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRemoveThreadMemberParams(args [2]string, argsEscaped bool, r *http.Request) (params RemoveThreadMemberParams, _ error) {
	// Decode path: thread_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "thread_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ThreadID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReplyToReviewParams is parameters of replyToReview operation.
type ReplyToReviewParams struct {
	ID int64
}

func unpackReplyToReviewParams(packed middleware.Parameters) (params ReplyToReviewParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeReplyToReviewParams(args [1]string, argsEscaped bool, r *http.Request) (params ReplyToReviewParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReportChatRoomParams is parameters of reportChatRoom operation.
type ReportChatRoomParams struct {
	ChatRoomID int64
}

func unpackReportChatRoomParams(packed middleware.Parameters) (params ReportChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "chat_room_id",
			In:   "path",
		}
		params.ChatRoomID = packed[key].(int64)
	}
	return params
}

func decodeReportChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params ReportChatRoomParams, _ error) {
	// Decode path: chat_room_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chat_room_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ChatRoomID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chat_room_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReportGroupParams is parameters of reportGroup operation.
type ReportGroupParams struct {
	GroupID int64
}

func unpackReportGroupParams(packed middleware.Parameters) (params ReportGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	return params
}

func decodeReportGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params ReportGroupParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReportPostParams is parameters of reportPost operation.
type ReportPostParams struct {
	PostID int64
}

func unpackReportPostParams(packed middleware.Parameters) (params ReportPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "post_id",
			In:   "path",
		}
		params.PostID = packed[key].(int64)
	}
	return params
}

func decodeReportPostParams(args [1]string, argsEscaped bool, r *http.Request) (params ReportPostParams, _ error) {
	// Decode path: post_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "post_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PostID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReportUserParams is parameters of reportUser operation.
type ReportUserParams struct {
	UserID int64
}

func unpackReportUserParams(packed middleware.Parameters) (params ReportUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeReportUserParams(args [1]string, argsEscaped bool, r *http.Request) (params ReportUserParams, _ error) {
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RepostParams is parameters of repost operation.
type RepostParams struct {
	XJwt OptString `json:",omitempty,omitzero"`
}

func unpackRepostParams(packed middleware.Parameters) (params RepostParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Jwt",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.XJwt = v.(OptString)
		}
	}
	return params
}

func decodeRepostParams(args [0]string, argsEscaped bool, r *http.Request) (params RepostParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Jwt.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotXJwtVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotXJwtVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.XJwt.SetTo(paramsDotXJwtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Jwt",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// RequestFollowParams is parameters of requestFollow operation.
type RequestFollowParams struct {
	TargetID int64
}

func unpackRequestFollowParams(packed middleware.Parameters) (params RequestFollowParams) {
	{
		key := middleware.ParameterKey{
			Name: "target_id",
			In:   "path",
		}
		params.TargetID = packed[key].(int64)
	}
	return params
}

func decodeRequestFollowParams(args [1]string, argsEscaped bool, r *http.Request) (params RequestFollowParams, _ error) {
	// Decode path: target_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "target_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TargetID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "target_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RequestGroupWalkthroughParams is parameters of requestGroupWalkthrough operation.
type RequestGroupWalkthroughParams struct {
	ID int64
}

func unpackRequestGroupWalkthroughParams(packed middleware.Parameters) (params RequestGroupWalkthroughParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeRequestGroupWalkthroughParams(args [1]string, argsEscaped bool, r *http.Request) (params RequestGroupWalkthroughParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// SearchGroupPostsParams is parameters of searchGroupPosts operation.
type SearchGroupPostsParams struct {
	ID              int64
	Keyword         OptString   `json:",omitempty,omitzero"`
	FromPostID      OptNilInt64 `json:",omitempty,omitzero"`
	Number          OptNilInt32 `json:",omitempty,omitzero"`
	OnlyThreadPosts OptBool     `json:",omitempty,omitzero"`
}

func unpackSearchGroupPostsParams(packed middleware.Parameters) (params SearchGroupPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "only_thread_posts",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OnlyThreadPosts = v.(OptBool)
		}
	}
	return params
}

func decodeSearchGroupPostsParams(args [1]string, argsEscaped bool, r *http.Request) (params SearchGroupPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: only_thread_posts.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "only_thread_posts",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOnlyThreadPostsVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotOnlyThreadPostsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OnlyThreadPosts.SetTo(paramsDotOnlyThreadPostsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "only_thread_posts",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SearchPostsParams is parameters of searchPosts operation.
type SearchPostsParams struct {
	Keyword        OptString   `json:",omitempty,omitzero"`
	PostOwnerScope OptString   `json:",omitempty,omitzero"`
	OnlyMedia      OptBool     `json:",omitempty,omitzero"`
	FromPostID     OptNilInt64 `json:",omitempty,omitzero"`
	Number         OptNilInt32 `json:",omitempty,omitzero"`
}

func unpackSearchPostsParams(packed middleware.Parameters) (params SearchPostsParams) {
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "post_owner_scope",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PostOwnerScope = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "only_media",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.OnlyMedia = v.(OptBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_post_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromPostID = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "number",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Number = v.(OptNilInt32)
		}
	}
	return params
}

func decodeSearchPostsParams(args [0]string, argsEscaped bool, r *http.Request) (params SearchPostsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: post_owner_scope.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "post_owner_scope",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPostOwnerScopeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPostOwnerScopeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PostOwnerScope.SetTo(paramsDotPostOwnerScopeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_owner_scope",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: only_media.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "only_media",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOnlyMediaVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotOnlyMediaVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OnlyMedia.SetTo(paramsDotOnlyMediaVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "only_media",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_post_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromPostIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromPostIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromPostID.SetTo(paramsDotFromPostIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_post_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: number.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNumberVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotNumberVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Number.SetTo(paramsDotNumberVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "number",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SearchUsersParams is parameters of searchUsers operation.
type SearchUsersParams struct {
	Gender             OptNilInt32  `json:",omitempty,omitzero"`
	Nickname           OptNilString `json:",omitempty,omitzero"`
	Title              OptNilString `json:",omitempty,omitzero"`
	Biography          OptNilString `json:",omitempty,omitzero"`
	FromTimestamp      OptNilInt64  `json:",omitempty,omitzero"`
	SimilarAge         OptNilBool   `json:",omitempty,omitzero"`
	NotRecentGomimushi OptNilBool   `json:",omitempty,omitzero"`
	RecentlyCreated    OptNilBool   `json:",omitempty,omitzero"`
	SamePrefecture     OptNilBool   `json:",omitempty,omitzero"`
	SaveRecentSearch   OptNilBool   `json:",omitempty,omitzero"`
}

func unpackSearchUsersParams(packed middleware.Parameters) (params SearchUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "gender",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Gender = v.(OptNilInt32)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "nickname",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Nickname = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "title",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Title = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "biography",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Biography = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from_timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FromTimestamp = v.(OptNilInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "similar_age",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SimilarAge = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "not_recent_gomimushi",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.NotRecentGomimushi = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "recently_created",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.RecentlyCreated = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "same_prefecture",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SamePrefecture = v.(OptNilBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "save_recent_search",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SaveRecentSearch = v.(OptNilBool)
		}
	}
	return params
}

func decodeSearchUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params SearchUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: gender.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "gender",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGenderVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotGenderVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Gender.SetTo(paramsDotGenderVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "gender",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: nickname.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "nickname",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNicknameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotNicknameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Nickname.SetTo(paramsDotNicknameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "nickname",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: title.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "title",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTitleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotTitleVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Title.SetTo(paramsDotTitleVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "title",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: biography.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "biography",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBiographyVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotBiographyVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Biography.SetTo(paramsDotBiographyVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "biography",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: from_timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFromTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FromTimestamp.SetTo(paramsDotFromTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from_timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: similar_age.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "similar_age",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSimilarAgeVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotSimilarAgeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SimilarAge.SetTo(paramsDotSimilarAgeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "similar_age",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: not_recent_gomimushi.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "not_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNotRecentGomimushiVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotNotRecentGomimushiVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.NotRecentGomimushi.SetTo(paramsDotNotRecentGomimushiVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "not_recent_gomimushi",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: recently_created.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "recently_created",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotRecentlyCreatedVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotRecentlyCreatedVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.RecentlyCreated.SetTo(paramsDotRecentlyCreatedVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "recently_created",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: same_prefecture.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "same_prefecture",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSamePrefectureVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotSamePrefectureVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SamePrefecture.SetTo(paramsDotSamePrefectureVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "same_prefecture",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: save_recent_search.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "save_recent_search",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSaveRecentSearchVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotSaveRecentSearchVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SaveRecentSearch.SetTo(paramsDotSaveRecentSearchVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "save_recent_search",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SendChatMessageParams is parameters of sendChatMessage operation.
type SendChatMessageParams struct {
	ID int64
}

func unpackSendChatMessageParams(packed middleware.Parameters) (params SendChatMessageParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeSendChatMessageParams(args [1]string, argsEscaped bool, r *http.Request) (params SendChatMessageParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// SetGroupTitleParams is parameters of setGroupTitle operation.
type SetGroupTitleParams struct {
	ID int64
}

func unpackSetGroupTitleParams(packed middleware.Parameters) (params SetGroupTitleParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeSetGroupTitleParams(args [1]string, argsEscaped bool, r *http.Request) (params SetGroupTitleParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// TakeOverGroupParams is parameters of takeOverGroup operation.
type TakeOverGroupParams struct {
	ID int64
}

func unpackTakeOverGroupParams(packed middleware.Parameters) (params TakeOverGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeTakeOverGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params TakeOverGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// TransferGroupParams is parameters of transferGroup operation.
type TransferGroupParams struct {
	ID int64
}

func unpackTransferGroupParams(packed middleware.Parameters) (params TransferGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeTransferGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params TransferGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnbanGroupUserParams is parameters of unbanGroupUser operation.
type UnbanGroupUserParams struct {
	ID     int64
	UserId int64
}

func unpackUnbanGroupUserParams(packed middleware.Parameters) (params UnbanGroupUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(int64)
	}
	return params
}

func decodeUnbanGroupUserParams(args [2]string, argsEscaped bool, r *http.Request) (params UnbanGroupUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnblockUserParams is parameters of unblockUser operation.
type UnblockUserParams struct {
	ID int64
}

func unpackUnblockUserParams(packed middleware.Parameters) (params UnblockUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnblockUserParams(args [1]string, argsEscaped bool, r *http.Request) (params UnblockUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnfollowUserParams is parameters of unfollowUser operation.
type UnfollowUserParams struct {
	ID int64
}

func unpackUnfollowUserParams(packed middleware.Parameters) (params UnfollowUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnfollowUserParams(args [1]string, argsEscaped bool, r *http.Request) (params UnfollowUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnhideChatsParams is parameters of unhideChats operation.
type UnhideChatsParams struct {
	ChatRoomIds OptInt64 `json:",omitempty,omitzero"`
}

func unpackUnhideChatsParams(packed middleware.Parameters) (params UnhideChatsParams) {
	{
		key := middleware.ParameterKey{
			Name: "chat_room_ids",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ChatRoomIds = v.(OptInt64)
		}
	}
	return params
}

func decodeUnhideChatsParams(args [0]string, argsEscaped bool, r *http.Request) (params UnhideChatsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: chat_room_ids.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chat_room_ids",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatRoomIdsVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotChatRoomIdsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatRoomIds.SetTo(paramsDotChatRoomIdsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chat_room_ids",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UnhideUsersParams is parameters of unhideUsers operation.
type UnhideUsersParams struct {
	UserIds []int64 `json:",omitempty"`
}

func unpackUnhideUsersParams(packed middleware.Parameters) (params UnhideUsersParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_ids[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserIds = v.([]int64)
		}
	}
	return params
}

func decodeUnhideUsersParams(args [0]string, argsEscaped bool, r *http.Request) (params UnhideUsersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: user_ids[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotUserIdsVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotUserIdsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.UserIds = append(params.UserIds, paramsDotUserIdsVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_ids[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UnlikePostParams is parameters of unlikePost operation.
type UnlikePostParams struct {
	ID int64
}

func unpackUnlikePostParams(packed middleware.Parameters) (params UnlikePostParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnlikePostParams(args [1]string, argsEscaped bool, r *http.Request) (params UnlikePostParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnmuteGroupUserParams is parameters of unmuteGroupUser operation.
type UnmuteGroupUserParams struct {
	ID     int64
	UserID int64
}

func unpackUnmuteGroupUserParams(packed middleware.Parameters) (params UnmuteGroupUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeUnmuteGroupUserParams(args [2]string, argsEscaped bool, r *http.Request) (params UnmuteGroupUserParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: user_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnpinChatRoomParams is parameters of unpinChatRoom operation.
type UnpinChatRoomParams struct {
	ID int64
}

func unpackUnpinChatRoomParams(packed middleware.Parameters) (params UnpinChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnpinChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params UnpinChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnpinGroupParams is parameters of unpinGroup operation.
type UnpinGroupParams struct {
	ID int64
}

func unpackUnpinGroupParams(packed middleware.Parameters) (params UnpinGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnpinGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params UnpinGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnpinGroupHighlightPostParams is parameters of unpinGroupHighlightPost operation.
type UnpinGroupHighlightPostParams struct {
	GroupID int64
	PostID  int64
}

func unpackUnpinGroupHighlightPostParams(packed middleware.Parameters) (params UnpinGroupHighlightPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "post_id",
			In:   "path",
		}
		params.PostID = packed[key].(int64)
	}
	return params
}

func decodeUnpinGroupHighlightPostParams(args [2]string, argsEscaped bool, r *http.Request) (params UnpinGroupHighlightPostParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: post_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "post_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PostID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "post_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UnpinGroupPostParams is parameters of unpinGroupPost operation.
type UnpinGroupPostParams struct {
	GroupID OptInt64 `json:",omitempty,omitzero"`
}

func unpackUnpinGroupPostParams(packed middleware.Parameters) (params UnpinGroupPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupID = v.(OptInt64)
		}
	}
	return params
}

func decodeUnpinGroupPostParams(args [0]string, argsEscaped bool, r *http.Request) (params UnpinGroupPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: group_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotGroupIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.GroupID.SetTo(paramsDotGroupIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UnpinReviewParams is parameters of unpinReview operation.
type UnpinReviewParams struct {
	ID int64
}

func unpackUnpinReviewParams(packed middleware.Parameters) (params UnpinReviewParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUnpinReviewParams(args [1]string, argsEscaped bool, r *http.Request) (params UnpinReviewParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateCallParams is parameters of updateCall operation.
type UpdateCallParams struct {
	CallID int64
}

func unpackUpdateCallParams(packed middleware.Parameters) (params UpdateCallParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	return params
}

func decodeUpdateCallParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateCallParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateCallUserParams is parameters of updateCallUser operation.
type UpdateCallUserParams struct {
	CallID int64
	UserID int64
}

func unpackUpdateCallUserParams(packed middleware.Parameters) (params UpdateCallUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "path",
		}
		params.CallID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeUpdateCallUserParams(args [2]string, argsEscaped bool, r *http.Request) (params UpdateCallUserParams, _ error) {
	// Decode path: call_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "call_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CallID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: user_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateChatRoomParams is parameters of updateChatRoom operation.
type UpdateChatRoomParams struct {
	ID int64
}

func unpackUpdateChatRoomParams(packed middleware.Parameters) (params UpdateChatRoomParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdateChatRoomParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateChatRoomParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateChatRoomNotificationSettingsParams is parameters of updateChatRoomNotificationSettings operation.
type UpdateChatRoomNotificationSettingsParams struct {
	ID int64
}

func unpackUpdateChatRoomNotificationSettingsParams(packed middleware.Parameters) (params UpdateChatRoomNotificationSettingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdateChatRoomNotificationSettingsParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateChatRoomNotificationSettingsParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateGroupParams is parameters of updateGroup operation.
type UpdateGroupParams struct {
	ID int64
}

func unpackUpdateGroupParams(packed middleware.Parameters) (params UpdateGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdateGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateGroupNotificationSettingsParams is parameters of updateGroupNotificationSettings operation.
type UpdateGroupNotificationSettingsParams struct {
	ID int64
}

func unpackUpdateGroupNotificationSettingsParams(packed middleware.Parameters) (params UpdateGroupNotificationSettingsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdateGroupNotificationSettingsParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateGroupNotificationSettingsParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdatePostParams is parameters of updatePost operation.
type UpdatePostParams struct {
	ID int64
}

func unpackUpdatePostParams(packed middleware.Parameters) (params UpdatePostParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdatePostParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdatePostParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateRelatedGroupsParams is parameters of updateRelatedGroups operation.
type UpdateRelatedGroupsParams struct {
	ID             int64
	RelatedGroupID []int64 `json:",omitempty"`
}

func unpackUpdateRelatedGroupsParams(packed middleware.Parameters) (params UpdateRelatedGroupsParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "related_group_id[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.RelatedGroupID = v.([]int64)
		}
	}
	return params
}

func decodeUpdateRelatedGroupsParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateRelatedGroupsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: related_group_id[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "related_group_id[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotRelatedGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						paramsDotRelatedGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.RelatedGroupID = append(params.RelatedGroupID, paramsDotRelatedGroupIDVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "related_group_id[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateThreadParams is parameters of updateThread operation.
type UpdateThreadParams struct {
	ID int64
}

func unpackUpdateThreadParams(packed middleware.Parameters) (params UpdateThreadParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeUpdateThreadParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateThreadParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ValidateCallActionSignatureParams is parameters of validateCallActionSignature operation.
type ValidateCallActionSignatureParams struct {
	CallID       OptInt64  `json:",omitempty,omitzero"`
	SenderUUID   OptString `json:",omitempty,omitzero"`
	ReceiverUUID OptString `json:",omitempty,omitzero"`
	Action       OptString `json:",omitempty,omitzero"`
	Timestamp    OptInt64  `json:",omitempty,omitzero"`
	Signature    OptString `json:",omitempty,omitzero"`
}

func unpackValidateCallActionSignatureParams(packed middleware.Parameters) (params ValidateCallActionSignatureParams) {
	{
		key := middleware.ParameterKey{
			Name: "call_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CallID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sender_uuid",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SenderUUID = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "receiver_uuid",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ReceiverUUID = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "action",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Action = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "timestamp",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Timestamp = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "signature",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Signature = v.(OptString)
		}
	}
	return params
}

func decodeValidateCallActionSignatureParams(args [0]string, argsEscaped bool, r *http.Request) (params ValidateCallActionSignatureParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: call_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "call_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCallIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotCallIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CallID.SetTo(paramsDotCallIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "call_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sender_uuid.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sender_uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSenderUUIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSenderUUIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SenderUUID.SetTo(paramsDotSenderUUIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sender_uuid",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: receiver_uuid.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "receiver_uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotReceiverUUIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotReceiverUUIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ReceiverUUID.SetTo(paramsDotReceiverUUIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "receiver_uuid",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: action.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "action",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotActionVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotActionVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Action.SetTo(paramsDotActionVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "action",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: timestamp.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTimestampVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotTimestampVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Timestamp.SetTo(paramsDotTimestampVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "timestamp",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: signature.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "signature",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSignatureVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSignatureVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Signature.SetTo(paramsDotSignatureVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "signature",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ViewPostVideoParams is parameters of viewPostVideo operation.
type ViewPostVideoParams struct {
	VideoId int64
}

func unpackViewPostVideoParams(packed middleware.Parameters) (params ViewPostVideoParams) {
	{
		key := middleware.ParameterKey{
			Name: "videoId",
			In:   "path",
		}
		params.VideoId = packed[key].(int64)
	}
	return params
}

func decodeViewPostVideoParams(args [1]string, argsEscaped bool, r *http.Request) (params ViewPostVideoParams, _ error) {
	// Decode path: videoId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "videoId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.VideoId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "videoId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// VisitGroupParams is parameters of visitGroup operation.
type VisitGroupParams struct {
	ID int64
}

func unpackVisitGroupParams(packed middleware.Parameters) (params VisitGroupParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeVisitGroupParams(args [1]string, argsEscaped bool, r *http.Request) (params VisitGroupParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// VoteSurveyParams is parameters of voteSurvey operation.
type VoteSurveyParams struct {
	ID int64
}

func unpackVoteSurveyParams(packed middleware.Parameters) (params VoteSurveyParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeVoteSurveyParams(args [1]string, argsEscaped bool, r *http.Request) (params VoteSurveyParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// WithdrawGroupDeputyParams is parameters of withdrawGroupDeputy operation.
type WithdrawGroupDeputyParams struct {
	GroupID int64
	UserID  int64
}

func unpackWithdrawGroupDeputyParams(packed middleware.Parameters) (params WithdrawGroupDeputyParams) {
	{
		key := middleware.ParameterKey{
			Name: "group_id",
			In:   "path",
		}
		params.GroupID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(int64)
	}
	return params
}

func decodeWithdrawGroupDeputyParams(args [2]string, argsEscaped bool, r *http.Request) (params WithdrawGroupDeputyParams, _ error) {
	// Decode path: group_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "group_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.GroupID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "group_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: user_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.UserID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// WithdrawGroupTransferParams is parameters of withdrawGroupTransfer operation.
type WithdrawGroupTransferParams struct {
	ID int64
}

func unpackWithdrawGroupTransferParams(packed middleware.Parameters) (params WithdrawGroupTransferParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeWithdrawGroupTransferParams(args [1]string, argsEscaped bool, r *http.Request) (params WithdrawGroupTransferParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
