load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "16e9fca53ed6bd4ff4ad76facc9b7b651a89db1689a2877d6fd7b82aa824e366",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
    ],
)

http_archive(
    name = "rules_proto",
    sha256 = "e017528fd1c91c5a33f15493e3a398181a9e821a804eb7ff5acdd1d2d6c2b18d",
    strip_prefix = "rules_proto-4.0.0-3.20.0",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0-3.20.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "501deb3d5695ab658e82f6f6f549ba681ea3ca2a5fb7911154b5aa45596183fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
    ],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "ae34c344514e08c23e90da0e2d6cb700fcd28e80c02e23e4d5715dddcb42f7b3",
    strip_prefix = "buildtools-4.2.2",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/4.2.2.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

### TOOLCHAINS
#### Golang
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.18.3")

gazelle_dependencies(go_sdk = "go_sdk")

go_repository(
    name = "com_github_segmentio_kafka_go",
    importpath = "github.com/segmentio/kafka-go",
    sum = "h1:Ohr+9E+kDv/Ld2UPJN9hnKZRd2qgiqCmI8v2e1qlfLM=",
    version = "v0.4.32",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:wKRjX6JRtDdrE9qwa4b/Cip7ACOshUI4smpCQanqjSY=",
    version = "v1.15.9",
)

go_repository(
    name = "com_github_pierrec_lz4_v4",
    importpath = "github.com/pierrec/lz4/v4",
    sum = "h1:MO0/ucJhngq7299dKLwIMtgTfbkoSPF6AoMYDd8Q4q0=",
    version = "v4.1.15",
)

go_repository(
    name = "io_gorm_driver_postgres",
    importpath = "gorm.io/driver/postgres",
    sum = "h1:8bEphSAB69t3odsCR4NDzt581iZEWQuRM27Cg6KgfPY=",
    version = "v1.3.8",
)

go_repository(
    name = "io_gorm_gorm",
    importpath = "gorm.io/gorm",
    sum = "h1:h8sGJ+biDgBA1AD1Ha9gFCx7h8npU7AsLdlkX0n2TpE=",
    version = "v1.23.8",
)

go_repository(
    name = "com_github_jackc_pgx_v4",
    importpath = "github.com/jackc/pgx/v4",
    sum = "h1:JzTglcal01DrghUqt+PmzWsZx/Yh7SC/CTQmSBMTd0Y=",
    version = "v4.16.1",
)

go_repository(
    name = "com_github_jackc_pgproto3_v2",
    importpath = "github.com/jackc/pgproto3/v2",
    sum = "h1:brH0pCGBDkBW07HWlN/oSBXrmo3WB0UvZd1pIuDcL8Y=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_jackc_pgtype",
    importpath = "github.com/jackc/pgtype",
    sum = "h1:u4uiGPz/1hryuXzyaBhSk6dnIyyG2683olG2OV+UUgs=",
    version = "v1.11.0",
)

go_repository(
    name = "com_github_jackc_pgconn",
    importpath = "github.com/jackc/pgconn",
    sum = "h1:rsDFzIpRk7xT4B8FufgpCCeyjdNpKyghZeSefViE5W8=",
    version = "v1.12.1",
)

go_repository(
    name = "com_github_jackc_pgio",
    importpath = "github.com/jackc/pgio",
    sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_chunkreader_v2",
    importpath = "github.com/jackc/chunkreader/v2",
    sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_jackc_pgservicefile",
    importpath = "github.com/jackc/pgservicefile",
    strip_prefix = "pgservicefile-master",
    urls = ["https://github.com/jackc/pgservicefile/archive/refs/heads/master.zip"],
)

go_repository(
    name = "com_github_jackc_pgpassfile",
    importpath = "github.com/jackc/pgpassfile",
    sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jinzhu_now",
    importpath = "github.com/jinzhu/now",
    sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
    version = "v1.1.5",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    importpath = "github.com/jinzhu/inflection",
    sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
    version = "v1.0.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
    version = "v1.21.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:dg6GjLku4EH+249NNmoIciG9N/jURbDG+pFlTkhzIC8=",
    version = "v1.8.0",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
    version = "v0.3.7",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:zuSxTR4o9y82ebqCUJYNGJbGPo6sKVl54f/TVDObg1c=",
    version = "v0.0.0-20220722155217-630584e8d5aa",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:9vYwv7OjYaky/tlAeD7C4oC9EsPTlaFl1H2jS++V+ME=",
    version = "v0.0.0-20220804214406-8e32c043e418",
)


