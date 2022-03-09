# Code Review Comments Summary

[Source](https://github.com/golang/go/wiki/CodeReviewComments)

* Comments documenting an item must start with the item's name and complete a full sentence
* `context.Context` must be the first parameter
* Careful when copying structs
* Don't use `math/rand` for crypto
* Prefer `var t []string` over `t := []string{}` when declaring an empty slice
* All exported names must have doc comments
* Don't use `panic` for error handling
* Don't start error messages with uppercase letters, as they might appear after other error messages
* Use [`Example`](https://go.dev/blog/examples) functions
* Make it clear when or whether a goroutine exits
* Always handle errors
* Avoid renaming packages and use [`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
* Package import for side effects should be done by `import _ "pkg/name"`
* Import dot can be useful in tests to avoid circular dependencies, otherwise don't use it
* Return `err` if something can fail, don't use in-band errors
* Happy path should have the least identation possible
* Acronyms should use `urlPath` or `URLPath` instead of `UrlPath`
* Put interfaces in a package that uses it, not in a package that implements it
* Implementors of an interface must return pointer-to-struct instead of the interface
* There is no rule for max characters per line, but try to keep it readable
* Adding names to result parameters can be useful for readability in case of multiple return values or same type returns
* [Naked Return](https://github.com/golang/go/wiki/CodeReviewComments#naked-returns)
* Package comments must be adjacent to its declaration
* The package name is also part of its exported types. E.g. prefer `http.Get()` over `http.HttpGet()`
* Avoid parameters as pointers when the value can be passed directly by value, like `string` and interfaces
* The name of a method's receiver should be a reflection of its identity (unlike `me`, `this`, `self`)
* When in doubt, use a pointer receiver ([read more](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type))
* Prefer functions that keep goroutines localized within a call (synchronous) over asynchronous functions since it is also easier to test
* Useful test failures
* Short variable names