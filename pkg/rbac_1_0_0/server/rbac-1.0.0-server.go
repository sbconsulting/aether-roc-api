// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"github.com/labstack/echo/v4"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// DELETE Generated from YANG model
	// (DELETE /rbac/v1.0.0/{target}/rbac)
	DeleteRbacV100targetRbac(ctx echo.Context, target types.Target) error
	// GET /rbac/v1.0.0/{target}/rbac Generated from YANG model
	// (GET /rbac/v1.0.0/{target}/rbac)
	GetRbacV100targetRbac(ctx echo.Context, target types.Target) error
	// POST Generated from YANG model
	// (POST /rbac/v1.0.0/{target}/rbac)
	PostRbacV100targetRbac(ctx echo.Context, target types.Target) error
	// DELETE Generated from YANG model
	// (DELETE /rbac/v1.0.0/{target}/rbac/group/{groupid})
	DeleteRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error
	// GET /rbac/v1.0.0/{target}/rbac/group Generated from YANG model
	// (GET /rbac/v1.0.0/{target}/rbac/group/{groupid})
	GetRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error
	// POST Generated from YANG model
	// (POST /rbac/v1.0.0/{target}/rbac/group/{groupid})
	PostRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error
	// DELETE Generated from YANG model
	// (DELETE /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid})
	DeleteRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error
	// GET /rbac/v1.0.0/{target}/rbac/group/role Generated from YANG model
	// (GET /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid})
	GetRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error
	// POST Generated from YANG model
	// (POST /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid})
	PostRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error
	// DELETE Generated from YANG model
	// (DELETE /rbac/v1.0.0/{target}/rbac/role/{roleid})
	DeleteRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error
	// GET /rbac/v1.0.0/{target}/rbac/role Generated from YANG model
	// (GET /rbac/v1.0.0/{target}/rbac/role/{roleid})
	GetRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error
	// POST Generated from YANG model
	// (POST /rbac/v1.0.0/{target}/rbac/role/{roleid})
	PostRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error
	// DELETE Generated from YANG model
	// (DELETE /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission)
	DeleteRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error
	// GET /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission Generated from YANG model
	// (GET /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission)
	GetRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error
	// POST Generated from YANG model
	// (POST /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission)
	PostRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteRbacV100targetRbac converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRbacV100targetRbac(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRbacV100targetRbac(ctx, target)
	return err
}

// GetRbacV100targetRbac converts echo context to params.
func (w *ServerInterfaceWrapper) GetRbacV100targetRbac(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRbacV100targetRbac(ctx, target)
	return err
}

// PostRbacV100targetRbac converts echo context to params.
func (w *ServerInterfaceWrapper) PostRbacV100targetRbac(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRbacV100targetRbac(ctx, target)
	return err
}

// DeleteRbacV100targetRbacGroup converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRbacV100targetRbacGroup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRbacV100targetRbacGroup(ctx, target, groupid)
	return err
}

// GetRbacV100targetRbacGroup converts echo context to params.
func (w *ServerInterfaceWrapper) GetRbacV100targetRbacGroup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRbacV100targetRbacGroup(ctx, target, groupid)
	return err
}

// PostRbacV100targetRbacGroup converts echo context to params.
func (w *ServerInterfaceWrapper) PostRbacV100targetRbacGroup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRbacV100targetRbacGroup(ctx, target, groupid)
	return err
}

// DeleteRbacV100targetRbacGroupRole converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRbacV100targetRbacGroupRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRbacV100targetRbacGroupRole(ctx, target, groupid, roleid)
	return err
}

// GetRbacV100targetRbacGroupRole converts echo context to params.
func (w *ServerInterfaceWrapper) GetRbacV100targetRbacGroupRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRbacV100targetRbacGroupRole(ctx, target, groupid, roleid)
	return err
}

// PostRbacV100targetRbacGroupRole converts echo context to params.
func (w *ServerInterfaceWrapper) PostRbacV100targetRbacGroupRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "groupid" -------------
	var groupid string

	groupid = ctx.Param("groupid")

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRbacV100targetRbacGroupRole(ctx, target, groupid, roleid)
	return err
}

// DeleteRbacV100targetRbacRole converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRbacV100targetRbacRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRbacV100targetRbacRole(ctx, target, roleid)
	return err
}

// GetRbacV100targetRbacRole converts echo context to params.
func (w *ServerInterfaceWrapper) GetRbacV100targetRbacRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRbacV100targetRbacRole(ctx, target, roleid)
	return err
}

// PostRbacV100targetRbacRole converts echo context to params.
func (w *ServerInterfaceWrapper) PostRbacV100targetRbacRole(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRbacV100targetRbacRole(ctx, target, roleid)
	return err
}

// DeleteRbacV100targetRbacRolePermission converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRbacV100targetRbacRolePermission(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRbacV100targetRbacRolePermission(ctx, target, roleid)
	return err
}

// GetRbacV100targetRbacRolePermission converts echo context to params.
func (w *ServerInterfaceWrapper) GetRbacV100targetRbacRolePermission(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRbacV100targetRbacRolePermission(ctx, target, roleid)
	return err
}

// PostRbacV100targetRbacRolePermission converts echo context to params.
func (w *ServerInterfaceWrapper) PostRbacV100targetRbacRolePermission(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "target" -------------
	var target types.Target

	target = types.Target(ctx.Param("target"))

	// ------------- Path parameter "roleid" -------------
	var roleid string

	roleid = ctx.Param("roleid")

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRbacV100targetRbacRolePermission(ctx, target, roleid)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/rbac/v1.0.0/:target/rbac", wrapper.DeleteRbacV100targetRbac)
	router.GET("/rbac/v1.0.0/:target/rbac", wrapper.GetRbacV100targetRbac)
	router.POST("/rbac/v1.0.0/:target/rbac", wrapper.PostRbacV100targetRbac)
	router.DELETE("/rbac/v1.0.0/:target/rbac/group/:groupid", wrapper.DeleteRbacV100targetRbacGroup)
	router.GET("/rbac/v1.0.0/:target/rbac/group/:groupid", wrapper.GetRbacV100targetRbacGroup)
	router.POST("/rbac/v1.0.0/:target/rbac/group/:groupid", wrapper.PostRbacV100targetRbacGroup)
	router.DELETE("/rbac/v1.0.0/:target/rbac/group/:groupid/role/:roleid", wrapper.DeleteRbacV100targetRbacGroupRole)
	router.GET("/rbac/v1.0.0/:target/rbac/group/:groupid/role/:roleid", wrapper.GetRbacV100targetRbacGroupRole)
	router.POST("/rbac/v1.0.0/:target/rbac/group/:groupid/role/:roleid", wrapper.PostRbacV100targetRbacGroupRole)
	router.DELETE("/rbac/v1.0.0/:target/rbac/role/:roleid", wrapper.DeleteRbacV100targetRbacRole)
	router.GET("/rbac/v1.0.0/:target/rbac/role/:roleid", wrapper.GetRbacV100targetRbacRole)
	router.POST("/rbac/v1.0.0/:target/rbac/role/:roleid", wrapper.PostRbacV100targetRbacRole)
	router.DELETE("/rbac/v1.0.0/:target/rbac/role/:roleid/permission", wrapper.DeleteRbacV100targetRbacRolePermission)
	router.GET("/rbac/v1.0.0/:target/rbac/role/:roleid/permission", wrapper.GetRbacV100targetRbacRolePermission)
	router.POST("/rbac/v1.0.0/:target/rbac/role/:roleid/permission", wrapper.PostRbacV100targetRbacRolePermission)

}