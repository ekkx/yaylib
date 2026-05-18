// Code generated; DO NOT EDIT.

package oas

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeAcceptChatRequestRequest(r *http.Request) (
	req *AcceptChatRequestReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request AcceptChatRequestReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "chat_room_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotChatRoomIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotChatRoomIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.ChatRoomIds = append(request.ChatRoomIds, requestDotChatRoomIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"chat_room_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeChangeEmailRequest(r *http.Request) (
	req *ChangeEmailReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ChangeEmailReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Email.SetTo(requestDotEmailVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email_grant_token",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailGrantTokenVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailGrantTokenVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.EmailGrantToken.SetTo(requestDotEmailGrantTokenVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email_grant_token\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Password.SetTo(requestDotPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"password\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeChangePasswordRequest(r *http.Request) (
	req *ChangePasswordReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ChangePasswordReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "current_password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCurrentPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCurrentPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CurrentPassword.SetTo(requestDotCurrentPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"current_password\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Password.SetTo(requestDotPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"password\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateChatRoomRequest(r *http.Request) (
	req *CreateChatRoomReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateChatRoomReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "background_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotBackgroundFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotBackgroundFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.BackgroundFilename.SetTo(requestDotBackgroundFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"background_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IconFilename.SetTo(requestDotIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Name.SetTo(requestDotNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"name\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "with_user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotWithUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotWithUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.WithUserIds = append(request.WithUserIds, requestDotWithUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"with_user_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateChatRoomV1Request(r *http.Request) (
	req *CreateChatRoomV1Req,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateChatRoomV1Req
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hima_chat",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHimaChatVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHimaChatVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HimaChat.SetTo(requestDotHimaChatVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hima_chat\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "matching_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMatchingIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotMatchingIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.MatchingID.SetTo(requestDotMatchingIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"matching_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "with_user_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotWithUserIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotWithUserIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.WithUserID.SetTo(requestDotWithUserIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"with_user_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateConferenceCallPostRequest(r *http.Request) (
	req *CreateConferenceCallPostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateConferenceCallPostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment2Filename.SetTo(requestDotAttachment2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment3Filename.SetTo(requestDotAttachment3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment4Filename.SetTo(requestDotAttachment4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_5_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment5FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment5FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment5Filename.SetTo(requestDotAttachment5FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_5_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_6_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment6FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment6FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment6Filename.SetTo(requestDotAttachment6FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_6_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_7_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment7FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment7FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment7Filename.SetTo(requestDotAttachment7FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_7_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_8_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment8FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment8FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment8Filename.SetTo(requestDotAttachment8FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_8_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_9_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment9FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment9FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment9Filename.SetTo(requestDotAttachment9FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_9_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachmentFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachmentFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AttachmentFilename.SetTo(requestDotAttachmentFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallTypeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCallTypeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallType.SetTo(requestDotCallTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_type\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "game_title",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGameTitleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGameTitleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GameTitle.SetTo(requestDotGameTitleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"game_title\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "joinable_by",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotJoinableByVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotJoinableByVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.JoinableBy.SetTo(requestDotJoinableByVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"joinable_by\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateGroupRequest(r *http.Request) (
	req *CreateGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_members_to_post_image_and_video",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowMembersToPostImageAndVideoVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowMembersToPostImageAndVideoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowMembersToPostImageAndVideo.SetTo(requestDotAllowMembersToPostImageAndVideoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_members_to_post_image_and_video\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_members_to_post_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowMembersToPostURLVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowMembersToPostURLVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowMembersToPostURL.SetTo(requestDotAllowMembersToPostURLVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_members_to_post_url\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_ownership_transfer",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowOwnershipTransferVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowOwnershipTransferVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowOwnershipTransfer.SetTo(requestDotAllowOwnershipTransferVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_ownership_transfer\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_thread_creation_by",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowThreadCreationByVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAllowThreadCreationByVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowThreadCreationBy.SetTo(requestDotAllowThreadCreationByVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_thread_creation_by\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_timeline_display",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallTimelineDisplayVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotCallTimelineDisplayVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallTimelineDisplay.SetTo(requestDotCallTimelineDisplayVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_timeline_display\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "cover_image_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCoverImageFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCoverImageFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CoverImageFilename.SetTo(requestDotCoverImageFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"cover_image_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDescriptionVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotDescriptionVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Description.SetTo(requestDotDescriptionVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"description\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "gender",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGenderVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotGenderVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Gender.SetTo(requestDotGenderVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"gender\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "generation_groups_limit",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGenerationGroupsLimitVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotGenerationGroupsLimitVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GenerationGroupsLimit.SetTo(requestDotGenerationGroupsLimitVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"generation_groups_limit\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupCategoryID.SetTo(requestDotGroupCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGroupIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupIconFilename.SetTo(requestDotGroupIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "guidelines",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGuidelinesVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGuidelinesVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Guidelines.SetTo(requestDotGuidelinesVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"guidelines\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_conference_call",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideConferenceCallVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideConferenceCallVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideConferenceCall.SetTo(requestDotHideConferenceCallVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_conference_call\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_from_game_eight",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideFromGameEightVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideFromGameEightVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideFromGameEight.SetTo(requestDotHideFromGameEightVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_from_game_eight\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_reported_posts",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideReportedPostsVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideReportedPostsVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideReportedPosts.SetTo(requestDotHideReportedPostsVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_reported_posts\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "is_private",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIsPrivateVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotIsPrivateVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IsPrivate.SetTo(requestDotIsPrivateVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"is_private\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "only_mobile_verified",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOnlyMobileVerifiedVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotOnlyMobileVerifiedVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OnlyMobileVerified.SetTo(requestDotOnlyMobileVerifiedVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"only_mobile_verified\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "only_verified_age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOnlyVerifiedAgeVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotOnlyVerifiedAgeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OnlyVerifiedAge.SetTo(requestDotOnlyVerifiedAgeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"only_verified_age\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "secret",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSecretVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotSecretVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Secret.SetTo(requestDotSecretVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"secret\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sub_category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSubCategoryIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSubCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SubCategoryID.SetTo(requestDotSubCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"sub_category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "topic",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTopicVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTopicVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Topic.SetTo(requestDotTopicVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"topic\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateMuteKeywordRequest(r *http.Request) (
	req *MuteKeywordRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request MuteKeywordRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, rawBody, close, errors.Wrap(err, "validate")
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreatePostRequest(r *http.Request) (
	req *CreatePostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreatePostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment2Filename.SetTo(requestDotAttachment2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment3Filename.SetTo(requestDotAttachment3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment4Filename.SetTo(requestDotAttachment4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_5_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment5FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment5FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment5Filename.SetTo(requestDotAttachment5FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_5_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_6_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment6FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment6FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment6Filename.SetTo(requestDotAttachment6FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_6_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_7_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment7FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment7FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment7Filename.SetTo(requestDotAttachment7FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_7_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_8_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment8FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment8FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment8Filename.SetTo(requestDotAttachment8FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_8_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_9_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment9FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment9FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment9Filename.SetTo(requestDotAttachment9FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_9_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachmentFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachmentFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AttachmentFilename.SetTo(requestDotAttachmentFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "choices[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChoicesVal []string
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotChoicesValVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								requestDotChoicesValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotChoicesVal = append(requestDotChoicesVal, requestDotChoicesValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.Choices.SetTo(requestDotChoicesVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"choices[]\"")
				}
				if err := func() error {
					if value, ok := request.Choices.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "in_reply_to",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotInReplyToVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotInReplyToVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.InReplyTo.SetTo(requestDotInReplyToVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"in_reply_to\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "mention_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMentionIdsVal []int64
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotMentionIdsValVal int64
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToInt64(val)
								if err != nil {
									return err
								}

								requestDotMentionIdsValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotMentionIdsVal = append(requestDotMentionIdsVal, requestDotMentionIdsValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.MentionIds.SetTo(requestDotMentionIdsVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"mention_ids[]\"")
				}
				if err := func() error {
					if value, ok := request.MentionIds.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPostTypeVal PostType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPostTypeVal = PostType(c)
						return nil
					}(); err != nil {
						return err
					}
					request.PostType.SetTo(requestDotPostTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_type\"")
				}
				if err := func() error {
					if value, ok := request.PostType.Get(); ok {
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
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "shared_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.SharedURL.Reset()
						if err := request.SharedURL.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"shared_url\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "video_file_name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotVideoFileNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotVideoFileNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.VideoFileName.SetTo(requestDotVideoFileNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"video_file_name\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateReviewRequest(r *http.Request) (
	req *CreateReviewReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateReviewReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "comment",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCommentVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCommentVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Comment.SetTo(requestDotCommentVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"comment\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.UserIds = append(request.UserIds, requestDotUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_ids[]\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateSharePostRequest(r *http.Request) (
	req *CreateSharePostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateSharePostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "shareable_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotShareableIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotShareableIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ShareableID.SetTo(requestDotShareableIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"shareable_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "shareable_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotShareableTypeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotShareableTypeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ShareableType.SetTo(requestDotShareableTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"shareable_type\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateThreadRequest(r *http.Request) (
	req *CreateGroupThreadRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request CreateGroupThreadRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateThreadPostRequest(r *http.Request) (
	req *CreateThreadPostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request CreateThreadPostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment2Filename.SetTo(requestDotAttachment2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment3Filename.SetTo(requestDotAttachment3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment4Filename.SetTo(requestDotAttachment4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_5_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment5FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment5FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment5Filename.SetTo(requestDotAttachment5FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_5_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_6_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment6FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment6FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment6Filename.SetTo(requestDotAttachment6FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_6_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_7_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment7FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment7FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment7Filename.SetTo(requestDotAttachment7FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_7_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_8_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment8FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment8FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment8Filename.SetTo(requestDotAttachment8FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_8_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_9_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment9FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment9FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment9Filename.SetTo(requestDotAttachment9FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_9_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachmentFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachmentFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AttachmentFilename.SetTo(requestDotAttachmentFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "choices[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChoicesVal []string
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotChoicesValVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								requestDotChoicesValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotChoicesVal = append(requestDotChoicesVal, requestDotChoicesValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.Choices.SetTo(requestDotChoicesVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"choices[]\"")
				}
				if err := func() error {
					if value, ok := request.Choices.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "in_reply_to",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotInReplyToVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotInReplyToVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.InReplyTo.SetTo(requestDotInReplyToVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"in_reply_to\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "mention_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMentionIdsVal []int64
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotMentionIdsValVal int64
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToInt64(val)
								if err != nil {
									return err
								}

								requestDotMentionIdsValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotMentionIdsVal = append(requestDotMentionIdsVal, requestDotMentionIdsValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.MentionIds.SetTo(requestDotMentionIdsVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"mention_ids[]\"")
				}
				if err := func() error {
					if value, ok := request.MentionIds.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPostTypeVal PostType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPostTypeVal = PostType(c)
						return nil
					}(); err != nil {
						return err
					}
					request.PostType.SetTo(requestDotPostTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_type\"")
				}
				if err := func() error {
					if value, ok := request.PostType.Get(); ok {
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
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "shared_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.SharedURL.Reset()
						if err := request.SharedURL.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"shared_url\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "video_file_name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotVideoFileNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotVideoFileNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.VideoFileName.SetTo(requestDotVideoFileNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"video_file_name\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeDeleteChatRoomsRequest(r *http.Request) (
	req *DeleteChatRoomsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request DeleteChatRoomsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "chat_room_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotChatRoomIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotChatRoomIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.ChatRoomIds = append(request.ChatRoomIds, requestDotChatRoomIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"chat_room_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeDeletePostsRequest(r *http.Request) (
	req *DeletePostsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request DeletePostsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "posts_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotPostsIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotPostsIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.PostsIds = append(request.PostsIds, requestDotPostsIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"posts_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeDeputizeGroupUsersMassRequest(r *http.Request) (
	req *DeputizeGroupUsersMassReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request DeputizeGroupUsersMassReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.UserIds = append(request.UserIds, requestDotUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_ids[]\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeDisableTwoFactorAuthRequest(r *http.Request) (
	req *DisableTwoFactorAuthReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request DisableTwoFactorAuthReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "code",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCodeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCodeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Code.SetTo(requestDotCodeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"code\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeEditUserRequest(r *http.Request) (
	req *EditUserReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request EditUserReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "biography",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotBiographyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotBiographyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Biography.SetTo(requestDotBiographyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"biography\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "country_code",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCountryCodeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCountryCodeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CountryCode.SetTo(requestDotCountryCodeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"country_code\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "cover_image_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCoverImageFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCoverImageFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CoverImageFilename.SetTo(requestDotCoverImageFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"cover_image_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "gender",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGenderVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotGenderVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Gender.SetTo(requestDotGenderVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"gender\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "nickname",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNicknameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotNicknameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Nickname.SetTo(requestDotNicknameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"nickname\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "prefecture",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPrefectureVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPrefectureVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Prefecture.SetTo(requestDotPrefectureVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"prefecture\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "profile_icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotProfileIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotProfileIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ProfileIconFilename.SetTo(requestDotProfileIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"profile_icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "username",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUsernameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUsernameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Username.SetTo(requestDotUsernameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"username\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeEditUserV2Request(r *http.Request) (
	req *EditUserV2Req,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request EditUserV2Req
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "is_private",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIsPrivateVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotIsPrivateVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IsPrivate.SetTo(requestDotIsPrivateVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"is_private\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "nickname",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNicknameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotNicknameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Nickname.SetTo(requestDotNicknameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"nickname\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeEnableTwoFactorAuthRequest(r *http.Request) (
	req *EnableTwoFactorAuthReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request EnableTwoFactorAuthReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "code",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCodeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCodeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Code.SetTo(requestDotCodeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"code\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTypeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTypeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Type.SetTo(requestDotTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"type\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeGenerateCallActionSignatureRequest(r *http.Request) (
	req *GenerateCallActionSignatureReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request GenerateCallActionSignatureReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "action",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotActionVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotActionVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Action.SetTo(requestDotActionVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"action\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "conference_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotConferenceIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotConferenceIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ConferenceID.SetTo(requestDotConferenceIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"conference_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "target_user_uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTargetUserUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTargetUserUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.TargetUserUUID.SetTo(requestDotTargetUserUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"target_user_uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeGetBlockedUsersRequest(r *http.Request) (
	req *SearchUsersRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request SearchUsersRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeGetChatableFollowingsRequest(r *http.Request) (
	req *SearchUsersRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request SearchUsersRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeGetRecommendedPostTagsRequest(r *http.Request) (
	req *GetRecommendedPostTagsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request GetRecommendedPostTagsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "save_recent_search",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSaveRecentSearchVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotSaveRecentSearchVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SaveRecentSearch.SetTo(requestDotSaveRecentSearchVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"save_recent_search\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "tag",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTagVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTagVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Tag.SetTo(requestDotTagVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"tag\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeHideChatsRequest(r *http.Request) (
	req *HideChatsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request HideChatsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "chat_room_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChatRoomIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotChatRoomIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ChatRoomID.SetTo(requestDotChatRoomIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"chat_room_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeHideUsersRequest(r *http.Request) (
	req *HideUsersReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request HideUsersReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotUserIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserID.SetTo(requestDotUserIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeInviteToCallRequest(r *http.Request) (
	req *InviteToCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request InviteToCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "chat_room_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChatRoomIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotChatRoomIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ChatRoomID.SetTo(requestDotChatRoomIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"chat_room_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "room_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotRoomIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotRoomIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.RoomID.SetTo(requestDotRoomIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"room_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "room_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotRoomURLVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotRoomURLVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.RoomURL.SetTo(requestDotRoomURLVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"room_url\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeInviteToChatRoomRequest(r *http.Request) (
	req *InviteToChatRoomReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request InviteToChatRoomReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "with_user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotWithUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotWithUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.WithUserIds = append(request.WithUserIds, requestDotWithUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"with_user_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeInviteToConferenceCallRequest(r *http.Request) (
	req *InviteToConferenceCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request InviteToConferenceCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.UserIds = append(request.UserIds, requestDotUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeInviteToGroupRequest(r *http.Request) (
	req *InviteToGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request InviteToGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.UserIds = append(request.UserIds, requestDotUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeKickFromChatRoomRequest(r *http.Request) (
	req *KickFromChatRoomReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request KickFromChatRoomReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "with_user_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotWithUserIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotWithUserIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.WithUserIds = append(request.WithUserIds, requestDotWithUserIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"with_user_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeKickFromConferenceCallRequest(r *http.Request) (
	req *KickFromConferenceCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request KickFromConferenceCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "ban",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotBanVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotBanVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Ban.SetTo(requestDotBanVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"ban\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLeaveAgoraChannelRequest(r *http.Request) (
	req *LeaveAgoraChannelReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request LeaveAgoraChannelReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "conference_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotConferenceIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotConferenceIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ConferenceID.SetTo(requestDotConferenceIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"conference_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotUserIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserID.SetTo(requestDotUserIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLeaveConferenceCallRequest(r *http.Request) (
	req *LeaveConferenceCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request LeaveConferenceCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_sid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallSidVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCallSidVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallSid.SetTo(requestDotCallSidVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_sid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "conference_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotConferenceIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotConferenceIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ConferenceID.SetTo(requestDotConferenceIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"conference_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "duration",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDurationVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotDurationVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Duration.SetTo(requestDotDurationVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"duration\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "total_users_in_call",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTotalUsersInCallVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotTotalUsersInCallVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.TotalUsersInCall.SetTo(requestDotTotalUsersInCallVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"total_users_in_call\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLikePostsRequest(r *http.Request) (
	req *LikePostsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request LikePostsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotPostIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotPostIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.PostIds = append(request.PostIds, requestDotPostIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLoginWithEmailRequest(r *http.Request) (
	req *LoginEmailUserRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request LoginEmailUserRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLogoutRequest(r *http.Request) (
	req *LogoutReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request LogoutReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeMovePostToThreadRequest(r *http.Request) (
	req *MovePostToThreadReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request MovePostToThreadReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "thread_icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotThreadIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotThreadIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ThreadIconFilename.SetTo(requestDotThreadIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"thread_icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "title",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTitleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTitleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Title.SetTo(requestDotTitleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"title\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeOauthTokenRequest(r *http.Request) (
	req *OauthTokenReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse form")
		}

		var request OauthTokenReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Email.SetTo(requestDotEmailVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "grant_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGrantTypeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGrantTypeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GrantType.SetTo(requestDotGrantTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"grant_type\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Password.SetTo(requestDotPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"password\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "refresh_token",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotRefreshTokenVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotRefreshTokenVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.RefreshToken.SetTo(requestDotRefreshTokenVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"refresh_token\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePinGroupRequest(r *http.Request) (
	req *PinGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request PinGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePinGroupPostRequest(r *http.Request) (
	req *PinGroupPostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request PinGroupPostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPostIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotPostIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.PostID.SetTo(requestDotPostIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePinPostRequest(r *http.Request) (
	req *PinPostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request PinPostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePinReviewRequest(r *http.Request) (
	req *PinReviewReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request PinReviewReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReadChatAttachmentsRequest(r *http.Request) (
	req *ReadAttachmentRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request ReadAttachmentRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, rawBody, close, errors.Wrap(err, "validate")
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReadChatVideosRequest(r *http.Request) (
	req *ReadChatVideosReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReadChatVideosReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "video_msg_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotVideoMsgIdsVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							requestDotVideoMsgIdsVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.VideoMsgIds = append(request.VideoMsgIds, requestDotVideoMsgIdsVal)
						return nil
					})
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"video_msg_ids[]\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReplyToReviewRequest(r *http.Request) (
	req *ReplyToReviewReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReplyToReviewReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "comment",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCommentVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCommentVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Comment.SetTo(requestDotCommentVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"comment\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "context",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotContextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotContextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Context.SetTo(requestDotContextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"context\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReportChatRoomRequest(r *http.Request) (
	req *ReportChatRoomReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReportChatRoomReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "opponent_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOpponentIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotOpponentIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OpponentID.SetTo(requestDotOpponentIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"opponent_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "reason",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotReasonVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotReasonVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Reason.SetTo(requestDotReasonVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"reason\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot2Filename.SetTo(requestDotScreenshot2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot3Filename.SetTo(requestDotScreenshot3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot4Filename.SetTo(requestDotScreenshot4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ScreenshotFilename.SetTo(requestDotScreenshotFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_filename\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReportGroupRequest(r *http.Request) (
	req *ReportGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReportGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "opponent_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOpponentIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotOpponentIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OpponentID.SetTo(requestDotOpponentIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"opponent_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "reason",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotReasonVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotReasonVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Reason.SetTo(requestDotReasonVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"reason\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot2Filename.SetTo(requestDotScreenshot2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot3Filename.SetTo(requestDotScreenshot3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot4Filename.SetTo(requestDotScreenshot4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ScreenshotFilename.SetTo(requestDotScreenshotFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_filename\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReportPostRequest(r *http.Request) (
	req *ReportPostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReportPostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "opponent_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOpponentIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotOpponentIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OpponentID.SetTo(requestDotOpponentIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"opponent_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "reason",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotReasonVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotReasonVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Reason.SetTo(requestDotReasonVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"reason\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot2Filename.SetTo(requestDotScreenshot2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot3Filename.SetTo(requestDotScreenshot3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot4Filename.SetTo(requestDotScreenshot4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ScreenshotFilename.SetTo(requestDotScreenshotFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_filename\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReportUserRequest(r *http.Request) (
	req *ReportUserReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ReportUserReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "reason",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotReasonVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotReasonVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Reason.SetTo(requestDotReasonVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"reason\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot2Filename.SetTo(requestDotScreenshot2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot3Filename.SetTo(requestDotScreenshot3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshot4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshot4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Screenshot4Filename.SetTo(requestDotScreenshot4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "screenshot_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotScreenshotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotScreenshotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ScreenshotFilename.SetTo(requestDotScreenshotFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"screenshot_filename\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeRepostRequest(r *http.Request) (
	req *RepostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request RepostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_2_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment2FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment2FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment2Filename.SetTo(requestDotAttachment2FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_2_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_3_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment3FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment3FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment3Filename.SetTo(requestDotAttachment3FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_3_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_4_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment4FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment4FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment4Filename.SetTo(requestDotAttachment4FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_4_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_5_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment5FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment5FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment5Filename.SetTo(requestDotAttachment5FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_5_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_6_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment6FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment6FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment6Filename.SetTo(requestDotAttachment6FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_6_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_7_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment7FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment7FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment7Filename.SetTo(requestDotAttachment7FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_7_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_8_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment8FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment8FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment8Filename.SetTo(requestDotAttachment8FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_8_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_9_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachment9FilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachment9FilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Attachment9Filename.SetTo(requestDotAttachment9FilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_9_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachmentFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachmentFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AttachmentFilename.SetTo(requestDotAttachmentFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "choices[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChoicesVal []string
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotChoicesValVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								requestDotChoicesValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotChoicesVal = append(requestDotChoicesVal, requestDotChoicesValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.Choices.SetTo(requestDotChoicesVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"choices[]\"")
				}
				if err := func() error {
					if value, ok := request.Choices.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "in_reply_to",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotInReplyToVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotInReplyToVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.InReplyTo.SetTo(requestDotInReplyToVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"in_reply_to\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "mention_ids[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMentionIdsVal []int64
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var requestDotMentionIdsValVal int64
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToInt64(val)
								if err != nil {
									return err
								}

								requestDotMentionIdsValVal = c
								return nil
							}(); err != nil {
								return err
							}
							requestDotMentionIdsVal = append(requestDotMentionIdsVal, requestDotMentionIdsValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					request.MentionIds.SetTo(requestDotMentionIdsVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"mention_ids[]\"")
				}
				if err := func() error {
					if value, ok := request.MentionIds.Get(); ok {
						if err := func() error {
							if value == nil {
								return errors.New("nil is invalid value")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPostIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotPostIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.PostID.SetTo(requestDotPostIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPostTypeVal PostType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPostTypeVal = PostType(c)
						return nil
					}(); err != nil {
						return err
					}
					request.PostType.SetTo(requestDotPostTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"post_type\"")
				}
				if err := func() error {
					if value, ok := request.PostType.Get(); ok {
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
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "shared_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.SharedURL.Reset()
						if err := request.SharedURL.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"shared_url\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "video_file_name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotVideoFileNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotVideoFileNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.VideoFileName.SetTo(requestDotVideoFileNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"video_file_name\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeRequestEmailVerificationRequest(r *http.Request) (
	req *RequestEmailVerificationReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request RequestEmailVerificationReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "device_uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDeviceUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotDeviceUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.DeviceUUID.SetTo(requestDotDeviceUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"device_uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Email.SetTo(requestDotEmailVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "intent",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIntentVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotIntentVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Intent.SetTo(requestDotIntentVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"intent\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "locale",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLocaleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLocaleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Locale.SetTo(requestDotLocaleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"locale\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeRequestFollowRequest(r *http.Request) (
	req *RequestFollowReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request RequestFollowReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "action",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotActionVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotActionVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Action.SetTo(requestDotActionVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"action\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeResetCountersRequest(r *http.Request) (
	req *ResetCountersReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ResetCountersReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "counter",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCounterVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCounterVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Counter.SetTo(requestDotCounterVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"counter\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeResetPasswordRequest(r *http.Request) (
	req *ResetPasswordReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ResetPasswordReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Email.SetTo(requestDotEmailVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email_grant_token",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailGrantTokenVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailGrantTokenVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.EmailGrantToken.SetTo(requestDotEmailGrantTokenVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email_grant_token\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Password.SetTo(requestDotPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"password\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeSendChatMessageRequest(r *http.Request) (
	req *SendChatMessageReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request SendChatMessageReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "attachment_file_name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAttachmentFileNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAttachmentFileNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AttachmentFileName.SetTo(requestDotAttachmentFileNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"attachment_file_name\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallTypeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCallTypeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallType.SetTo(requestDotCallTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_type\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "gif_image_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGIFImageIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGIFImageIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GIFImageID.SetTo(requestDotGIFImageIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"gif_image_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMessageTypeVal MessageType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotMessageTypeVal = MessageType(c)
						return nil
					}(); err != nil {
						return err
					}
					request.MessageType.SetTo(requestDotMessageTypeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_type\"")
				}
				if err := func() error {
					if value, ok := request.MessageType.Get(); ok {
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
					return req, rawBody, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "parent_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotParentIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotParentIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ParentID.SetTo(requestDotParentIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"parent_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sticker_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotStickerIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotStickerIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.StickerID.SetTo(requestDotStickerIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"sticker_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sticker_pack_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotStickerPackIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotStickerPackIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.StickerPackID.SetTo(requestDotStickerPackIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"sticker_pack_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "video_file_name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotVideoFileNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotVideoFileNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.VideoFileName.SetTo(requestDotVideoFileNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"video_file_name\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeSetGroupTitleRequest(r *http.Request) (
	req *SetGroupTitleReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request SetGroupTitleReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "title",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTitleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTitleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Title.SetTo(requestDotTitleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"title\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeStartConferenceCallRequest(r *http.Request) (
	req *StartConferenceCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request StartConferenceCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_sid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallSidVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCallSidVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallSid.SetTo(requestDotCallSidVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_sid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "conference_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotConferenceIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotConferenceIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ConferenceID.SetTo(requestDotConferenceIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"conference_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTransferGroupRequest(r *http.Request) (
	req *TransferGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request TransferGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotUserIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserID.SetTo(requestDotUserIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateAdditionalNotificationSettingRequest(r *http.Request) (
	req *UpdateAdditionalNotificationSettingReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateAdditionalNotificationSettingReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "mode",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotModeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotModeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Mode.SetTo(requestDotModeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"mode\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "on",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOnVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotOnVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.On.SetTo(requestDotOnVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"on\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateCallRequest(r *http.Request) (
	req *UpdateCallReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateCallReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCategoryIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CategoryID.SetTo(requestDotCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "game_title",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGameTitleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGameTitleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GameTitle.SetTo(requestDotGameTitleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"game_title\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "joinable_by",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotJoinableByVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotJoinableByVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.JoinableBy.SetTo(requestDotJoinableByVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"joinable_by\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateCallUserRequest(r *http.Request) (
	req *UpdateCallUserReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateCallUserReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "role",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotRoleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotRoleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Role.SetTo(requestDotRoleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"role\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateChatRoomRequest(r *http.Request) (
	req *UpdateChatRoomReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateChatRoomReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "background_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotBackgroundFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotBackgroundFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.BackgroundFilename.SetTo(requestDotBackgroundFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"background_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IconFilename.SetTo(requestDotIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Name.SetTo(requestDotNameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"name\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateChatRoomNotificationSettingsRequest(r *http.Request) (
	req *UpdateChatRoomNotificationSettingsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateChatRoomNotificationSettingsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "notification_chat",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNotificationChatVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotNotificationChatVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.NotificationChat.SetTo(requestDotNotificationChatVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"notification_chat\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateGroupRequest(r *http.Request) (
	req *UpdateGroupReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateGroupReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_members_to_post_image_and_video",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowMembersToPostImageAndVideoVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowMembersToPostImageAndVideoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowMembersToPostImageAndVideo.SetTo(requestDotAllowMembersToPostImageAndVideoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_members_to_post_image_and_video\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_members_to_post_url",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowMembersToPostURLVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowMembersToPostURLVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowMembersToPostURL.SetTo(requestDotAllowMembersToPostURLVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_members_to_post_url\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_ownership_transfer",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowOwnershipTransferVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotAllowOwnershipTransferVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowOwnershipTransfer.SetTo(requestDotAllowOwnershipTransferVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_ownership_transfer\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "allow_thread_creation_by",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAllowThreadCreationByVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAllowThreadCreationByVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.AllowThreadCreationBy.SetTo(requestDotAllowThreadCreationByVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"allow_thread_creation_by\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "call_timeline_display",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCallTimelineDisplayVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotCallTimelineDisplayVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CallTimelineDisplay.SetTo(requestDotCallTimelineDisplayVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"call_timeline_display\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "cover_image_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCoverImageFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCoverImageFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CoverImageFilename.SetTo(requestDotCoverImageFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"cover_image_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDescriptionVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotDescriptionVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Description.SetTo(requestDotDescriptionVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"description\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "gender",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGenderVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotGenderVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Gender.SetTo(requestDotGenderVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"gender\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "generation_groups_limit",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGenerationGroupsLimitVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotGenerationGroupsLimitVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GenerationGroupsLimit.SetTo(requestDotGenerationGroupsLimitVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"generation_groups_limit\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupCategoryIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupCategoryID.SetTo(requestDotGroupCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGroupIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupIconFilename.SetTo(requestDotGroupIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "guidelines",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGuidelinesVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotGuidelinesVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Guidelines.SetTo(requestDotGuidelinesVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"guidelines\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_conference_call",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideConferenceCallVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideConferenceCallVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideConferenceCall.SetTo(requestDotHideConferenceCallVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_conference_call\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_from_game_eight",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideFromGameEightVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideFromGameEightVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideFromGameEight.SetTo(requestDotHideFromGameEightVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_from_game_eight\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "hide_reported_posts",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotHideReportedPostsVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotHideReportedPostsVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.HideReportedPosts.SetTo(requestDotHideReportedPostsVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"hide_reported_posts\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "is_private",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIsPrivateVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotIsPrivateVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IsPrivate.SetTo(requestDotIsPrivateVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"is_private\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "only_mobile_verified",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOnlyMobileVerifiedVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotOnlyMobileVerifiedVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OnlyMobileVerified.SetTo(requestDotOnlyMobileVerifiedVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"only_mobile_verified\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "only_verified_age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOnlyVerifiedAgeVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotOnlyVerifiedAgeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OnlyVerifiedAge.SetTo(requestDotOnlyVerifiedAgeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"only_verified_age\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "secret",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSecretVal bool
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToBool(val)
						if err != nil {
							return err
						}

						requestDotSecretVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Secret.SetTo(requestDotSecretVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"secret\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sub_category_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSubCategoryIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSubCategoryIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SubCategoryID.SetTo(requestDotSubCategoryIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"sub_category_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "topic",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTopicVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTopicVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Topic.SetTo(requestDotTopicVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"topic\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateGroupNotificationSettingsRequest(r *http.Request) (
	req *UpdateGroupNotificationSettingsReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateGroupNotificationSettingsReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "notification_group_join",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNotificationGroupJoinVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotNotificationGroupJoinVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.NotificationGroupJoin.SetTo(requestDotNotificationGroupJoinVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"notification_group_join\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "notification_group_message_tag_all",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNotificationGroupMessageTagAllVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotNotificationGroupMessageTagAllVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.NotificationGroupMessageTagAll.SetTo(requestDotNotificationGroupMessageTagAllVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"notification_group_message_tag_all\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "notification_group_post",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNotificationGroupPostVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotNotificationGroupPostVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.NotificationGroupPost.SetTo(requestDotNotificationGroupPostVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"notification_group_post\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "notification_group_request",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotNotificationGroupRequestVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotNotificationGroupRequestVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.NotificationGroupRequest.SetTo(requestDotNotificationGroupRequestVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"notification_group_request\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateLanguageRequest(r *http.Request) (
	req *UpdateLanguageReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateLanguageReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateLoginRequest(r *http.Request) (
	req *UpdateLoginReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateLoginReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "current_password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCurrentPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotCurrentPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.CurrentPassword.SetTo(requestDotCurrentPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"current_password\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Email.SetTo(requestDotEmailVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email_grant_token",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotEmailGrantTokenVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotEmailGrantTokenVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.EmailGrantToken.SetTo(requestDotEmailGrantTokenVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"email_grant_token\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "password",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotPasswordVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotPasswordVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Password.SetTo(requestDotPasswordVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"password\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdatePostRequest(r *http.Request) (
	req *UpdatePostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdatePostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "api_key",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotAPIKeyVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotAPIKeyVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.APIKey.SetTo(requestDotAPIKeyVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"api_key\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "color",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotColorVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotColorVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Color.SetTo(requestDotColorVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"color\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "font_size",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFontSizeVal int32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt32(val)
						if err != nil {
							return err
						}

						requestDotFontSizeVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FontSize.SetTo(requestDotFontSizeVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"font_size\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "language",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotLanguageVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotLanguageVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Language.SetTo(requestDotLanguageVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"language\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "message_tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}
					if err := func(d *jx.Decoder) error {
						request.MessageTags.Reset()
						if err := request.MessageTags.Decode(d); err != nil {
							return err
						}
						return nil
					}(jx.DecodeStr(val)); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"message_tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "signed_info",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSignedInfoVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotSignedInfoVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.SignedInfo.SetTo(requestDotSignedInfoVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"signed_info\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "timestamp",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTimestampVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotTimestampVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Timestamp.SetTo(requestDotTimestampVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"timestamp\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateThreadRequest(r *http.Request) (
	req *UpdateThreadReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request UpdateThreadReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "thread_icon_filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotThreadIconFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotThreadIconFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ThreadIconFilename.SetTo(requestDotThreadIconFilenameVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"thread_icon_filename\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "title",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTitleVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTitleVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Title.SetTo(requestDotTitleVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"title\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateUserInterestsRequest(r *http.Request) (
	req *CommonIdsRequest,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
		if err != nil {
			return req, rawBody, close, err
		}

		// Reset the body to allow for downstream reading.
		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		if len(buf) == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}

		rawBody = append(rawBody, buf...)
		d := jx.DecodeBytes(buf)

		var request CommonIdsRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, rawBody, close, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, rawBody, close, errors.Wrap(err, "validate")
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeValidatePostRequest(r *http.Request) (
	req *ValidatePostReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request ValidatePostReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "group_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotGroupIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotGroupIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.GroupID.SetTo(requestDotGroupIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"group_id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "text",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotTextVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotTextVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Text.SetTo(requestDotTextVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"text\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "thread_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotThreadIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotThreadIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ThreadID.SetTo(requestDotThreadIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"thread_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeVoteSurveyRequest(r *http.Request) (
	req *VoteSurveyReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request VoteSurveyReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "choice_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotChoiceIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotChoiceIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ChoiceID.SetTo(requestDotChoiceIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"choice_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeWithdrawGroupTransferRequest(r *http.Request) (
	req *WithdrawGroupTransferReq,
	rawBody []byte,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = errors.Join(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = errors.Join(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, rawBody, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, rawBody, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, rawBody, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request WithdrawGroupTransferReq
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "user_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIDVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						requestDotUserIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserID.SetTo(requestDotUserIDVal)
					return nil
				}); err != nil {
					return req, rawBody, close, errors.Wrap(err, "decode \"user_id\"")
				}
			}
		}
		return &request, rawBody, close, nil
	default:
		return req, rawBody, close, validate.InvalidContentType(ct)
	}
}
