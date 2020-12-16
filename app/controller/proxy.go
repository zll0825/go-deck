package controller

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/service"
	"go-deck/app/util"
)

func Proxy(ctx *gin.Context) {

	s := util.Service{}
	p := service.NewProxy(ctx, &s)
	p.Run()
}