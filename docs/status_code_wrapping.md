# Wrapping business logic status code



## Usage examples

```go
package main

import (
	"context"
	"errors"
	"strconv"

	"bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	TinyErrCodeAccessTokenNotRegistered        = 7001
	TinyErrCodeUnableToSign                    = 7002
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type tokenStorageService interface {
	IsTokenExistsByUUID(tokenUUID string, walletUUID string) (bool, error)
}

type signerService interface {
	SignData(dataForSign []byte, walletUUID string) ([]byte, error)
}

type pbResponse struct {
	WalletUUID string
	SignedData []byte
}

type processor struct {
	// tokenStorage - implementation of token storage service-component from another package... 
	tokenStorage tokenStorageService
	// signer - implementation of data signer service-component from another package...
	signer signerService
}

func (p *processor) ProcessSignRequest(ctx context.Context,
	tokeUUID string,
	walletUUID string,
	dataForSign []byte,
) (*pbResponse, error) {
	signedData, err := p.processSignRequest(ctx, tokeUUID, walletUUID, dataForSign)
	if err != nil {
		switch tinyerrors.ErrorGetCode(err) {
		case TinyErrCodeAccessTokenNotRegistered:
			return nil, status.New(codes.PermissionDenied, "Token not found").WithDetails(
				&errdetails.ErrorInfo{
					Reason: pb.ErrorReason_TOKEN_NOT_FOUND.String(), 
					Domain: domain,
					Metadata: map[string]string {
						"internal_error_code": strconv.Itoa(TinyErrCodeAccessTokenNotRegistered),
                    },
				},
			)
			
		case TinyErrCodeUnableToSign:
			return nil, status.New(codes.Internal, "Unable to sign data").WithDetails(
				&errdetails.ErrorInfo{
					Reason: pb.ErrorReason_SIGNATURE_FLOW_FAILED.String(),
					Domain: domain,
					Metadata: map[string]string {
						"internal_error_code": strconv.Itoa(TinyErrCodeUnableToSign),
					},
				},
			)
		}
		
    }
	
	resp := &pbResponse{
		WalletUUID: walletUUID,
		SignedData: signedData,
	}
	
	return resp, nil
}

func (p *processor) processSignRequest(ctx context.Context,
	tokeUUID string,
	walletUUID string,
	dataForSign []byte,
) ([]byte, error) {
	isExists, err := p.tokenStorage.IsTokenExistsByUUID(tokeUUID, walletUUID)
	if err != nil {
		// wrapping errors from another package 
		return nil, tinyerrors.ErrorOnly(err)
	}

	if !isExists {
		// wrapping errors from another package with internal status code
		return nil, tinyerrors.ErrorWithCode(ErrTokenNotFound, TinyErrCodeAccessTokenNotRegistered)
	}

	signedData, err := p.signer.SignData(dataForSign, walletUUID)
	if err != nil {
		// pseudo-wrapping errors from another package. Because this error already has status-code
		return nil, tinyerrors.ErrorNoWrap(err)
	}
	
	return signedData, nil
}
```