package dao

import (
	"context"
	"fmt"
	"github.com/obase/demo/api"
	"testing"
)

func TestAdd(t *testing.T) {
	rsp, err := Add(context.Background(), &api.Student{
		Name:  "test",
		Age:   44,
		Score: 1.5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", rsp)
}
