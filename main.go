package main

import (
  "github.com/gin-gonic/gin"
)

func main(){
  r := gin.Default()

  //r.Any("/*pth", controllers.ReverseProxy)    

  r.GET("/", func(c *gin.Context){
    c.String(200, c.Request.Host) 
  })

  r.Run()
}
