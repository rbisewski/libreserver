package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
)

// index ... return the index.html or equivalent
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: add proper checks here
		// input validation
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// select based on the file type
		pieces := strings.Split(r.URL.Path, ".")
		filetype := "txt"
		if len(pieces) > 1 {
			filetype = pieces[len(pieces)-1]
		}

		// handle different file types
		switch filetype {

		case "htm":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		case "html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		case "xhtml":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

		case "txt":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		}

		// handle the header
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		// TODO: add logic here to read / print file contents
		fmt.Fprintln(w, "this is a test, consider adding the "+
			"ability to read text / html files!")
	})
}

// healthz ... return whether or not the server is available
func healthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// logging ... print out a series of log messages to the attached console
func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// tracing ... obtain a trace of the request
func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
