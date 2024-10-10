# bc-wallet-common-lib-tinyerrors

## Description
Small error-wrapper library. This library used as default error wrapper in applications or another libraries. 
Direct usage of this library as target dependency is allowed by crypto-bundle code-style. 
Crypto-bundle code-style allows direct usage in applications and another libraries. This allowance is important because 
main error-formatter library - [bc-wallet-common-lib-errors](https://github.com/crypto-bundle/bc-wallet-common-lib-errors) - 
does not allow usage target usage in another libraries, just in application.  

The ideal future for this library is to be at version v1.0.0 forever and never change.

Repository 'bc-wallet-common-lib-tinyerrors' instantly in read-only mode. Code of this library should never change.
Restriction on read-only mode can be removed only if you need change description of main error-formatter interface 
and standard implementation. 

In case of code change first you need - disable read-only(archive) mode in repository settings.

## Error formatter service interface
Main 'errfmt' interface described in [common.go](/pkg/tinyerrors/common.go) file.

The interface requires the implementation of the following functions:

* `ErrorWithCode(err error, code int) error` - error wrapper function which not modified error text, just for storing internal error code. 
This function fully depend on implementation of error formatter service. In standard error formatter service-component this function has not storing code value,
just return origin error.
* `ErrWithCode(err error, code int) error` - it's just alias for `ErrorWithCode`. 
* `ErrorGetCode(err error) int` - function for extract code from error. Function return -1 if code not stored in error.
Behavior of this function depends on implementation.
* `ErrGetCode(err error) int` - just alias for `ErrorGetCode`.
* `ErrorNoWrap(err error) error` - function pseudo-wrapper. This function **should not** modify passed error - this rule independent for implementation.
All implementation of `ErrorNoWrap` function **must follow this rule**.
* `ErrNoWrap(err error) error` - just alias for `ErrorNoWrap`.
* `ErrorNoWrapOrNil(err error) error` - same with `ErrorNoWrap` function. Please dont use this function - deprecated.
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

### ErrorWithCode, ErrWithCode, ErrorGetCode, ErrGetCode

### ErrorNoWrap, ErrNoWrap, ErrorNoWrapOrNil, ErrNoWrapOrNil

### ErrorOnly

### Error

### NewError

### NewErrorf

## Contributors

* Maintainer - [@gudron (Alex V Kotelnikov)](https://github.com/gudron)

## Licence

**bc-wallet-common-lib-tinyerrors** is licensed the [MIT NON-AI](./LICENSE) License.