package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("shadowclone", HandleShadowClone)
	r.Run()
}
