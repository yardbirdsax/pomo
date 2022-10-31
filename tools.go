//go:build tools
package tools

import (
	_ "github.com/posener/goreadme/cmd/goreadme"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/goreleaser/goreleaser"
)