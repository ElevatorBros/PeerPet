BINARY_NAME=PeerPet
FLAGS=-ldflags "-s -w"

build:
	go build ${FLAGS} -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME} $(arg)

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go get github.com/schollz/croc
	go get github.com/bennicholls/burl-E/reximage
	go get github.com/rivo/tview
	go get github.com/codegoalie/golibnotify
	go get github.com/joho/godotenv
	go get github.com/schollz/croc/v9

