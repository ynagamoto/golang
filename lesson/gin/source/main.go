package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "local.packages/modules"
)

func main() {
  r := gin.Default()
  // path := "/go/gin"
  // r.LoadHTMLGlob(path+"/templates/*.html")
  r.LoadHTMLGlob("./templates/*.html")

  // path
  r.GET("", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
  })
  r.GET("/json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "msg": "hoge",
    })
  })
  r.GET("/get", func(c *gin.Context) {
    c.HTML(http.StatusOK, "get.html", gin.H{
      "msg": c.Query("msg"),
    })
  })
  r.GET("/post", func(c *gin.Context) {
    c.HTML(http.StatusOK, "post.html", gin.H{})
  })
  r.POST("/post", func(c *gin.Context) {
    c.HTML(http.StatusOK, "post.html", gin.H{
      "name": c.PostForm("name"),
      "passwd": c.PostForm("passwd"),
    })
  })

  r.GET("/useradd", func(c *gin.Context) {
    c.HTML(http.StatusOK, "useradd.html", gin.H{
      "users": modules.GetUsers(),
    })
  })
  r.POST("/useradd", func(c *gin.Context) {
    modules.UserAdd(c.PostForm("name"), c.PostForm("passwd"))
    c.HTML(http.StatusOK, "useradd.html", gin.H{
      "users": modules.GetUsers(),
    })
  })

  r.Run() // listen and serve on 0.0.0.0:8080
}
