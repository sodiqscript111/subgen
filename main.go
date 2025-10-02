package main

import (
	"subgen/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 700 << 20

	r.POST("/upload", api.UploadFile)

	r.Run(":8080") // start server
}
