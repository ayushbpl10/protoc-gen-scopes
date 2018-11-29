protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/example.proto
protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./  ./scope/scope.proto

packr && go build && protoc -I /usr/local/include -I  ./  --plugin=protoc-gen-scopes=protoc-gen-scopes  --scopes_out=:./example  ./example/example.proto
