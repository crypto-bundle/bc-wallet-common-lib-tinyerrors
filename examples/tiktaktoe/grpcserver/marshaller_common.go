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
	"tiktaktoe/types"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	"github.com/asaskevich/govalidator/v11"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type commonMarshaller struct {
}

func (m *commonMarshaller) marshallValidationError(handlerName string, err error) error {
	switch tinyerrors.ErrorGetCode(err) {
	case types.TinyErrorValidationFailed:
		var errs govalidator.Errors

		isGoValidatorErrors := errors.As(err, &errs)
		if !isGoValidatorErrors {
			return m.marshallShortValidationError(handlerName, err)
		}

		return m.marshallFullValidationError(handlerName, errs)

	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (m *commonMarshaller) marshallFullValidationError(handlerName string, errs govalidator.Errors) error {
	metaData := make(map[string]string, len(errs)+2)
	metaData["internal_error_status_code"] = types.TinyErrorValidationFailed.Itoa()
	metaData["internal_error_status_i18"] = types.TinyErrorValidationFailed.I18n()

	respErrStatus := status.New(codes.InvalidArgument, types.TinyErrorValidationFailed.String())

	for _, e := range errs {
		var goValErr govalidator.Error

		isGoValError := errors.As(e, &goValErr)
		if !isGoValError {
			continue
		}

		respErrStatus, _ = respErrStatus.WithDetails(&errdetails.ErrorInfo{
			Reason:   goValErr.Error(),
			Domain:   app.ApplicationDomain.WithSubDomains(goValErr.Name, handlerName),
			Metadata: metaData,
		})

		metaData[goValErr.Name] = goValErr.Error()
	}

	return respErrStatus.Err()
}

func (m *commonMarshaller) marshallShortValidationError(handlerName string, err error) error {
	respStatus, _ := status.New(codes.InvalidArgument, types.TinyErrorValidationFailed.String()).
		WithDetails(&errdetails.ErrorInfo{
			Reason: err.Error(),
			Domain: app.ApplicationDomain.WithSubDomain(handlerName),
			Metadata: map[string]string{
				"internal_error_status_code": types.TinyErrorValidationFailed.Itoa(),
				"internal_error_status_i18":  types.TinyErrorValidationFailed.I18n(),
			},
		})

	return respStatus.Err()
}
