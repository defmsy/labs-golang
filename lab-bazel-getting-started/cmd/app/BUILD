load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "app",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
    deps = ["@com_github_gin_gonic_gin//:go_default_library"],
)
