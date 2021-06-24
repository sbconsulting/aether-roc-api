// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	legacyrouter "github.com/getkin/kin-openapi/routers/legacy"
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/middleware/openapi3mw"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Passing a field like description with no length should give an error
// in validation because the minLength=1
func Test_ValidateRequestGivesError(t *testing.T) {
	e := echo.New()
	testDelete := `{
    "default-target": "connectivity-service-v2",
    "Deletes": {
    },
    "Updates": {
        "qos-profile-2.1.0": {
            "qos-profile": [
                {
                    "id": "sed",
                    "description": ""
                }
            ]
        }
    },
    "Extensions": {
        "model-version-101": "2.1.0",
        "model-type-102": "Aether"
    }
}`
	req := httptest.NewRequest(http.MethodPatch, "/aether-roc-api", strings.NewReader(testDelete))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.Assert(t, c != nil)
	h := func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	}
	assert.Assert(t, h != nil)
	openAPIDefinition, err := GetSwagger()
	assert.NilError(t, err)

	openapi3Router, err := legacyrouter.NewRouter(openAPIDefinition)
	assert.NilError(t, err)

	_, err = openapi3mw.ValidateRequest(c, openapi3Router)
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, 400)
	expectError := `Error at "/Updates/qos-profile-2.1.0/qos-profile/0/description": minimum string length is 1
Schema:
  {
    "description": "description of this profile",
    "maxLength": 100,
    "minLength": 1,
    "title": "description",
    "type": "string"
  }

Value:
  ""
`
	expectErrorEscaped := fmt.Sprintf("%#v\n", expectError)
	assert.Equal(t, expectErrorEscaped, rec.Body.String())
}

// Passing a description with a length of 0 in a delete is OK though
// because in the API, this is our way of telling the gNMI to delete it
func Test_ValidateRequestDeleteDesc(t *testing.T) {
	e := echo.New()
	testDelete := `{
    "default-target": "connectivity-service-v2",
    "Updates": {
    },
    "Deletes": {
        "qos-profile-2.1.0": {
            "qos-profile": [
                {
                    "id": "sed",
                    "description": ""
                }
            ]
        }
    },
    "Extensions": {
        "model-version-101": "2.1.0",
        "model-type-102": "Aether"
    }
}`
	req := httptest.NewRequest(http.MethodPatch, "/aether-roc-api", strings.NewReader(testDelete))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.Assert(t, c != nil)
	h := func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	}
	assert.Assert(t, h != nil)
	openAPIDefinition, err := GetSwagger()
	assert.NilError(t, err)

	openapi3Router, err := legacyrouter.NewRouter(openAPIDefinition)
	assert.NilError(t, err)

	_, err = openapi3mw.ValidateRequest(c, openapi3Router)
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, 200)
}