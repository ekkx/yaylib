// Code generated; DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"go.opentelemetry.io/otel/trace"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func (c *codeRecorder) Unwrap() http.ResponseWriter {
	return c.ResponseWriter
}

// handleAcceptChatRequestRequest handles acceptChatRequest operation.
//
// POST /v1/chat_rooms/accept_chat_request
func (s *Server) handleAcceptChatRequestRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptChatRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/accept_chat_request"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), AcceptChatRequestOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: AcceptChatRequestOperation,
			ID:   "acceptChatRequest",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeAcceptChatRequestRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *AcceptChatRequestNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    AcceptChatRequestOperation,
			OperationSummary: "",
			OperationID:      "acceptChatRequest",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *AcceptChatRequestReq
			Params   = struct{}
			Response = *AcceptChatRequestNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.AcceptChatRequest(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.AcceptChatRequest(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAcceptChatRequestResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAcceptGroupJoinRequestRequest handles acceptGroupJoinRequest operation.
//
// POST /v1/groups/{id}/accept/{userId}
func (s *Server) handleAcceptGroupJoinRequestRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptGroupJoinRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/accept/{userId}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), AcceptGroupJoinRequestOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: AcceptGroupJoinRequestOperation,
			ID:   "acceptGroupJoinRequest",
		}
	)
	params, err := decodeAcceptGroupJoinRequestParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *AcceptGroupJoinRequestNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    AcceptGroupJoinRequestOperation,
			OperationSummary: "",
			OperationID:      "acceptGroupJoinRequest",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = AcceptGroupJoinRequestParams
			Response = *AcceptGroupJoinRequestNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAcceptGroupJoinRequestParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.AcceptGroupJoinRequest(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.AcceptGroupJoinRequest(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAcceptGroupJoinRequestResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAcceptGroupTransferRequest handles acceptGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer
func (s *Server) handleAcceptGroupTransferRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/transfer"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), AcceptGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: AcceptGroupTransferOperation,
			ID:   "acceptGroupTransfer",
		}
	)
	params, err := decodeAcceptGroupTransferParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *AcceptGroupTransferNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    AcceptGroupTransferOperation,
			OperationSummary: "",
			OperationID:      "acceptGroupTransfer",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = AcceptGroupTransferParams
			Response = *AcceptGroupTransferNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAcceptGroupTransferParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.AcceptGroupTransfer(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.AcceptGroupTransfer(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAcceptGroupTransferResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAddThreadMemberRequest handles addThreadMember operation.
//
// POST /v1/threads/{thread_id}/members/{id}
func (s *Server) handleAddThreadMemberRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("addThreadMember"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/threads/{thread_id}/members/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), AddThreadMemberOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: AddThreadMemberOperation,
			ID:   "addThreadMember",
		}
	)
	params, err := decodeAddThreadMemberParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *AddThreadMemberNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    AddThreadMemberOperation,
			OperationSummary: "",
			OperationID:      "addThreadMember",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "thread_id",
					In:   "path",
				}: params.ThreadID,
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = AddThreadMemberParams
			Response = *AddThreadMemberNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAddThreadMemberParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.AddThreadMember(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.AddThreadMember(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAddThreadMemberResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleAgreePolicyRequest handles agreePolicy operation.
//
// POST /v1/users/policy_agreements/{type}
func (s *Server) handleAgreePolicyRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("agreePolicy"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/policy_agreements/{type}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), AgreePolicyOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: AgreePolicyOperation,
			ID:   "agreePolicy",
		}
	)
	params, err := decodeAgreePolicyParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *AgreePolicyNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    AgreePolicyOperation,
			OperationSummary: "",
			OperationID:      "agreePolicy",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "type",
					In:   "path",
				}: params.Type,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = AgreePolicyParams
			Response = *AgreePolicyNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAgreePolicyParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.AgreePolicy(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.AgreePolicy(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAgreePolicyResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleBanGroupUserRequest handles banGroupUser operation.
//
// POST /v1/groups/{id}/ban/{userId}
func (s *Server) handleBanGroupUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("banGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/ban/{userId}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), BanGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: BanGroupUserOperation,
			ID:   "banGroupUser",
		}
	)
	params, err := decodeBanGroupUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BanGroupUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    BanGroupUserOperation,
			OperationSummary: "",
			OperationID:      "banGroupUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = BanGroupUserParams
			Response = *BanGroupUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackBanGroupUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.BanGroupUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.BanGroupUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeBanGroupUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleBlockUserRequest handles blockUser operation.
//
// POST /v1/users/{id}/block
func (s *Server) handleBlockUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("blockUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/{id}/block"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), BlockUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: BlockUserOperation,
			ID:   "blockUser",
		}
	)
	params, err := decodeBlockUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BlockUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    BlockUserOperation,
			OperationSummary: "",
			OperationID:      "blockUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = BlockUserParams
			Response = *BlockUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackBlockUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.BlockUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.BlockUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeBlockUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleBulkInviteToCallRequest handles bulkInviteToCall operation.
