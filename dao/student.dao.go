package dao

import (
	"context"
	"github.com/obase/demo/api"
	"github.com/obase/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mdb = mongodb.Must("demo") // 启动时强制检查是否已经配置相关资源

func Add(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	_, err = mdb.ReplaceId("student", req.Name, req, options.Replace().SetUpsert(true))
	rsp = req
	return
}
func Upd(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	_, err = mdb.UpdateId("student", req.Name, bson.M{"$set": req})
	return
}
func Del(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	nvl, err := mdb.FindId("student", req.Name, &rsp)
	if nvl || err != nil {
		return
	}
	_, err = mdb.DeleteId("student", req.Name)
	return
}
func Get(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	_, err = mdb.FindId("student", req.Name, &rsp)
	return
}

func Page(ctx context.Context, req *api.PageReq) (rsp *api.PageRsp, err error) {
	rsp = new(api.PageRsp)
	con := make(bson.M)
	ops := options.Find()

	if req.Search != "" {
		con["name"] = bson.M{"$regex": req.Search}
	}
	if req.Sortby != "" {
		var sort int
		if req.Desc {
			sort = -1
		} else {
			sort = 1
		}
		ops.SetSort(bson.M{req.Sortby: sort})
	}
	if req.Page >= 0 {
		ops.SetSkip(int64(req.Page * req.Size))
		ops.SetLimit(int64(req.Size))
	}

	if rsp.Total, err = mdb.Count("stduent", con); err != nil {
		return
	}

	err = mdb.Find("student", con, &rsp.Data, ops)
	return
}
