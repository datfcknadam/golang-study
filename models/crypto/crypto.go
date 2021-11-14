package crypto

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EncryptFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad file"})
		return
	}
	of, err := file.Open()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad file"})
		return
	}
	defer of.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(of)
	for sc.Scan() {
		for _, byte := range md5.Sum(sc.Bytes()) {
			wr.WriteByte(byte)
		}
	}
	arrStr := strings.Split(file.Filename, ".")
	ext := arrStr[len(arrStr)-1]
	b := []byte(ext)

	for _, byte := range b {
		wr.WriteByte(byte)
	}

	wr.WriteString(string(rune(len(b))))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+file.Filename)
	_, err = wr.WriteTo(c.Writer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad file"})
		log.Println(err)
		return
	}
}