//
// POST /v1/calls/{call_id}/bulk_invite
func (s *Server) handleBulkInviteToCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("bulkInviteToCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}/bulk_invite"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), BulkInviteToCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: BulkInviteToCallOperation,
			ID:   "bulkInviteToCall",
		}
	)
	params, err := decodeBulkInviteToCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BulkInviteToCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    BulkInviteToCallOperation,
			OperationSummary: "",
			OperationID:      "bulkInviteToCall",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = BulkInviteToCallParams
			Response = *BulkInviteToCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackBulkInviteToCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.BulkInviteToCall(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.BulkInviteToCall(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeBulkInviteToCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleBumpCallRequest handles bumpCall operation.
//
// POST /v1/calls/{call_id}/bump
func (s *Server) handleBumpCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("bumpCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}/bump"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), BumpCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: BumpCallOperation,
			ID:   "bumpCall",
		}
	)
	params, err := decodeBumpCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BumpCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    BumpCallOperation,
			OperationSummary: "",
			OperationID:      "bumpCall",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
				{
					Name: "participant_limit",
					In:   "query",
				}: params.ParticipantLimit,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = BumpCallParams
			Response = *BumpCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackBumpCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.BumpCall(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.BumpCall(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeBumpCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCancelGroupTransferRequest handles cancelGroupTransfer operation.
//
// DELETE /v1/groups/{id}/transfer
func (s *Server) handleCancelGroupTransferRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("cancelGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/transfer"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CancelGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CancelGroupTransferOperation,
			ID:   "cancelGroupTransfer",
		}
	)
	params, err := decodeCancelGroupTransferParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *CancelGroupTransferNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CancelGroupTransferOperation,
			OperationSummary: "",
			OperationID:      "cancelGroupTransfer",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = CancelGroupTransferParams
			Response = *CancelGroupTransferNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCancelGroupTransferParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.CancelGroupTransfer(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.CancelGroupTransfer(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCancelGroupTransferResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleChangeEmailRequest handles changeEmail operation.
//
// PUT /v1/users/change_email
func (s *Server) handleChangeEmailRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("changeEmail"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/change_email"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ChangeEmailOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ChangeEmailOperation,
			ID:   "changeEmail",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeChangeEmailRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LoginUpdateResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ChangeEmailOperation,
			OperationSummary: "",
			OperationID:      "changeEmail",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ChangeEmailReq
			Params   = struct{}
			Response = *LoginUpdateResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ChangeEmail(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.ChangeEmail(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeChangeEmailResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleChangePasswordRequest handles changePassword operation.
//
// PUT /v1/users/change_password
func (s *Server) handleChangePasswordRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("changePassword"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/change_password"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ChangePasswordOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ChangePasswordOperation,
			ID:   "changePassword",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeChangePasswordRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LoginUpdateResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ChangePasswordOperation,
			OperationSummary: "",
			OperationID:      "changePassword",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ChangePasswordReq
			Params   = struct{}
			Response = *LoginUpdateResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ChangePassword(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.ChangePassword(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeChangePasswordResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateBookmarkRequest handles createBookmark operation.
//
// PUT /v1/users/{user_id}/bookmarks/{id}
func (s *Server) handleCreateBookmarkRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createBookmark"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/{user_id}/bookmarks/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateBookmarkOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateBookmarkOperation,
			ID:   "createBookmark",
		}
	)
	params, err := decodeCreateBookmarkParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BookmarkPostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateBookmarkOperation,
			OperationSummary: "",
			OperationID:      "createBookmark",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = CreateBookmarkParams
			Response = *BookmarkPostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCreateBookmarkParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateBookmark(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateBookmark(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateBookmarkResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateChatRoomRequest handles createChatRoom operation.
//
// POST /v3/chat_rooms/new
func (s *Server) handleCreateChatRoomRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/chat_rooms/new"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateChatRoomOperation,
			ID:   "createChatRoom",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateChatRoomRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreateChatRoomResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateChatRoomOperation,
			OperationSummary: "",
			OperationID:      "createChatRoom",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateChatRoomReq
			Params   = struct{}
			Response = *CreateChatRoomResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateChatRoom(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateChatRoom(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateChatRoomV1Request handles createChatRoomV1 operation.
//
// POST /v1/chat_rooms/new
func (s *Server) handleCreateChatRoomV1Request(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createChatRoomV1"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/new"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateChatRoomV1Operation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateChatRoomV1Operation,
			ID:   "createChatRoomV1",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateChatRoomV1Request(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreateChatRoomResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateChatRoomV1Operation,
			OperationSummary: "",
			OperationID:      "createChatRoomV1",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateChatRoomV1Req
			Params   = struct{}
			Response = *CreateChatRoomResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateChatRoomV1(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateChatRoomV1(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateChatRoomV1Response(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateConferenceCallPostRequest handles createConferenceCallPost operation.
//
// POST /v2/posts/new_conference_call
func (s *Server) handleCreateConferenceCallPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createConferenceCallPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/posts/new_conference_call"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateConferenceCallPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateConferenceCallPostOperation,
			ID:   "createConferenceCallPost",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateConferenceCallPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreatePostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateConferenceCallPostOperation,
			OperationSummary: "",
			OperationID:      "createConferenceCallPost",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateConferenceCallPostReq
			Params   = struct{}
			Response = *CreatePostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateConferenceCallPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateConferenceCallPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateConferenceCallPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateGroupRequest handles createGroup operation.
//
// POST /v3/groups/new
func (s *Server) handleCreateGroupRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/groups/new"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateGroupOperation,
			ID:   "createGroup",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreateGroupResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateGroupOperation,
			OperationSummary: "",
			OperationID:      "createGroup",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateGroupReq
			Params   = struct{}
			Response = *CreateGroupResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateGroup(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateGroup(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateMuteKeywordRequest handles createMuteKeyword operation.
//
// POST /v1/hidden/words
func (s *Server) handleCreateMuteKeywordRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createMuteKeyword"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/hidden/words"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateMuteKeywordOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateMuteKeywordOperation,
			ID:   "createMuteKeyword",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateMuteKeywordRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreateMuteKeywordResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateMuteKeywordOperation,
			OperationSummary: "",
			OperationID:      "createMuteKeyword",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *MuteKeywordRequest
			Params   = struct{}
			Response = *CreateMuteKeywordResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateMuteKeyword(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateMuteKeyword(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateMuteKeywordResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreatePostRequest handles createPost operation.
//
// POST /v3/posts/new
func (s *Server) handleCreatePostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/posts/new"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreatePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreatePostOperation,
			ID:   "createPost",
		}
	)
	params, err := decodeCreatePostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreatePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Post
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreatePostOperation,
			OperationSummary: "",
			OperationID:      "createPost",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "X-Jwt",
					In:   "header",
				}: params.XJwt,
			},
			Raw: r,
		}

		type (
			Request  = *CreatePostReq
			Params   = CreatePostParams
			Response = *Post
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCreatePostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreatePost(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreatePost(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreatePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateReviewRequest handles createReview operation.
//
// POST /v1/users/reviews
func (s *Server) handleCreateReviewRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/reviews"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateReviewOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateReviewOperation,
			ID:   "createReview",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateReviewRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreateReviewNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateReviewOperation,
			OperationSummary: "",
			OperationID:      "createReview",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateReviewReq
			Params   = struct{}
			Response = *CreateReviewNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.CreateReview(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.CreateReview(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateReviewResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateSharePostRequest handles createSharePost operation.
//
// POST /v2/posts/new_share_post
func (s *Server) handleCreateSharePostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSharePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/posts/new_share_post"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateSharePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateSharePostOperation,
			ID:   "createSharePost",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateSharePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Post
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateSharePostOperation,
			OperationSummary: "",
			OperationID:      "createSharePost",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateSharePostReq
			Params   = struct{}
			Response = *Post
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateSharePost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateSharePost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateSharePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateThreadRequest handles createThread operation.
//
// POST /v1/threads
func (s *Server) handleCreateThreadRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createThread"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/threads"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateThreadOperation,
			ID:   "createThread",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateThreadRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ThreadInfo
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateThreadOperation,
			OperationSummary: "",
			OperationID:      "createThread",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CreateGroupThreadRequest
			Params   = struct{}
			Response = *ThreadInfo
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateThread(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateThread(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleCreateThreadPostRequest handles createThreadPost operation.
//
// POST /v1/threads/{id}/posts
func (s *Server) handleCreateThreadPostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createThreadPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/threads/{id}/posts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), CreateThreadPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: CreateThreadPostOperation,
			ID:   "createThreadPost",
		}
	)
	params, err := decodeCreateThreadPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeCreateThreadPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Post
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CreateThreadPostOperation,
			OperationSummary: "",
			OperationID:      "createThreadPost",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "X-Jwt",
					In:   "header",
				}: params.XJwt,
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *CreateThreadPostReq
			Params   = CreateThreadPostParams
			Response = *Post
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCreateThreadPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateThreadPost(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateThreadPost(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateThreadPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeclineGroupJoinRequestRequest handles declineGroupJoinRequest operation.
//
// POST /v1/groups/{id}/decline/{userId}
func (s *Server) handleDeclineGroupJoinRequestRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("declineGroupJoinRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/decline/{userId}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeclineGroupJoinRequestOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeclineGroupJoinRequestOperation,
			ID:   "declineGroupJoinRequest",
		}
	)
	params, err := decodeDeclineGroupJoinRequestParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeclineGroupJoinRequestNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeclineGroupJoinRequestOperation,
			OperationSummary: "",
			OperationID:      "declineGroupJoinRequest",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeclineGroupJoinRequestParams
			Response = *DeclineGroupJoinRequestNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeclineGroupJoinRequestParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeclineGroupJoinRequest(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeclineGroupJoinRequest(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeclineGroupJoinRequestResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteAllPostsRequest handles deleteAllPosts operation.
//
// POST /v1/posts/delete_all_post
func (s *Server) handleDeleteAllPostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteAllPosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/posts/delete_all_post"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteAllPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *DeleteAllPostsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteAllPostsOperation,
			OperationSummary: "",
			OperationID:      "deleteAllPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *DeleteAllPostsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteAllPosts(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteAllPosts(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteAllPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteBookmarkRequest handles deleteBookmark operation.
//
// DELETE /v1/users/{user_id}/bookmarks/{id}
func (s *Server) handleDeleteBookmarkRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteBookmark"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/users/{user_id}/bookmarks/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteBookmarkOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteBookmarkOperation,
			ID:   "deleteBookmark",
		}
	)
	params, err := decodeDeleteBookmarkParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteBookmarkNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteBookmarkOperation,
			OperationSummary: "",
			OperationID:      "deleteBookmark",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteBookmarkParams
			Response = *DeleteBookmarkNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteBookmarkParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteBookmark(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteBookmark(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteBookmarkResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteChatMessageRequest handles deleteChatMessage operation.
//
// DELETE /v1/chat_rooms/{room_id}/messages/{message_id}/delete
func (s *Server) handleDeleteChatMessageRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteChatMessage"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/{room_id}/messages/{message_id}/delete"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteChatMessageOperation,
			ID:   "deleteChatMessage",
		}
	)
	params, err := decodeDeleteChatMessageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteChatMessageNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteChatMessageOperation,
			OperationSummary: "",
			OperationID:      "deleteChatMessage",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "room_id",
					In:   "path",
				}: params.RoomID,
				{
					Name: "message_id",
					In:   "path",
				}: params.MessageID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteChatMessageParams
			Response = *DeleteChatMessageNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteChatMessageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteChatMessage(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteChatMessage(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteChatMessageResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteChatRoomsRequest handles deleteChatRooms operation.
//
// POST /v1/chat_rooms/mass_destroy
func (s *Server) handleDeleteChatRoomsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteChatRooms"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/mass_destroy"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteChatRoomsOperation,
			ID:   "deleteChatRooms",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeDeleteChatRoomsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *DeleteChatRoomsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteChatRoomsOperation,
			OperationSummary: "",
			OperationID:      "deleteChatRooms",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *DeleteChatRoomsReq
			Params   = struct{}
			Response = *DeleteChatRoomsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteChatRooms(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteChatRooms(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteChatRoomsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteFootprintRequest handles deleteFootprint operation.
//
// DELETE /v2/users/{user_id}/footprints/{footprint_id}
func (s *Server) handleDeleteFootprintRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteFootprint"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v2/users/{user_id}/footprints/{footprint_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteFootprintOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteFootprintOperation,
			ID:   "deleteFootprint",
		}
	)
	params, err := decodeDeleteFootprintParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteFootprintNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteFootprintOperation,
			OperationSummary: "",
			OperationID:      "deleteFootprint",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
				{
					Name: "footprint_id",
					In:   "path",
				}: params.FootprintID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteFootprintParams
			Response = *DeleteFootprintNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteFootprintParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteFootprint(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteFootprint(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteFootprintResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteMuteKeywordRequest handles deleteMuteKeyword operation.
//
// DELETE /v1/hidden/words
func (s *Server) handleDeleteMuteKeywordRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteMuteKeyword"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/hidden/words"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteMuteKeywordOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteMuteKeywordOperation,
			ID:   "deleteMuteKeyword",
		}
	)
	params, err := decodeDeleteMuteKeywordParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteMuteKeywordNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteMuteKeywordOperation,
			OperationSummary: "",
			OperationID:      "deleteMuteKeyword",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "ids[]",
					In:   "query",
				}: params.Ids,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteMuteKeywordParams
			Response = *DeleteMuteKeywordNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteMuteKeywordParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteMuteKeyword(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteMuteKeyword(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteMuteKeywordResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteMyReviewsRequest handles deleteMyReviews operation.
//
// DELETE /v1/users/reviews
func (s *Server) handleDeleteMyReviewsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteMyReviews"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/users/reviews"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteMyReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteMyReviewsOperation,
			ID:   "deleteMyReviews",
		}
	)
	params, err := decodeDeleteMyReviewsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteMyReviewsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteMyReviewsOperation,
			OperationSummary: "",
			OperationID:      "deleteMyReviews",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "review_ids[]",
					In:   "query",
				}: params.ReviewIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteMyReviewsParams
			Response = *DeleteMyReviewsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteMyReviewsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteMyReviews(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteMyReviews(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteMyReviewsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeletePostsRequest handles deletePosts operation.
//
// POST /v2/posts/mass_destroy
func (s *Server) handleDeletePostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deletePosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/posts/mass_destroy"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeletePostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeletePostsOperation,
			ID:   "deletePosts",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeDeletePostsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *DeletePostsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeletePostsOperation,
			OperationSummary: "",
			OperationID:      "deletePosts",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *DeletePostsReq
			Params   = struct{}
			Response = *DeletePostsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeletePosts(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.DeletePosts(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeletePostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteThreadRequest handles deleteThread operation.
//
// DELETE /v1/threads/{id}
func (s *Server) handleDeleteThreadRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteThread"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/threads/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeleteThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeleteThreadOperation,
			ID:   "deleteThread",
		}
	)
	params, err := decodeDeleteThreadParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeleteThreadNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeleteThreadOperation,
			OperationSummary: "",
			OperationID:      "deleteThread",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteThreadParams
			Response = *DeleteThreadNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteThreadParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteThread(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteThread(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeputizeGroupUsersRequest handles deputizeGroupUsers operation.
//
// PUT /v1/groups/{id}/deputize
func (s *Server) handleDeputizeGroupUsersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deputizeGroupUsers"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/deputize"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeputizeGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeputizeGroupUsersOperation,
			ID:   "deputizeGroupUsers",
		}
	)
	params, err := decodeDeputizeGroupUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *DeputizeGroupUsersNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeputizeGroupUsersOperation,
			OperationSummary: "",
			OperationID:      "deputizeGroupUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeputizeGroupUsersParams
			Response = *DeputizeGroupUsersNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeputizeGroupUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeputizeGroupUsers(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeputizeGroupUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeputizeGroupUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeputizeGroupUsersMassRequest handles deputizeGroupUsersMass operation.
//
// POST /v3/groups/{group_id}/deputize/mass
func (s *Server) handleDeputizeGroupUsersMassRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deputizeGroupUsersMass"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/groups/{group_id}/deputize/mass"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DeputizeGroupUsersMassOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DeputizeGroupUsersMassOperation,
			ID:   "deputizeGroupUsersMass",
		}
	)
	params, err := decodeDeputizeGroupUsersMassParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeDeputizeGroupUsersMassRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *DeputizeGroupUsersMassNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DeputizeGroupUsersMassOperation,
			OperationSummary: "",
			OperationID:      "deputizeGroupUsersMass",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
			},
			Raw: r,
		}

		type (
			Request  = *DeputizeGroupUsersMassReq
			Params   = DeputizeGroupUsersMassParams
			Response = *DeputizeGroupUsersMassNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeputizeGroupUsersMassParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeputizeGroupUsersMass(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeputizeGroupUsersMass(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeputizeGroupUsersMassResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDisableTwoFactorAuthRequest handles disableTwoFactorAuth operation.
//
// PUT /v1/users/secret/disable
func (s *Server) handleDisableTwoFactorAuthRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("disableTwoFactorAuth"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/secret/disable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), DisableTwoFactorAuthOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: DisableTwoFactorAuthOperation,
			ID:   "disableTwoFactorAuth",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeDisableTwoFactorAuthRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *DisableTwoFactorAuthNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    DisableTwoFactorAuthOperation,
			OperationSummary: "",
			OperationID:      "disableTwoFactorAuth",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *DisableTwoFactorAuthReq
			Params   = struct{}
			Response = *DisableTwoFactorAuthNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DisableTwoFactorAuth(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.DisableTwoFactorAuth(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDisableTwoFactorAuthResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleEditUserRequest handles editUser operation.
//
// POST /v3/users/edit
func (s *Server) handleEditUserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("editUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/users/edit"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), EditUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: EditUserOperation,
			ID:   "editUser",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeEditUserRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *EditUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    EditUserOperation,
			OperationSummary: "",
			OperationID:      "editUser",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *EditUserReq
			Params   = struct{}
			Response = *EditUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.EditUser(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.EditUser(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeEditUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleEditUserV2Request handles editUserV2 operation.
//
// POST /v2/users/edit
func (s *Server) handleEditUserV2Request(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("editUserV2"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/edit"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), EditUserV2Operation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: EditUserV2Operation,
			ID:   "editUserV2",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeEditUserV2Request(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *EditUserV2NoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    EditUserV2Operation,
			OperationSummary: "",
			OperationID:      "editUserV2",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *EditUserV2Req
			Params   = struct{}
			Response = *EditUserV2NoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.EditUserV2(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.EditUserV2(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeEditUserV2Response(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleEnableTwoFactorAuthRequest handles enableTwoFactorAuth operation.
//
// PUT /v1/users/secret/enable
func (s *Server) handleEnableTwoFactorAuthRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("enableTwoFactorAuth"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/secret/enable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), EnableTwoFactorAuthOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: EnableTwoFactorAuthOperation,
			ID:   "enableTwoFactorAuth",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeEnableTwoFactorAuthRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TwoStepAuthEnabledResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    EnableTwoFactorAuthOperation,
			OperationSummary: "",
			OperationID:      "enableTwoFactorAuth",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *EnableTwoFactorAuthReq
			Params   = struct{}
			Response = *TwoStepAuthEnabledResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.EnableTwoFactorAuth(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.EnableTwoFactorAuth(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeEnableTwoFactorAuthResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleFireGroupUserRequest handles fireGroupUser operation.
//
// POST /v1/groups/{group_id}/fire/{user_id}
func (s *Server) handleFireGroupUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("fireGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/fire/{user_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), FireGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: FireGroupUserOperation,
			ID:   "fireGroupUser",
		}
	)
	params, err := decodeFireGroupUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FireGroupUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    FireGroupUserOperation,
			OperationSummary: "",
			OperationID:      "fireGroupUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = FireGroupUserParams
			Response = *FireGroupUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackFireGroupUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.FireGroupUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.FireGroupUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFireGroupUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleFollowUserRequest handles followUser operation.
//
// POST /v2/users/{id}/follow
func (s *Server) handleFollowUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("followUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/{id}/follow"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), FollowUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: FollowUserOperation,
			ID:   "followUser",
		}
	)
	params, err := decodeFollowUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FollowUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    FollowUserOperation,
			OperationSummary: "",
			OperationID:      "followUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = FollowUserParams
			Response = *FollowUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackFollowUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.FollowUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.FollowUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFollowUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleFollowUsersRequest handles followUsers operation.
//
// POST /v2/users/follow
func (s *Server) handleFollowUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("followUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/follow"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), FollowUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: FollowUsersOperation,
			ID:   "followUsers",
		}
	)
	params, err := decodeFollowUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FollowUsersNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    FollowUsersOperation,
			OperationSummary: "",
			OperationID:      "followUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_ids[]",
					In:   "query",
				}: params.UserIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = FollowUsersParams
			Response = *FollowUsersNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackFollowUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.FollowUsers(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.FollowUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFollowUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGenerateCallActionSignatureRequest handles generateCallActionSignature operation.
//
// POST /v1/calls/action_signature/generate
func (s *Server) handleGenerateCallActionSignatureRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("generateCallActionSignature"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/action_signature/generate"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GenerateCallActionSignatureOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GenerateCallActionSignatureOperation,
			ID:   "generateCallActionSignature",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeGenerateCallActionSignatureRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CallActionSignatureResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GenerateCallActionSignatureOperation,
			OperationSummary: "",
			OperationID:      "generateCallActionSignature",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *GenerateCallActionSignatureReq
			Params   = struct{}
			Response = *CallActionSignatureResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GenerateCallActionSignature(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.GenerateCallActionSignature(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGenerateCallActionSignatureResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetActiveCallPostRequest handles getActiveCallPost operation.
//
// GET /v1/posts/active_call
func (s *Server) handleGetActiveCallPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getActiveCallPost"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/posts/active_call"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetActiveCallPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetActiveCallPostOperation,
			ID:   "getActiveCallPost",
		}
	)
	params, err := decodeGetActiveCallPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetActiveCallPostOperation,
			OperationSummary: "",
			OperationID:      "getActiveCallPost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "query",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetActiveCallPostParams
			Response = *PostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetActiveCallPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetActiveCallPost(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetActiveCallPost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetActiveCallPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetActiveFollowingsRequest handles getActiveFollowings operation.
//
// GET /v1/users/active_followings
func (s *Server) handleGetActiveFollowingsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getActiveFollowings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/active_followings"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetActiveFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetActiveFollowingsOperation,
			ID:   "getActiveFollowings",
		}
	)
	params, err := decodeGetActiveFollowingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ActiveFollowingsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetActiveFollowingsOperation,
			OperationSummary: "",
			OperationID:      "getActiveFollowings",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "only_online",
					In:   "query",
				}: params.OnlyOnline,
				{
					Name: "from_loggedin_at",
					In:   "query",
				}: params.FromLoggedinAt,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetActiveFollowingsParams
			Response = *ActiveFollowingsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetActiveFollowingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetActiveFollowings(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetActiveFollowings(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetActiveFollowingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetAdditionalNotificationSettingRequest handles getAdditionalNotificationSetting operation.
//
// GET /v1/users/additonal_notification_setting
func (s *Server) handleGetAdditionalNotificationSettingRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAdditionalNotificationSetting"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/additonal_notification_setting"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetAdditionalNotificationSettingOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *AdditionalSettingsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetAdditionalNotificationSettingOperation,
			OperationSummary: "",
			OperationID:      "getAdditionalNotificationSetting",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *AdditionalSettingsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetAdditionalNotificationSetting(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetAdditionalNotificationSetting(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetAdditionalNotificationSettingResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetAgoraRtmTokenRequest handles getAgoraRtmToken operation.
//
// GET /v3/calls/{call_id}/agora_rtm_token
func (s *Server) handleGetAgoraRtmTokenRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAgoraRtmToken"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v3/calls/{call_id}/agora_rtm_token"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetAgoraRtmTokenOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetAgoraRtmTokenOperation,
			ID:   "getAgoraRtmToken",
		}
	)
	params, err := decodeGetAgoraRtmTokenParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RtmTokenResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetAgoraRtmTokenOperation,
			OperationSummary: "",
			OperationID:      "getAgoraRtmToken",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetAgoraRtmTokenParams
			Response = *RtmTokenResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetAgoraRtmTokenParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetAgoraRtmToken(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetAgoraRtmToken(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetAgoraRtmTokenResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetAppConfigRequest handles getAppConfig operation.
//
// GET /api/apps/{app}
func (s *Server) handleGetAppConfigRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAppConfig"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/apps/{app}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetAppConfigOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetAppConfigOperation,
			ID:   "getAppConfig",
		}
	)
	params, err := decodeGetAppConfigParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ApplicationConfigResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetAppConfigOperation,
			OperationSummary: "",
			OperationID:      "getAppConfig",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "app",
					In:   "path",
				}: params.App,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetAppConfigParams
			Response = *ApplicationConfigResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetAppConfigParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetAppConfig(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetAppConfig(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetAppConfigResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetBannedWordsRequest handles getBannedWords operation.
//
// GET /{countryApiValue}/api/v2/banned_words
func (s *Server) handleGetBannedWordsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBannedWords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/{countryApiValue}/api/v2/banned_words"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetBannedWordsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetBannedWordsOperation,
			ID:   "getBannedWords",
		}
	)
	params, err := decodeGetBannedWordsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *BanWordsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBannedWordsOperation,
			OperationSummary: "",
			OperationID:      "getBannedWords",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "countryApiValue",
					In:   "path",
				}: params.CountryApiValue,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetBannedWordsParams
			Response = *BanWordsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBannedWordsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBannedWords(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBannedWords(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBannedWordsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetBlockedUserIdsRequest handles getBlockedUserIds operation.
//
// GET /v1/users/block_ids
func (s *Server) handleGetBlockedUserIdsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBlockedUserIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/block_ids"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetBlockedUserIdsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *BlockedUserIdsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBlockedUserIdsOperation,
			OperationSummary: "",
			OperationID:      "getBlockedUserIds",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *BlockedUserIdsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBlockedUserIds(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBlockedUserIds(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBlockedUserIdsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetBlockedUsersRequest handles getBlockedUsers operation.
//
// POST /v2/users/blocked
func (s *Server) handleGetBlockedUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBlockedUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/blocked"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetBlockedUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetBlockedUsersOperation,
			ID:   "getBlockedUsers",
		}
	)
	params, err := decodeGetBlockedUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeGetBlockedUsersRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *BlockedUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBlockedUsersOperation,
			OperationSummary: "",
			OperationID:      "getBlockedUsers",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
			},
			Raw: r,
		}

		type (
			Request  = *SearchUsersRequest
			Params   = GetBlockedUsersParams
			Response = *BlockedUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBlockedUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBlockedUsers(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBlockedUsers(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBlockedUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetBookmarkedPostsRequest handles getBookmarkedPosts operation.
//
// GET /v1/users/{user_id}/bookmarks
func (s *Server) handleGetBookmarkedPostsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBookmarkedPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/{user_id}/bookmarks"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetBookmarkedPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetBookmarkedPostsOperation,
			ID:   "getBookmarkedPosts",
		}
	)
	params, err := decodeGetBookmarkedPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBookmarkedPostsOperation,
			OperationSummary: "",
			OperationID:      "getBookmarkedPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetBookmarkedPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBookmarkedPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBookmarkedPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBookmarkedPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBookmarkedPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetBucketPresignedUrlsRequest handles getBucketPresignedUrls operation.
//
// GET /v1/buckets/presigned_urls
func (s *Server) handleGetBucketPresignedUrlsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBucketPresignedUrls"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/buckets/presigned_urls"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetBucketPresignedUrlsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetBucketPresignedUrlsOperation,
			ID:   "getBucketPresignedUrls",
		}
	)
	params, err := decodeGetBucketPresignedUrlsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PresignedUrlsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBucketPresignedUrlsOperation,
			OperationSummary: "",
			OperationID:      "getBucketPresignedUrls",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "file_names[]",
					In:   "query",
				}: params.FileNames,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetBucketPresignedUrlsParams
			Response = *PresignedUrlsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBucketPresignedUrlsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBucketPresignedUrls(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBucketPresignedUrls(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBucketPresignedUrlsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetCallBgmsRequest handles getCallBgms operation.
//
// GET /v1/calls/bgm
func (s *Server) handleGetCallBgmsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallBgms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/calls/bgm"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetCallBgmsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *BgmsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetCallBgmsOperation,
			OperationSummary: "",
			OperationID:      "getCallBgms",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *BgmsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetCallBgms(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetCallBgms(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetCallBgmsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetCallFollowersTimelineRequest handles getCallFollowersTimeline operation.
//
// GET /v2/posts/call_followers_timeline
func (s *Server) handleGetCallFollowersTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallFollowersTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/call_followers_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetCallFollowersTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetCallFollowersTimelineOperation,
			ID:   "getCallFollowersTimeline",
		}
	)
	params, err := decodeGetCallFollowersTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetCallFollowersTimelineOperation,
			OperationSummary: "",
			OperationID:      "getCallFollowersTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "category_id",
					In:   "query",
				}: params.CategoryID,
				{
					Name: "call_type",
					In:   "query",
				}: params.CallType,
				{
					Name: "include_circle_call",
					In:   "query",
				}: params.IncludeCircleCall,
				{
					Name: "exclude_recent_gomimushi",
					In:   "query",
				}: params.ExcludeRecentGomimushi,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetCallFollowersTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetCallFollowersTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetCallFollowersTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetCallFollowersTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetCallFollowersTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetCallGiftHistoryRequest handles getCallGiftHistory operation.
//
// GET /v1/calls/{call_id}/gift_transactions
func (s *Server) handleGetCallGiftHistoryRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallGiftHistory"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}/gift_transactions"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetCallGiftHistoryOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetCallGiftHistoryOperation,
			ID:   "getCallGiftHistory",
		}
	)
	params, err := decodeGetCallGiftHistoryParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *CallGiftHistoryResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetCallGiftHistoryOperation,
			OperationSummary: "",
			OperationID:      "getCallGiftHistory",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetCallGiftHistoryParams
			Response = *CallGiftHistoryResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetCallGiftHistoryParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetCallGiftHistory(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetCallGiftHistory(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetCallGiftHistoryResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetCallTimelineRequest handles getCallTimeline operation.
//
// GET /v2/posts/call_timeline
func (s *Server) handleGetCallTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/call_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetCallTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetCallTimelineOperation,
			ID:   "getCallTimeline",
		}
	)
	params, err := decodeGetCallTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetCallTimelineOperation,
			OperationSummary: "",
			OperationID:      "getCallTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "category_id",
					In:   "query",
				}: params.CategoryID,
				{
					Name: "call_type",
					In:   "query",
				}: params.CallType,
				{
					Name: "include_circle_call",
					In:   "query",
				}: params.IncludeCircleCall,
				{
					Name: "cross_generation",
					In:   "query",
				}: params.CrossGeneration,
				{
					Name: "exclude_recent_gomimushi",
					In:   "query",
				}: params.ExcludeRecentGomimushi,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetCallTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetCallTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetCallTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetCallTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetCallTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatMessagesRequest handles getChatMessages operation.
//
// GET /v2/chat_rooms/{id}/messages
func (s *Server) handleGetChatMessagesRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatMessages"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}/messages"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatMessagesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetChatMessagesOperation,
			ID:   "getChatMessages",
		}
	)
	params, err := decodeGetChatMessagesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *MessagesResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatMessagesOperation,
			OperationSummary: "",
			OperationID:      "getChatMessages",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from_message_id",
					In:   "query",
				}: params.FromMessageID,
				{
					Name: "to_message_id",
					In:   "query",
				}: params.ToMessageID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetChatMessagesParams
			Response = *MessagesResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetChatMessagesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatMessages(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatMessages(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatMessagesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatRequestCountRequest handles getChatRequestCount operation.
//
// GET /v1/chat_rooms/total_chat_request
func (s *Server) handleGetChatRequestCountRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRequestCount"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/total_chat_request"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatRequestCountOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *TotalChatRequestResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatRequestCountOperation,
			OperationSummary: "",
			OperationID:      "getChatRequestCount",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *TotalChatRequestResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatRequestCount(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatRequestCount(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatRequestCountResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatRequestsRequest handles getChatRequests operation.
//
// GET /v1/chat_rooms/request_list
func (s *Server) handleGetChatRequestsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRequests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/request_list"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatRequestsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetChatRequestsOperation,
			ID:   "getChatRequests",
		}
	)
	params, err := decodeGetChatRequestsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ChatRoomsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatRequestsOperation,
			OperationSummary: "",
			OperationID:      "getChatRequests",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetChatRequestsParams
			Response = *ChatRoomsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetChatRequestsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatRequests(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatRequests(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatRequestsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatRoomRequest handles getChatRoom operation.
//
// GET /v2/chat_rooms/{id}
func (s *Server) handleGetChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRoom"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetChatRoomOperation,
			ID:   "getChatRoom",
		}
	)
	params, err := decodeGetChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ChatRoomResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatRoomOperation,
			OperationSummary: "",
			OperationID:      "getChatRoom",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetChatRoomParams
			Response = *ChatRoomResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatRoom(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatRoom(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatUnreadStatusRequest handles getChatUnreadStatus operation.
//
// GET /v1/chat_rooms/unread_status
func (s *Server) handleGetChatUnreadStatusRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatUnreadStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/unread_status"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatUnreadStatusOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetChatUnreadStatusOperation,
			ID:   "getChatUnreadStatus",
		}
	)
	params, err := decodeGetChatUnreadStatusParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnreadStatusResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatUnreadStatusOperation,
			OperationSummary: "",
			OperationID:      "getChatUnreadStatus",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_time",
					In:   "query",
				}: params.FromTime,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetChatUnreadStatusParams
			Response = *UnreadStatusResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetChatUnreadStatusParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatUnreadStatus(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatUnreadStatus(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatUnreadStatusResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetChatableFollowingsRequest handles getChatableFollowings operation.
//
// POST /v1/users/followings/chatable
func (s *Server) handleGetChatableFollowingsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatableFollowings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/followings/chatable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetChatableFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetChatableFollowingsOperation,
			ID:   "getChatableFollowings",
		}
	)
	params, err := decodeGetChatableFollowingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeGetChatableFollowingsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *FollowUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetChatableFollowingsOperation,
			OperationSummary: "",
			OperationID:      "getChatableFollowings",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_follow_id",
					In:   "query",
				}: params.FromFollowID,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "order_by",
					In:   "query",
				}: params.OrderBy,
			},
			Raw: r,
		}

		type (
			Request  = *SearchUsersRequest
			Params   = GetChatableFollowingsParams
			Response = *FollowUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetChatableFollowingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetChatableFollowings(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetChatableFollowings(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetChatableFollowingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetConferenceCallRequest handles getConferenceCall operation.
//
// GET /v2/calls/conferences/{call_id}
func (s *Server) handleGetConferenceCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getConferenceCall"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/calls/conferences/{call_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetConferenceCallOperation,
			ID:   "getConferenceCall",
		}
	)
	params, err := decodeGetConferenceCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ConferenceCallResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetConferenceCallOperation,
			OperationSummary: "",
			OperationID:      "getConferenceCall",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetConferenceCallParams
			Response = *ConferenceCallResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetConferenceCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetConferenceCall(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetConferenceCall(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetConferenceCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetConversationRequest handles getConversation operation.
//
// GET /v2/conversations/{id}
func (s *Server) handleGetConversationRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getConversation"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/conversations/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetConversationOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetConversationOperation,
			ID:   "getConversation",
		}
	)
	params, err := decodeGetConversationParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetConversationOperation,
			OperationSummary: "",
			OperationID:      "getConversation",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
				{
					Name: "thread_id",
					In:   "query",
				}: params.ThreadID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "reverse",
					In:   "query",
				}: params.Reverse,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetConversationParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetConversationParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetConversation(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetConversation(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetConversationResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetDefaultSettingsRequest handles getDefaultSettings operation.
//
// GET /v1/users/default_settings
func (s *Server) handleGetDefaultSettingsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getDefaultSettings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/default_settings"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetDefaultSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *DefaultSettingsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetDefaultSettingsOperation,
			OperationSummary: "",
			OperationID:      "getDefaultSettings",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *DefaultSettingsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetDefaultSettings(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetDefaultSettings(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetDefaultSettingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFollowRequestCountRequest handles getFollowRequestCount operation.
//
// GET /v2/users/follow_requests_count
func (s *Server) handleGetFollowRequestCountRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowRequestCount"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/follow_requests_count"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFollowRequestCountOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *FollowRequestCountResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFollowRequestCountOperation,
			OperationSummary: "",
			OperationID:      "getFollowRequestCount",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *FollowRequestCountResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFollowRequestCount(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFollowRequestCount(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFollowRequestCountResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFollowRequestsRequest handles getFollowRequests operation.
//
// GET /v2/users/follow_requests
func (s *Server) handleGetFollowRequestsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowRequests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/follow_requests"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFollowRequestsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetFollowRequestsOperation,
			ID:   "getFollowRequests",
		}
	)
	params, err := decodeGetFollowRequestsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersByTimestampResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFollowRequestsOperation,
			OperationSummary: "",
			OperationID:      "getFollowRequests",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetFollowRequestsParams
			Response = *UsersByTimestampResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetFollowRequestsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFollowRequests(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFollowRequests(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFollowRequestsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFollowingTimelineRequest handles getFollowingTimeline operation.
//
// GET /v2/posts/following_timeline
func (s *Server) handleGetFollowingTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowingTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/following_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFollowingTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetFollowingTimelineOperation,
			ID:   "getFollowingTimeline",
		}
	)
	params, err := decodeGetFollowingTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFollowingTimelineOperation,
			OperationSummary: "",
			OperationID:      "getFollowingTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "only_root",
					In:   "query",
				}: params.OnlyRoot,
				{
					Name: "order_by",
					In:   "query",
				}: params.OrderBy,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "mxn",
					In:   "query",
				}: params.Mxn,
				{
					Name: "reduce_selfie",
					In:   "query",
				}: params.ReduceSelfie,
				{
					Name: "custom_generation_range",
					In:   "query",
				}: params.CustomGenerationRange,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetFollowingTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetFollowingTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFollowingTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFollowingTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFollowingTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFollowingsBornTodayRequest handles getFollowingsBornToday operation.
//
// GET /v1/users/following_born_today
func (s *Server) handleGetFollowingsBornTodayRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowingsBornToday"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/following_born_today"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFollowingsBornTodayOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetFollowingsBornTodayOperation,
			ID:   "getFollowingsBornToday",
		}
	)
	params, err := decodeGetFollowingsBornTodayParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFollowingsBornTodayOperation,
			OperationSummary: "",
			OperationID:      "getFollowingsBornToday",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "birthdate",
					In:   "query",
				}: params.Birthdate,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetFollowingsBornTodayParams
			Response = *UsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetFollowingsBornTodayParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFollowingsBornToday(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFollowingsBornToday(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFollowingsBornTodayResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFootprintsRequest handles getFootprints operation.
//
// GET /v3/users/footprints
func (s *Server) handleGetFootprintsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFootprints"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v3/users/footprints"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFootprintsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetFootprintsOperation,
			ID:   "getFootprints",
		}
	)
	params, err := decodeGetFootprintsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FootprintsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFootprintsOperation,
			OperationSummary: "",
			OperationID:      "getFootprints",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "mode",
					In:   "query",
				}: params.Mode,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetFootprintsParams
			Response = *FootprintsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetFootprintsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFootprints(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFootprints(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFootprintsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetFreshUserRequest handles getFreshUser operation.
//
// GET /v2/users/fresh/{id}
func (s *Server) handleGetFreshUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFreshUser"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/fresh/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetFreshUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetFreshUserOperation,
			ID:   "getFreshUser",
		}
	)
	params, err := decodeGetFreshUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetFreshUserOperation,
			OperationSummary: "",
			OperationID:      "getFreshUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetFreshUserParams
			Response = *UserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetFreshUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetFreshUser(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetFreshUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetFreshUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupRequest handles getGroup operation.
//
// GET /v1/groups/{id}
func (s *Server) handleGetGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroup"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupOperation,
			ID:   "getGroup",
		}
	)
	params, err := decodeGetGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupOperation,
			OperationSummary: "",
			OperationID:      "getGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupParams
			Response = *GroupResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroup(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupBanListRequest handles getGroupBanList operation.
//
// GET /v1/groups/{id}/ban_list
func (s *Server) handleGetGroupBanListRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupBanList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/ban_list"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupBanListOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupBanListOperation,
			ID:   "getGroupBanList",
		}
	)
	params, err := decodeGetGroupBanListParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupBanListOperation,
			OperationSummary: "",
			OperationID:      "getGroupBanList",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "page",
					In:   "query",
				}: params.Page,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupBanListParams
			Response = *UsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupBanListParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupBanList(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupBanList(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupBanListResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupCreateQuotaRequest handles getGroupCreateQuota operation.
//
// GET /v1/groups/created_quota
func (s *Server) handleGetGroupCreateQuotaRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupCreateQuota"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/created_quota"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupCreateQuotaOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *CreateQuotaResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupCreateQuotaOperation,
			OperationSummary: "",
			OperationID:      "getGroupCreateQuota",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *CreateQuotaResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupCreateQuota(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupCreateQuota(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupCreateQuotaResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupGiftHistoryRequest handles getGroupGiftHistory operation.
//
// GET /v1/groups/{group_id}/gift_history
func (s *Server) handleGetGroupGiftHistoryRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupGiftHistory"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/gift_history"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupGiftHistoryOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupGiftHistoryOperation,
			ID:   "getGroupGiftHistory",
		}
	)
	params, err := decodeGetGroupGiftHistoryParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupGiftHistoryResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupGiftHistoryOperation,
			OperationSummary: "",
			OperationID:      "getGroupGiftHistory",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupGiftHistoryParams
			Response = *GroupGiftHistoryResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupGiftHistoryParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupGiftHistory(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupGiftHistory(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupGiftHistoryResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupGiftTransactionsRequest handles getGroupGiftTransactions operation.
//
// GET /v1/groups/{group_id}/gift_transactions
func (s *Server) handleGetGroupGiftTransactionsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/gift_transactions"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupGiftTransactionsOperation,
			ID:   "getGroupGiftTransactions",
		}
	)
	params, err := decodeGetGroupGiftTransactionsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftTransactionsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupGiftTransactionsOperation,
			OperationSummary: "",
			OperationID:      "getGroupGiftTransactions",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupGiftTransactionsParams
			Response = *GiftTransactionsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupGiftTransactionsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupGiftTransactions(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupGiftTransactions(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupGiftTransactionsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupHighlightsRequest handles getGroupHighlights operation.
//
// GET /v1/groups/{group_id}/highlights
func (s *Server) handleGetGroupHighlightsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupHighlights"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/highlights"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupHighlightsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupHighlightsOperation,
			ID:   "getGroupHighlights",
		}
	)
	params, err := decodeGetGroupHighlightsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupHighlightsOperation,
			OperationSummary: "",
			OperationID:      "getGroupHighlights",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupHighlightsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupHighlightsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupHighlights(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupHighlights(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupHighlightsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupMemberRequest handles getGroupMember operation.
//
// GET /v1/groups/{id}/members/{userId}
func (s *Server) handleGetGroupMemberRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupMember"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/members/{userId}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupMemberOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupMemberOperation,
			ID:   "getGroupMember",
		}
	)
	params, err := decodeGetGroupMemberParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupUserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupMemberOperation,
			OperationSummary: "",
			OperationID:      "getGroupMember",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupMemberParams
			Response = *GroupUserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupMemberParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupMember(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupMember(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupMemberResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupMembersRequest handles getGroupMembers operation.
//
// GET /v2/groups/{id}/members
func (s *Server) handleGetGroupMembersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupMembers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/groups/{id}/members"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupMembersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupMembersOperation,
			ID:   "getGroupMembers",
		}
	)
	params, err := decodeGetGroupMembersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupMembersOperation,
			OperationSummary: "",
			OperationID:      "getGroupMembers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupMembersParams
			Response = *GroupUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupMembersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupMembers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupMembers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupMembersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupNotificationSettingsRequest handles getGroupNotificationSettings operation.
//
// GET /v2/notification_settings/groups/{id}
func (s *Server) handleGetGroupNotificationSettingsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/notification_settings/groups/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupNotificationSettingsOperation,
			ID:   "getGroupNotificationSettings",
		}
	)
	params, err := decodeGetGroupNotificationSettingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupNotificationSettingsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupNotificationSettingsOperation,
			OperationSummary: "",
			OperationID:      "getGroupNotificationSettings",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupNotificationSettingsParams
			Response = *GroupNotificationSettingsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupNotificationSettingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupNotificationSettings(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupNotificationSettings(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupNotificationSettingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupReceivedGiftSendersRequest handles getGroupReceivedGiftSenders operation.
//
// GET /v1/groups/{group_id}/received_gifts/{gift_id}/senders
func (s *Server) handleGetGroupReceivedGiftSendersRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupReceivedGiftSenders"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/received_gifts/{gift_id}/senders"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupReceivedGiftSendersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupReceivedGiftSendersOperation,
			ID:   "getGroupReceivedGiftSenders",
		}
	)
	params, err := decodeGetGroupReceivedGiftSendersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftSendersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupReceivedGiftSendersOperation,
			OperationSummary: "",
			OperationID:      "getGroupReceivedGiftSenders",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "gift_id",
					In:   "path",
				}: params.GiftID,
				{
					Name: "user_id",
					In:   "query",
				}: params.UserID,
				{
					Name: "timestamp",
					In:   "query",
				}: params.Timestamp,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupReceivedGiftSendersParams
			Response = *GiftSendersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupReceivedGiftSendersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupReceivedGiftSenders(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupReceivedGiftSenders(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupReceivedGiftSendersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupTimelineRequest handles getGroupTimeline operation.
//
// GET /v2/posts/group_timeline
func (s *Server) handleGetGroupTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/group_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupTimelineOperation,
			ID:   "getGroupTimeline",
		}
	)
	params, err := decodeGetGroupTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupTimelineOperation,
			OperationSummary: "",
			OperationID:      "getGroupTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "reverse",
					In:   "query",
				}: params.Reverse,
				{
					Name: "post_type",
					In:   "query",
				}: params.PostType,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "only_root",
					In:   "query",
				}: params.OnlyRoot,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetGroupUnreadStatusRequest handles getGroupUnreadStatus operation.
//
// GET /v1/groups/unread_status
func (s *Server) handleGetGroupUnreadStatusRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupUnreadStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/unread_status"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetGroupUnreadStatusOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetGroupUnreadStatusOperation,
			ID:   "getGroupUnreadStatus",
		}
	)
	params, err := decodeGetGroupUnreadStatusParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnreadStatusResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetGroupUnreadStatusOperation,
			OperationSummary: "",
			OperationID:      "getGroupUnreadStatus",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_time",
					In:   "query",
				}: params.FromTime,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetGroupUnreadStatusParams
			Response = *UnreadStatusResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetGroupUnreadStatusParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetGroupUnreadStatus(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetGroupUnreadStatus(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetGroupUnreadStatusResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetHimaUsersRequest handles getHimaUsers operation.
//
// GET /v2/users/hima_users
func (s *Server) handleGetHimaUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getHimaUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/hima_users"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetHimaUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetHimaUsersOperation,
			ID:   "getHimaUsers",
		}
	)
	params, err := decodeGetHimaUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *HimaUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetHimaUsersOperation,
			OperationSummary: "",
			OperationID:      "getHimaUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_hima_id",
					In:   "query",
				}: params.FromHimaID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetHimaUsersParams
			Response = *HimaUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetHimaUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetHimaUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetHimaUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetHimaUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetInvitableCallUsersRequest handles getInvitableCallUsers operation.
//
// GET /v1/calls/{call_id}/users/invitable
func (s *Server) handleGetInvitableCallUsersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getInvitableCallUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}/users/invitable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetInvitableCallUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetInvitableCallUsersOperation,
			ID:   "getInvitableCallUsers",
		}
	)
	params, err := decodeGetInvitableCallUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersByTimestampResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetInvitableCallUsersOperation,
			OperationSummary: "",
			OperationID:      "getInvitableCallUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "user[nickname]",
					In:   "query",
				}: params.UserNickname,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetInvitableCallUsersParams
			Response = *UsersByTimestampResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetInvitableCallUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetInvitableCallUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetInvitableCallUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetInvitableCallUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetInvitableGroupUsersRequest handles getInvitableGroupUsers operation.
