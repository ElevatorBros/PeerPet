package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"peer.pet/src/client"
	"peer.pet/src/common"
	"peer.pet/src/server"
)

func MyGetPet(c *gin.Context) {
	server.GetPetGIN(c)
}

func MyPostPet(c *gin.Context) {
	server.PostPetGIN(c)
}
func main() {
	isServer := flag.Bool("s", false, "Run as server")
	flag.Parse()
	if *isServer {
		Server()
	} else {
		Client()
	}
}

func Server() {
	r := gin.Default()
	r.GET("/:pet", MyGetPet)
	// FIXED THIS
	r.POST("/:pet", MyPostPet)

	if err := r.RunUnix("http.sock"); err != nil {
		panic(err)
	}

}

func Client() {
	// FIXED THIS
	pet := common.NewPet("pet")
	err := client.PostPet(pet)
	if err == nil {

		host := false
		if len(os.Args) == 2 && os.Args[1] == "-S" {
			host = true
		}

		if host {
			client.GetHostKey()
		} else {
			client.GetHostKey()
		}

		client.EnterCombat(host)

		//client.RunGUI()

		//reads stored json files to pet array
		//pets := ReadPets()

		////prints pet array data
		//for _, thing := range pets {
		//	thing.Print()
		//}
	}
}
