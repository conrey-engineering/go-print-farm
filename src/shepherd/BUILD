load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

package(default_visibility = ["//visibility:public"])

gazelle(name = "gazelle")

go_binary(
    name = "shepherd",
    embed = [":shepherd_lib"],
)

go_library(
    name = "shepherd_lib",
    srcs = [
        "main.go",
        "printer_type.go",
    ],
    importpath = "github.com/conrey-engineering/go-print-farm/src/shepherd",
    deps = [
        "//src/protobufs:printer_proto_go",
        "@com_github_google_uuid//:uuid",
        "@com_github_segmentio_kafka_go//:kafka-go",
        "@io_gorm_driver_postgres//:postgres",
        "@io_gorm_gorm//:gorm",
        "@org_uber_go_zap//:zap",
    ],
)
