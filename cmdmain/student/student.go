package student

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/obase/cmd"
	"github.com/obase/demo/api"
	"github.com/obase/demo/biz"
)

var (
	method string
	params string
)

func Init() int {
	flag.StringVar(&method, "method", "", "请求方法名称")
	flag.StringVar(&params, "params", "", "请求参数列表")
	return 0
}

func Exec() int {
	if method != "" {
		service := new(biz.StudentServiceService)

		var (
			err error
			rsp interface{}
		)
		switch method {
		case "add":
			var req api.Student
			if err = json.Unmarshal([]byte(params), &req); err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			}
			rsp, err = service.Add(context.Background(), &req)
		case "del":
			var req api.Student
			if err = json.Unmarshal([]byte(params), &req); err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			}
			rsp, err = service.Del(context.Background(), &req)
		case "upd":
			var req api.Student
			if err = json.Unmarshal([]byte(params), &req); err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			}
			rsp, err = service.Upd(context.Background(), &req)
		case "get":
			var req api.Student
			if err = json.Unmarshal([]byte(params), &req); err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			}
			rsp, err = service.Get(context.Background(), &req)
		case "page":
			var req api.PageReq
			if err = json.Unmarshal([]byte(params), &req); err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			}
			rsp, err = service.Page(context.Background(), &req)
		}

		if rsp != nil {
			if err != nil {
				cmd.PrintError("parse params error: %v", err)
				return 1
			} else {
				bs, _ := json.Marshal(rsp)
				cmd.PrintInfo("result is: %s", bs)
			}
		}
	}
	return 0
}
