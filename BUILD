load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/codemanufaktur/web-cli",
)

go_library(
    name = "web-cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/codemanufaktur/web-cli",
    visibility = ["//visibility:private"],
    deps = ["//cmd"],
)

go_binary(
    name = "web-cli",
    embed = [":web-cli_lib"],
    visibility = ["//visibility:public"],
)
