package main

import (
	"github.com/obase/demo/api"
	"github.com/obase/demo/biz"
	"github.com/obase/demo/hdl"
	"github.com/obase/log"
	"github.com/obase/pbapi"
)

func main() {
	server := pbapi.NewServer()
	// 注册proto服务
	server.RegisterService(api.RegisterStudentServiceServerHandler, new(biz.StudentServiceService))
	// 注册普通的http服务
	server.GET("/index", hdl.Index)
	if err := server.Serve(); err != nil {
		log.Errorf("server serve error: %v", err)
	}
}
