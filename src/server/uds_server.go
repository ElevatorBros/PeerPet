package server

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"peer.pet/src/common"
)

func GetPetGIN(c *gin.Context) {
	data, err := os.ReadFile(folder_path + "/pet.json")
	errorCheck(c, err)

	c.JSON(http.StatusOK, gin.H{
		"pet": Secret(data),
	})
}

func PostPetGIN(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	err = os.WriteFile(folder_path+"/temp", data, fs.FileMode(0644))
	if err != nil {
		println(err.Error())
	}

	pet := new(common.Pet)
	err = json.Unmarshal(data, pet)

	err = os.WriteFile(folder_path+"/pet.json", Secret(data), fs.FileMode(0644))

	c.AbortWithStatus(http.StatusOK)
}

func errorCheck(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
