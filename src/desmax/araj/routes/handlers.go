// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package routes

import (
	"fmt"
	"desmax/araj/middleware"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
)

func Index(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	conf, _ := middleware.GetConfig(ctx)
	fmt.Fprint(ctx, "Welcome to ", conf.AppName)
}
