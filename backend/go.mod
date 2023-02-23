module main

go 1.16

require (
	github.com/gin-gonic/contrib v0.0.0-20221130124618-7e01895a63f2 // indirect
	github.com/gin-gonic/gin v1.8.2
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/ugorji/go/codec v1.2.10 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	myfunctions v1.0.0 // local function
)

// A replace directive replaces the contents of a specific version of a module,
// or all versions of a module, with contents found elsewhere.
// The replacement may be specified with either another module path and version, or a platform-specific file path.
replace myfunctions v1.0.0 => /myfunctions
