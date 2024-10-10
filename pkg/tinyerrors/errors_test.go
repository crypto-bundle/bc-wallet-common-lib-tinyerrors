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
	"testing"
)

func TestErrorFormatting(t *testing.T) {
	t.Run("ErrorWithCode", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrorWithCode(errors.New("test error"), 15)
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("ErrWithCode", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrWithCode(errors.New("test error"), 25)
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("ErrorGetCode", func(t *testing.T) {
		const (
			expectedResult = "test error"
			expectedCode   = 42
		)

		err := ErrorWithCode(errors.New("test error"), expectedCode)
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}

		if code := ErrorGetCode(err); code != expectedCode {
			t.Errorf("error code not equal with expected. current: %d, expected: %d",
				code, expectedCode)
		}
	})

	t.Run("ErrGetCode", func(t *testing.T) {
		const (
			expectedResult = "test error"
			expectedCode   = 69
		)

		err := ErrWithCode(errors.New("test error"), expectedCode)
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}

		if code := ErrGetCode(err); code != expectedCode {
			t.Errorf("error code not equal with expected. current: %d, expected: %d",
				code, expectedCode)
		}
	})

	t.Run("ErrorNoWrap", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrorNoWrap(errors.New("test error"))
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("ErrorNoWrapOrNil", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrorNoWrapOrNil(nil)
		if err != nil {
			t.Errorf("error should be nil. current: %e",
				err)
		}

		errTwo := ErrorNoWrapOrNil(errors.New(expectedResult))
		if errTwo.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				errTwo.Error(), expectedResult)
		}
	})

	t.Run("ErrorNoWrapOrNil", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrNoWrapOrNil(nil)
		if err != nil {
			t.Errorf("error should be nil. current: %e",
				err)
		}

		errTwo := ErrNoWrapOrNil(errors.New(expectedResult))
		if errTwo.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				errTwo.Error(), expectedResult)
		}
	})

	t.Run("ErrNoWrap", func(t *testing.T) {
		const expectedResult = "test error"

		err := ErrNoWrap(errors.New("test error"))
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("error only", func(t *testing.T) {
		const expectedResult = "test error -> abc"

		err := ErrorOnly(errors.New("test error"), "abc")
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("common error", func(t *testing.T) {
		const expectedResult = "test error -> efg"

		err := Error(errors.New("test error"), "efg")
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("new error", func(t *testing.T) {
		const expectedResult = "test, error - some text"

		err := NewError("test", "error - some text")
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("new errorf", func(t *testing.T) {
		const expectedResult = "test arg - with value 100500"

		err := NewErrorf("test %s - with value %d", "arg", 100500)
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("common errorf", func(t *testing.T) {
		const expectedResult = "test error -> test arg"

		err := Errorf(errors.New("test error"), "test %s", "arg")
		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})

	t.Run("error no wrap", func(t *testing.T) {
		const expectedResult = "test error"
		var errExpected = errors.New(expectedResult)

		err := ErrorNoWrap(errExpected)
		if !errors.Is(err, errExpected) {
			t.Errorf("error not equal with expected %s, %s", err.Error(), errExpected.Error())
		}

		if err.Error() != expectedResult {
			t.Errorf("error text not equal with expected. current: %s, expected: %s",
				err.Error(), expectedResult)
		}
	})
}
