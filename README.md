Context
=======

Context is api-compatible with go 1.7's context package, and with the older
 golang.org/x/net/context package.  It uses build flags to control which
 backing context implementation it uses.  On go 1.7, it will delegate
 to the package in the SDK.  On earlier versions of go, it will use the
 golang.org/x/net/context package.