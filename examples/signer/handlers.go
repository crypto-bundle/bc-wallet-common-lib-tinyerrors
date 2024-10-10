/*
 *
 *
 * MIT NON-AI License
 *
 * Copyright (c) 2024-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package main

import (
	"context"
	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
)

type grpcService struct {
	UnimplementedSignerApiServer

	signHandlerSvc *signRequestHandle
}

func NewGrpcService(signHandlerSvc *signRequestHandle) *grpcService {
	return &grpcService{
		signHandlerSvc: signHandlerSvc,
	}
}

func (h *grpcService) SignData(ctx context.Context, req *SignDataRequest) (*SignDataResponse, error) {
	return h.signHandlerSvc.ProcessRequest(ctx, req.TokenUUID, req.WalletUUID, req.DataForSign)
}

type signRequestHandle struct {
	// tokenStorageSvc - implementation of token storage service-component from another package...
	tokenStorageSvc tokenStorageService
	// signerSvc - implementation of data signer service-component from another package...
	signerSvc     signerService
	marshallerSvc marshallerService
}

func (p *signRequestHandle) ProcessRequest(ctx context.Context,
	tokeUUID string,
	walletUUID string,
	dataForSign []byte,
) (*SignDataResponse, error) {
	signedData, err := p.processSignRequest(ctx, tokeUUID, walletUUID, dataForSign)
	if err != nil {
		return nil, p.marshallerSvc.MarshallErrorResponse(err)
	}

	return p.marshallerSvc.MarshallSignResponse(walletUUID, signedData), nil
}

func (p *signRequestHandle) processSignRequest(_ context.Context,
	tokeUUID string,
	walletUUID string,
	dataForSign []byte,
) ([]byte, error) {
	isExists, err := p.tokenStorageSvc.IsTokenExistsByUUID(tokeUUID, walletUUID)
	if err != nil {
		// example of wrapping errors from another package, no wrapping error status code here - internal error
		return nil, tinyerrors.ErrorOnly(err)
	}

	if !isExists {
		// wrapping errors from another package with internal status code
		return nil, tinyerrors.ErrorWithCode(ErrTokenNotFound, TinyErrCodeAccessTokenNotRegistered.Int())
	}

	isTokenExpired, err := p.tokenStorageSvc.IsTokenExpiredByUUID(tokeUUID, walletUUID)
	if err != nil {
		// example of wrapping errors from another package, no wrapping error status code here - internal error
		return nil, tinyerrors.ErrorOnly(err)
	}

	if !isTokenExpired {
		// wrapping errors from another package with internal status code
		return nil, tinyerrors.ErrorWithCode(ErrTokenExpired, TinyErrCodeAccessTokenExpired.Int())
	}

	signedData, err := p.signerSvc.SignData(dataForSign, walletUUID)
	if err != nil {
		// pseudo-wrapping errors from another package. Because this error already has status-code
		return nil, tinyerrors.ErrorNoWrap(err)
	}

	return signedData, nil
}
