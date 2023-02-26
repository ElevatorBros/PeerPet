package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"

	"peer.pet/src/common"
)

type ReadPetJson struct {
	Pet common.Pet
}

var udsClient = http.Client{
	Transport: &http.Transport{
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("unix", "http.sock")
		},
	},
}

func ReadPet() (*common.Pet, error) {
	resp, err := udsClient.Get("http://unix/pet")
	if err != nil {
		return nil, err
	}
	respJson := &ReadPetJson{}
	var data []byte
	_, err = resp.Body.Read(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, respJson); err != nil {
		return nil, err
	}
	return &respJson.Pet, nil
}

func PostPet(pet *common.Pet) error {
	petJson, err := pet.Jsonify()
	if err != nil {
		return err
	}
	reader := bytes.NewReader(petJson)
	if err != nil {
		return err
	}
    resp, err := udsClient.Post("http://unix/pet", "application/json", reader)
	if err != nil {
		return err
	}
	var data []byte
	if _, err := resp.Body.Read(data); err != nil {
		return err
	}
	return nil
}
