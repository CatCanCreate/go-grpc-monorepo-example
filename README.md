# Golang gRPC monorepo example



#### Пример генерации Golang кода из proto файлов:
1) `cd shippy-service-consignment`
2) `protoc ../proto/consignment/*.proto --go_out=proto/consignment --go-grpc_out=proto/consignment --proto_path=../proto/consignment/`

#### or
1) `cd shippy-cli-consignment`
2) `protoc ../proto/consignment/*.proto --go_out=proto/consignment --go-grpc_out=proto/consignment --proto_path=../proto/consignment/`

### Разбор генерации protobuf -> golang

+ `protoc` - компилятор protobuf (`brew install protobuf`)
+ `../proto/consignment/*.proto` - путь до proto файла/файлов
+ `--go_out=plugins=grpc:proto/consignment` - компилятор proto -> golang + выходная директория
+ `--grpc_out=plugins=grpc:proto/consignment` - компилятор proto -> golang gRPC + выходная директория
+ `--proto_path=../proto/consignment/` - путь к директории с proto файлами (издержки производства)

### Линтер:
1) `cd shippy-service-consignment`
2) `golangci-lint run`


### Альтернатива cli - https://github.com/ktr0731/evans
1) `evans proto/consignment/consignment.proto -p 50051` - REPL mode


### TODO
1) Решить, хорошая ли идея использовать `pb` как domain (DTO) сущности (импортировать по всему сервису)
