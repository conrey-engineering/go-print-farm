load("@bazel_tools//tools/build_defs/repo:utils.bzl", _maybe = "maybe")
load("@bazel_gazelle//:deps.bzl", "go_repository")

def api_dependencies():
    _maybe(
        go_repository,
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
        version = "v1.8.0",
    )
