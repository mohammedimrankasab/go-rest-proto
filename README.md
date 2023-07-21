# go-rest-proto
go http server with proto file

use below command to generate the pb files

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto
