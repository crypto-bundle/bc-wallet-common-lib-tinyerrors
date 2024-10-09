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

package tinyerrors

import (
	"errors"
	"fmt"
	"strings"
)

type FmtService struct{}

func (s *FmtService) ErrorNoWrapOrNil(err error) error {
	if err != nil {
		return err
	}

	return nil
}

func (s *FmtService) ErrGetCode(err error) int {
	return s.ErrorGetCode(err)
}

func (s *FmtService) ErrorGetCode(err error) int {
	var ccErr *codeContainsError
	if errors.As(err, &ccErr) {
		return ccErr.code
	}

	return -1
}

func (s *FmtService) ErrWithCode(err error, code int) error {
	return s.ErrorWithCode(err, code)
}

func (s *FmtService) ErrorWithCode(err error, code int) error {
	return &codeContainsError{
		Err:  err,
		code: code,
	}
}

func (s *FmtService) ErrNoWrap(err error) error {
	return s.ErrorNoWrap(err)
}

func (s *FmtService) ErrorNoWrap(err error) error {
	if err == nil {
		return nil
	}

	return err
}

// ErrorOnly combines given error with details, WITHOUT function name...
func (s *FmtService) ErrorOnly(err error, details ...string) error {
	if err == nil {
		return nil
	}

	if len(details) == 0 {
		return err
	}

	return fmt.Errorf("%w -> %s", err, strings.Join(details, ", "))
}

// Error combines given error with details and finishes with caller func name...
func (s *FmtService) Error(err error, details ...string) error {
	return s.ErrorOnly(err, details...)
}

// NewError returns error by combining given details and finishes with caller func name...
//
//nolint:err113
func (s *FmtService) NewError(details ...string) error {
	return fmt.Errorf("%s", strings.Join(details, ", "))
}

// NewErrorf returns error by combining given details and finishes with caller func name, printf formatting...
//
//nolint:err113
func (s *FmtService) NewErrorf(format string, args ...interface{}) error {
	return fmt.Errorf(
		"%s",
		strings.Join([]string{fmt.Sprintf(format, args...)}, ", "),
	)
}

// Errorf combines given error with details and finishes with caller func name, printf formatting...
func (s *FmtService) Errorf(err error, format string, args ...interface{}) error {
	return s.ErrorOnly(err, fmt.Sprintf(format, args...))
}