//
// GET /v1/groups/{group_id}/users/invitable
func (s *Server) handleGetInvitableGroupUsersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getInvitableGroupUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/users/invitable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetInvitableGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetInvitableGroupUsersOperation,
			ID:   "getInvitableGroupUsers",
		}
	)
	params, err := decodeGetInvitableGroupUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersByTimestampResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetInvitableGroupUsersOperation,
			OperationSummary: "",
			OperationID:      "getInvitableGroupUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "user[nickname]",
					In:   "query",
				}: params.UserNickname,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetInvitableGroupUsersParams
			Response = *UsersByTimestampResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetInvitableGroupUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetInvitableGroupUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetInvitableGroupUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetInvitableGroupUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetJoinedGroupStatusesRequest handles getJoinedGroupStatuses operation.
//
// GET /v1/groups/joined_statuses
func (s *Server) handleGetJoinedGroupStatusesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getJoinedGroupStatuses"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/joined_statuses"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetJoinedGroupStatusesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetJoinedGroupStatusesOperation,
			ID:   "getJoinedGroupStatuses",
		}
	)
	params, err := decodeGetJoinedGroupStatusesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response GetJoinedGroupStatusesOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetJoinedGroupStatusesOperation,
			OperationSummary: "",
			OperationID:      "getJoinedGroupStatuses",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "ids[]",
					In:   "query",
				}: params.Ids,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetJoinedGroupStatusesParams
			Response = GetJoinedGroupStatusesOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetJoinedGroupStatusesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetJoinedGroupStatuses(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetJoinedGroupStatuses(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetJoinedGroupStatusesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetJoinedThreadStatusesRequest handles getJoinedThreadStatuses operation.
//
// GET /v1/threads/joined_statuses
func (s *Server) handleGetJoinedThreadStatusesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getJoinedThreadStatuses"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/threads/joined_statuses"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetJoinedThreadStatusesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetJoinedThreadStatusesOperation,
			ID:   "getJoinedThreadStatuses",
		}
	)
	params, err := decodeGetJoinedThreadStatusesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response GetJoinedThreadStatusesOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetJoinedThreadStatusesOperation,
			OperationSummary: "",
			OperationID:      "getJoinedThreadStatuses",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "ids[]",
					In:   "query",
				}: params.Ids,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetJoinedThreadStatusesParams
			Response = GetJoinedThreadStatusesOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetJoinedThreadStatusesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetJoinedThreadStatuses(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetJoinedThreadStatuses(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetJoinedThreadStatusesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetMainChatRoomsRequest handles getMainChatRooms operation.
//
// GET /v1/chat_rooms/main_list
func (s *Server) handleGetMainChatRoomsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMainChatRooms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/main_list"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetMainChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetMainChatRoomsOperation,
			ID:   "getMainChatRooms",
		}
	)
	params, err := decodeGetMainChatRoomsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ChatRoomsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetMainChatRoomsOperation,
			OperationSummary: "",
			OperationID:      "getMainChatRooms",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetMainChatRoomsParams
			Response = *ChatRoomsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetMainChatRoomsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetMainChatRooms(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetMainChatRooms(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetMainChatRoomsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetMyPostsRequest handles getMyPosts operation.
//
// GET /v2/posts/mine
func (s *Server) handleGetMyPostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMyPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/mine"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetMyPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetMyPostsOperation,
			ID:   "getMyPosts",
		}
	)
	params, err := decodeGetMyPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetMyPostsOperation,
			OperationSummary: "",
			OperationID:      "getMyPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "include_group_post",
					In:   "query",
				}: params.IncludeGroupPost,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetMyPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetMyPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetMyPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetMyPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetMyPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetMyReviewsRequest handles getMyReviews operation.
