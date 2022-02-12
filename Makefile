build:
	go build -o bin/routines routines/routines.go

protos:
	protoc --go_out=proto-files ./proto-files/*.proto

run:
	./bin/routines

build_and_run: build run