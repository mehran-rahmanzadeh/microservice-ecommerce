protoc domain/proto/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=. &&
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative domain/proto/register.proto &&
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative domain/proto/login.proto &&
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative domain/proto/profile.proto