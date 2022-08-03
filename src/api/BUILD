load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

package(default_visibility = ["//visibility:public"])

gazelle(name = "gazelle")

go_binary(
    name = "api",
    embed = [":api_lib"],
)

go_library(
    name = "api_lib",
    srcs = [
        # "kafka_connector.go",
        "kafka_connector.go",
        "kafka_messages.go",
        "main.go",
    ],
    importpath = "github.com/conrey-engineering/go-print-farm/src/api",
    deps = [
        "//src/protobufs:printer_proto_go",
        "@com_github_gorilla_mux//:mux",
        "@com_github_segmentio_kafka_go//:kafka-go",
        "@org_uber_go_zap//:zap",
    ],
)

go_test(
    name = "api_test",
    srcs = ["kafka_messages_test.go"],
    embed = [":api_lib"],
)
