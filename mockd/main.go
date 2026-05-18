// Command mockserver is a behavior-contract HTTP server for the
// cross-language SDK test suites. It serves every route the API spec
// defines and, when a request carries an X-Yay-Mock-Scenario header,
// produces the deterministic failure/latency behavior that header
// selects. See README.md for the scenario grammar (a cross-language
// contract).
package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ekkx/yaylib/mockd/internal/mock"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address (use host:0 for an ephemeral port)")
	readyFile := flag.String("ready-file", "", "if set, the resolved host:port is written here once listening (lets a harness use :0 without a port race)")
	flag.Parse()

	srv, err := mock.New()
	if err != nil {
		log.Fatalf("build mock server: %v", err)
	}

	// Listen explicitly so an ephemeral port (host:0) resolves to a
	// concrete address the caller can discover via -ready-file.
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen %s: %v", *addr, err)
	}
	resolved := ln.Addr().String()
	if *readyFile != "" {
		if err := os.WriteFile(*readyFile, []byte(resolved), 0o600); err != nil {
			log.Fatalf("write ready-file %s: %v", *readyFile, err)
		}
	}

	httpSrv := &http.Server{
		Handler:           srv,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("mockserver listening on %s", resolved)
		if err := httpSrv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("serve: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
		os.Exit(1)
	}
}
