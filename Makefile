.PHONY:setup
setup:
	go mod tidy

.PHONY:proto-gen
proto-gen:
	protoc -I=./src/pkg/domain --go_out=./src/pkg/domain ./src/pkg/domain/proto/*.proto

.PHONY:run
run:
	go run main.go

