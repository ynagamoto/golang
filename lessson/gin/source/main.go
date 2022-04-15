package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  path := "/go/gin"
  r.LoadHTMLGlob(path+"/templates/*.html")

  // path
  r.GET("", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
  })
  r.GET("/hoge", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "hoge",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080
}
