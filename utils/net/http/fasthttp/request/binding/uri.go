// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
)

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) BindUri(c fiber.Ctx, obj interface{}) error {
	if err := binding.MappingByPtr(obj, (*ctxSource)(c), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
