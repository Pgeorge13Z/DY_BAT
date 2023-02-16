// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	router "DY_BAT/biz/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	// idl 中定义的路由
	router.GeneratedRegister(r)

	// 用户自定义路由
	customizedRegister(r)
}
