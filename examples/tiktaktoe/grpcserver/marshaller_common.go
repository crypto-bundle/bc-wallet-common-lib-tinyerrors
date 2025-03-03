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

package grpcserver

import (
	"errors"

	"tiktaktoe/app"
	pb "tiktaktoe/pkg"
	"tiktaktoe/types"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	validate "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type commonMarshaller struct {
	serviceName string
}

func (m *commonMarshaller) getGRPCStatusByErrorCode(errCode tinyerrors.TinyErrCode) codes.Code {
	switch errCode {
	case types.TinyErrCodeMatchAlreadyRegistered:
		return codes.AlreadyExists
	case types.TinyErrCodeMatchNotRegistered:
		return codes.NotFound
	case types.TinyErrFieldAlreadyTaken:
		return codes.InvalidArgument
	case types.TinyErrFieldPositionOutOfMap:
		return codes.InvalidArgument
	case types.TinyErrAllFieldsTaken:
		return codes.InvalidArgument
	case types.TinyErrNotYourMovementOrder:
		return codes.InvalidArgument
	case types.TinyErrorUnableToCreateBattlefield:
		return codes.Internal
	case types.TinyErrorAccessTokensNotFound:
		return codes.PermissionDenied
	case types.TinyErrorAccessTokenAlreadyExists:
		return codes.InvalidArgument
	default:
		return codes.Unknown
	}
}

func (m *commonMarshaller) MarshallError(handlerName string, err error) error {
	_, isValidationError := tinyerrors.ErrorCodeIsOneOf(err,
		types.TinyErrorValidationFailed, types.TinyErrorValidationInternal)
	if isValidationError {
		return m.marshallValidationError(handlerName, err)
	}

	return m.marshallCommonError(handlerName, err)
}

func (m *commonMarshaller) marshallCommonError(handlerName string,
	err error,
) error {
	errCode := tinyerrors.ErrorGetCode(err)
	gRPCStatusCode := m.getGRPCStatusByErrorCode(errCode)

	respStatus, _ := status.New(gRPCStatusCode, errCode.String()).
		WithDetails(&errdetails.ErrorInfo{
			Reason: err.Error(),
			Domain: app.ApplicationDomain.WithSubDomains(handlerName, m.serviceName),
			Metadata: map[string]string{
				"error_status_code": errCode.Itoa(),
				"error_status_i18":  errCode.I18n(),
			},
		})

	return respStatus.Err()
}

func (m *commonMarshaller) marshallValidationError(handlerName string, err error) error {
	switch tinyerrors.ErrorGetCode(err) {
	case types.TinyErrorValidationInternal:
		return m.marshallValidationInternalError(handlerName, err)

	case types.TinyErrorValidationFailed:
		var errs validate.Errors

		isGoValidatorErrors := errors.As(err, &errs)
		if !isGoValidatorErrors {
			return m.marshallShortValidationError(handlerName, err)
		}

		return m.marshallFullValidationError(handlerName, errs)

	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (m *commonMarshaller) marshallFullValidationError(handlerName string, errs validate.Errors) error {
	respErrStatus := status.New(codes.InvalidArgument, types.TinyErrorValidationFailed.String())

	for _, e := range errs {
		var goValErr validate.Error

		isGoValError := errors.As(e, &goValErr)
		if !isGoValError {
			continue
		}

		respErrStatus, _ = respErrStatus.WithDetails(&errdetails.ErrorInfo{
			Reason: goValErr.Error(),
			Domain: app.ApplicationDomain.WithSubDomains(handlerName, m.serviceName),
			Metadata: map[string]string{
				"validation_error_i18":     goValErr.Code(),
				"validation_error_message": goValErr.Message(),
			},
		})
	}

	return respErrStatus.Err()
}

func (m *commonMarshaller) marshallValidationInternalError(handlerName string, err error) error {
	respStatus, _ := status.New(codes.InvalidArgument, types.TinyErrorValidationInternal.String()).
		WithDetails(&errdetails.ErrorInfo{
			Reason: err.Error(),
			Domain: app.ApplicationDomain.WithSubDomains(handlerName, m.serviceName),
			Metadata: map[string]string{
				"error_status_code": types.TinyErrorValidationInternal.Itoa(),
				"error_status_i18":  types.TinyErrorValidationInternal.I18n(),
			},
		})

	return respStatus.Err()
}

func (m *commonMarshaller) marshallShortValidationError(handlerName string, err error) error {
	respStatus, _ := status.New(codes.InvalidArgument, types.TinyErrorValidationFailed.String()).
		WithDetails(&errdetails.ErrorInfo{
			Reason: err.Error(),
			Domain: app.ApplicationDomain.WithSubDomains(handlerName, m.serviceName),
			Metadata: map[string]string{
				"error_status_code": types.TinyErrorValidationFailed.Itoa(),
				"error_status_i18":  types.TinyErrorValidationFailed.I18n(),
			},
		})

	return respStatus.Err()
}

func newCommonMarshaller() *commonMarshaller {
	return &commonMarshaller{
		serviceName: pb.GameApi_ServiceDesc.ServiceName,
	}
}
