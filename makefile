exefile = service-test
revision = $(shell git log | head -n 1 | cut -f 2 -d ' ')

clean: 
	rm -f $(exefile)

build: 
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X main.revision=$(revision)"

gen:
	protoc --proto_path=idl \
	--go_out=plugins=grpc:$$GOPATH/src \
	idl/dto/dto.proto

	protoc --proto_path=idl \
	-I$$GOPATH/src/github.com/vanhieudn01/service-test \
	--go_out=plugins=grpc:./grpc-gen \
	idl/service.proto