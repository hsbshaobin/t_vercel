package handler

import (
	"adCheck/cmd/api/service"
	"adCheck/cmd/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateCode(context *gin.Context) {
	var bo model.GeneralCodeBO
	if err := context.ShouldBindJSON(&bo); err != nil {
		context.JSON(http.StatusOK, model.Failed(http.StatusBadRequest, "参数绑定异常", err.Error()))
		return
	}
	context.JSON(http.StatusOK, service.GenerateCode(context, bo))
}

func CheckCode(context *gin.Context) {
	var bo model.CheckCodeBO
	if err := context.ShouldBindJSON(&bo); err != nil {
		context.JSON(http.StatusOK, model.Failed(http.StatusBadRequest, "参数绑定异常", err.Error()))
		return
	}
	context.JSON(http.StatusOK, service.CheckCode(context, bo))
}
