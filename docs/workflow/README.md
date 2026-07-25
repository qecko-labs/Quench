# Build workflow

A typical Quench workflow is straightforward:

1. Initialize or configure a project.
2. Point the build system at a source tree or source file.
3. Let the builder discover inputs and compile them.
4. Link the resulting objects into the final binary.
5. Optionally run scripts, hooks, or dependency builds after the main build.

## Typical project flow

- `qh init` creates a starter configuration and ignore file.
- `qh -dir src -out app` builds a project from a source directory.
- `qh -watch` turns the build into a continuously refreshing loop while editing.
- `qh -clean` removes generated artifacts.

## Why this workflow works

The workflow is designed to stay close to the actual build graph. A user can see what was built, where the outputs went, and which toolchain settings influenced the result.
