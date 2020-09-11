package main

import (
	"github.com/obase/cmd"
	"github.com/obase/demo/cmdmain/student"
)

func main() {
	cmd.Add("stduent", student.Init, student.Exec, "学生查询命令")
	cmd.Exec("demo")
}
