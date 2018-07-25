load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# GO INIT
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.12.1/rules_go-0.12.1.tar.gz"],
    sha256 = "8b68d0630d63d95dacc0016c3bb4b76154fe34fca93efd65d1c366de3fcb4294",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz"],
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# SCALA COMPILER
rules_scala_version = "64faf06a4932a9a1d3378b6ba1a6d77479cefef3"  # update this as needed

http_archive(
    name = "io_bazel_rules_scala",
    url = "https://github.com/bazelbuild/rules_scala/archive/%s.zip" % rules_scala_version,
    type = "zip",
    strip_prefix = "rules_scala-%s" % rules_scala_version,
)

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_repositories")

scala_repositories((
    "2.12.6",
    {
        "scala_compiler": "3023b07cc02f2b0217b2c04f8e636b396130b3a8544a8dfad498a19c3e57a863",
        "scala_library": "f81d7144f0ce1b8123335b72ba39003c4be2870767aca15dd0888ba3dab65e98",
        "scala_reflect": "ffa70d522fc9f9deec14358aa674e6dd75c9dfa39d4668ef15bb52f002ce99fa",
    },
))

load("@io_bazel_rules_scala//scala:toolchains.bzl", "scala_register_toolchains")

scala_register_toolchains()

# DOCKER RULES
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "6dede2c65ce86289969b907f343a1382d33c14fbce5e30dd17bb59bb55bb6593",
    strip_prefix = "rules_docker-0.4.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.4.0.tar.gz"],
)

# DOCKER BASE IMAGES
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)

# This is NOT needed when going through the language lang_image
# "repositories" function(s).
container_repositories()

container_pull(
    name = "alpine_base",
    registry = "index.docker.io",
    repository = "library/alpine",
    tag = "3.8",
)

