# GoMock

## Generate Mocks

How to generate a Mock for an interface.
The following command will generate Mocks for interfaces located in `internal/core/port/fruit.go`.
Generated Mocks will be created in `internal/core/port/mock/fruit.go` with the package name `mock`.

```shell
mockgen -source=internal/core/port/fruit.go -destination=internal/core/port/mock/fruit.go -package=mock
```
