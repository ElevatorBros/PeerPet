package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"peer.pet/src/common"
)

type PostPetJson struct {
	Pet common.Pet
}

func GetPet(c *gin.Context) {
	folder, err := os.Open(folder_path)
	errorCheck(c, err)
	data, err := os.ReadFile(folder.Name() + "/pet.json")
	errorCheck(c, err)
	c.JSON(http.StatusOK, gin.H{
		"pet": data,
	})
}

func PostPet(c *gin.Context) {
	var data []byte
	_, err := c.Request.Body.Read(data)
	errorCheck(c, err)
	jsonReq := &PostPetJson{}
	err = json.Unmarshal(data, jsonReq)
	errorCheck(c, err)
	petJson, err := jsonReq.Pet.Jsonify()
	errorCheck(c, err)
	err = os.WriteFile(folder_path+"/pet.json", Secret(petJson), fs.FileMode(0644))
	errorCheck(c, err)
	c.AbortWithStatus(http.StatusOK)
}

func errorCheck(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
