package yaylib

import (
	"context"
	"log/slog"
)

// The SDK is a library: with no logger configured it produces zero
// output. The default Client.Logger is an *slog.Logger over this
// handler, whose Enabled always reports false — so slog never formats
// a record and the no-op path stays allocation-cheap. A caller opts in
// with WithLogger.
//
// Records carry a stable "event" key (snake_case); the event value is
// the cross-language contract (see PORTING.md §12.2). Banned values —
// tokens, password, API key, signed_info/version, X-Jwt, the
// Authorization header, request/response bodies — MUST NEVER be passed
// to the logger, at any level.
type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (discardHandler) WithAttrs([]slog.Attr) slog.Handler        { return discardHandler{} }
func (discardHandler) WithGroup(string) slog.Handler             { return discardHandler{} }

// newDiscardLogger returns the silent default logger.
func newDiscardLogger() *slog.Logger { return slog.New(discardHandler{}) }
