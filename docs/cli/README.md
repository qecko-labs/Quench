# CLI usage

The CLI is intentionally thin. It is responsible for parsing user intent, selecting configuration, and handing the actual work to the build engine.

## Typical commands

- `qh` or `qh -dir .` builds the project in the current directory.
- `qh -asm path/to/file.asm` assembles a single file.
- `qh -cc path/to/file.c` compiles a single C-like translation unit.
- `qh -dir src -out app` builds from a specific directory and writes a named binary.
- `qh -watch` rebuilds automatically when sources change.
- `qh -clean` removes generated build artifacts.
- `qh -json` produces machine-readable output for scripting and CI.

## Important design point

The CLI is not the actual compiler. It is the control plane. The real compilation work lives in the builder and backend integrations. That separation is deliberate: it keeps the user experience consistent even when the underlying toolchain changes.
