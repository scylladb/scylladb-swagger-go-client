# ScyllaDB Swagger Go client

[![CI](https://github.com/scylladb/scylladb-swagger-go-client/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/scylladb/scylladb-swagger-go-client/actions/workflows/go.yaml?query=branch%3Amaster)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

This repository contains a Swagger client implementation for ScyllaDB products in the Go programming language. 

## Installation

To use the ScyllaDB Swagger client in your Go project, you need to have Go installed and set up on your system. 
Once Go is set up, you can install the client by running the following command:

```bash
go get -u github.com/scylladb/scylladb-swagger-go-client
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	api "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	scyllav1client "github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v1/client"
	scyllav1operations "github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v1/client/operations"
)

func main() {
	// Initialize the ScyllaDB client
	runtime := api.New("127.0.0.1:10000", scyllav1client.DefaultBasePath, scyllav1client.DefaultSchemes)
	scyllaClient := scyllav1client.New(runtime, strfmt.Default)

	// Get all hosts
	resp, err := scyllaClient.Operations.StorageServiceHostIDGet(&scyllav1operations.StorageServiceHostIDGetParams{})
	if err != nil {
		log.Fatalf("can't get cluster hosts: %s", err)
	}

	for _, kv := range resp.Payload {
		fmt.Println("Address:", kv.Key)
		fmt.Println("HostID:", kv.Value)
	}
}

```
