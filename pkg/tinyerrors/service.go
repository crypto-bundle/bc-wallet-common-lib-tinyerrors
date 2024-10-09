package tinyerrors

import (
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

func (s *FmtService) ErrorGetCode(_ error) int {
	return -1
}

func (s *FmtService) ErrWithCode(err error, code int) error {
	return s.ErrorWithCode(err, code)
}

func (s *FmtService) ErrorWithCode(err error, _ int) error {
	return err
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
