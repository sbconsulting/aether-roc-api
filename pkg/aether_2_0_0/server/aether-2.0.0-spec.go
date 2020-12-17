// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//4uHRzZ/fAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}