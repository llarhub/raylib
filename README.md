# raylib

[![GoDoc](https://pkg.go.dev/badge/github.com/llarhub/raylib.svg)](https://pkg.go.dev/github.com/llarhub/raylib)
[![GitHub release](https://img.shields.io/github/v/tag/llarhub/raylib.svg?label=release)](https://github.com/llarhub/raylib/releases)
[![LLGo](https://img.shields.io/badge/powered_by-LLGo-green.svg)](https://github.com/xgo-dev/llgo)
[![LLGo bindings](https://img.shields.io/badge/llarhub-LLGo_Bindings-brightgreen.svg)](https://github.com/xgo-dev/llgo/blob/main/doc/LLGo-Bindings.md)
[![XGo](https://img.shields.io/badge/project-XGo-blue.svg)](https://github.com/goplus/xgo)

LLGo bindings for raylib 6.0.0. The generated package exposes the public `raylib.h` API, camera helpers from `rcamera.h`, and raylib value types and constants. Functions that take a raylib struct as their first argument are available as methods where LLGo can represent that form.

## Installation

Install LLGo, raylib, and `pkg-config`:

```bash
brew install raylib pkg-config
```

The binding links with `pkg-config --libs raylib` and therefore uses the raylib installation selected by your `PKG_CONFIG_PATH`.

## Usage

```go
package main

import (
	"github.com/goplus/lib/c"
	"github.com/llarhub/raylib"
)

func main() {
	raylib.InitWindow(800, 450, c.Str("raylib + LLGo"))
	defer raylib.CloseWindow()
	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RAYWHITE)
		raylib.EndDrawing()
	}
}
```

The package also supports headless data operations such as `raylib.GenImageColor`; call the generated `Image.Unload` method when an image is no longer needed.

## Verification

Run the standard Go check and the LLGo tests from the repository root:

```bash
go test ./...
llgo test ./...
```

The LLGo test exercises image generation, resource release, generated constants, and value type layout. The binding is generated from the raylib headers installed on the build machine, so keep the raylib API version aligned with the headers and library used at link time.
