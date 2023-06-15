

> goctl rpc protoc --proto_path=.:$GOPATH/src --go_out=./user_query --go_opt=paths=source_relative --go-grpc_out=./user_query --go-grpc_opt=paths=source_relative --zrpc_out=. --style=go_zero ./grpc_user_query.proto