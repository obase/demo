package demo

import (
	"context"
	"fmt"
	"github.com/obase/demo/api"
	"google.golang.org/grpc"
	"testing"
)

/*
grpc客户湍测试
 */
func TestDemo(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8100", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewStudentServiceClient(conn)

	req := &api.Student{
		Name:  "小王",
		Age:   25,
		Score: 12.5,
	}

	rsp, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rsp)
}

/*
http客户端测试: postman

curl -XPOST http://127.0.0.1:8000/api/student/add \
-d'
{
    "name": "小张",
    "age": 30,
    "score": 3.3
}
'
----------------------
结果
{
    "code": 0,
    "data": {
        "name": "小张",
        "age": 30,
        "score": 3.3
    },
    "tag": "StudentService.Add"
}
 */


/*
websocket客户端测试: chrome + smart websocket client
必须注意: "wbskNotCheckOrigin: true", 否则同域检查会失效

ws://localhost:8000/api/student/add

{
    "name": "小李",
    "age": 30,
    "score": 3.3
}
 */