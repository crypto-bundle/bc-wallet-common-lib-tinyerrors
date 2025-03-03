module signer

go 1.22

replace github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors => ../..

require (
	github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors v0.0.0-00010101000000-000000000000
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
