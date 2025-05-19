# Bazel getting started

Bazel lab for a Go project.

## Commands

### Build

```shell
bazel build //cmd/app
```

### Run

```shell
bazel run //cmd/app
```

```shell
# After build
bazel-bin/cmd/app/app_/app
```

### Test

```shell
bazel test //internal/core/service:all
```
