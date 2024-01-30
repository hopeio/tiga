// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	if err := MapForm(obj, FormSource(values)); err != nil {
		return err
	}
	return Validate(obj)
}

func (queryBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	values := ctx.Request.URL.Query()
	args := Args{FormSource(ctx.Request.Form), FormSource(values), paramSource(ctx.Params)}
	if err := MapForm(obj, args); err != nil {
		return err
	}
	return Validate(obj)
}
