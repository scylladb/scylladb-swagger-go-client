// Copyright (C) 2023 ScyllaDB

//go:build tools
// +build tools

// Force go mod to download and vendor code that isn't depended upon.
package dependencymagnet

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)
