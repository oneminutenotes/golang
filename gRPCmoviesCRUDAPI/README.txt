Friday, 22 April 2022

Build a CRUD API with gRPC and Golang
Steps ($:represents terminal commands)
$mkdir go-movies-crud
$cd go-movies-crud
$mkdir protos
$touch moviesapp.proto
Update movies.proto
$cd ..
$go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$go mod init moviesapp.com/grpc
$export PATH="$PATH:$(go env GOPATH)/bin"
$protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/moviesapp.proto
$go mod tidy
$touch server/main.go
$touch client/main.go
Update server/main.go
$gofmt -s -w server/main.go
$go run server/main.go
#go run client/main.go

