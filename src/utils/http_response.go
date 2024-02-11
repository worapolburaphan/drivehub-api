package utils

import "github.com/gin-gonic/gin"

func NewResponse(result interface{}) gin.H {
	return gin.H{"result": result}
}
