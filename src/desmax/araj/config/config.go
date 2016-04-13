// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package config

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"desmax/araj/common"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

)

type Configuration struct {
	AppName     string
	ServerPort  string
	Address  string
	Development bool
	Certfile string
	KeyFile string
}

func GetConfiguration() Configuration {

	development:=true
	address := os.Getenv(common.ADDRESS)
	serverPort := os.Getenv(common.PORT)
	appName := os.Getenv(common.APPNAME)
	dev := os.Getenv(common.DEVELOPMENT)
	certFile := os.Getenv(common.CERTFILE)
	keyFile := os.Getenv(common.KEYFILE)
	development, _ = strconv.ParseBool(dev)

	conf := Configuration{
		Development : development,
		Address : address,
		ServerPort : serverPort,
		AppName : appName,
		Certfile : certFile,
		KeyFile : keyFile,
	}

	if conf.Development {
		fmt.Println("Conf:", conf)
	}
	return conf
}

func HttpServer(router *fasthttprouter.Router,conf Configuration){
	log.Fatal(fasthttp.ListenAndServe(":" + conf.ServerPort, router.Handler))
}

func HttpsServer(router *fasthttprouter.Router, conf Configuration){
	log.Fatal(fasthttp.ListenAndServeTLS(conf.Address + ":"+conf.ServerPort, conf.Certfile, conf.KeyFile, router.Handler))
}



