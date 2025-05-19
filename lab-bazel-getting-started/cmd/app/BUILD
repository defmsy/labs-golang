load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "app_lib",
    srcs = ["main.go"],
    importpath = "github.com/defmsy/labs-golang/lab-bazel-getting-started/cmd/app",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_gin_gonic_gin//:gin",
    ],
)

go_binary(
    name = "app",
    embed = [":app_lib"],
    visibility = ["//visibility:public"],
)
