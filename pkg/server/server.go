// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// GET TURN credentials
	// (GET /)
	GetTurnAuth(w http.ResponseWriter, r *http.Request, params GetTurnAuthParams)
	// GET ICE credentials
	// (GET /ice)
	GetIceAuth(w http.ResponseWriter, r *http.Request, params GetIceAuthParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetTurnAuth operation middleware
func (siw *ServerInterfaceWrapper) GetTurnAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTurnAuthParams

	// ------------- Required query parameter "service" -------------

	if paramValue := r.URL.Query().Get("service"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "service"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "service", r.URL.Query(), &params.Service)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "service", Err: err})
		return
	}

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", r.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	// ------------- Optional query parameter "ttl" -------------

	err = runtime.BindQueryParameter("form", true, false, "ttl", r.URL.Query(), &params.Ttl)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "ttl", Err: err})
		return
	}

	// ------------- Optional query parameter "key" -------------

	err = runtime.BindQueryParameter("form", true, false, "key", r.URL.Query(), &params.Key)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	// ------------- Optional query parameter "namespace" -------------

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Optional query parameter "gateway" -------------

	err = runtime.BindQueryParameter("form", true, false, "gateway", r.URL.Query(), &params.Gateway)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "gateway", Err: err})
		return
	}

	// ------------- Optional query parameter "listener" -------------

	err = runtime.BindQueryParameter("form", true, false, "listener", r.URL.Query(), &params.Listener)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listener", Err: err})
		return
	}

	// ------------- Optional query parameter "publicIp" -------------

	err = runtime.BindQueryParameter("form", true, false, "publicIp", r.URL.Query(), &params.PublicIP)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "publicIp", Err: err})
		return
	}

	// ------------- Optional query parameter "publicPort" -------------

	err = runtime.BindQueryParameter("form", true, false, "publicPort", r.URL.Query(), &params.PublicPort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "publicPort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTurnAuth(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetIceAuth operation middleware
func (siw *ServerInterfaceWrapper) GetIceAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetIceAuthParams

	// ------------- Optional query parameter "service" -------------

	err = runtime.BindQueryParameter("form", true, false, "service", r.URL.Query(), &params.Service)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "service", Err: err})
		return
	}

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", r.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	// ------------- Optional query parameter "ttl" -------------

	err = runtime.BindQueryParameter("form", true, false, "ttl", r.URL.Query(), &params.Ttl)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "ttl", Err: err})
		return
	}

	// ------------- Optional query parameter "iceTransportPolicy" -------------

	err = runtime.BindQueryParameter("form", true, false, "iceTransportPolicy", r.URL.Query(), &params.IceTransportPolicy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "iceTransportPolicy", Err: err})
		return
	}

	// ------------- Optional query parameter "key" -------------

	err = runtime.BindQueryParameter("form", true, false, "key", r.URL.Query(), &params.Key)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	// ------------- Optional query parameter "namespace" -------------

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Optional query parameter "gateway" -------------

	err = runtime.BindQueryParameter("form", true, false, "gateway", r.URL.Query(), &params.Gateway)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "gateway", Err: err})
		return
	}

	// ------------- Optional query parameter "listener" -------------

	err = runtime.BindQueryParameter("form", true, false, "listener", r.URL.Query(), &params.Listener)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listener", Err: err})
		return
	}

	// ------------- Optional query parameter "publicIp" -------------

	err = runtime.BindQueryParameter("form", true, false, "publicIp", r.URL.Query(), &params.PublicIP)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "publicIp", Err: err})
		return
	}

	// ------------- Optional query parameter "publicPort" -------------

	err = runtime.BindQueryParameter("form", true, false, "publicPort", r.URL.Query(), &params.PublicPort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "publicPort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetIceAuth(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/", wrapper.GetTurnAuth).Methods("GET")

	r.HandleFunc(options.BaseURL+"/ice", wrapper.GetIceAuth).Methods("GET")

	return r
}
