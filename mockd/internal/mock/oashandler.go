package mock

import (
	"context"
	"net/http"

	"github.com/ekkx/yaylib/mockd/internal/oas"
)

// oasHandler implements the typed router's Handler interface by
// embedding its no-op base handler: every operation returns
// http.ErrNotImplemented. We deliberately do not hand-write 200+
// typed handlers — the typed router still parses the path and
// parameters (so unknown routes 404 and the SDKs exercise real
// routing), and genericErrorHandler turns the not-implemented signal
// into an empty 200 JSON body. Endpoints that parity tests need a
// real body for (the token endpoint) are owned by the scenario
// middleware ahead of this handler.
type oasHandler struct {
	oas.UnimplementedHandler
}

// genericErrorHandler converts any router-level error into a
// permissive empty 200 JSON object. A stub server must not reject
// otherwise-valid SDK calls just because request decoding or response
// typing is imperfect; behavioral fidelity lives in the scenario
// middleware, not here. The original error is surfaced in a debug
// header for troubleshooting.
func genericErrorHandler(_ context.Context, w http.ResponseWriter, _ *http.Request, err error) {
	if err != nil {
		w.Header().Set("X-Mock-Router-Err", err.Error())
	}
	writeJSON(w, http.StatusOK, struct{}{})
}

// newOASServer builds the typed router with the permissive error
// handler installed.
func newOASServer() (http.Handler, error) {
	return oas.NewServer(oasHandler{}, oas.WithErrorHandler(genericErrorHandler))
}
