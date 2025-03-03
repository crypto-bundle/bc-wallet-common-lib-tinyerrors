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
	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type marshaller struct{}

func (m *marshaller) MarshallSignResponse(walletUUID string, signedData []byte) *SignDataResponse {
	return &SignDataResponse{
		WalletUUID: walletUUID,
		SignedData: signedData,
	}
}

func (m *marshaller) MarshallErrorResponse(err error) error {
	// example of extract error status code from error via tinyerrors.ErrorGetCode function
	switch TinyErrStatusCode(tinyerrors.ErrorGetCode(err)) {
	case TinyErrCodeAccessTokenNotRegistered:
		return m.marshallErrorTokenNotFound(err)

	case TinyErrCodeAccessTokenExpired:
		return m.marshallErrorTokenExpired(err)

	case TinyErrCodeUnableToSign:
		return m.marshallErrorUnableToSign(err)

	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (m *marshaller) marshallErrorTokenNotFound(err error) error {
	respErrStatus, _ := status.New(codes.PermissionDenied, "Token not found").
		WithDetails(&errdetails.ErrorInfo{
			Reason: ErrorReasons_ACCESS_TOKEN_NOT_FOUND.String(),
			Domain: Domain,
			Metadata: map[string]string{
				"internal_error_status_code": TinyErrCodeAccessTokenNotRegistered.Itoa(),
				"internal_error_status_text": TinyErrCodeAccessTokenNotRegistered.String(),
				"error_message":              err.Error(),
			},
		})

	return respErrStatus.Err()
}

func (m *marshaller) marshallErrorTokenExpired(err error) error {
	respErrStatus, _ := status.New(codes.PermissionDenied, "Token expired").
		WithDetails(&errdetails.ErrorInfo{
			Reason: ErrorReasons_ACCESS_TOKEN_EXPIRED.String(),
			Domain: Domain,
			Metadata: map[string]string{
				"internal_error_status_code": TinyErrCodeAccessTokenExpired.Itoa(),
				"internal_error_status_text": TinyErrCodeAccessTokenExpired.String(),
				"error_message":              err.Error(),
			},
		})

	return respErrStatus.Err()
}

func (m *marshaller) marshallErrorUnableToSign(err error) error {
	respErrStatus, _ := status.New(codes.Internal, "Unable to sign data").
		WithDetails(&errdetails.ErrorInfo{
			Reason: ErrorReasons_SIGNATURE_FLOW_FAILED.String(),
			Domain: Domain,
			Metadata: map[string]string{
				"internal_error_status_code": TinyErrCodeUnableToSign.Itoa(),
				"internal_error_status_text": TinyErrCodeUnableToSign.String(),
				"error_message":              err.Error(),
			},
		})

	return respErrStatus.Err()
}
