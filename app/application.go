package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//test comment
func StartApplication() {
	mapUrls()
	router.Run(":3001")
}
