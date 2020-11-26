package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func MakeSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

func MakeFail(c *gin.Context, status int, err error) {
	c.JSON(status, &Response{
		Code: status,
		Msg:  err.Error(),
		Data: nil,
	})
}
