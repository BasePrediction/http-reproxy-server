package controllers

import (
  "fmt"
  "net/http"
  "net/http/httputil"
  "net/url"

	"github.com/gin-gonic/gin"
)


func ReverseProxy(c *gin.Context){
  addr, _addr_err := url.Parse("http://localhost:8080/")
  if _addr_err != nil {
    fmt.Println("*Error:", _addr_err)
  }

  rp := httputil.NewSingleHostReverseProxy(addr)
  rp.Director = func(r *http.Request){
    r.Header = c.Request.Header
    r.Host = addr.Host
    r.URL.Scheme = addr.Scheme
    r.URL.Host = addr.Host
    r.URL.Path = c.Param("pth")
  }

  rp.ServeHTTP(c.Writer, c.Request)

}
