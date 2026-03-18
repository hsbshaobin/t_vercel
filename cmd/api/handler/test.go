package handler

import (
	"base_go_web/cmd/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(context *gin.Context) {
	//if err := context.ShouldBindJSON(&bo); err != nil {
	//	context.JSON(http.StatusOK, model.Failed(http.StatusBadRequest, "参数绑定异常", err.Error()))
	//	return
	//}
	context.JSON(http.StatusOK, service.Test())
}
