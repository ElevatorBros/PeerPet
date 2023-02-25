package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/schollz/croc/v9/src/comm"
	"github.com/schollz/croc/v9/src/tcp"
)

type Relay struct {
	ip_port  string
	password string
}

func (r *Relay) joinRoom(shared_secret string) (*comm.Comm, error) {
	comm, _, _, err := tcp.ConnectToTCPServer(r.ip_port, r.password, shared_secret)
	if err != nil {
		return nil, err
	}
	return comm, err
}

func NewRelay() *Relay {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env file")
	}
	ip_port := os.Getenv("RELAY_IP")
	if ip_port == "" {
		log.Fatal("RELAY_IP not net in .env")
	}
	pass := os.Getenv("RELAY_PASSWORD")
	if pass == "" {
		log.Fatal("RELAY_PASSWORD not net in .env")
	}
	return &Relay{
		ip_port,
		pass,
	}
}
