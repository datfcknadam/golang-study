package routes

import (
	mCrypto "golang-study/models/crypto"

	"github.com/gin-gonic/gin"
)

func (r routes) addCrypto(rg *gin.RouterGroup) {
	name := rg.Group("/crypto")

	name.POST("/", mCrypto.EncryptFile)
}
