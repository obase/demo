package biz

import (
	"context"
	"fmt"
	"github.com/obase/demo/api"
	"testing"
)

func TestStudentServiceService_Add(t *testing.T) {
	service := new(StudentServiceService)
	rsp, err := service.Add(context.Background(), &api.Student{
		Name:  "test",
		Age:   44,
		Score: 1.5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", rsp)
}
