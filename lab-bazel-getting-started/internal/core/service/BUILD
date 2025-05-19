load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "service",
    srcs = ["fruit.go"],
    importpath = "github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/service",
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["fruit_test.go"],
    embed = [":service"],
)
