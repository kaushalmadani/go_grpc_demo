protoc --go_out=./output modules/*.proto


protoc --go_out=./output --go-grpc_out=./output ./modules/role/*.proto

protoc --go_out=./output **--go-grpc_opt=require_unimplemented_servers=false** --go-grpc_out=./output ./modules/user/*.proto