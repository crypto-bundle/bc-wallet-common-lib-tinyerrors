package tinyerrors

//nolint:gochecknoglobals // it's ok
var DefaultErrorFormatterSvc = errFmtSvc

//nolint:gochecknoglobals // it's ok
var errFmtSvc ErrorFormatterService = new(FmtService)

// Default returns the standard error formatted service-component...
func Default() ErrorFormatterService { return errFmtSvc }

// SetDefault re-assign default errFmtSvc variable(default error formatter service) with by passed argument...
func SetDefault(fmtSvc ErrorFormatterService) {
	errFmtSvc = fmtSvc
}

func ErrorNoWrapOrNil(err error) error {
	if err != nil {
		return err
	}

	return nil
}

func ErrorWithCode(err error, code int) error {
	return errFmtSvc.ErrorWithCode(err, code)
}

func ErrWithCode(err error, code int) error {
	return errFmtSvc.ErrWithCode(err, code)
}

func ErrorGetCode(err error) int {
	return errFmtSvc.ErrorGetCode(err)
}

func ErrGetCode(err error) int {
	return errFmtSvc.ErrGetCode(err)
}

func ErrorNoWrap(err error) error {
	return errFmtSvc.ErrorNoWrap(err)
}

func ErrNoWrap(err error) error {
	return errFmtSvc.ErrNoWrap(err)
}

func ErrorOnly(err error, details ...string) error {
	return errFmtSvc.ErrorOnly(err, details...)
}

func Error(err error, details ...string) error {
	return errFmtSvc.Error(err, details...)
}

func Errorf(err error, format string, args ...interface{}) error {
	return errFmtSvc.Errorf(err, format, args...)
}

func NewError(details ...string) error {
	return errFmtSvc.NewError(details...)
}

func NewErrorf(format string, args ...interface{}) error {
	return errFmtSvc.NewErrorf(format, args...)
}
