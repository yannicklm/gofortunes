all: build

build:
	go build gf-add.go
	go build gf-get.go
	go build gf-srv.go
	go build gf-client.go

check:
	go test gofortunes/fortunes

clean:
	go clean
