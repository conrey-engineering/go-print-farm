load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_binary(
    name = "main",
    deps = [
        "//src/protobufs:printer_proto_go",
    ],
    srcs = [
        "main.go",
    ],
)
