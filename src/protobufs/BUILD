# gazelle:proto package
# gazelle:proto_group go_package
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "printer_proto",
    srcs = [
        "Printer.proto",
    ],
    deps = ["@com_google_protobuf//:timestamp_proto"],
)

go_proto_library(
    name = "printer_proto_go",
    importpath = "github.com/conrey-engineering/go-print-farm/src/protobufs/printer",
    proto = ":printer_proto",
)

proto_library(
    name = "heartbeat_proto",
    srcs = [
        "Heartbeat.proto",
    ],
    deps = [
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "heartbeat_go_proto",
    importpath = "github.com/conrey-engineering/go-print-farm/src/protobufs/heartbeat",
    proto = ":heartbeat_proto",
)

exports_files(["*.proto"])
