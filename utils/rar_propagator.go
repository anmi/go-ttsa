package utils

import (
	"context"
	"net/http"
	"net/url"
)

var RequestKey struct{}
var ResponseWriterKey struct{}

type RARPropagator struct {
	Next http.Handler
}

func ContextWithRequest(
	ctx context.Context,
	r *http.Request,
	rw http.ResponseWriter,
) context.Context {
	return context.WithValue(
		context.WithValue(ctx, RequestKey, r),
		ResponseWriterKey,
		&rw,
	)
}

var request_key struct {
	Request int
}
var response_key struct {
	Response int
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, ok := ctx.Value(request_key).(*http.Request)
	if !ok {
		panic("No request in context")
	}

	return r
}

func ResponseWriterFromContext(ctx context.Context) *http.ResponseWriter {
	r, ok := ctx.Value(response_key).(*http.ResponseWriter)
	if !ok {
		panic("No ResponseWriter in context")
	}

	return r
}

func (p RARPropagator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, request_key, r)
	ctx = context.WithValue(ctx, response_key, &w)
	if r.URL.Path[:4] == "/api" {
		r.URL = &url.URL{
			Scheme:      r.URL.Scheme,
			Opaque:      r.URL.Opaque,
			User:        r.URL.User,
			Host:        r.URL.Host,
			Path:        r.URL.Path[4:],
			RawPath:     r.URL.RawPath,
			OmitHost:    r.URL.OmitHost,
			ForceQuery:  r.URL.ForceQuery,
			RawQuery:    r.URL.RawQuery,
			Fragment:    r.URL.Fragment,
			RawFragment: r.URL.RawFragment,
		}
	}
	p.Next.ServeHTTP(
		w,
		r.WithContext(ctx),
	)
}
