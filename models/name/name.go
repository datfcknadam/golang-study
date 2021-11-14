package name

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type fullname struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

var fullNames = []fullname{}

func AddFullName(c *gin.Context) {
	var newFullName fullname

	if err := c.ShouldBindJSON(&newFullName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad json :("})
	}

	fullNames = append(fullNames, newFullName)

	c.JSON(http.StatusOK, gin.H{"message": "welcome to list " + newFullName.Name + "!"})
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, fullNames)
}

func GetByWord(c *gin.Context) {
	word := c.Param("word")

	for _, n := range fullNames {
		if n.Name == word || n.Patronymic == word || n.Surname == word {
			c.JSON(http.StatusOK, n)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "fullname by word - " + word + " - not found"})
}
