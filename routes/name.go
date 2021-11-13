package routes

import (
	mName "golang-study/models/name"

	"github.com/gin-gonic/gin"
)

func (r routes) addName(rg *gin.RouterGroup) {
	name := rg.Group("/name")

	name.POST("/add", mName.AddFullName)
	name.GET("/all", mName.GetAll)
	name.GET("/word/:word", mName.GetByWord)
}
