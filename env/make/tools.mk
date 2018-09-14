.PHONY: protobuf
protobuf:
	protoc -Ipkg/transport/grpc/ \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	       --go_out=plugins=grpc:pkg/transport/grpc \
	       license.proto
	protoc -Ipkg/transport/grpc/ \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	       --grpc-gateway_out=logtostderr=true:pkg/transport/grpc \
	       license.proto


.PHONY: test
test: test-control test-service

.PHONY: test-control
test-control:
	go test -race -tags 'cli ctl' -v .

.PHONY: test-service
test-service:
	go test -race -v ./...

.PHONY: test-integration
test-integration:
	go test -tags integration -v ./env/test/...
