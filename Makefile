all: build check

build:
	go build gf-add.go
	go build gf-get.go

check:
	go test gofortunes/fortunes

clean:
	go clean
