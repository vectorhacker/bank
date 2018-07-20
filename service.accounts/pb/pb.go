package accounts

//go:generate protoc -I . command.proto --go_out=plugins=grpc:.
//go:generate protoc -I . query.proto --go_out=plugins=grpc:.
