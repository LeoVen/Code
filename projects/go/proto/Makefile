.PHONY: run proto

run:
	go run cmd/example/main.go

proto:
	protoc -I=. --go_out=./internal ./proto/addressbook.proto
