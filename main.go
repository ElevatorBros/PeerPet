package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"peer.pet/src/client"
	"peer.pet/src/common"
	"peer.pet/src/server"
)

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
	r.GET("/pet", server.GetPet)
	r.PUT("/pet", server.PostPet)

	if err := r.RunUnix("http.sock"); err != nil {
		panic(err)
	}

}

func Client() {
	//creates pet named john
	pet := common.NewPet("pet")
	//makes sure storing directory exists and returns path
	//folder_path := server.CreateDataDir()

	//serializes pet to json
	err := client.PostPet(pet)
	if err != nil {
		panic(err)
	}
	client.RunGUI()

	//reads stored json files to pet array
	//pets := ReadPets()

	////prints pet array data
	//for _, thing := range pets {
	//	thing.Print()
	//}
	client.RunGUI()
}
