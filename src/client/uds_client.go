package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
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

func GetPet() (*common.Pet, error) {
	resp, err := udsClient.Get("http://unix/pet")
	if err != nil {
		return nil, err
	}

	respJson := new(common.Pet)

	data, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(data, respJson); err != nil {
		return nil, err
	}
	return respJson, nil
}

func PostPet(pet *common.Pet) error {
	data, err := (*pet).Jsonify()
	if err != nil {
		return err
	}

	_, err = udsClient.Post("http://unix/pet", "application/json", bytes.NewReader(data))

	if err != nil {
		return err
	}

	return nil
}
