# Bazel

## Commands

### Clean cache

```shell
bazel clean --expunge
```

### Mod Tidy

```shell
bazel mod tidy
```

### Generate BUILD files

```shell
bazel run //:gazelle
```

### Build

```shell
bazel build //cmd/app
```

### Run built binary

```shell
# After build
bazel-bin/cmd/app/app_/app
```

### Run

```shell
bazel run //cmd/app
```

### Test

```shell
bazel test //internal/core/service:all
```
