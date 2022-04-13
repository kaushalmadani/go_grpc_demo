protoc --go_out=./output modules/*.proto


protoc --go_out=./output --go-grpc_out=./output ./modules/role/*.proto

protoc --go_out=./output --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./output ./packages/user/*.proto

sudo mv ~/Downloads/protoc-gen-grpc-web-1.3.1-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
chmod +x /usr/local/bin/protoc-gen-grpc-web

protoc --go_out=./output --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./output ./packages/user/*.proto
protoc -I./packages/user --js_out=import_style=commonjs:./output/web --grpc-web_out=import_style=typescript,mode=grpcwebtext:./output/web ./packages/user/*.proto

// dashboard
protoc --go_out=./output --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./output ./packages/dashboard/*.proto
protoc -I./packages/dashboard --js_out=import_style=commonjs:./output/web --grpc-web_out=import_style=typescript,mode=grpcwebtext:./output/web ./packages/dashboard/*.proto