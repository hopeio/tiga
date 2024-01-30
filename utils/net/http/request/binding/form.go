// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	if err := Decode(obj, req.Form); err != nil {
		return err
	}
	return Validate(obj)
}

func (formBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	args := Args{FormSource(ctx.Request.Form), paramSource(ctx.Params)}
	if err := MapForm(obj, args); err != nil {
		return err
	}
	return Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := Decode(obj, req.PostForm); err != nil {
		return err
	}
	return Validate(obj)
}

func (formPostBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseForm(); err != nil {
		return err
	}

	args := Args{FormSource(ctx.Request.Form), paramSource(ctx.Params)}
	if err := MapForm(obj, args); err != nil {
		return err
	}
	return Validate(obj)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := MappingByPtr(obj, (*multipartRequest)(req), Tag); err != nil {
		return err
	}

	return Validate(obj)
}

func (formMultipartBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := MappingByPtr(obj, (*multipartRequest)(ctx.Request), Tag); err != nil {
		return err
	}

	return Validate(obj)
}
