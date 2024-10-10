lint:
	golangci-lint run --config .golangci.yml -v ./...

signer: signer_proto signer_gomod

signer_proto:
	protoc -I ./examples/signer/ \
		--go_out=./examples/signer/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./examples/signer/ \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./examples/signer/ \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=source_relative \
		./examples/signer/*.proto

signer_gomod:
	go mod tidy -C ./examples/signer
	go mod vendor -C ./examples/signer

.PHONY: signer_proto lint