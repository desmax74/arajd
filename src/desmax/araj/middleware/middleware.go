// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package middleware

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"desmax/araj/config"
	"fmt"
	"desmax/araj/common"
)

/* CONFIGURATION  */

func ConfigurationMiddleware(h fasthttprouter.Handle, conf config.Configuration) fasthttprouter.Handle  {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		SetConfig(ctx, conf)
		h(ctx, ps)
	})
}


func GetConfig(ctx *fasthttp.RequestCtx) (config.Configuration, error) {
	configuration := ctx.UserValue(common.CONFIG)
	if configuration != nil {
		return configuration.(config.Configuration), nil
	}else {
		return config.Configuration{}, fmt.Errorf("Config not found")
	}
}

func SetConfig(ctx *fasthttp.RequestCtx, configuration config.Configuration) {
	ctx.SetUserValue(common.CONFIG, configuration)
}

/* END CONFIGURATION */



