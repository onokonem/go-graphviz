# go-graphviz [![Go](https://github.com/onokonem/go-graphviz/workflows/Go/badge.svg)](https://github.com/onokonem/go-graphviz/actions) [![GoDoc](https://godoc.org/github.com/onokonem/go-graphviz?status.svg)](https://pkg.go.dev/github.com/onokonem/go-graphviz) 

Go bindings for Graphviz

<img src="https://user-images.githubusercontent.com/209884/90976476-64e84000-e578-11ea-9596-fb4a7d3b11a6.png" width="400px"></img>

# Features

Graphviz version is [here](./graphviz.version)

- Pure Go Library
- No need to install Graphviz library ( ~`brew install graphviz`~ or ~`apt-get install graphviz`~ )
  - The Graphviz library has been converted to WebAssembly (WASM) and embedded it, so it works consistently across all environments
- Supports encoding/decoding for DOT language
- Supports custom renderer for custom format
- Supports setting graph properties in a type-safe manner

## Supported Layout

`circo` `dot` `fdp` `neato` `nop` `nop1` `nop2` `osage` `patchwork` `sfdp` `twopi`

## Supported Format

`dot` `svg` `png` `jpg`

The above are the formats supported by default. You can also add custom formats.

# Installation

```bash
$ go get github.com/onokonem/go-graphviz
```

# Synopsis

## 1. Write DOT Graph in Go

```go
package main

import (
  "bytes"
  "context"
  "fmt"
  "log"

  "github.com/onokonem/go-graphviz"
)

func main() {
  ctx := context.Background()
  g, err := graphviz.New(ctx)
  if err != nil { panic(err )}

  graph, err := g.Graph()
  if err != nil { panic(err) }
  defer func() {
    if err := graph.Close(); err != nil { panic(err) }
    g.Close()
  }()
  n, err := graph.CreateNodeByName("n")
  if err != nil { panic(err) }

  m, err := graph.CreateNodeByName("m")
  if err != nil { panic(err) }

  e, err := graph.CreateEdgeByName("e", n, m)
  if err != nil { panic(err) }
  e.SetLabel("e")

  var buf bytes.Buffer
  if err := g.Render(ctx, graph, "dot", &buf); err != nil {
    log.Fatal(err)
  }
  fmt.Println(buf.String())
}
```

## 2. Parse DOT Graph

```go
path := "/path/to/dot.gv"
b, err := os.ReadFile(path)
if err != nil { panic(err) }
graph, err := graphviz.ParseBytes(b)
```

## 3. Render Graph

```go
ctx := context.Background()
g, err := graphviz.New(ctx)
if err != nil { panic(err) }

graph, err := g.Graph()
if err != nil { panic(err) }

// create your graph

// 1. write encoded PNG data to buffer
var buf bytes.Buffer
if err := g.Render(ctx, graph, graphviz.PNG, &buf); err != nil { panic(err) }

// 2. get as image.Image instance
image, err := g.RenderImage(ctx, graph)
if err != nil { panic(err) }

// 3. write to file directly
if err := g.RenderFilename(ctx, graph, graphviz.PNG, "/path/to/graph.png"); err != nil { panic(err) }
```

# Tool

## `dot`

`dot` is a standalone command that renders graphs using the Graphviz library embedded in this project. Every [release](https://github.com/onokonem/go-graphviz/releases) is versioned after the Graphviz library it wraps (release `vX.Y.Z` bundles Graphviz `X.Y.Z`, see [`graphviz.version`](./graphviz.version)) and publishes precompiled binaries for macOS, Windows, and Linux on both `amd64` and `arm64`.

### Installation

#### Precompiled binary (recommended)

Download the archive for your platform from the [latest release](https://github.com/onokonem/go-graphviz/releases/latest), unpack it, and put the binary on your `PATH`.

Each release publishes the following artifacts (replace `<version>` with the Graphviz version wrapped by the release, e.g. `16.1.0` for release `v16.1.0`):

| Platform | Architecture | Archive | Binary |
| --- | --- | --- | --- |
| macOS | amd64 | `go-graphviz_<version>_darwin_amd64.tar.gz` | `dot` |
| macOS | arm64 | `go-graphviz_<version>_darwin_arm64.tar.gz` | `dot` |
| Linux | amd64 | `go-graphviz_<version>_linux_amd64.tar.gz` | `dot` |
| Linux | arm64 | `go-graphviz_<version>_linux_arm64.tar.gz` | `dot` |
| Windows | amd64 | `go-graphviz_<version>_windows_amd64.zip` | `dot.exe` |
| Windows | arm64 | `go-graphviz_<version>_windows_arm64.zip` | `dot.exe` |

Every archive contains the binary, `README.md`, `LICENSE`, and `graphviz.version`. The `checksums.txt` file with SHA-256 checksums of all artifacts is published alongside them.

**macOS** (example for Apple Silicon; adjust the version in the URLs for the release you are installing):

```bash
curl -LO https://github.com/onokonem/go-graphviz/releases/download/v16.1.0/go-graphviz_16.1.0_darwin_arm64.tar.gz
tar -xzf go-graphviz_16.1.0_darwin_arm64.tar.gz
sudo mv dot /usr/local/bin/
```

**Linux** (example for `amd64`):

```bash
curl -LO https://github.com/onokonem/go-graphviz/releases/download/v16.1.0/go-graphviz_16.1.0_linux_amd64.tar.gz
tar -xzf go-graphviz_16.1.0_linux_amd64.tar.gz
sudo mv dot /usr/local/bin/
```

**Windows** (PowerShell, example for `amd64`):

```powershell
Invoke-WebRequest -OutFile dot.zip https://github.com/onokonem/go-graphviz/releases/download/v16.1.0/go-graphviz_16.1.0_windows_amd64.zip
Expand-Archive dot.zip -DestinationPath dot
Move-Item dot\dot.exe $env:LOCALAPPDATA\Microsoft\WindowsApps\
```

Verify the installation:

```bash
$ dot --version
16.1.0
```

#### From source

```bash
# install the latest released module version (v0.x tag, wraps the same
# Graphviz release as the binaries above)
$ go install github.com/onokonem/go-graphviz/cmd/dot@latest
```

or build from a clone:

```bash
$ git clone https://github.com/onokonem/go-graphviz.git
$ cd go-graphviz/cmd/dot
$ go build -o dot .
```

### Usage

```
Usage:
  dot [OPTIONS]

Application Options:
  -T=            specify output format ( currently supported: dot svg png jpg )
                 (default: dot)
  -K=            specify layout engine ( currently supported: circo dot fdp
                 neato nop nop1 nop2 osage patchwork sfdp twopi )
  -o=            specify output file name
  -v, --version  show version

Help Options:
  -h, --help     Show this help message
```

# How it works

1. Generates bindings between Go and C from [Protocol Buffers file](./internal/wasm/bind.proto).
2. Builds graphviz.wasm on the [docker container](./internal/wasm/build/Dockerfile).
3. Uses Graphviz functionality from a sub-packages ( `cdt` `cgraph` `gvc` ) via the `internal/wasm` package. 
4. `graphviz` package provides facade interface for all sub packages.

# License

MIT

This library embeds and uses `graphviz.wasm`, which is generated based on the original source code of Graphviz. Therefore, the `graphviz.wasm` follows [the license adopted by Graphviz](https://graphviz.org/license) ( Eclipse Public License ).