//
// GET /v1/users/reviews/mine
func (s *Server) handleGetMyReviewsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMyReviews"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/reviews/mine"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetMyReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetMyReviewsOperation,
			ID:   "getMyReviews",
		}
	)
	params, err := decodeGetMyReviewsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ReviewsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetMyReviewsOperation,
			OperationSummary: "",
			OperationID:      "getMyReviews",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetMyReviewsParams
			Response = *ReviewsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetMyReviewsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetMyReviews(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetMyReviews(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetMyReviewsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPhoneStatusRequest handles getPhoneStatus operation.
//
// GET /v1/calls/phone_status/{opponent_id}
func (s *Server) handleGetPhoneStatusRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPhoneStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/calls/phone_status/{opponent_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPhoneStatusOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPhoneStatusOperation,
			ID:   "getPhoneStatus",
		}
	)
	params, err := decodeGetPhoneStatusParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *CallStatusResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPhoneStatusOperation,
			OperationSummary: "",
			OperationID:      "getPhoneStatus",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "opponent_id",
					In:   "path",
				}: params.OpponentID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPhoneStatusParams
			Response = *CallStatusResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPhoneStatusParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPhoneStatus(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPhoneStatus(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPhoneStatusResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPolicyAgreementsRequest handles getPolicyAgreements operation.
//
// GET /v1/users/policy_agreements
func (s *Server) handleGetPolicyAgreementsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPolicyAgreements"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/policy_agreements"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPolicyAgreementsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *PolicyAgreementsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPolicyAgreementsOperation,
			OperationSummary: "",
			OperationID:      "getPolicyAgreements",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *PolicyAgreementsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPolicyAgreements(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPolicyAgreements(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPolicyAgreementsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPopularWordsRequest handles getPopularWords operation.
//
// GET /{countryApiValue}/api/apps/{app}/popular_words
func (s *Server) handleGetPopularWordsRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPopularWords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/{countryApiValue}/api/apps/{app}/popular_words"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPopularWordsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPopularWordsOperation,
			ID:   "getPopularWords",
		}
	)
	params, err := decodeGetPopularWordsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PopularWordsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPopularWordsOperation,
			OperationSummary: "",
			OperationID:      "getPopularWords",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "countryApiValue",
					In:   "path",
				}: params.CountryApiValue,
				{
					Name: "app",
					In:   "path",
				}: params.App,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPopularWordsParams
			Response = *PopularWordsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPopularWordsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPopularWords(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPopularWords(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPopularWordsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostRequest handles getPost operation.
//
// GET /v2/posts/{id}
func (s *Server) handleGetPostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPost"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostOperation,
			ID:   "getPost",
		}
	)
	params, err := decodeGetPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostOperation,
			OperationSummary: "",
			OperationID:      "getPost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "Cache-Control",
					In:   "header",
				}: params.CacheControl,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostParams
			Response = *PostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPost(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostGiftTransactionsRequest handles getPostGiftTransactions operation.
//
// GET /v1/posts/{post_id}/gift_transactions
func (s *Server) handleGetPostGiftTransactionsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/posts/{post_id}/gift_transactions"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostGiftTransactionsOperation,
			ID:   "getPostGiftTransactions",
		}
	)
	params, err := decodeGetPostGiftTransactionsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftTransactionsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostGiftTransactionsOperation,
			OperationSummary: "",
			OperationID:      "getPostGiftTransactions",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "post_id",
					In:   "path",
				}: params.PostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostGiftTransactionsParams
			Response = *GiftTransactionsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostGiftTransactionsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostGiftTransactions(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostGiftTransactions(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostGiftTransactionsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostLikersRequest handles getPostLikers operation.
//
// GET /v1/posts/{id}/likers
func (s *Server) handleGetPostLikersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostLikers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/posts/{id}/likers"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostLikersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostLikersOperation,
			ID:   "getPostLikers",
		}
	)
	params, err := decodeGetPostLikersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostLikersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostLikersOperation,
			OperationSummary: "",
			OperationID:      "getPostLikers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostLikersParams
			Response = *PostLikersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostLikersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostLikers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostLikers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostLikersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostRepostsRequest handles getPostReposts operation.
//
// GET /v2/posts/{id}/reposts
func (s *Server) handleGetPostRepostsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostReposts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/{id}/reposts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostRepostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostRepostsOperation,
			ID:   "getPostReposts",
		}
	)
	params, err := decodeGetPostRepostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostRepostsOperation,
			OperationSummary: "",
			OperationID:      "getPostReposts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostRepostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostRepostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostReposts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostReposts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostRepostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostUrlMetadataRequest handles getPostUrlMetadata operation.
//
// GET /v2/posts/url_metadata
func (s *Server) handleGetPostUrlMetadataRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostUrlMetadata"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/url_metadata"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostUrlMetadataOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostUrlMetadataOperation,
			ID:   "getPostUrlMetadata",
		}
	)
	params, err := decodeGetPostUrlMetadataParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *SharedUrl
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostUrlMetadataOperation,
			OperationSummary: "",
			OperationID:      "getPostUrlMetadata",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "url",
					In:   "query",
				}: params.URL,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostUrlMetadataParams
			Response = *SharedUrl
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostUrlMetadataParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostUrlMetadata(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostUrlMetadata(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostUrlMetadataResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostsByIdsRequest handles getPostsByIds operation.
//
// GET /v2/posts/multiple
func (s *Server) handleGetPostsByIdsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostsByIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/multiple"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostsByIdsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostsByIdsOperation,
			ID:   "getPostsByIds",
		}
	)
	params, err := decodeGetPostsByIdsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostsByIdsOperation,
			OperationSummary: "",
			OperationID:      "getPostsByIds",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "post_ids[]",
					In:   "query",
				}: params.PostIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostsByIdsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostsByIdsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostsByIds(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostsByIds(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostsByIdsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPostsByTagRequest handles getPostsByTag operation.
//
// GET /v2/posts/tags/{tag}
func (s *Server) handleGetPostsByTagRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostsByTag"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/tags/{tag}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetPostsByTagOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPostsByTagOperation,
			ID:   "getPostsByTag",
		}
	)
	params, err := decodeGetPostsByTagParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPostsByTagOperation,
			OperationSummary: "",
			OperationID:      "getPostsByTag",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "tag",
					In:   "path",
				}: params.Tag,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPostsByTagParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPostsByTagParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPostsByTag(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPostsByTag(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPostsByTagResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetReceivedGiftSendersRequest handles getReceivedGiftSenders operation.
//
// GET /v1/received_gifts/{gift_id}/senders
func (s *Server) handleGetReceivedGiftSendersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getReceivedGiftSenders"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/received_gifts/{gift_id}/senders"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetReceivedGiftSendersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetReceivedGiftSendersOperation,
			ID:   "getReceivedGiftSenders",
		}
	)
	params, err := decodeGetReceivedGiftSendersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftSendersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetReceivedGiftSendersOperation,
			OperationSummary: "",
			OperationID:      "getReceivedGiftSenders",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "gift_id",
					In:   "path",
				}: params.GiftID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetReceivedGiftSendersParams
			Response = *GiftSendersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetReceivedGiftSendersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetReceivedGiftSenders(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetReceivedGiftSenders(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetReceivedGiftSendersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetReceivedGiftTransactionRequest handles getReceivedGiftTransaction operation.
//
// GET /v1/received_gifts/{gift_transaction_uuid}
func (s *Server) handleGetReceivedGiftTransactionRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getReceivedGiftTransaction"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/received_gifts/{gift_transaction_uuid}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetReceivedGiftTransactionOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetReceivedGiftTransactionOperation,
			ID:   "getReceivedGiftTransaction",
		}
	)
	params, err := decodeGetReceivedGiftTransactionParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftReceivedTransactionResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetReceivedGiftTransactionOperation,
			OperationSummary: "",
			OperationID:      "getReceivedGiftTransaction",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "gift_transaction_uuid",
					In:   "path",
				}: params.GiftTransactionUUID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetReceivedGiftTransactionParams
			Response = *GiftReceivedTransactionResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetReceivedGiftTransactionParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetReceivedGiftTransaction(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetReceivedGiftTransaction(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetReceivedGiftTransactionResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRecentEngagementPostsRequest handles getRecentEngagementPosts operation.
//
// GET /v2/posts/recent_engagement
func (s *Server) handleGetRecentEngagementPostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecentEngagementPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/recent_engagement"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRecentEngagementPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRecentEngagementPostsOperation,
			ID:   "getRecentEngagementPosts",
		}
	)
	params, err := decodeGetRecentEngagementPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRecentEngagementPostsOperation,
			OperationSummary: "",
			OperationID:      "getRecentEngagementPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRecentEngagementPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRecentEngagementPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRecentEngagementPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRecentEngagementPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRecentEngagementPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRecommendedFollowUsersRequest handles getRecommendedFollowUsers operation.
//
// GET /v1/users/{id}/follow_recommended
func (s *Server) handleGetRecommendedFollowUsersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedFollowUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/{id}/follow_recommended"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRecommendedFollowUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRecommendedFollowUsersOperation,
			ID:   "getRecommendedFollowUsers",
		}
	)
	params, err := decodeGetRecommendedFollowUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRecommendedFollowUsersOperation,
			OperationSummary: "",
			OperationID:      "getRecommendedFollowUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "page",
					In:   "query",
				}: params.Page,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRecommendedFollowUsersParams
			Response = *UsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRecommendedFollowUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRecommendedFollowUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRecommendedFollowUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRecommendedFollowUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRecommendedPostTagsRequest handles getRecommendedPostTags operation.
//
// POST /v1/posts/recommended_tag
func (s *Server) handleGetRecommendedPostTagsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedPostTags"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/posts/recommended_tag"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRecommendedPostTagsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRecommendedPostTagsOperation,
			ID:   "getRecommendedPostTags",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeGetRecommendedPostTagsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PostTagsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRecommendedPostTagsOperation,
			OperationSummary: "",
			OperationID:      "getRecommendedPostTags",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *GetRecommendedPostTagsReq
			Params   = struct{}
			Response = *PostTagsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRecommendedPostTags(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRecommendedPostTags(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRecommendedPostTagsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRecommendedTimelineRequest handles getRecommendedTimeline operation.
//
// GET /v2/posts/recommended_timeline
func (s *Server) handleGetRecommendedTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/recommended_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRecommendedTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRecommendedTimelineOperation,
			ID:   "getRecommendedTimeline",
		}
	)
	params, err := decodeGetRecommendedTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRecommendedTimelineOperation,
			OperationSummary: "",
			OperationID:      "getRecommendedTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "experiment_num",
					In:   "query",
				}: params.ExperimentNum,
				{
					Name: "variant_num",
					In:   "query",
				}: params.VariantNum,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRecommendedTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRecommendedTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRecommendedTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRecommendedTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRecommendedTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRelatableGroupsRequest handles getRelatableGroups operation.
//
// GET /v1/groups/{id}/relatable
func (s *Server) handleGetRelatableGroupsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRelatableGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/relatable"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRelatableGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRelatableGroupsOperation,
			ID:   "getRelatableGroups",
		}
	)
	params, err := decodeGetRelatableGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupsRelatedResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRelatableGroupsOperation,
			OperationSummary: "",
			OperationID:      "getRelatableGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRelatableGroupsParams
			Response = *GroupsRelatedResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRelatableGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRelatableGroups(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRelatableGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRelatableGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRelatedGroupsRequest handles getRelatedGroups operation.
//
// GET /v1/groups/{id}/related
func (s *Server) handleGetRelatedGroupsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/related"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRelatedGroupsOperation,
			ID:   "getRelatedGroups",
		}
	)
	params, err := decodeGetRelatedGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupsRelatedResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRelatedGroupsOperation,
			OperationSummary: "",
			OperationID:      "getRelatedGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRelatedGroupsParams
			Response = *GroupsRelatedResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRelatedGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRelatedGroups(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRelatedGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRelatedGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetResetCountersRequest handles getResetCounters operation.
//
// GET /v1/users/reset_counters
func (s *Server) handleGetResetCountersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getResetCounters"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/reset_counters"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetResetCountersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *RefreshCounterRequestsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetResetCountersOperation,
			OperationSummary: "",
			OperationID:      "getResetCounters",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *RefreshCounterRequestsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetResetCounters(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetResetCounters(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetResetCountersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetRootPostsRequest handles getRootPosts operation.
//
// GET /v2/conversations/root_posts
func (s *Server) handleGetRootPostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRootPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/conversations/root_posts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetRootPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetRootPostsOperation,
			ID:   "getRootPosts",
		}
	)
	params, err := decodeGetRootPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetRootPostsOperation,
			OperationSummary: "",
			OperationID:      "getRootPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "ids[]",
					In:   "query",
				}: params.Ids,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetRootPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetRootPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetRootPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetRootPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetRootPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetThreadRequest handles getThread operation.
//
// GET /v1/threads/{id}
func (s *Server) handleGetThreadRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getThread"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/threads/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetThreadOperation,
			ID:   "getThread",
		}
	)
	params, err := decodeGetThreadParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ThreadInfo
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetThreadOperation,
			OperationSummary: "",
			OperationID:      "getThread",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetThreadParams
			Response = *ThreadInfo
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetThreadParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetThread(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetThread(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetThreadPostsRequest handles getThreadPosts operation.
//
// GET /v1/threads/{id}/posts
func (s *Server) handleGetThreadPostsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getThreadPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/threads/{id}/posts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetThreadPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetThreadPostsOperation,
			ID:   "getThreadPosts",
		}
	)
	params, err := decodeGetThreadPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetThreadPostsOperation,
			OperationSummary: "",
			OperationID:      "getThreadPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "post_type",
					In:   "query",
				}: params.PostType,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetThreadPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetThreadPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetThreadPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetThreadPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetThreadPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTimelineRequest handles getTimeline operation.
//
// GET /v2/posts/{noreply_mode}timeline
func (s *Server) handleGetTimelineRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/{noreply_mode}timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetTimelineOperation,
			ID:   "getTimeline",
		}
	)
	params, err := decodeGetTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetTimelineOperation,
			OperationSummary: "",
			OperationID:      "getTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "noreply_mode",
					In:   "path",
				}: params.NoreplyMode,
				{
					Name: "order_by",
					In:   "query",
				}: params.OrderBy,
				{
					Name: "experiment_older_age_rules",
					In:   "query",
				}: params.ExperimentOlderAgeRules,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "mxn",
					In:   "query",
				}: params.Mxn,
				{
					Name: "en",
					In:   "query",
				}: params.En,
				{
					Name: "vn",
					In:   "query",
				}: params.Vn,
				{
					Name: "reduce_selfie",
					In:   "query",
				}: params.ReduceSelfie,
				{
					Name: "custom_generation_range",
					In:   "query",
				}: params.CustomGenerationRange,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTwoFactorAuthRequestInfoRequest handles getTwoFactorAuthRequestInfo operation.
