// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"desmax/araj/common"
	"desmax/araj/config"
	"desmax/araj/middleware"
	"desmax/araj/routes"
	"github.com/buaazp/fasthttprouter"
)

var conf config.Configuration

func init() {
	fmt.Println(common.Banner)
	conf = config.GetConfiguration()
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", middleware.ConfigurationMiddleware(routes.Index, conf))
	//config.HttpsServer(router, conf)
	config.HttpServer(router, conf)
}

