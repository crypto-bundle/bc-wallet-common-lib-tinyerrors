module tiktaktoe

go 1.22

replace github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors => ../..

require (
	github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
