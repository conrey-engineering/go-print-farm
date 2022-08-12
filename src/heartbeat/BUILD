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
        "heartbeat.go",
        "main.go",
        "poll.go",
        "printer.go",
    ],
    importpath = "github.com/conrey-engineering/go-print-farm/src/heartbeat",
    deps = [
        "//lib/kafka",
        "//lib/printers/octopi",
        "//lib/tracing:tracing_lib",
        "//src/protobufs:heartbeat_go_proto",
        "//src/protobufs:printer_proto_go",
        "@com_github_segmentio_kafka_go//:kafka-go",
        "@io_opentelemetry_go_otel//:otel",
        # "@io_opentelemetry_go_otel//:otel",
        # "@io_opentelemetry_go_contrib//instrumentation/net/http/otelhttp",
        # "@io_opentelemetry_go_otel//:otel",
        # "@io_opentelemetry_go_contrib//instrumentation/net/http/otelhttp",
        # "@io_opentelemetry_go_otel//:otel",
        # "@io_opentelemetry_go_otel//attribute",
        "@org_uber_go_zap//:zap",
    ],
)
