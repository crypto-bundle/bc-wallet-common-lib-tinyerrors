package tinyerrors

import (
	"errors"
	"testing"
)

func TestErrorFormatting(t *testing.T) {
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
