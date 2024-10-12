# bc-wallet-common-lib-tinyerrors

## Description
Small error-wrapper library. This library used as default error wrapper in applications or another libraries. 
Direct usage of this library as target dependency is allowed by crypto-bundle code-style. 
Crypto-bundle code-style allows direct usage in applications and another libraries. This allowance is important because 
main error-formatter library - [bc-wallet-common-lib-errors](https://github.com/crypto-bundle/bc-wallet-common-lib-errors) - 
does not allow usage target usage in another libraries, just in application.  

The ideal future for this library is to be at version v1.0.0 forever and never change.

Repository `bc-wallet-common-lib-tinyerrors` instantly in read-only mode. Code of this library should never change.
Restriction on read-only mode can be removed only if you need change description of main error-formatter interface 
and standard implementation. 

In case of code change first you need - disable read-only(archive) mode in repository settings.

## Error formatter service interface
Main `errfmt` interface described in [common.go](/pkg/tinyerrors/common.go) file. All implementations must follow interface requirements.

The interface requires the implementation of the these functions:
* `ErrorWithCode(err error, code int) error` - error wrapper function which not modified error text, just for storing internal error code in error. 
This function fully depend on implementation of error formatter service.
* `ErrWithCode(err error, code int) error` - it's just alias for `ErrorWithCode`. 
* `ErrorGetCode(err error) int` - function for extract code from error. Function return -1 if code not stored in error.
Behavior of this function depends on implementation.
* `ErrGetCode(err error) int` - just alias for `ErrorGetCode`.
* `ErrorNoWrap(err error) error` - function pseudo-wrapper. This function **should not** modify passed error - this rule independent for implementation.
All implementation of `ErrorNoWrap` function **must follow this rule**.
* `ErrNoWrap(err error) error` - just alias for `ErrorNoWrap`.
* `ErrorNoWrapOrNil(err error) error` - same with `ErrorNoWrap` function. Please don't use this function - deprecated.
* `ErrNoWrapOrNil(err error) error` - same with `ErrorNoWrapOrNil` function.
* `ErrorOnly(err error, details ...string) error` - error wrapper function. This function wrap origin error in new error.
Behavior of this function depends on implementation. Passed additional `details` **must** be added to the error text.
* `Error(err error, details ...string) error` - usually same with `ErrorOnly`. This function also error wrapper function,
but behavior depends on implementation of errfmt service-component.
* `Errorf(err error, format string, args ...interface{}) error` - error wrapper function. This function wrapped origin error, 
but error message will be formatted in passed format. Specific behavior depends on implementation of errfmt service-component.
* `NewError(details ...string) error` - function to create a new error with additional `details`. 
Passed additional `details` **must be** added to the error text. Behavior depends on implementation.
* `NewErrorf(format string, args ...interface{}) error` - function to create a new error with specific format. 
Error text format pass to function as argument. Compiled by specific format message **should** exit in error text. Behavior depends on implementation.

Default implementation described in [service.go](/pkg/tinyerrors/service.go) and assigned to public variable - `DefaultErrorFormatterSvc`.
If you need u can change default error formatter service by calling `SetDefault(fmtSvc ErrorFormatterService)` function. 
If the capabilities of the standard formatter are not enough for you - use the implementations from the 
[bc-wallet-common-lib-errors](https://github.com/crypto-bundle/bc-wallet-common-lib-errors) repository, all changes that expand the capabilities and 
functionality of the formatter must be committed there.

### ErrorWithCode, ErrWithCode, ErrorGetCode, ErrGetCode
Main purpose of these functions - wrap business logic error status-code in error. This use-case relevant as communication option between application layers -
you don't need use `errors.Is` and import errors from another application layers and sub-package, all you need - it can just compare `int` values. 
This function fully depend on implementation of error formatter service. Standard implementation of `tinyerrors` package, presented in [errors.go](/pkg/tinyerrors/errors.go),
storing code value in non-exported struct `codeContainsError` [types.go](/pkg/tinyerrors/types.go), which wrap origin error.

Full information about these functions with programming code examples you can see in [docs/status_code_wrapping.md](/docs/status_code_wrapping.md) file.
Also, examples of error status-code wrapping presented in:
* [signer application](/examples/signer) - gRPC-server application for sign user-data, witch contains example of error status-code wrapping.
  * [signer/marshaller.go](/examples/signer/marshaller.go)
  * [signer/handlers.go](/examples/signer/handlers.go)
* [unit tests](/pkg/tinyerrors/errors_test.go) - Unit-tests for all error formatter methods

### ErrorNoWrap, ErrNoWrap, ErrorNoWrapOrNil, ErrNoWrapOrNil

### ErrorOnly, Error
Typical usage of these functions - wrap existing error to new error with additional details.

More information about errors wrap and re-wrap you can read in [docs/error_wrap.md](/docs/error_wrap.md). 
Error wrapping is simply to use, but be careful with re-wrap already wrapped error.

You can see examples of error wrapping in example application `coinflip`, which located in:
* [coinflip application](/examples/coinflip) - HTTP-server application which make coin-flip on each HTTP-request.
  * [coinflip/presenter.go](/examples/coinflip/presenter.go)
  * [coinflip/handlers.go](/examples/coinflip/handlers.go)

Also, examples of error wrapping located in [unit-tests](/pkg/tinyerrors/errors_test.go)

Format for new error message depend on implementation of error fmt service. Another implementation of these functions and  you can see in
[bc-wallet-common-lib-errors](https://github.com/crypto-bundle/bc-wallet-common-lib-errors)

### NewError, NewErrorf

## Contributors

* Maintainer - [@gudron (Alex V Kotelnikov)](https://github.com/gudron)

## Licence

**bc-wallet-common-lib-tinyerrors** is licensed the [MIT NON-AI](./LICENSE) License.