.PHONY:setup
setup:
	go mod tidy

.PHONY:proto-gen
proto-gen:
	protoc -I=./ --go_out=./ ./proto/person.proto

.PHONY:run
run:
	go run main.go

