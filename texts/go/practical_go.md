# Practical Go Notes

[Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html)

# 1 Guiding Principles

Examples ([Source](https://www.youtube.com/watch?v=2wZ1pCpJUIM))

|                  |                 |              |
| ---------------- | --------------- | ------------ |
| Approachability  | Integrity       | Robustness   |
| Availability     | Maintainability | Safety       |
| Compatibility    | Measurability   | Security     |
| Composability    | Operability     | Simplicity   |
| Debuggability    | Performance     | Stability    |
| Expressiveness   | Portability     | Thoroughness |
| Extensibility    | Resiliency      | Transparency |
| Interoperability | Rigor           | Velocity     |

## Go

* Clarity
* Simplicity
* Productivity

# 2 Identifiers

* Choose identifiers for clarity, not brevity
* Identifier length
  * Depends on the context and for how long the variable is used
* A variable's name should describe its content
* Consistent naming style
  * `CamelCase`
* Consistent declaration style
  * When declaring but not initializing: `var phone Phone`
  * When declaring and initializing: `phone := NewPhone(12345)`
  * Prefer `thing := &Thing{}` over `thing := new(Thing)`
  * If type needs to be specified: `var rune uint32 = 0x80`
* If any, follow pre-stablished styles in the existing code base
* Remeber to name exported identifiers based on the package name, since the latter will always be written before the name
  * `const ROUTE = ".."` instead of `const API_ROUTE = ".."`
  * With `api` as the package name, `api.ROUTE` is better than `api.API_ROUTE` since the latter has redundancy

# 3 Comments

Every comment should be __only one__ of these:

* __What__ the thing does
* __How__ the thing does what it does
* __Why__ it is the way it is

* Comments on variables and constants should describe their contents, not their purpose
  * Variable name: purpose
  * Variable comment: content
  * See also: `http.StatusXXX` constants
* Uninitialized variables should describe who is responsible for initializing it
* Sometimes you’ll find a better name for a variable hiding in a comment
* __Always document public symbols__
  * [godoc](https://go.dev/blog/godoc)
  * [Google Style Guide](https://google.github.io/styleguide/)
    * Any public function that is not both obvious and short must be commented
    * Any function in a library must be commented regardless of length or complexity
  * No need to comment on methods that implement an interface. The comment describing the method at the interface should be sufficient
* Raise issues with `TODO` for refactor
* *How can I improve the code so that this comment isn't needed?*

# 4 Package Design

* A package should be a unit, that is, atomic and indivisible
* The public API of a package describes __what__ it does, not how
* __A good package starts with its name__
  * Name the package based on what it provides, not what it contains
* __Good package names should be unique__
  * See: `io/ioutil` and `net/http/httputil` packages
* Avoid package names that are too generic, like a utility package
* Utility packages should use the plural form
  * Like a utility package for `Thing` should be `Things`
* __A public Identifier includes its package name__
  * Use `api.Initialize()` instead of `api.InitializeApi()`
  * It gets redundant since the exported identifier is always accompanied by the package name
* __Early return over nesting__
  * [Align the happy path to the left edge](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88)
* Make the zero value useful
* Avoid package level state
  * A change of one package should have a low probability of affecting another
  * Loose coupling
    * Use interfaces to describe behaviour
    * Avoid global state

# 5 Project Structure

* Avoid combining multiple purposes into a single project
* Consider fewer, larger packages
  * "A java package is equivalent to a single `.go` file
  * A Go package is equivalent to a whole .NET assembly
* Different files should be responsible for different areas of the package
  * E.g.: in `net/http` package, there are `messages.go`, `client.go` and `server.go` files
* Use `internal` packages to reduce API surface
* Keep package `main` as small as possible
  * Hard to test code in `main`
  * Business logic should be out of package `main`
  * Use it only to parse flags, open connections to databases, loggers, and such, then hand off execution to a high level object

# 6 API Design

* Design APIs that are hard to misuse
  * Be wary of functions that take two or more parameters of the same type
  * APIs with multiple parameters of the same type are hard to use
* Design APIs for their default use case
  * Don't require parameters where the caller doesn't care about
  * Discourage the use o `nil` as a parameter
  * Don’t mix `nil` and non `nil`-able parameters in the same function signature
  * __Clear is better than concise__
  * Prefer var args `values ...int` over slice `values []int`
    * Avoids you having to 'box' a single value in a slice
    * Slices accept `nil` or an empty slice and this can indicate that something is wrong at the caller
* Let functions define the behaviour they require
  * __Interface segregation principle__: no code should be forced to depend on methods it does not use
* Understanding `nil`
  * `nil` cannot be compared with itself for equality or inequality
  * `nil` may appear on either side of a binary operation
  * If you assign `nil` to a pointer the pointer will not point to anything
  * If you assign `nil` to an interface, the interface will not contain anything
  * If you assign `nil` to a slice and the slice will have zero len and zero cap and no longer point an underlying array
  * Be ware of the typed `nil`
  * If a method returns an interface, always return `nil` explicitly
    * Do `return nil` instead of `return err`

# 7 Error Handling

* Exceptions obfuscate control flow
* Errors are just values
* `panic` and `recover` have very specific uses and semantics
* Errors should be opaque
  * Prevents coupling
  * Prevents having to break API compatibility
  * Exception to the rule: `io.EOF`
* Assert errors for behaviour, not type
  * Don’t assert an error value is a specific type, but rather assert that the value implements a particular behaviour
  * The caller should not presume anything about the state of the `error` value
  * If a function returns `error`, you cannot make assumptions about the state of any other value until the error is checked
  * However, errors can be checked against known public types in order to retry an operation
* Never use `nil` to indicate failure
* Don't check for `nil` receivers. Instead, check it before calling the method (but there can be cases where the receiver can be `nil`)
  * If the receiver can't be `nil`, but it is, then consider:
    * Letting it panic or doing so manually with a better message
    * Returning an `error` (but this makes every method having to return an `error`)
* Don't panic
  * `panic` must be the last resort
  * Avoid `log.Fatal` and `log.Panic`
* Eliminate error handling by eliminating errors
* __Only handle an error once__

# 8 Testing

* Write software in a way that is testable
* Unit tests describe what the package promises to do
* Test scope (files ending with `XXX_test.go`)
* Code coverage
  * `go test -coverprofile=c.out`
  * `go tool cover -func=c.out`
* Table driven testing
  * Common fields:
    * `name` for the test name
    * `want` the wanted value, compared to `got`
    * `input` to the function being tested
* Sub tests
* Comparing `want` to `got`
  * Use `%#v` formatting for better output
  * Even better is to use [go-cmp](https://github.com/google/go-cmp)
* Prefer internal tests to external tests to avoid dot imports, but place `Example` test functions in an external test file

| Type          | Code File | Package | Test File      | Test Package |
| ------------- | --------- | ------- | -------------- | ------------ |
| Internal Test | `http.go` | `http`  | `http_test.go` | `http`       |
| External Test | `http.go` | `http`  | `http_test.go` | `http_test`  |

* Tests can assure that you can ship the master branch
* Tests give you confidence to change someone else's code

# 9 Concurrency

* Channel Axioms
  * A send to a `nil` channel, which is the default value, blocks forever
  * A receiver from a `nil` channel blocks forever
  * A send to a closed channel panics
  * A receive from a closed channel returns the zero value immediately
* Prefer channels with a size of zero or one
* Leave concurrency to the caller
* If a function starts a goroutine, provide ways to the caller to eventually stop it
* Never start a goroutine without knowing when it will stop
