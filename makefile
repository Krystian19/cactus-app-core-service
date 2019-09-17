run : build
	./app

gen :
	protoc --go_out=plugins=grpc,paths=source_relative:. proto/*.proto

build : **/*.go
	go build -o app -i bin/*.go