//
// GET /v1/users/secret
func (s *Server) handleGetTwoFactorAuthRequestInfoRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTwoFactorAuthRequestInfo"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/secret"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetTwoFactorAuthRequestInfoOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *TwoStepAuthRequestInfoResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetTwoFactorAuthRequestInfoOperation,
			OperationSummary: "",
			OperationID:      "getTwoFactorAuthRequestInfo",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *TwoStepAuthRequestInfoResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTwoFactorAuthRequestInfo(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTwoFactorAuthRequestInfo(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetTwoFactorAuthRequestInfoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTwoFactorAuthStatusRequest handles getTwoFactorAuthStatus operation.
//
// GET /v1/users/secret/status
func (s *Server) handleGetTwoFactorAuthStatusRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTwoFactorAuthStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/secret/status"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetTwoFactorAuthStatusOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *TwoFAStatusResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetTwoFactorAuthStatusOperation,
			OperationSummary: "",
			OperationID:      "getTwoFactorAuthStatus",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *TwoFAStatusResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTwoFactorAuthStatus(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTwoFactorAuthStatus(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetTwoFactorAuthStatusResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUpdatedChatRoomsRequest handles getUpdatedChatRooms operation.
//
// GET /v2/chat_rooms/update
func (s *Server) handleGetUpdatedChatRoomsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUpdatedChatRooms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/update"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUpdatedChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUpdatedChatRoomsOperation,
			ID:   "getUpdatedChatRooms",
		}
	)
	params, err := decodeGetUpdatedChatRoomsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ChatRoomsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUpdatedChatRoomsOperation,
			OperationSummary: "",
			OperationID:      "getUpdatedChatRooms",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_time",
					In:   "query",
				}: params.FromTime,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUpdatedChatRoomsParams
			Response = *ChatRoomsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUpdatedChatRoomsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUpdatedChatRooms(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUpdatedChatRooms(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUpdatedChatRoomsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserRequest handles getUser operation.
//
// GET /v2/users/{id}
func (s *Server) handleGetUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUser"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserOperation,
			ID:   "getUser",
		}
	)
	params, err := decodeGetUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserOperation,
			OperationSummary: "",
			OperationID:      "getUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserParams
			Response = *UserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUser(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserActivitiesRequest handles getUserActivities operation.
//
// GET /api/v2/user_activities
func (s *Server) handleGetUserActivitiesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserActivities"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v2/user_activities"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserActivitiesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserActivitiesOperation,
			ID:   "getUserActivities",
		}
	)
	params, err := decodeGetUserActivitiesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ActivitiesResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserActivitiesOperation,
			OperationSummary: "",
			OperationID:      "getUserActivities",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserActivitiesParams
			Response = *ActivitiesResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserActivitiesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserActivities(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserActivities(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserActivitiesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserActivitiesV1Request handles getUserActivitiesV1 operation.
//
// GET /api/user_activities
func (s *Server) handleGetUserActivitiesV1Request(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserActivitiesV1"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/user_activities"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserActivitiesV1Operation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserActivitiesV1Operation,
			ID:   "getUserActivitiesV1",
		}
	)
	params, err := decodeGetUserActivitiesV1Params(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ActivitiesResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserActivitiesV1Operation,
			OperationSummary: "",
			OperationID:      "getUserActivitiesV1",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "important",
					In:   "query",
				}: params.Important,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserActivitiesV1Params
			Response = *ActivitiesResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserActivitiesV1Params,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserActivitiesV1(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserActivitiesV1(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserActivitiesV1Response(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserByQrRequest handles getUserByQr operation.
//
// GET /v1/users/qr_codes/{qr}
func (s *Server) handleGetUserByQrRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserByQr"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/qr_codes/{qr}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserByQrOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserByQrOperation,
			ID:   "getUserByQr",
		}
	)
	params, err := decodeGetUserByQrParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserByQrOperation,
			OperationSummary: "",
			OperationID:      "getUserByQr",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "qr",
					In:   "path",
				}: params.Qr,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserByQrParams
			Response = *UserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserByQrParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserByQr(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserByQr(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserByQrResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserCustomDefinitionsRequest handles getUserCustomDefinitions operation.
//
// GET /v1/users/custom_definitions
func (s *Server) handleGetUserCustomDefinitionsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserCustomDefinitions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/custom_definitions"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserCustomDefinitionsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *UserCustomDefinitionsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserCustomDefinitionsOperation,
			OperationSummary: "",
			OperationID:      "getUserCustomDefinitions",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *UserCustomDefinitionsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserCustomDefinitions(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserCustomDefinitions(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserCustomDefinitionsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserFollowersRequest handles getUserFollowers operation.
//
// GET /v3/users/{id}/followers
func (s *Server) handleGetUserFollowersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserFollowers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v3/users/{id}/followers"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserFollowersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserFollowersOperation,
			ID:   "getUserFollowers",
		}
	)
	params, err := decodeGetUserFollowersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FollowUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserFollowersOperation,
			OperationSummary: "",
			OperationID:      "getUserFollowers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "followed_by_me",
					In:   "query",
				}: params.FollowedByMe,
				{
					Name: "user[nickname]",
					In:   "query",
				}: params.UserNickname,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserFollowersParams
			Response = *FollowUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserFollowersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserFollowers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserFollowers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserFollowersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserFollowingsRequest handles getUserFollowings operation.
//
// GET /v3/users/{id}/followings
func (s *Server) handleGetUserFollowingsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserFollowings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v3/users/{id}/followings"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserFollowingsOperation,
			ID:   "getUserFollowings",
		}
	)
	params, err := decodeGetUserFollowingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *FollowUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserFollowingsOperation,
			OperationSummary: "",
			OperationID:      "getUserFollowings",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "order_by",
					In:   "query",
				}: params.OrderBy,
				{
					Name: "user[nickname]",
					In:   "query",
				}: params.UserNickname,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserFollowingsParams
			Response = *FollowUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserFollowingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserFollowings(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserFollowings(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserFollowingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserGiftTransactionsRequest handles getUserGiftTransactions operation.
//
// GET /v1/users/{user_id}/gift_transactions
func (s *Server) handleGetUserGiftTransactionsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/{user_id}/gift_transactions"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserGiftTransactionsOperation,
			ID:   "getUserGiftTransactions",
		}
	)
	params, err := decodeGetUserGiftTransactionsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftTransactionsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserGiftTransactionsOperation,
			OperationSummary: "",
			OperationID:      "getUserGiftTransactions",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserGiftTransactionsParams
			Response = *GiftTransactionsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserGiftTransactionsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserGiftTransactions(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserGiftTransactions(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserGiftTransactionsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserGroupListRequest handles getUserGroupList operation.
//
// GET /v1/groups/user_group_list
func (s *Server) handleGetUserGroupListRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserGroupList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/user_group_list"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserGroupListOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserGroupListOperation,
			ID:   "getUserGroupList",
		}
	)
	params, err := decodeGetUserGroupListParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserGroupListOperation,
			OperationSummary: "",
			OperationID:      "getUserGroupList",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "user_id",
					In:   "query",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserGroupListParams
			Response = *GroupsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserGroupListParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserGroupList(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserGroupList(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserGroupListResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserInfoRequest handles getUserInfo operation.
//
// GET /v2/users/info/{id}
func (s *Server) handleGetUserInfoRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserInfo"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/info/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserInfoOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserInfoOperation,
			ID:   "getUserInfo",
		}
	)
	params, err := decodeGetUserInfoParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserInfoOperation,
			OperationSummary: "",
			OperationID:      "getUserInfo",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserInfoParams
			Response = *UserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserInfoParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserInfo(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserInfo(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserInfoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserInterestsRequest handles getUserInterests operation.
//
// GET /v1/users/interests
func (s *Server) handleGetUserInterestsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserInterests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/interests"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserInterestsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *UserInterestsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserInterestsOperation,
			OperationSummary: "",
			OperationID:      "getUserInterests",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *UserInterestsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserInterests(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserInterests(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserInterestsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserPresignedUrlRequest handles getUserPresignedUrl operation.
//
// GET /v1/users/presigned_url
func (s *Server) handleGetUserPresignedUrlRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserPresignedUrl"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/presigned_url"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserPresignedUrlOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserPresignedUrlOperation,
			ID:   "getUserPresignedUrl",
		}
	)
	params, err := decodeGetUserPresignedUrlParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PresignedUrlResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserPresignedUrlOperation,
			OperationSummary: "",
			OperationID:      "getUserPresignedUrl",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "video_file_name",
					In:   "query",
				}: params.VideoFileName,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserPresignedUrlParams
			Response = *PresignedUrlResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserPresignedUrlParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserPresignedUrl(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserPresignedUrl(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserPresignedUrlResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserReviewsRequest handles getUserReviews operation.
//
// GET /v2/users/reviews/{id}
func (s *Server) handleGetUserReviewsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserReviews"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/reviews/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserReviewsOperation,
			ID:   "getUserReviews",
		}
	)
	params, err := decodeGetUserReviewsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ReviewsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserReviewsOperation,
			OperationSummary: "",
			OperationID:      "getUserReviews",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserReviewsParams
			Response = *ReviewsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserReviewsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserReviews(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserReviews(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserReviewsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserTimelineRequest handles getUserTimeline operation.
//
// GET /v2/posts/user_timeline
func (s *Server) handleGetUserTimelineRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/user_timeline"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUserTimelineOperation,
			ID:   "getUserTimeline",
		}
	)
	params, err := decodeGetUserTimelineParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserTimelineOperation,
			OperationSummary: "",
			OperationID:      "getUserTimeline",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "query",
				}: params.UserID,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "post_type",
					In:   "query",
				}: params.PostType,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUserTimelineParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUserTimelineParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserTimeline(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserTimeline(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserTimelineResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUserTimestampRequest handles getUserTimestamp operation.
//
// GET /v2/users/timestamp
func (s *Server) handleGetUserTimestampRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserTimestamp"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/users/timestamp"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUserTimestampOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *UserTimestampResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUserTimestampOperation,
			OperationSummary: "",
			OperationID:      "getUserTimestamp",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *UserTimestampResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUserTimestamp(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUserTimestamp(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUserTimestampResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetUsersByIdsRequest handles getUsersByIds operation.
//
// GET /v1/users/list_id
func (s *Server) handleGetUsersByIdsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUsersByIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/list_id"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetUsersByIdsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetUsersByIdsOperation,
			ID:   "getUsersByIds",
		}
	)
	params, err := decodeGetUsersByIdsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetUsersByIdsOperation,
			OperationSummary: "",
			OperationID:      "getUsersByIds",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "X-Jwt",
					In:   "header",
				}: params.XJwt,
				{
					Name: "user_ids[]",
					In:   "query",
				}: params.UserIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetUsersByIdsParams
			Response = *UsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetUsersByIdsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetUsersByIds(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetUsersByIds(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetUsersByIdsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetWebSocketTokenRequest handles getWebSocketToken operation.
//
// GET /v1/users/ws_token
func (s *Server) handleGetWebSocketTokenRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getWebSocketToken"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/ws_token"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), GetWebSocketTokenOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *WebSocketTokenResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetWebSocketTokenOperation,
			OperationSummary: "",
			OperationID:      "getWebSocketToken",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *WebSocketTokenResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetWebSocketToken(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetWebSocketToken(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetWebSocketTokenResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleHideChatsRequest handles hideChats operation.
//
// POST /v1/hidden/chats
func (s *Server) handleHideChatsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("hideChats"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/hidden/chats"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), HideChatsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: HideChatsOperation,
			ID:   "hideChats",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeHideChatsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *HideChatsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    HideChatsOperation,
			OperationSummary: "",
			OperationID:      "hideChats",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *HideChatsReq
			Params   = struct{}
			Response = *HideChatsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.HideChats(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.HideChats(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeHideChatsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleHideUsersRequest handles hideUsers operation.
//
// POST /v1/hidden/users
func (s *Server) handleHideUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("hideUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/hidden/users"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), HideUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: HideUsersOperation,
			ID:   "hideUsers",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeHideUsersRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *HideUsersNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    HideUsersOperation,
			OperationSummary: "",
			OperationID:      "hideUsers",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *HideUsersReq
			Params   = struct{}
			Response = *HideUsersNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.HideUsers(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.HideUsers(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeHideUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInviteToCallRequest handles inviteToCall operation.
//
// POST /v2/calls/invite
func (s *Server) handleInviteToCallRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/calls/invite"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), InviteToCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: InviteToCallOperation,
			ID:   "inviteToCall",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeInviteToCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *InviteToCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    InviteToCallOperation,
			OperationSummary: "",
			OperationID:      "inviteToCall",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *InviteToCallReq
			Params   = struct{}
			Response = *InviteToCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InviteToCall(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.InviteToCall(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInviteToCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInviteToChatRoomRequest handles inviteToChatRoom operation.
//
// POST /v2/chat_rooms/{id}/invite
func (s *Server) handleInviteToChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}/invite"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), InviteToChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: InviteToChatRoomOperation,
			ID:   "inviteToChatRoom",
		}
	)
	params, err := decodeInviteToChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeInviteToChatRoomRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *InviteToChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    InviteToChatRoomOperation,
			OperationSummary: "",
			OperationID:      "inviteToChatRoom",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *InviteToChatRoomReq
			Params   = InviteToChatRoomParams
			Response = *InviteToChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInviteToChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InviteToChatRoom(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.InviteToChatRoom(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInviteToChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInviteToConferenceCallRequest handles inviteToConferenceCall operation.
//
// POST /v1/calls/conference_calls/{call_id}/invite
func (s *Server) handleInviteToConferenceCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/conference_calls/{call_id}/invite"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), InviteToConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: InviteToConferenceCallOperation,
			ID:   "inviteToConferenceCall",
		}
	)
	params, err := decodeInviteToConferenceCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeInviteToConferenceCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *InviteToConferenceCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    InviteToConferenceCallOperation,
			OperationSummary: "",
			OperationID:      "inviteToConferenceCall",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
			},
			Raw: r,
		}

		type (
			Request  = *InviteToConferenceCallReq
			Params   = InviteToConferenceCallParams
			Response = *InviteToConferenceCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInviteToConferenceCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InviteToConferenceCall(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.InviteToConferenceCall(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInviteToConferenceCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInviteToGroupRequest handles inviteToGroup operation.
//
// POST /v1/groups/{id}/invite
func (s *Server) handleInviteToGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/invite"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), InviteToGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: InviteToGroupOperation,
			ID:   "inviteToGroup",
		}
	)
	params, err := decodeInviteToGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeInviteToGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *InviteToGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    InviteToGroupOperation,
			OperationSummary: "",
			OperationID:      "inviteToGroup",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *InviteToGroupReq
			Params   = InviteToGroupParams
			Response = *InviteToGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInviteToGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InviteToGroup(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.InviteToGroup(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInviteToGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleJoinGroupRequest handles joinGroup operation.
//
// POST /v1/groups/{id}/join
func (s *Server) handleJoinGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("joinGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/join"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), JoinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: JoinGroupOperation,
			ID:   "joinGroup",
		}
	)
	params, err := decodeJoinGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *JoinGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    JoinGroupOperation,
			OperationSummary: "",
			OperationID:      "joinGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = JoinGroupParams
			Response = *JoinGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackJoinGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.JoinGroup(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.JoinGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeJoinGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleKickFromChatRoomRequest handles kickFromChatRoom operation.
//
// POST /v2/chat_rooms/{id}/kick
func (s *Server) handleKickFromChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("kickFromChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}/kick"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), KickFromChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: KickFromChatRoomOperation,
			ID:   "kickFromChatRoom",
		}
	)
	params, err := decodeKickFromChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeKickFromChatRoomRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *KickFromChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    KickFromChatRoomOperation,
			OperationSummary: "",
			OperationID:      "kickFromChatRoom",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *KickFromChatRoomReq
			Params   = KickFromChatRoomParams
			Response = *KickFromChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackKickFromChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.KickFromChatRoom(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.KickFromChatRoom(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeKickFromChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleKickFromConferenceCallRequest handles kickFromConferenceCall operation.
//
// POST /v3/calls/conference_calls/{call_id}/kick
func (s *Server) handleKickFromConferenceCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("kickFromConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/calls/conference_calls/{call_id}/kick"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), KickFromConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: KickFromConferenceCallOperation,
			ID:   "kickFromConferenceCall",
		}
	)
	params, err := decodeKickFromConferenceCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeKickFromConferenceCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CallActionSignatureResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    KickFromConferenceCallOperation,
			OperationSummary: "",
			OperationID:      "kickFromConferenceCall",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
			},
			Raw: r,
		}

		type (
			Request  = *KickFromConferenceCallReq
			Params   = KickFromConferenceCallParams
			Response = *CallActionSignatureResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackKickFromConferenceCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.KickFromConferenceCall(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.KickFromConferenceCall(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeKickFromConferenceCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLeaveAgoraChannelRequest handles leaveAgoraChannel operation.
//
// POST /v1/calls/leave_agora_channel
func (s *Server) handleLeaveAgoraChannelRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveAgoraChannel"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/leave_agora_channel"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LeaveAgoraChannelOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LeaveAgoraChannelOperation,
			ID:   "leaveAgoraChannel",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeLeaveAgoraChannelRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LeaveAgoraChannelNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LeaveAgoraChannelOperation,
			OperationSummary: "",
			OperationID:      "leaveAgoraChannel",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LeaveAgoraChannelReq
			Params   = struct{}
			Response = *LeaveAgoraChannelNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.LeaveAgoraChannel(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.LeaveAgoraChannel(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLeaveAgoraChannelResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLeaveConferenceCallRequest handles leaveConferenceCall operation.
//
// POST /v1/calls/leave_conference_call
func (s *Server) handleLeaveConferenceCallRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/calls/leave_conference_call"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LeaveConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LeaveConferenceCallOperation,
			ID:   "leaveConferenceCall",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeLeaveConferenceCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LeaveConferenceCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LeaveConferenceCallOperation,
			OperationSummary: "",
			OperationID:      "leaveConferenceCall",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LeaveConferenceCallReq
			Params   = struct{}
			Response = *LeaveConferenceCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.LeaveConferenceCall(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.LeaveConferenceCall(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLeaveConferenceCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLeaveGroupRequest handles leaveGroup operation.
//
// DELETE /v1/groups/{id}/leave
func (s *Server) handleLeaveGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveGroup"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/leave"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LeaveGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LeaveGroupOperation,
			ID:   "leaveGroup",
		}
	)
	params, err := decodeLeaveGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *LeaveGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LeaveGroupOperation,
			OperationSummary: "",
			OperationID:      "leaveGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = LeaveGroupParams
			Response = *LeaveGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackLeaveGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.LeaveGroup(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.LeaveGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLeaveGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLikePostsRequest handles likePosts operation.
//
// POST /v2/posts/like
func (s *Server) handleLikePostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("likePosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/posts/like"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LikePostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LikePostsOperation,
			ID:   "likePosts",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeLikePostsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LikePostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LikePostsOperation,
			OperationSummary: "",
			OperationID:      "likePosts",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LikePostsReq
			Params   = struct{}
			Response = *LikePostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.LikePosts(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.LikePosts(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLikePostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGameAppsRequest handles listGameApps operation.
//
// GET /v1/games/apps
func (s *Server) handleListGameAppsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGameApps"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/games/apps"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGameAppsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGameAppsOperation,
			ID:   "listGameApps",
		}
	)
	params, err := decodeListGameAppsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GamesResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGameAppsOperation,
			OperationSummary: "",
			OperationID:      "listGameApps",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from_id",
					In:   "query",
				}: params.FromID,
				{
					Name: "ids[]",
					In:   "query",
				}: params.Ids,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGameAppsParams
			Response = *GamesResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGameAppsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGameApps(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGameApps(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGameAppsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGameWalkthroughsRequest handles listGameWalkthroughs operation.
//
// GET /v1/games/apps/{app_id}/walkthroughs
func (s *Server) handleListGameWalkthroughsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGameWalkthroughs"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/games/apps/{app_id}/walkthroughs"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGameWalkthroughsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGameWalkthroughsOperation,
			ID:   "listGameWalkthroughs",
		}
	)
	params, err := decodeListGameWalkthroughsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response []Walkthrough
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGameWalkthroughsOperation,
			OperationSummary: "",
			OperationID:      "listGameWalkthroughs",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "app_id",
					In:   "path",
				}: params.AppID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGameWalkthroughsParams
			Response = []Walkthrough
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGameWalkthroughsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGameWalkthroughs(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGameWalkthroughs(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGameWalkthroughsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGenresRequest handles listGenres operation.
//
// GET /v1/genres
func (s *Server) handleListGenresRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGenres"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/genres"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGenresOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGenresOperation,
			ID:   "listGenres",
		}
	)
	params, err := decodeListGenresParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GenresResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGenresOperation,
			OperationSummary: "",
			OperationID:      "listGenres",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGenresParams
			Response = *GenresResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGenresParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGenres(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGenres(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGenresResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGiftsRequest handles listGifts operation.
//
// GET /v2/gifts
func (s *Server) handleListGiftsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGifts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/gifts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGiftsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGiftsOperation,
			ID:   "listGifts",
		}
	)
	params, err := decodeListGiftsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGiftsOperation,
			OperationSummary: "",
			OperationID:      "listGifts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "currency",
					In:   "query",
				}: params.Currency,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from",
					In:   "query",
				}: params.From,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGiftsParams
			Response = *GiftsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGiftsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGifts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGifts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGiftsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGroupCategoriesRequest handles listGroupCategories operation.
//
// GET /v1/groups/categories
func (s *Server) handleListGroupCategoriesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGroupCategories"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/groups/categories"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGroupCategoriesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGroupCategoriesOperation,
			ID:   "listGroupCategories",
		}
	)
	params, err := decodeListGroupCategoriesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupCategoriesResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGroupCategoriesOperation,
			OperationSummary: "",
			OperationID:      "listGroupCategories",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGroupCategoriesParams
			Response = *GroupCategoriesResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGroupCategoriesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGroupCategories(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGroupCategories(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGroupCategoriesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListGroupsRequest handles listGroups operation.
//
// GET /v2/groups
func (s *Server) handleListGroupsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/groups"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListGroupsOperation,
			ID:   "listGroups",
		}
	)
	params, err := decodeListGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListGroupsOperation,
			OperationSummary: "",
			OperationID:      "listGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_category_id",
					In:   "query",
				}: params.GroupCategoryID,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "sub_category_id",
					In:   "query",
				}: params.SubCategoryID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListGroupsParams
			Response = *GroupsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListGroups(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListHiddenChatsRequest handles listHiddenChats operation.
//
// GET /v1/hidden/chats
func (s *Server) handleListHiddenChatsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listHiddenChats"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/hidden/chats"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListHiddenChatsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListHiddenChatsOperation,
			ID:   "listHiddenChats",
		}
	)
	params, err := decodeListHiddenChatsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ChatRoomsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListHiddenChatsOperation,
			OperationSummary: "",
			OperationID:      "listHiddenChats",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListHiddenChatsParams
			Response = *ChatRoomsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListHiddenChatsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListHiddenChats(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListHiddenChats(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListHiddenChatsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListHiddenUsersRequest handles listHiddenUsers operation.
//
// GET /v1/hidden/users
func (s *Server) handleListHiddenUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listHiddenUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/hidden/users"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListHiddenUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListHiddenUsersOperation,
			ID:   "listHiddenUsers",
		}
	)
	params, err := decodeListHiddenUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *HiddenResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListHiddenUsersOperation,
			OperationSummary: "",
			OperationID:      "listHiddenUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListHiddenUsersParams
			Response = *HiddenResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListHiddenUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListHiddenUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListHiddenUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListHiddenUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListMuteKeywordsRequest handles listMuteKeywords operation.
//
// GET /v1/hidden/words
func (s *Server) handleListMuteKeywordsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMuteKeywords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/hidden/words"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListMuteKeywordsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *MuteKeywordResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListMuteKeywordsOperation,
			OperationSummary: "",
			OperationID:      "listMuteKeywords",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *MuteKeywordResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListMuteKeywords(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListMuteKeywords(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListMuteKeywordsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListMutedGroupUsersRequest handles listMutedGroupUsers operation.
//
// GET /v1/group_mute/{id}/muted_users
func (s *Server) handleListMutedGroupUsersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMutedGroupUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/group_mute/{id}/muted_users"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListMutedGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListMutedGroupUsersOperation,
			ID:   "listMutedGroupUsers",
		}
	)
	params, err := decodeListMutedGroupUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupMuteUsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListMutedGroupUsersOperation,
			OperationSummary: "",
			OperationID:      "listMutedGroupUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "cursor",
					In:   "query",
				}: params.Cursor,
				{
					Name: "size",
					In:   "query",
				}: params.Size,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListMutedGroupUsersParams
			Response = *GroupMuteUsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListMutedGroupUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListMutedGroupUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListMutedGroupUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListMutedGroupUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListMyGroupsRequest handles listMyGroups operation.
//
// GET /v2/groups/mine
func (s *Server) handleListMyGroupsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMyGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/groups/mine"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListMyGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListMyGroupsOperation,
			ID:   "listMyGroups",
		}
	)
	params, err := decodeListMyGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListMyGroupsOperation,
			OperationSummary: "",
			OperationID:      "listMyGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListMyGroupsParams
			Response = *GroupsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListMyGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListMyGroups(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListMyGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListMyGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListReceivedGiftsRequest handles listReceivedGifts operation.
//
// GET /v2/received_gifts
func (s *Server) handleListReceivedGiftsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listReceivedGifts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/received_gifts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListReceivedGiftsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *GiftReceivedResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListReceivedGiftsOperation,
			OperationSummary: "",
			OperationID:      "listReceivedGifts",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *GiftReceivedResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListReceivedGifts(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListReceivedGifts(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListReceivedGiftsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListReceivedGiftsV1Request handles listReceivedGiftsV1 operation.
//
// GET /v1/received_gifts
func (s *Server) handleListReceivedGiftsV1Request(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listReceivedGiftsV1"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/received_gifts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListReceivedGiftsV1Operation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListReceivedGiftsV1Operation,
			ID:   "listReceivedGiftsV1",
		}
	)
	params, err := decodeListReceivedGiftsV1Params(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GiftReceivedResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListReceivedGiftsV1Operation,
			OperationSummary: "",
			OperationID:      "listReceivedGiftsV1",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListReceivedGiftsV1Params
			Response = *GiftReceivedResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListReceivedGiftsV1Params,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListReceivedGiftsV1(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListReceivedGiftsV1(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListReceivedGiftsV1Response(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListStickerPacksRequest handles listStickerPacks operation.
//
// GET /v2/sticker_packs
func (s *Server) handleListStickerPacksRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listStickerPacks"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/sticker_packs"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListStickerPacksOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *StickerPacksResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListStickerPacksOperation,
			OperationSummary: "",
			OperationID:      "listStickerPacks",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *StickerPacksResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListStickerPacks(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListStickerPacks(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListStickerPacksResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListThreadsRequest handles listThreads operation.
//
// GET /v1/threads
func (s *Server) handleListThreadsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listThreads"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/threads"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ListThreadsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ListThreadsOperation,
			ID:   "listThreads",
		}
	)
	params, err := decodeListThreadsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *GroupThreadListResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ListThreadsOperation,
			OperationSummary: "",
			OperationID:      "listThreads",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "join_status",
					In:   "query",
				}: params.JoinStatus,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListThreadsParams
			Response = *GroupThreadListResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListThreadsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListThreads(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListThreads(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListThreadsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLoginWithEmailRequest handles loginWithEmail operation.
//
// POST /v3/users/login_with_email
func (s *Server) handleLoginWithEmailRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("loginWithEmail"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/users/login_with_email"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LoginWithEmailOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LoginWithEmailOperation,
			ID:   "loginWithEmail",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeLoginWithEmailRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LoginUserResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LoginWithEmailOperation,
			OperationSummary: "",
			OperationID:      "loginWithEmail",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LoginEmailUserRequest
			Params   = struct{}
			Response = *LoginUserResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.LoginWithEmail(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.LoginWithEmail(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLoginWithEmailResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleLogoutRequest handles logout operation.
//
// POST /v1/users/logout
func (s *Server) handleLogoutRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("logout"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/logout"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), LogoutOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: LogoutOperation,
			ID:   "logout",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeLogoutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LogoutNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    LogoutOperation,
			OperationSummary: "",
			OperationID:      "logout",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LogoutReq
			Params   = struct{}
			Response = *LogoutNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.Logout(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.Logout(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLogoutResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMovePostToSpecificThreadRequest handles movePostToSpecificThread operation.
//
// PUT /v3/posts/{id}/move_to_thread/{thread_id}
func (s *Server) handleMovePostToSpecificThreadRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("movePostToSpecificThread"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v3/posts/{id}/move_to_thread/{thread_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), MovePostToSpecificThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: MovePostToSpecificThreadOperation,
			ID:   "movePostToSpecificThread",
		}
	)
	params, err := decodeMovePostToSpecificThreadParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ThreadInfo
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    MovePostToSpecificThreadOperation,
			OperationSummary: "",
			OperationID:      "movePostToSpecificThread",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "thread_id",
					In:   "path",
				}: params.ThreadID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = MovePostToSpecificThreadParams
			Response = *ThreadInfo
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackMovePostToSpecificThreadParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MovePostToSpecificThread(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.MovePostToSpecificThread(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMovePostToSpecificThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMovePostToThreadRequest handles movePostToThread operation.
//
// POST /v3/posts/{id}/move_to_thread
func (s *Server) handleMovePostToThreadRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("movePostToThread"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/posts/{id}/move_to_thread"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), MovePostToThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: MovePostToThreadOperation,
			ID:   "movePostToThread",
		}
	)
	params, err := decodeMovePostToThreadParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeMovePostToThreadRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ThreadInfo
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    MovePostToThreadOperation,
			OperationSummary: "",
			OperationID:      "movePostToThread",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *MovePostToThreadReq
			Params   = MovePostToThreadParams
			Response = *ThreadInfo
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackMovePostToThreadParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MovePostToThread(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.MovePostToThread(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMovePostToThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMuteGroupUserRequest handles muteGroupUser operation.
//
// POST /v1/group_mute/{id}/mute/{user_id}
func (s *Server) handleMuteGroupUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("muteGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/group_mute/{id}/mute/{user_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), MuteGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: MuteGroupUserOperation,
			ID:   "muteGroupUser",
		}
	)
	params, err := decodeMuteGroupUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *MuteGroupUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    MuteGroupUserOperation,
			OperationSummary: "",
			OperationID:      "muteGroupUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = MuteGroupUserParams
			Response = *MuteGroupUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackMuteGroupUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.MuteGroupUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.MuteGroupUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMuteGroupUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleOauthTokenRequest handles oauthToken operation.
//
// POST /api/v1/oauth/token
func (s *Server) handleOauthTokenRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("oauthToken"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/oauth/token"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), OauthTokenOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: OauthTokenOperation,
			ID:   "oauthToken",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeOauthTokenRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TokenResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    OauthTokenOperation,
			OperationSummary: "",
			OperationID:      "oauthToken",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *OauthTokenReq
			Params   = struct{}
			Response = *TokenResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.OauthToken(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.OauthToken(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOauthTokenResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinChatRoomRequest handles pinChatRoom operation.
//
// POST /v1/chat_rooms/{id}/pinned
func (s *Server) handlePinChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/{id}/pinned"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinChatRoomOperation,
			ID:   "pinChatRoom",
		}
	)
	params, err := decodePinChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PinChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinChatRoomOperation,
			OperationSummary: "",
			OperationID:      "pinChatRoom",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = PinChatRoomParams
			Response = *PinChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPinChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinChatRoom(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.PinChatRoom(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinGroupRequest handles pinGroup operation.
//
// POST /v1/pinned/groups
func (s *Server) handlePinGroupRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/pinned/groups"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinGroupOperation,
			ID:   "pinGroup",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodePinGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PinGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinGroupOperation,
			OperationSummary: "",
			OperationID:      "pinGroup",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *PinGroupReq
			Params   = struct{}
			Response = *PinGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinGroup(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.PinGroup(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinGroupHighlightPostRequest handles pinGroupHighlightPost operation.
//
// PUT /v1/groups/{group_id}/highlights/{post_id}
func (s *Server) handlePinGroupHighlightPostRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroupHighlightPost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/highlights/{post_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinGroupHighlightPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinGroupHighlightPostOperation,
			ID:   "pinGroupHighlightPost",
		}
	)
	params, err := decodePinGroupHighlightPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PinGroupHighlightPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinGroupHighlightPostOperation,
			OperationSummary: "",
			OperationID:      "pinGroupHighlightPost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "post_id",
					In:   "path",
				}: params.PostID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = PinGroupHighlightPostParams
			Response = *PinGroupHighlightPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPinGroupHighlightPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinGroupHighlightPost(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.PinGroupHighlightPost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinGroupHighlightPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinGroupPostRequest handles pinGroupPost operation.
//
// PUT /v2/posts/group_pinned_post
func (s *Server) handlePinGroupPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroupPost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v2/posts/group_pinned_post"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinGroupPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinGroupPostOperation,
			ID:   "pinGroupPost",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodePinGroupPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PinGroupPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinGroupPostOperation,
			OperationSummary: "",
			OperationID:      "pinGroupPost",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *PinGroupPostReq
			Params   = struct{}
			Response = *PinGroupPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinGroupPost(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.PinGroupPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinGroupPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinPostRequest handles pinPost operation.
//
// POST /v1/pinned/posts
func (s *Server) handlePinPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/pinned/posts"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinPostOperation,
			ID:   "pinPost",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodePinPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PinPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinPostOperation,
			OperationSummary: "",
			OperationID:      "pinPost",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *PinPostReq
			Params   = struct{}
			Response = *PinPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinPost(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.PinPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePinReviewRequest handles pinReview operation.
//
// POST /v1/pinned/reviews
func (s *Server) handlePinReviewRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/pinned/reviews"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PinReviewOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: PinReviewOperation,
			ID:   "pinReview",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodePinReviewRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PinReviewNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PinReviewOperation,
			OperationSummary: "",
			OperationID:      "pinReview",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *PinReviewReq
			Params   = struct{}
			Response = *PinReviewNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PinReview(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.PinReview(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePinReviewResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePingAliveRequest handles pingAlive operation.
//
// POST /v1/users/alive
func (s *Server) handlePingAliveRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pingAlive"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/alive"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), PingAliveOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *PingAliveNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    PingAliveOperation,
			OperationSummary: "",
			OperationID:      "pingAlive",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *PingAliveNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.PingAlive(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.PingAlive(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePingAliveResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReadChatAttachmentsRequest handles readChatAttachments operation.
//
// POST /v1/chat_rooms/{id}/attachments_read
func (s *Server) handleReadChatAttachmentsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatAttachments"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/{id}/attachments_read"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReadChatAttachmentsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReadChatAttachmentsOperation,
			ID:   "readChatAttachments",
		}
	)
	params, err := decodeReadChatAttachmentsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReadChatAttachmentsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReadChatAttachmentsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReadChatAttachmentsOperation,
			OperationSummary: "",
			OperationID:      "readChatAttachments",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *ReadAttachmentRequest
			Params   = ReadChatAttachmentsParams
			Response = *ReadChatAttachmentsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReadChatAttachmentsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReadChatAttachments(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReadChatAttachments(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReadChatAttachmentsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReadChatMessageRequest handles readChatMessage operation.
//
// POST /v2/chat_rooms/{id}/messages/{message_id}/read
func (s *Server) handleReadChatMessageRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatMessage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}/messages/{message_id}/read"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReadChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReadChatMessageOperation,
			ID:   "readChatMessage",
		}
	)
	params, err := decodeReadChatMessageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ReadChatMessageNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReadChatMessageOperation,
			OperationSummary: "",
			OperationID:      "readChatMessage",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "message_id",
					In:   "path",
				}: params.MessageID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ReadChatMessageParams
			Response = *ReadChatMessageNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReadChatMessageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReadChatMessage(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReadChatMessage(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReadChatMessageResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReadChatVideosRequest handles readChatVideos operation.
//
// POST /v1/chat_rooms/{id}/videos_read
func (s *Server) handleReadChatVideosRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatVideos"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/{id}/videos_read"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReadChatVideosOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReadChatVideosOperation,
			ID:   "readChatVideos",
		}
	)
	params, err := decodeReadChatVideosParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReadChatVideosRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReadChatVideosNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReadChatVideosOperation,
			OperationSummary: "",
			OperationID:      "readChatVideos",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *ReadChatVideosReq
			Params   = ReadChatVideosParams
			Response = *ReadChatVideosNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReadChatVideosParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReadChatVideos(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReadChatVideos(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReadChatVideosResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveChatRoomBackgroundRequest handles removeChatRoomBackground operation.
//
// DELETE /v2/chat_rooms/{id}/background
func (s *Server) handleRemoveChatRoomBackgroundRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeChatRoomBackground"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v2/chat_rooms/{id}/background"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveChatRoomBackgroundOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveChatRoomBackgroundOperation,
			ID:   "removeChatRoomBackground",
		}
	)
	params, err := decodeRemoveChatRoomBackgroundParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveChatRoomBackgroundNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveChatRoomBackgroundOperation,
			OperationSummary: "",
			OperationID:      "removeChatRoomBackground",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveChatRoomBackgroundParams
			Response = *RemoveChatRoomBackgroundNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveChatRoomBackgroundParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveChatRoomBackground(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveChatRoomBackground(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveChatRoomBackgroundResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveCoverImageRequest handles removeCoverImage operation.
//
// POST /v2/users/remove_cover_image
func (s *Server) handleRemoveCoverImageRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeCoverImage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/remove_cover_image"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveCoverImageOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *RemoveCoverImageNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveCoverImageOperation,
			OperationSummary: "",
			OperationID:      "removeCoverImage",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *RemoveCoverImageNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveCoverImage(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveCoverImage(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveCoverImageResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveGroupCoverRequest handles removeGroupCover operation.
//
// DELETE /v3/groups/{id}/cover
func (s *Server) handleRemoveGroupCoverRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupCover"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v3/groups/{id}/cover"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveGroupCoverOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveGroupCoverOperation,
			ID:   "removeGroupCover",
		}
	)
	params, err := decodeRemoveGroupCoverParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveGroupCoverNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveGroupCoverOperation,
			OperationSummary: "",
			OperationID:      "removeGroupCover",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveGroupCoverParams
			Response = *RemoveGroupCoverNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveGroupCoverParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveGroupCover(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveGroupCover(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveGroupCoverResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveGroupDeputiesRequest handles removeGroupDeputies operation.
//
// DELETE /v1/groups/{id}/deputize
func (s *Server) handleRemoveGroupDeputiesRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupDeputies"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/deputize"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveGroupDeputiesOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveGroupDeputiesOperation,
			ID:   "removeGroupDeputies",
		}
	)
	params, err := decodeRemoveGroupDeputiesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveGroupDeputiesNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveGroupDeputiesOperation,
			OperationSummary: "",
			OperationID:      "removeGroupDeputies",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveGroupDeputiesParams
			Response = *RemoveGroupDeputiesNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveGroupDeputiesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveGroupDeputies(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveGroupDeputies(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveGroupDeputiesResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveGroupIconRequest handles removeGroupIcon operation.
//
// DELETE /v3/groups/{id}/icon
func (s *Server) handleRemoveGroupIconRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupIcon"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v3/groups/{id}/icon"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveGroupIconOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveGroupIconOperation,
			ID:   "removeGroupIcon",
		}
	)
	params, err := decodeRemoveGroupIconParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveGroupIconNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveGroupIconOperation,
			OperationSummary: "",
			OperationID:      "removeGroupIcon",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveGroupIconParams
			Response = *RemoveGroupIconNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveGroupIconParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveGroupIcon(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveGroupIcon(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveGroupIconResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveProfilePhotoRequest handles removeProfilePhoto operation.
//
// POST /v2/users/remove_profile_photo
func (s *Server) handleRemoveProfilePhotoRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeProfilePhoto"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/remove_profile_photo"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveProfilePhotoOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *RemoveProfilePhotoNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveProfilePhotoOperation,
			OperationSummary: "",
			OperationID:      "removeProfilePhoto",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *RemoveProfilePhotoNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveProfilePhoto(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveProfilePhoto(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveProfilePhotoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveRelatedGroupsRequest handles removeRelatedGroups operation.
//
// DELETE /v1/groups/{id}/related
func (s *Server) handleRemoveRelatedGroupsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/related"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveRelatedGroupsOperation,
			ID:   "removeRelatedGroups",
		}
	)
	params, err := decodeRemoveRelatedGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveRelatedGroupsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveRelatedGroupsOperation,
			OperationSummary: "",
			OperationID:      "removeRelatedGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "related_group_id[]",
					In:   "query",
				}: params.RelatedGroupID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveRelatedGroupsParams
			Response = *RemoveRelatedGroupsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveRelatedGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveRelatedGroups(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveRelatedGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveRelatedGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRemoveThreadMemberRequest handles removeThreadMember operation.
//
// DELETE /v1/threads/{thread_id}/members/{id}
func (s *Server) handleRemoveThreadMemberRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeThreadMember"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/threads/{thread_id}/members/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RemoveThreadMemberOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RemoveThreadMemberOperation,
			ID:   "removeThreadMember",
		}
	)
	params, err := decodeRemoveThreadMemberParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RemoveThreadMemberNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RemoveThreadMemberOperation,
			OperationSummary: "",
			OperationID:      "removeThreadMember",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "thread_id",
					In:   "path",
				}: params.ThreadID,
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RemoveThreadMemberParams
			Response = *RemoveThreadMemberNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRemoveThreadMemberParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RemoveThreadMember(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RemoveThreadMember(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRemoveThreadMemberResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReplyToReviewRequest handles replyToReview operation.
//
// POST /v2/users/reviews/{id}
func (s *Server) handleReplyToReviewRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("replyToReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/reviews/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReplyToReviewOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReplyToReviewOperation,
			ID:   "replyToReview",
		}
	)
	params, err := decodeReplyToReviewParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReplyToReviewRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReplyToReviewNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReplyToReviewOperation,
			OperationSummary: "",
			OperationID:      "replyToReview",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *ReplyToReviewReq
			Params   = ReplyToReviewParams
			Response = *ReplyToReviewNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReplyToReviewParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReplyToReview(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReplyToReview(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReplyToReviewResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReportChatRoomRequest handles reportChatRoom operation.
//
// POST /v3/chat_rooms/{chat_room_id}/report
func (s *Server) handleReportChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/chat_rooms/{chat_room_id}/report"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReportChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReportChatRoomOperation,
			ID:   "reportChatRoom",
		}
	)
	params, err := decodeReportChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReportChatRoomRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReportChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReportChatRoomOperation,
			OperationSummary: "",
			OperationID:      "reportChatRoom",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "chat_room_id",
					In:   "path",
				}: params.ChatRoomID,
			},
			Raw: r,
		}

		type (
			Request  = *ReportChatRoomReq
			Params   = ReportChatRoomParams
			Response = *ReportChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReportChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReportChatRoom(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReportChatRoom(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReportChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReportGroupRequest handles reportGroup operation.
//
// POST /v3/groups/{group_id}/report
func (s *Server) handleReportGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/groups/{group_id}/report"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReportGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReportGroupOperation,
			ID:   "reportGroup",
		}
	)
	params, err := decodeReportGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReportGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReportGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReportGroupOperation,
			OperationSummary: "",
			OperationID:      "reportGroup",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
			},
			Raw: r,
		}

		type (
			Request  = *ReportGroupReq
			Params   = ReportGroupParams
			Response = *ReportGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReportGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReportGroup(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReportGroup(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReportGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReportPostRequest handles reportPost operation.
//
// POST /v3/posts/{post_id}/report
func (s *Server) handleReportPostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/posts/{post_id}/report"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReportPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReportPostOperation,
			ID:   "reportPost",
		}
	)
	params, err := decodeReportPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReportPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReportPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReportPostOperation,
			OperationSummary: "",
			OperationID:      "reportPost",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "post_id",
					In:   "path",
				}: params.PostID,
			},
			Raw: r,
		}

		type (
			Request  = *ReportPostReq
			Params   = ReportPostParams
			Response = *ReportPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReportPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReportPost(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReportPost(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReportPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReportUserRequest handles reportUser operation.
//
// POST /v3/users/{user_id}/report
func (s *Server) handleReportUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/users/{user_id}/report"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ReportUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ReportUserOperation,
			ID:   "reportUser",
		}
	)
	params, err := decodeReportUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeReportUserRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ReportUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ReportUserOperation,
			OperationSummary: "",
			OperationID:      "reportUser",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = *ReportUserReq
			Params   = ReportUserParams
			Response = *ReportUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReportUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ReportUser(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.ReportUser(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReportUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRepostRequest handles repost operation.
//
// POST /v3/posts/repost
func (s *Server) handleRepostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("repost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/posts/repost"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RepostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RepostOperation,
			ID:   "repost",
		}
	)
	params, err := decodeRepostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeRepostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CreatePostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RepostOperation,
			OperationSummary: "",
			OperationID:      "repost",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "X-Jwt",
					In:   "header",
				}: params.XJwt,
			},
			Raw: r,
		}

		type (
			Request  = *RepostReq
			Params   = RepostParams
			Response = *CreatePostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRepostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Repost(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Repost(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRepostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRequestEmailVerificationRequest handles requestEmailVerification operation.
//
// POST /v1/email_verification_urls
func (s *Server) handleRequestEmailVerificationRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestEmailVerification"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/email_verification_urls"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RequestEmailVerificationOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RequestEmailVerificationOperation,
			ID:   "requestEmailVerification",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeRequestEmailVerificationRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *CommonUrlResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RequestEmailVerificationOperation,
			OperationSummary: "",
			OperationID:      "requestEmailVerification",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *RequestEmailVerificationReq
			Params   = struct{}
			Response = *CommonUrlResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.RequestEmailVerification(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.RequestEmailVerification(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRequestEmailVerificationResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRequestFollowRequest handles requestFollow operation.
//
// POST /v2/users/{target_id}/follow_request
func (s *Server) handleRequestFollowRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestFollow"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/{target_id}/follow_request"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RequestFollowOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RequestFollowOperation,
			ID:   "requestFollow",
		}
	)
	params, err := decodeRequestFollowParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeRequestFollowRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *RequestFollowNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RequestFollowOperation,
			OperationSummary: "",
			OperationID:      "requestFollow",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "target_id",
					In:   "path",
				}: params.TargetID,
			},
			Raw: r,
		}

		type (
			Request  = *RequestFollowReq
			Params   = RequestFollowParams
			Response = *RequestFollowNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRequestFollowParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RequestFollow(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.RequestFollow(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRequestFollowResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRequestGroupWalkthroughRequest handles requestGroupWalkthrough operation.
//
// POST /v1/groups/{id}/request_walkthrough
func (s *Server) handleRequestGroupWalkthroughRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestGroupWalkthrough"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/request_walkthrough"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), RequestGroupWalkthroughOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RequestGroupWalkthroughOperation,
			ID:   "requestGroupWalkthrough",
		}
	)
	params, err := decodeRequestGroupWalkthroughParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *RequestGroupWalkthroughNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RequestGroupWalkthroughOperation,
			OperationSummary: "",
			OperationID:      "requestGroupWalkthrough",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RequestGroupWalkthroughParams
			Response = *RequestGroupWalkthroughNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRequestGroupWalkthroughParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.RequestGroupWalkthrough(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.RequestGroupWalkthrough(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRequestGroupWalkthroughResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleResendConfirmEmailRequest handles resendConfirmEmail operation.
//
// POST /v2/users/resend_confirm_email
func (s *Server) handleResendConfirmEmailRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resendConfirmEmail"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/resend_confirm_email"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ResendConfirmEmailOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *ResendConfirmEmailNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ResendConfirmEmailOperation,
			OperationSummary: "",
			OperationID:      "resendConfirmEmail",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *ResendConfirmEmailNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ResendConfirmEmail(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.ResendConfirmEmail(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeResendConfirmEmailResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleResetCountersRequest handles resetCounters operation.
//
// POST /v1/users/reset_counters
func (s *Server) handleResetCountersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resetCounters"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/reset_counters"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ResetCountersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ResetCountersOperation,
			ID:   "resetCounters",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeResetCountersRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ResetCountersNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ResetCountersOperation,
			OperationSummary: "",
			OperationID:      "resetCounters",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ResetCountersReq
			Params   = struct{}
			Response = *ResetCountersNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ResetCounters(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.ResetCounters(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeResetCountersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleResetPasswordRequest handles resetPassword operation.
//
// PUT /v1/users/reset_password
func (s *Server) handleResetPasswordRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resetPassword"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/reset_password"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ResetPasswordOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ResetPasswordOperation,
			ID:   "resetPassword",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeResetPasswordRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ResetPasswordNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ResetPasswordOperation,
			OperationSummary: "",
			OperationID:      "resetPassword",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ResetPasswordReq
			Params   = struct{}
			Response = *ResetPasswordNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ResetPassword(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.ResetPassword(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeResetPasswordResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSearchGroupPostsRequest handles searchGroupPosts operation.
//
// GET /v2/groups/{id}/posts/search
func (s *Server) handleSearchGroupPostsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchGroupPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/groups/{id}/posts/search"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SearchGroupPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SearchGroupPostsOperation,
			ID:   "searchGroupPosts",
		}
	)
	params, err := decodeSearchGroupPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SearchGroupPostsOperation,
			OperationSummary: "",
			OperationID:      "searchGroupPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
				{
					Name: "only_thread_posts",
					In:   "query",
				}: params.OnlyThreadPosts,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchGroupPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchGroupPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SearchGroupPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SearchGroupPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSearchGroupPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSearchPostsRequest handles searchPosts operation.
//
// GET /v2/posts/search
func (s *Server) handleSearchPostsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/posts/search"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SearchPostsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SearchPostsOperation,
			ID:   "searchPosts",
		}
	)
	params, err := decodeSearchPostsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *PostsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SearchPostsOperation,
			OperationSummary: "",
			OperationID:      "searchPosts",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
				{
					Name: "post_owner_scope",
					In:   "query",
				}: params.PostOwnerScope,
				{
					Name: "only_media",
					In:   "query",
				}: params.OnlyMedia,
				{
					Name: "from_post_id",
					In:   "query",
				}: params.FromPostID,
				{
					Name: "number",
					In:   "query",
				}: params.Number,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchPostsParams
			Response = *PostsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchPostsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SearchPosts(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SearchPosts(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSearchPostsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSearchUsersRequest handles searchUsers operation.
//
// GET /v1/users/search
func (s *Server) handleSearchUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/users/search"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SearchUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SearchUsersOperation,
			ID:   "searchUsers",
		}
	)
	params, err := decodeSearchUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UsersResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SearchUsersOperation,
			OperationSummary: "",
			OperationID:      "searchUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "gender",
					In:   "query",
				}: params.Gender,
				{
					Name: "nickname",
					In:   "query",
				}: params.Nickname,
				{
					Name: "title",
					In:   "query",
				}: params.Title,
				{
					Name: "biography",
					In:   "query",
				}: params.Biography,
				{
					Name: "from_timestamp",
					In:   "query",
				}: params.FromTimestamp,
				{
					Name: "similar_age",
					In:   "query",
				}: params.SimilarAge,
				{
					Name: "not_recent_gomimushi",
					In:   "query",
				}: params.NotRecentGomimushi,
				{
					Name: "recently_created",
					In:   "query",
				}: params.RecentlyCreated,
				{
					Name: "same_prefecture",
					In:   "query",
				}: params.SamePrefecture,
				{
					Name: "save_recent_search",
					In:   "query",
				}: params.SaveRecentSearch,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchUsersParams
			Response = *UsersResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SearchUsers(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SearchUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSearchUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSendChatMessageRequest handles sendChatMessage operation.
//
// POST /v3/chat_rooms/{id}/messages/new
func (s *Server) handleSendChatMessageRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sendChatMessage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/chat_rooms/{id}/messages/new"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SendChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SendChatMessageOperation,
			ID:   "sendChatMessage",
		}
	)
	params, err := decodeSendChatMessageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeSendChatMessageRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *MessageResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SendChatMessageOperation,
			OperationSummary: "",
			OperationID:      "sendChatMessage",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *SendChatMessageReq
			Params   = SendChatMessageParams
			Response = *MessageResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSendChatMessageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SendChatMessage(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SendChatMessage(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSendChatMessageResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSetGroupTitleRequest handles setGroupTitle operation.
//
// POST /v1/groups/{id}/set_title
func (s *Server) handleSetGroupTitleRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setGroupTitle"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/set_title"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SetGroupTitleOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SetGroupTitleOperation,
			ID:   "setGroupTitle",
		}
	)
	params, err := decodeSetGroupTitleParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeSetGroupTitleRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *SetGroupTitleNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SetGroupTitleOperation,
			OperationSummary: "",
			OperationID:      "setGroupTitle",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *SetGroupTitleReq
			Params   = SetGroupTitleParams
			Response = *SetGroupTitleNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSetGroupTitleParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.SetGroupTitle(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.SetGroupTitle(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSetGroupTitleResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSetHimaRequest handles setHima operation.
//
// POST /v1/users/hima
func (s *Server) handleSetHimaRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setHima"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/hima"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), SetHimaOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var rawBody []byte

	var response *SetHimaNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SetHimaOperation,
			OperationSummary: "",
			OperationID:      "setHima",
			Body:             nil,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *SetHimaNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.SetHima(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.SetHima(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSetHimaResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleStartConferenceCallRequest handles startConferenceCall operation.
//
// POST /v2/calls/start_conference_call
func (s *Server) handleStartConferenceCallRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("startConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/calls/start_conference_call"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), StartConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: StartConferenceCallOperation,
			ID:   "startConferenceCall",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeStartConferenceCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ConferenceCallResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    StartConferenceCallOperation,
			OperationSummary: "",
			OperationID:      "startConferenceCall",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *StartConferenceCallReq
			Params   = struct{}
			Response = *ConferenceCallResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.StartConferenceCall(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.StartConferenceCall(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeStartConferenceCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTakeOverGroupRequest handles takeOverGroup operation.
//
// POST /v1/groups/{id}/take_over
func (s *Server) handleTakeOverGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("takeOverGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/take_over"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TakeOverGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TakeOverGroupOperation,
			ID:   "takeOverGroup",
		}
	)
	params, err := decodeTakeOverGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *TakeOverGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TakeOverGroupOperation,
			OperationSummary: "",
			OperationID:      "takeOverGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = TakeOverGroupParams
			Response = *TakeOverGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackTakeOverGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TakeOverGroup(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.TakeOverGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTakeOverGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTransferGroupRequest handles transferGroup operation.
//
// POST /v3/groups/{id}/transfer
func (s *Server) handleTransferGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("transferGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/groups/{id}/transfer"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TransferGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TransferGroupOperation,
			ID:   "transferGroup",
		}
	)
	params, err := decodeTransferGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeTransferGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TransferGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TransferGroupOperation,
			OperationSummary: "",
			OperationID:      "transferGroup",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *TransferGroupReq
			Params   = TransferGroupParams
			Response = *TransferGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackTransferGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TransferGroup(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.TransferGroup(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTransferGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnbanGroupUserRequest handles unbanGroupUser operation.
//
// POST /v1/groups/{id}/unban/{userId}
func (s *Server) handleUnbanGroupUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unbanGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/unban/{userId}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnbanGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnbanGroupUserOperation,
			ID:   "unbanGroupUser",
		}
	)
	params, err := decodeUnbanGroupUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnbanGroupUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnbanGroupUserOperation,
			OperationSummary: "",
			OperationID:      "unbanGroupUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "userId",
					In:   "path",
				}: params.UserId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnbanGroupUserParams
			Response = *UnbanGroupUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnbanGroupUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnbanGroupUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnbanGroupUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnbanGroupUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnblockUserRequest handles unblockUser operation.
//
// POST /v2/users/{id}/unblock
func (s *Server) handleUnblockUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unblockUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/{id}/unblock"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnblockUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnblockUserOperation,
			ID:   "unblockUser",
		}
	)
	params, err := decodeUnblockUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnblockUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnblockUserOperation,
			OperationSummary: "",
			OperationID:      "unblockUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnblockUserParams
			Response = *UnblockUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnblockUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnblockUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnblockUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnblockUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnfollowUserRequest handles unfollowUser operation.
//
// POST /v2/users/{id}/unfollow
func (s *Server) handleUnfollowUserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unfollowUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/users/{id}/unfollow"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnfollowUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnfollowUserOperation,
			ID:   "unfollowUser",
		}
	)
	params, err := decodeUnfollowUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnfollowUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnfollowUserOperation,
			OperationSummary: "",
			OperationID:      "unfollowUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnfollowUserParams
			Response = *UnfollowUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnfollowUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnfollowUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnfollowUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnfollowUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnhideChatsRequest handles unhideChats operation.
//
// DELETE /v1/hidden/chats
func (s *Server) handleUnhideChatsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unhideChats"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/hidden/chats"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnhideChatsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnhideChatsOperation,
			ID:   "unhideChats",
		}
	)
	params, err := decodeUnhideChatsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnhideChatsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnhideChatsOperation,
			OperationSummary: "",
			OperationID:      "unhideChats",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "chat_room_ids",
					In:   "query",
				}: params.ChatRoomIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnhideChatsParams
			Response = *UnhideChatsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnhideChatsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnhideChats(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnhideChats(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnhideChatsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnhideUsersRequest handles unhideUsers operation.
//
// DELETE /v1/hidden/users
func (s *Server) handleUnhideUsersRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unhideUsers"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/hidden/users"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnhideUsersOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnhideUsersOperation,
			ID:   "unhideUsers",
		}
	)
	params, err := decodeUnhideUsersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnhideUsersNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnhideUsersOperation,
			OperationSummary: "",
			OperationID:      "unhideUsers",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "user_ids[]",
					In:   "query",
				}: params.UserIds,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnhideUsersParams
			Response = *UnhideUsersNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnhideUsersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnhideUsers(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnhideUsers(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnhideUsersResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnlikePostRequest handles unlikePost operation.
//
// POST /v1/posts/{id}/unlike
func (s *Server) handleUnlikePostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unlikePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/posts/{id}/unlike"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnlikePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnlikePostOperation,
			ID:   "unlikePost",
		}
	)
	params, err := decodeUnlikePostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnlikePostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnlikePostOperation,
			OperationSummary: "",
			OperationID:      "unlikePost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnlikePostParams
			Response = *UnlikePostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnlikePostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnlikePost(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnlikePost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnlikePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnmuteGroupUserRequest handles unmuteGroupUser operation.
//
// DELETE /v1/group_mute/{id}/unmute/{user_id}
func (s *Server) handleUnmuteGroupUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unmuteGroupUser"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/group_mute/{id}/unmute/{user_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnmuteGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnmuteGroupUserOperation,
			ID:   "unmuteGroupUser",
		}
	)
	params, err := decodeUnmuteGroupUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnmuteGroupUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnmuteGroupUserOperation,
			OperationSummary: "",
			OperationID:      "unmuteGroupUser",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnmuteGroupUserParams
			Response = *UnmuteGroupUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnmuteGroupUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnmuteGroupUser(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnmuteGroupUser(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnmuteGroupUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnpinChatRoomRequest handles unpinChatRoom operation.
//
// DELETE /v1/chat_rooms/{id}/pinned
func (s *Server) handleUnpinChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinChatRoom"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/chat_rooms/{id}/pinned"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnpinChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnpinChatRoomOperation,
			ID:   "unpinChatRoom",
		}
	)
	params, err := decodeUnpinChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnpinChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnpinChatRoomOperation,
			OperationSummary: "",
			OperationID:      "unpinChatRoom",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnpinChatRoomParams
			Response = *UnpinChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnpinChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnpinChatRoom(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnpinChatRoom(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnpinChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnpinGroupRequest handles unpinGroup operation.
//
// DELETE /v1/pinned/groups/{id}
func (s *Server) handleUnpinGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroup"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/pinned/groups/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnpinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnpinGroupOperation,
			ID:   "unpinGroup",
		}
	)
	params, err := decodeUnpinGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnpinGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnpinGroupOperation,
			OperationSummary: "",
			OperationID:      "unpinGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnpinGroupParams
			Response = *UnpinGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnpinGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnpinGroup(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnpinGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnpinGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnpinGroupHighlightPostRequest handles unpinGroupHighlightPost operation.
//
// DELETE /v1/groups/{group_id}/highlights/{post_id}
func (s *Server) handleUnpinGroupHighlightPostRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroupHighlightPost"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/highlights/{post_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnpinGroupHighlightPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnpinGroupHighlightPostOperation,
			ID:   "unpinGroupHighlightPost",
		}
	)
	params, err := decodeUnpinGroupHighlightPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnpinGroupHighlightPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnpinGroupHighlightPostOperation,
			OperationSummary: "",
			OperationID:      "unpinGroupHighlightPost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "post_id",
					In:   "path",
				}: params.PostID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnpinGroupHighlightPostParams
			Response = *UnpinGroupHighlightPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnpinGroupHighlightPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnpinGroupHighlightPost(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnpinGroupHighlightPost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnpinGroupHighlightPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnpinGroupPostRequest handles unpinGroupPost operation.
//
// DELETE /v2/posts/group_pinned_post
func (s *Server) handleUnpinGroupPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroupPost"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v2/posts/group_pinned_post"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnpinGroupPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnpinGroupPostOperation,
			ID:   "unpinGroupPost",
		}
	)
	params, err := decodeUnpinGroupPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnpinGroupPostNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnpinGroupPostOperation,
			OperationSummary: "",
			OperationID:      "unpinGroupPost",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "query",
				}: params.GroupID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnpinGroupPostParams
			Response = *UnpinGroupPostNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnpinGroupPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnpinGroupPost(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnpinGroupPost(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnpinGroupPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUnpinReviewRequest handles unpinReview operation.
//
// DELETE /v1/pinned/reviews/{id}
func (s *Server) handleUnpinReviewRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinReview"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/pinned/reviews/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UnpinReviewOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UnpinReviewOperation,
			ID:   "unpinReview",
		}
	)
	params, err := decodeUnpinReviewParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UnpinReviewNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UnpinReviewOperation,
			OperationSummary: "",
			OperationID:      "unpinReview",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UnpinReviewParams
			Response = *UnpinReviewNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUnpinReviewParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UnpinReview(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UnpinReview(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUnpinReviewResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateAdditionalNotificationSettingRequest handles updateAdditionalNotificationSetting operation.
//
// POST /v1/users/additonal_notification_setting
func (s *Server) handleUpdateAdditionalNotificationSettingRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateAdditionalNotificationSetting"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/users/additonal_notification_setting"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateAdditionalNotificationSettingOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateAdditionalNotificationSettingOperation,
			ID:   "updateAdditionalNotificationSetting",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateAdditionalNotificationSettingRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateAdditionalNotificationSettingNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateAdditionalNotificationSettingOperation,
			OperationSummary: "",
			OperationID:      "updateAdditionalNotificationSetting",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UpdateAdditionalNotificationSettingReq
			Params   = struct{}
			Response = *UpdateAdditionalNotificationSettingNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateAdditionalNotificationSetting(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateAdditionalNotificationSetting(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateAdditionalNotificationSettingResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateCallRequest handles updateCall operation.
//
// PUT /v1/calls/{call_id}
func (s *Server) handleUpdateCallRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateCall"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateCallOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateCallOperation,
			ID:   "updateCall",
		}
	)
	params, err := decodeUpdateCallParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateCallRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateCallNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateCallOperation,
			OperationSummary: "",
			OperationID:      "updateCall",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateCallReq
			Params   = UpdateCallParams
			Response = *UpdateCallNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateCallParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateCall(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateCall(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateCallResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateCallUserRequest handles updateCallUser operation.
//
// PUT /v1/calls/{call_id}/users/{user_id}
func (s *Server) handleUpdateCallUserRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateCallUser"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/calls/{call_id}/users/{user_id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateCallUserOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateCallUserOperation,
			ID:   "updateCallUser",
		}
	)
	params, err := decodeUpdateCallUserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateCallUserRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateCallUserNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateCallUserOperation,
			OperationSummary: "",
			OperationID:      "updateCallUser",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "path",
				}: params.CallID,
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateCallUserReq
			Params   = UpdateCallUserParams
			Response = *UpdateCallUserNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateCallUserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateCallUser(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateCallUser(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateCallUserResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateChatRoomRequest handles updateChatRoom operation.
//
// POST /v3/chat_rooms/{id}/edit
func (s *Server) handleUpdateChatRoomRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/chat_rooms/{id}/edit"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateChatRoomOperation,
			ID:   "updateChatRoom",
		}
	)
	params, err := decodeUpdateChatRoomParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateChatRoomRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateChatRoomNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateChatRoomOperation,
			OperationSummary: "",
			OperationID:      "updateChatRoom",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateChatRoomReq
			Params   = UpdateChatRoomParams
			Response = *UpdateChatRoomNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateChatRoomParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateChatRoom(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateChatRoom(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateChatRoomResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateChatRoomNotificationSettingsRequest handles updateChatRoomNotificationSettings operation.
//
// POST /v2/notification_settings/chat_rooms/{id}
func (s *Server) handleUpdateChatRoomNotificationSettingsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateChatRoomNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/notification_settings/chat_rooms/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateChatRoomNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateChatRoomNotificationSettingsOperation,
			ID:   "updateChatRoomNotificationSettings",
		}
	)
	params, err := decodeUpdateChatRoomNotificationSettingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateChatRoomNotificationSettingsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *NotificationSettingResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateChatRoomNotificationSettingsOperation,
			OperationSummary: "",
			OperationID:      "updateChatRoomNotificationSettings",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateChatRoomNotificationSettingsReq
			Params   = UpdateChatRoomNotificationSettingsParams
			Response = *NotificationSettingResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateChatRoomNotificationSettingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateChatRoomNotificationSettings(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateChatRoomNotificationSettings(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateChatRoomNotificationSettingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateGroupRequest handles updateGroup operation.
//
// POST /v3/groups/{id}/update
func (s *Server) handleUpdateGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/groups/{id}/update"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateGroupOperation,
			ID:   "updateGroup",
		}
	)
	params, err := decodeUpdateGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateGroupRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *GroupResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateGroupOperation,
			OperationSummary: "",
			OperationID:      "updateGroup",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateGroupReq
			Params   = UpdateGroupParams
			Response = *GroupResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateGroup(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateGroup(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateGroupNotificationSettingsRequest handles updateGroupNotificationSettings operation.
//
// POST /v2/notification_settings/groups/{id}
func (s *Server) handleUpdateGroupNotificationSettingsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateGroupNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/notification_settings/groups/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateGroupNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateGroupNotificationSettingsOperation,
			ID:   "updateGroupNotificationSettings",
		}
	)
	params, err := decodeUpdateGroupNotificationSettingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateGroupNotificationSettingsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *AdditionalSettingsResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateGroupNotificationSettingsOperation,
			OperationSummary: "",
			OperationID:      "updateGroupNotificationSettings",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateGroupNotificationSettingsReq
			Params   = UpdateGroupNotificationSettingsParams
			Response = *AdditionalSettingsResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateGroupNotificationSettingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateGroupNotificationSettings(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateGroupNotificationSettings(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateGroupNotificationSettingsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateLanguageRequest handles updateLanguage operation.
//
// PUT /v1/users/language
func (s *Server) handleUpdateLanguageRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateLanguage"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/language"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateLanguageOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateLanguageOperation,
			ID:   "updateLanguage",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateLanguageRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateLanguageNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateLanguageOperation,
			OperationSummary: "",
			OperationID:      "updateLanguage",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UpdateLanguageReq
			Params   = struct{}
			Response = *UpdateLanguageNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateLanguage(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateLanguage(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateLanguageResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateLoginRequest handles updateLogin operation.
//
// POST /v3/users/login_update
func (s *Server) handleUpdateLoginRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateLogin"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v3/users/login_update"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateLoginOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateLoginOperation,
			ID:   "updateLogin",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateLoginRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *LoginUpdateResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateLoginOperation,
			OperationSummary: "",
			OperationID:      "updateLogin",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UpdateLoginReq
			Params   = struct{}
			Response = *LoginUpdateResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateLogin(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateLogin(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateLoginResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdatePostRequest handles updatePost operation.
//
// PUT /v3/posts/{id}
func (s *Server) handleUpdatePostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updatePost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v3/posts/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdatePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdatePostOperation,
			ID:   "updatePost",
		}
	)
	params, err := decodeUpdatePostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdatePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Post
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdatePostOperation,
			OperationSummary: "",
			OperationID:      "updatePost",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdatePostReq
			Params   = UpdatePostParams
			Response = *Post
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdatePostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdatePost(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdatePost(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdatePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateRelatedGroupsRequest handles updateRelatedGroups operation.
//
// PUT /v1/groups/{id}/related
func (s *Server) handleUpdateRelatedGroupsRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/related"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateRelatedGroupsOperation,
			ID:   "updateRelatedGroups",
		}
	)
	params, err := decodeUpdateRelatedGroupsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *UpdateRelatedGroupsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateRelatedGroupsOperation,
			OperationSummary: "",
			OperationID:      "updateRelatedGroups",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
				{
					Name: "related_group_id[]",
					In:   "query",
				}: params.RelatedGroupID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UpdateRelatedGroupsParams
			Response = *UpdateRelatedGroupsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateRelatedGroupsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateRelatedGroups(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateRelatedGroups(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateRelatedGroupsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateThreadRequest handles updateThread operation.
//
// PUT /v1/threads/{id}
func (s *Server) handleUpdateThreadRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateThread"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/threads/{id}"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateThreadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateThreadOperation,
			ID:   "updateThread",
		}
	)
	params, err := decodeUpdateThreadParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateThreadRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateThreadNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateThreadOperation,
			OperationSummary: "",
			OperationID:      "updateThread",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateThreadReq
			Params   = UpdateThreadParams
			Response = *UpdateThreadNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateThreadParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateThread(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateThread(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateThreadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateUserInterestsRequest handles updateUserInterests operation.
//
// PUT /v1/users/interests
func (s *Server) handleUpdateUserInterestsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateUserInterests"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/users/interests"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), UpdateUserInterestsOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UpdateUserInterestsOperation,
			ID:   "updateUserInterests",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeUpdateUserInterestsRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UpdateUserInterestsNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UpdateUserInterestsOperation,
			OperationSummary: "",
			OperationID:      "updateUserInterests",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *CommonIdsRequest
			Params   = struct{}
			Response = *UpdateUserInterestsNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UpdateUserInterests(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.UpdateUserInterests(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateUserInterestsResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleValidateCallActionSignatureRequest handles validateCallActionSignature operation.
//
// GET /v2/calls/action_signature/validate
func (s *Server) handleValidateCallActionSignatureRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("validateCallActionSignature"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v2/calls/action_signature/validate"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ValidateCallActionSignatureOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ValidateCallActionSignatureOperation,
			ID:   "validateCallActionSignature",
		}
	)
	params, err := decodeValidateCallActionSignatureParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ValidateCallActionSignatureNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ValidateCallActionSignatureOperation,
			OperationSummary: "",
			OperationID:      "validateCallActionSignature",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "call_id",
					In:   "query",
				}: params.CallID,
				{
					Name: "sender_uuid",
					In:   "query",
				}: params.SenderUUID,
				{
					Name: "receiver_uuid",
					In:   "query",
				}: params.ReceiverUUID,
				{
					Name: "action",
					In:   "query",
				}: params.Action,
				{
					Name: "timestamp",
					In:   "query",
				}: params.Timestamp,
				{
					Name: "signature",
					In:   "query",
				}: params.Signature,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ValidateCallActionSignatureParams
			Response = *ValidateCallActionSignatureNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackValidateCallActionSignatureParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ValidateCallActionSignature(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ValidateCallActionSignature(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeValidateCallActionSignatureResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleValidatePostRequest handles validatePost operation.
//
// POST /v1/posts/validate
func (s *Server) handleValidatePostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("validatePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/posts/validate"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ValidatePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ValidatePostOperation,
			ID:   "validatePost",
		}
	)

	var rawBody []byte
	request, rawBody, close, err := s.decodeValidatePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ValidationPostResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ValidatePostOperation,
			OperationSummary: "",
			OperationID:      "validatePost",
			Body:             request,
			RawBody:          rawBody,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ValidatePostReq
			Params   = struct{}
			Response = *ValidationPostResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ValidatePost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.ValidatePost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeValidatePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleViewPostVideoRequest handles viewPostVideo operation.
//
// POST /v1/posts/videos/{videoId}/view
func (s *Server) handleViewPostVideoRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("viewPostVideo"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/posts/videos/{videoId}/view"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), ViewPostVideoOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ViewPostVideoOperation,
			ID:   "viewPostVideo",
		}
	)
	params, err := decodeViewPostVideoParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *ViewPostVideoNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ViewPostVideoOperation,
			OperationSummary: "",
			OperationID:      "viewPostVideo",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "videoId",
					In:   "path",
				}: params.VideoId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ViewPostVideoParams
			Response = *ViewPostVideoNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackViewPostVideoParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ViewPostVideo(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ViewPostVideo(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeViewPostVideoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleVisitGroupRequest handles visitGroup operation.
//
// POST /v1/groups/{id}/visit
func (s *Server) handleVisitGroupRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("visitGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/visit"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), VisitGroupOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: VisitGroupOperation,
			ID:   "visitGroup",
		}
	)
	params, err := decodeVisitGroupParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *VisitGroupNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    VisitGroupOperation,
			OperationSummary: "",
			OperationID:      "visitGroup",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = VisitGroupParams
			Response = *VisitGroupNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackVisitGroupParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.VisitGroup(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.VisitGroup(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeVisitGroupResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleVoteSurveyRequest handles voteSurvey operation.
//
// POST /v2/surveys/{id}/vote
func (s *Server) handleVoteSurveyRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("voteSurvey"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v2/surveys/{id}/vote"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), VoteSurveyOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: VoteSurveyOperation,
			ID:   "voteSurvey",
		}
	)
	params, err := decodeVoteSurveyParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeVoteSurveyRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *VoteSurveyResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    VoteSurveyOperation,
			OperationSummary: "",
			OperationID:      "voteSurvey",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *VoteSurveyReq
			Params   = VoteSurveyParams
			Response = *VoteSurveyResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackVoteSurveyParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.VoteSurvey(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.VoteSurvey(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeVoteSurveyResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleWithdrawGroupDeputyRequest handles withdrawGroupDeputy operation.
//
// PUT /v1/groups/{group_id}/deputize/{user_id}/withdraw
func (s *Server) handleWithdrawGroupDeputyRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("withdrawGroupDeputy"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{group_id}/deputize/{user_id}/withdraw"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), WithdrawGroupDeputyOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: WithdrawGroupDeputyOperation,
			ID:   "withdrawGroupDeputy",
		}
	)
	params, err := decodeWithdrawGroupDeputyParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte

	var response *WithdrawGroupDeputyNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    WithdrawGroupDeputyOperation,
			OperationSummary: "",
			OperationID:      "withdrawGroupDeputy",
			Body:             nil,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "group_id",
					In:   "path",
				}: params.GroupID,
				{
					Name: "user_id",
					In:   "path",
				}: params.UserID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = WithdrawGroupDeputyParams
			Response = *WithdrawGroupDeputyNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackWithdrawGroupDeputyParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.WithdrawGroupDeputy(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.WithdrawGroupDeputy(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeWithdrawGroupDeputyResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleWithdrawGroupTransferRequest handles withdrawGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer/withdraw
func (s *Server) handleWithdrawGroupTransferRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("withdrawGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/groups/{id}/transfer/withdraw"),
	}
	// Add attributes from config.
	otelAttrs = append(otelAttrs, s.cfg.Attributes...)

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), WithdrawGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: WithdrawGroupTransferOperation,
			ID:   "withdrawGroupTransfer",
		}
	)
	params, err := decodeWithdrawGroupTransferParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var rawBody []byte
	request, rawBody, close, err := s.decodeWithdrawGroupTransferRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *WithdrawGroupTransferNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    WithdrawGroupTransferOperation,
			OperationSummary: "",
			OperationID:      "withdrawGroupTransfer",
			Body:             request,
			RawBody:          rawBody,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *WithdrawGroupTransferReq
			Params   = WithdrawGroupTransferParams
			Response = *WithdrawGroupTransferNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackWithdrawGroupTransferParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.WithdrawGroupTransfer(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.WithdrawGroupTransfer(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeWithdrawGroupTransferResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
