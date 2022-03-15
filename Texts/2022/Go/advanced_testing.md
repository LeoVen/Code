# Advanced Testing with Go Summary

Summary based on [GopherCon 2017: Mitchell Hashimoto - Advanced Testing with Go
](https://www.youtube.com/watch?v=8hQG7QlcLBk)

* Test Methodology
* Writing Testable Code

## Test Methodology

* Subtests
* Table-driven tests
* Test Fixtures
* Golden Files
* Test helpers
* Don't mock `net.Conn`
* Subprocessing

## Writing Testable Code

* Avoid global state
* Repeat yourself
* Test exported API
* Don't underpackage or overpackage
* Configurability
* Interfaces
* Complex struct comparison
  * Test string
  * Reflection
* Expose API with ways of testing the application
