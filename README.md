# deckr 🃏

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://godoc.org/github.com/mgjules/deckr)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=for-the-badge)](LICENSE)

Deckr provides a REST API to simulate a deck of cards.


## Contents

- [deckr 🃏](#deckr-)
  - [Contents](#contents)
  - [Requirements](#requirements)
  - [Mage Targets](#mage-targets)
    - [Example](#example)
      - [Generate docs](#generate-docs)
      - [Generate stubs from proto files](#generate-stubs-from-proto-files)
      - [Run tests with race detector](#run-tests-with-race-detector)
      - [Build deckr for MacOS M1](#build-deckr-for-macos-m1)
  - [Install](#install)
  - [Usage](#usage)
    - [REST API server](#rest-api-server)
  - [REST API documentation](#rest-api-documentation)
  - [License](#license)

## Requirements

- [Go 1.18+](https://golang.org/doc/install)
- [Mage](https://github.com/magefile/mage) - replacement for Makefile in Go.
    
    ```shell
    $ go install github.com/magefile/mage@latest
    ```

- [Golangci-lint](https://github.com/golangci/golangci-lint) - Fast Go linters runner.
      
    ```shell
    $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    ```
    
- [Ginkgo](https://github.com/onsi/ginkgo) - Expressive testing framework.
        
    ```shell
    $ go install github.com/onsi/ginkgo/v2/ginkgo@latest
    ```
    
- [Swag](https://github.com/swaggo/swag) - Generate REST API documentation.
        
    ```shell
    $ go install github.com/swaggo/swag/cmd/swag@latest
    ```

- [Buf](https://github.com/bufbuild/buf) - A new way of working with Protocol Buffers.

   ```shell
   $ go install github.com/bufbuild/buf/cmd/buf@latest
   ```

- [Protoc](https://developers.google.com/protocol-buffers) - Protocol Buffers compiler.

   For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).

- [Protoc-gen-go](https://github.com/golang/protobuf/protoc-gen-go) - Go code generator for Protocol Buffers.
  
   ```shell
   $ go install github.com/golang/protobuf/protoc-gen-go@latest
   ```

- [Protoc-gen-go-grpc](https://google.golang.org/grpc/cmd/protoc-gen-go-grpc) - Go code generator for gRPC.
  
   ```shell
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
    
- [Docker](https://www.docker.com) - Containerization.

## Mage Targets

```shell
$ mage -l
```

```
Targets:
  build:all           Builds for all supported popular OS/Arch
  build:linuxAmd64    Builds for Linux 64bit
  build:linuxArm64    Builds for Linux ARM 64bit
  build:macOSAmd64    Builds for MacOS 64bit
  build:macOSArm64    Builds for MacOS M1
  build:winAmd64      Builds for Windows 64bit
  docs                Generates docs
  lint                Run golangci linters
  proto               Generate stubs from proto files
  test                Run tests
  testRace            Run tests with race detector
  tidy                Run go mod tidy
```

### Example

#### Generate docs

```shell
$ mage -v docs
```

#### Generate stubs from proto files

```shell
$ mage -v proto
```

#### Run tests with race detector

```shell
$ mage -v testRace
```

#### Build deckr for MacOS M1

```shell
$ mage -v build:macOSArm64
```

## Install

- You can install using the [latest released binary](https://github.com/mgjules/deckr/releases/latest).

- **OR** using Go:

    ```shell
    $ go install github.com/mgjules/deckr@latest
    ```

- **OR** bulding from source:

    Example (MacOS M1):

    ```shell
    $ mage -v build:macOSArm64
    ```

## Usage

```shell
$ deckr --help
```

```
NAME:
   deckr - A REST API for playing with a deck of cards

USAGE:
   deckr [global options] command [command options] [arguments...]

DESCRIPTION:
   Deckr exposes a REST API for playing with a deck of cards of your choice.

AUTHOR:
   Michaël Giovanni Jules <julesmichaelgiovanni@gmail.com>

COMMANDS:
   serve       Starts the REST API server.
   version, v  Shows the version
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

COPYRIGHT:
   (c) 2022 Michaël Giovanni Jules
```

### REST API server

```shell
$ deckr serve --help
```

```
NAME:
   deckr serve - Starts the REST API server.

USAGE:
   deckr serve [command options] [arguments...]

OPTIONS:
   --debug              whether running in PROD or DEBUG mode (default: false) [$DECKR_DEBUG]
   --server-uri value   URI of server (default: "http://localhost:9000") [$DECKR_SERVER_URI]
   --storage-uri value  URI of storage (default: "inmemory://") [$DECKR_STORAGE_URI]
   --help, -h           show help (default: false)
```

## REST API documentation

The REST API documentation is generated using Swag and is available at `/swagger/index.html`.

## License

Deckr is Apache 2.0 licensed.