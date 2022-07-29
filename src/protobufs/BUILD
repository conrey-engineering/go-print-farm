# gazelle:proto package
# gazelle:proto_group go_package
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "printer_proto",
    srcs = ["Printer.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "printer_proto_go",
    importpath = "example.com/repo/protobufs",
    proto = ":printer_proto",
)
