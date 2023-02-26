package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/schollz/croc/v9/src/comm"
	"github.com/schollz/croc/v9/src/tcp"
)

type Relay struct {
	host     string
	port     string
	password string
}

func (r *Relay) JoinRoom(shared_secret string) (*comm.Comm, error) {
	ipaddr, err := net.LookupHost(r.host)
	if err != nil || len(ipaddr) < 1 {
		log.Fatal(err)
	}
	comm, _, _, err := tcp.ConnectToTCPServer(net.JoinHostPort(ipaddr[0], r.port), r.password, shared_secret)
	if err != nil {
		return nil, err
	}
	return comm, err
}

func NewRelay() *Relay {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env file")
	}
	return &Relay{
		readEnv("RELAY_HOST"),
		readEnv("RELAY_PORT"),
		readEnv("RELAY_PASS"),
	}
}

func readEnv(envname string) string {
	value := os.Getenv(envname)
	if value == "" {
		log.Fatal(envname + " not set in .env")
	}
	return value
}
