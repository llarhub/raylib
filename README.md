# raylib binding inputs

This branch stores the source materials used to generate the LLGo raylib binding on the `main` branch. The checked-in headers are from raylib 6.0.0:

- `c/include/raylib.h` — core windowing, drawing, textures, models, audio, and input API
- `c/include/raymath.h` — raylib math types used by the umbrella input
- `c/include/rcamera.h` — camera helper declarations
- `c/include/raylib_all.h` — one translation unit that includes the three public headers

Install LLGo, raylib, and `pkg-config` before regenerating. On macOS with Homebrew:

```bash
brew install raylib pkg-config
```

From this worktree's `c` directory, generate into the main worktree with:

```bash
PATH=/usr/local/go/bin:$PATH GOMODCACHE=$(go env GOMODCACHE) llcppg ../..
```

`c/llcppg.cfg` links the generated package through `pkg-config --libs raylib` (with `-lraylib` as a fallback). Verify the generated package and the source-material test with:

```bash
llgo test ./...
```

The generated package is the public artifact on `main`; the input set intentionally covers the public raylib headers and does not expose raylib's private implementation headers.