go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:wrwAbqJqf+ncEK1F/bXTYpgO6zXIgQXi/2ppBgmYI9g=",
    version = "v1.44.70",
)

# go_repository(
#     name = "com_github_aws_aws_sdk_go_aws_client",
#     importpath = "github.com/aws/aws-sdk-go/aws/client",
#     sum = "h1:3A3DEizrCK6dAbBoRGh8KmoZij7She9snclG1ixY/xQ=",
#     version = "v1.44.70",
# )

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_conrey_engineering_go_octoprint",
    importpath = "github.com/conrey-engineering/go-octoprint",
    urls = ["https://github.com/conrey-engineering/go-octoprint/archive/refs/heads/main.zip"],
    strip_prefix = "go-octoprint-main",
)

go_repository(
    name = "io_opentelemetry_go_contrib",
    importpath = "go.opentelemetry.io/contrib",
    version = "v1.9.0",
    sum = "h1:2KAoCVu4OMI9TYoSWvcV7+UbbIPOi4623S77nV+M/Ks=",
)

# go_repository(
#     name = "io_opentelemetry_go_otel",
#     importpath = "go.opentelemetry.io/otel",
#     version = "v1.9.0",
#     sum = "h1:8WZNQFIB2a71LnANS9JeyidJKKGOOremcUtb/OtHISw=",
# )

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0"
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_trace",
    importpath = "go.opentelemetry.io/otel/sdk/trace",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/sdk/trace"
)

# go_repository(
#     name = "io_opentelemetry_go_otel_attribute",
#     importpath = "go.opentelemetry.io/otel/attribute",
#     urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
#     strip_prefix = "opentelemetry-go-1.9.0/attribute"
# )

# go_repository(
#     name = "io_opentelemetry_go_otel_internal",
#     importpath = "go.opentelemetry.io/otel/internal",
#     urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
#     strip_prefix = "opentelemetry-go-1.9.0/internal"
# )

go_repository(
    name = "io_opentelemetry_go_otel_sdk_instrumentation",
    importpath = "go.opentelemetry.io/otel/sdk/instrumentation",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/sdk/instrumentation"
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_resource",
    importpath = "go.opentelemetry.io/otel/sdk/resource",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/sdk/resource"
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_internal",
    importpath = "go.opentelemetry.io/otel/sdk/internal",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/sdk/internal"
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_jaeger",
    importpath = "go.opentelemetry.io/otel/exporters/jaeger",
    # version = "v1.9.0",
    # sum = "h1:gAEgEVGDWwFjcis9jJTOJqZNxDzoZfR12WNIxr7g9Ww=",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/exporters/jaeger"
)

go_repository(
    name = "io_opentelemetry_go_otel_semconv_v1_12_0",
    importpath = "go.opentelemetry.io/otel/semconv/v1.12.0",
    # version = "v1.9.0",
    # sum = "h1:gAEgEVGDWwFjcis9jJTOJqZNxDzoZfR12WNIxr7g9Ww=",
    urls = ["https://github.com/open-telemetry/opentelemetry-go/archive/refs/tags/v1.9.0.zip"],
    strip_prefix = "opentelemetry-go-1.9.0/semconv/v1.12.0"
)

go_repository(
    name = "com_github_go_logr_stdr",
    importpath = "github.com/go-logr/stdr",
    version = "v1.2.2",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    version = "v1.2.3",
    sum = "h1:2DntVwHkVopvECVRSlL5PSo9eG+cAkDCuckLubN+rq0=",
)

#### Protobuf
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

rules_proto_dependencies()

rules_proto_toolchains()

protobuf_deps()

load("//src/api:deps.bzl", "api_dependencies")

api_dependencies()