container_pull(
    name = "java_base",
    registry = "index.docker.io",
    repository = "library/openjdk",
    tag = "8u171-jre-alpine3.8",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# JAVA DEPS
maven_jar(
    name = "com_geteventstore_eventstore_client",
    artifact = "com.geteventstore:eventstore-client_2.12:5.0.8",
)

maven_jar(
    name = "org_apache_kafka_kafka_clients",
    artifact = "org.apache.kafka:kafka-clients:1.1.0",
)

maven_jar(
    name = "com_typesafe_akka_akka_actor",
    artifact = "com.typesafe.akka:akka-actor_2.12:2.5.13",
)

maven_jar(
    name = "com_typesafe_akka_akka_stream",
    artifact = "com.typesafe.akka:akka-stream_2.12:2.5.13",
)

maven_jar(
    name = "com_typesafe_akka_akka_testkit",
    artifact = "com.typesafe.akka:akka-testkit_2.12:2.5.13",
)

maven_jar(
    name = "com_typesafe_config",
    artifact = "com.typesafe:config:1.3.3",
)

maven_jar(
    name = "com_typesafe_akka_akka_http",
    artifact = "com.typesafe.akka:akka-http_2.12:10.1.3",
)

maven_jar(
    name = "com_typesafe_akka_akka_http_testkit",
    artifact = "com.typesafe.akka:akka-http-testkit_2.12:10.1.3",
)

maven_jar(
    name = "com_typesafe_akka_akka_http_core",
    artifact = "com.typesafe.akka:akka-http-core_2.12:10.1.3",
)

maven_jar(
    name = "com_typesafe_akka_akka_parsing",
    artifact = "com.typesafe.akka:akka-parsing_2.12:10.1.3",
)

maven_jar(
    name = "org_reactivestreams_reactive_streams",
    artifact = "org.reactivestreams:reactive-streams:1.0.2",
)

maven_jar(
    name = "com_google_protobuf_protobuf_java",
    artifact = "com.google.protobuf:protobuf-java:3.6.0",
)

maven_jar(
    name = "joda_time_joda_time",
    artifact = "joda-time:joda-time:2.10",
)

maven_jar(
    name = "org_joda_joda_convert",
    artifact = "org.joda:joda-convert:2.1",
)

maven_jar(
    name = "org_json_json",
    artifact = "org.json:json:20180130",
)

maven_jar(
    name = "org_slf4j_slf4j_api",
    artifact = "org.slf4j:slf4j-api:1.7.25",
)

maven_jar(
    name = "org_slf4j_slf4j_simple",
    artifact = "org.slf4j:slf4j-simple:1.7.25",
)

# GO DEPS
go_repository(
    name = "com_github_bsm_sarama_cluster",
    commit = "cf455bc755fe41ac9bb2861e7a961833d9c2ecc3",
    importpath = "github.com/bsm/sarama-cluster",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "346938d642f2ec3594ed81d874461961cd0faa76",
    importpath = "github.com/davecgh/go-spew",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    commit = "ea41b0fad31007accc7f806884dcdf3da98b79ce",
    importpath = "github.com/eapache/go-resiliency",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    commit = "040cc1a32f578808623071247fdbd5cc43f37f5f",
    importpath = "github.com/eapache/go-xerial-snappy",
)

go_repository(
    name = "com_github_eapache_queue",
    commit = "44cc805cf13205b55f69e14bcb69867d1ae92f98",
    importpath = "github.com/eapache/queue",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "b4deda0973fb4c70b50d226b1af49f3da59f5265",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_golang_snappy",
    commit = "2e65f85255dbc3072edf28d6b5b8efc472979f5a",
    importpath = "github.com/golang/snappy",
)

go_repository(
    name = "com_github_hashicorp_consul",
    commit = "39f93f011e591c842acc8053a7f5972aa6e592fd",
    importpath = "github.com/hashicorp/consul",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    commit = "d5fe4b57a186c716b0e00b8c301cbd9b4182694d",
    importpath = "github.com/hashicorp/go-cleanhttp",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    commit = "6bb64b370b90e7ef1fa532be9e591a81c3493e00",
    importpath = "github.com/hashicorp/go-rootcerts",
)

go_repository(
    name = "com_github_hashicorp_serf",
    commit = "d6574a5bb1226678d7010325fb6c985db20ee458",
    importpath = "github.com/hashicorp/serf",
)

go_repository(
    name = "com_github_jinzhu_gorm",
    commit = "6ed508ec6a4ecb3531899a69cbc746ccf65a4166",
    importpath = "github.com/jinzhu/gorm",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    commit = "04140366298a54a039076d798123ffa108fff46c",
    importpath = "github.com/jinzhu/inflection",
)

go_repository(
    name = "com_github_lib_pq",
    commit = "90697d60dd844d5ef6ff15135d0203f65d2f53b8",
    importpath = "github.com/lib/pq",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    commit = "3864e76763d94a6df2f9960b16a20a33da9f9a66",
    importpath = "github.com/mitchellh/go-homedir",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "f15292f7a699fcc1a38a80977f80a046874ba8ac",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_pierrec_lz4",
    commit = "1958fd8fff7f115e79725b1288e0b878b3e06b00",
    importpath = "github.com/pierrec/lz4",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "645ef00459ed84a119197bfb8d8205042c6df63d",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    commit = "e2704e165165ec55d062f5919b4b29494e9fa790",
    importpath = "github.com/rcrowley/go-metrics",
)

go_repository(
    name = "com_github_satori_go_uuid",
    commit = "36e9d2ebbde5e3f13ab2e25625fd453271d6522e",
    importpath = "github.com/satori/go.uuid",
)

go_repository(
    name = "com_github_shopify_sarama",
    commit = "35324cf48e33d8260e1c7c18854465a904ade249",
    importpath = "github.com/Shopify/sarama",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "fedd2861243fd1a8152376292b921b394c7bef7e",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "168a6198bcb0ef175f7dacec0b8691fc141dc9b8",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_net",
    commit = "a680a1efc54dd51c040b3b5ce4939ea3cf2ea0d1",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)
