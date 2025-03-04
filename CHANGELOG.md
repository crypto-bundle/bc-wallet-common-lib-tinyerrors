# Change Log

## [v0.0.3] - 04.03.2025
### Changed
* Added support of Go 1.23

## [v0.0.2] - 03.03.2025
### Fix
* Fixed Unwrap method
* Added unit-tests for basic tiny-error type(`codeContainsError`) for implementation of Error/Unwrap interface 

## [v0.0.1] - 03.03.2025
### Added
* Create lib-tinyerrors library
* Added linters checks
* Added unit-tests for default error-formatter
* Added library description in [README.md](/README.md) file
* Added license info:
  * [LICENSE](/LICENSE) file with **MIT NON-AI License**
  * License banner to all *.go files
* Added examples of library usage:
  * Error code wrapping [signer/](/examples/signer)
  * Error wrapping [coinflip](/examples/coinflip)
  * Error re-wrapping [tiktaktoe](/examples/tiktaktoe)
  * New error creation [pingpong](/examples/pingpong)