package hdl

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context)  {
	ctx.String(200, `这是首页`)
}
