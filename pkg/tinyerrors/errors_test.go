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
