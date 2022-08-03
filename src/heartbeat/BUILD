load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_binary(
    name = "heartbeat",
    embed = [":heartbeat_lib"],
)

go_library(
    name = "heartbeat_lib",
    srcs = [
        "base_printer.go",
        "main.go",
        "octopi_printer.go",
        "printer.go",
    ],
    importpath = "github.com/conrey-engineering/go-print-farm/src/heartbeat",
    deps = [
        "//src/protobufs:heartbeat_go_proto",
        "//src/protobufs:printer_proto_go",
        "@com_github_segmentio_kafka_go//:kafka-go",
        "@org_uber_go_zap//:zap",
    ],
)
