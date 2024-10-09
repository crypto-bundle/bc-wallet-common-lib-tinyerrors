package tinyerrors

//nolint:gochecknoglobals // it's ok
var DefaultErrorFormatterSvc = errFmtSvc

//nolint:gochecknoglobals // it's ok
var errFmtSvc ErrorFormatterService = new(FmtService)

// Default returns the standard error formatted service-component...
func Default() ErrorFormatterService { return DefaultErrorFormatterSvc }

// SetDefault re-assign default errFmtSvc variable(default error formatter service) with by passed argument...
func SetDefault(fmtSvc ErrorFormatterService) {
	DefaultErrorFormatterSvc = fmtSvc
}

// ErrorNoWrapOrNil - pseudo-wrapper function which return error in origin state, or nil if origin error is nil.
// Deprecated: Please don't use this function - candidate for deletion. Use ErrorNoWrap or ErrNoWrap instead of it.
// Function removed from interface ErrorFormatterService requirements...
func ErrorNoWrapOrNil(err error) error {
	if err != nil {
		return err
	}

	return nil
}

func ErrorWithCode(err error, code int) error {
	return DefaultErrorFormatterSvc.ErrorWithCode(err, code)
}

func ErrWithCode(err error, code int) error {
	return DefaultErrorFormatterSvc.ErrWithCode(err, code)
}

func ErrorGetCode(err error) int {
	return DefaultErrorFormatterSvc.ErrorGetCode(err)
}

func ErrGetCode(err error) int {
	return DefaultErrorFormatterSvc.ErrGetCode(err)
}

func ErrorNoWrap(err error) error {
	return DefaultErrorFormatterSvc.ErrorNoWrap(err)
}

func ErrNoWrap(err error) error {
	return DefaultErrorFormatterSvc.ErrNoWrap(err)
}

func ErrorOnly(err error, details ...string) error {
	return DefaultErrorFormatterSvc.ErrorOnly(err, details...)
}

func Error(err error, details ...string) error {
	return DefaultErrorFormatterSvc.Error(err, details...)
}

func Errorf(err error, format string, args ...interface{}) error {
	return DefaultErrorFormatterSvc.Errorf(err, format, args...)
}

func NewError(details ...string) error {
	return DefaultErrorFormatterSvc.NewError(details...)
}

func NewErrorf(format string, args ...interface{}) error {
	return DefaultErrorFormatterSvc.NewErrorf(format, args...)
}
