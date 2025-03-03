lint:
	golangci-lint run --config .golangci.yml -v ./pkg/tinyerrors

tiktaktoe: tiktaktoe_proto tiktaktoe_gomod

tiktaktoe_gomod:
	go mod tidy -C ./examples/tiktaktoe
	go mod vendor -C ./examples/tiktaktoe

tiktaktoe_proto:
	protoc -I ./examples/tiktaktoe/pkg/ \
		--go_out=./examples/tiktaktoe/pkg/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./examples/tiktaktoe/pkg/ \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./examples/tiktaktoe/pkg/ \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=source_relative \
		./examples/tiktaktoe/pkg/*.proto

coinflip: coinflip_gomod

coinflip_gomod:
	go mod tidy -C ./examples/coinflip
	go mod vendor -C ./examples/coinflip

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