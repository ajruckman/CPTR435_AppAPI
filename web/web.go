package web

import (
    . "github.com/ajruckman/xlib"
    "github.com/gin-gonic/gin"

    "api/web/endpoint"
)

func Serve() {
    r := gin.Default()

    r.GET("/get", root)
    r.POST("/reading/create", endpoint.Reading_Create)
    r.GET("/reading/list", endpoint.Reading_List)

    err := r.Run(":80")
    Err(err)
}
