package controllers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// GetAssets
func GetAssets() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./web/build", true))
}
