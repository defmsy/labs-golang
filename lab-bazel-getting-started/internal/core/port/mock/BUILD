load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "port_mock",
    srcs = ["fruit.go"],
    importpath = "github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/port/mock",
    visibility = ["//visibility:public"],
    deps = ["@org_uber_go_mock//gomock"],
)
