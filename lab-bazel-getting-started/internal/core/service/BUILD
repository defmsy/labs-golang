load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "service",
    srcs = ["fruit.go"],
    importpath = "github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/service",
    visibility = ["//visibility:public"],
    deps = ["//internal/core/port"],
)

go_test(
    name = "service_test",
    srcs = ["fruit_test.go"],
    deps = [
        ":service",
        "//internal/core/port/mock:port_mock",
        "@org_uber_go_mock//gomock",
    ],
)
