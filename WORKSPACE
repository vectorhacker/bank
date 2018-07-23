load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
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
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()


# docker rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "6dede2c65ce86289969b907f343a1382d33c14fbce5e30dd17bb59bb55bb6593",
    strip_prefix = "rules_docker-0.4.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.4.0.tar.gz"],
)


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

## Java deps
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

# Scala compiler
rules_scala_version="64faf06a4932a9a1d3378b6ba1a6d77479cefef3" # update this as needed

http_archive(
    name = "io_bazel_rules_scala",
    url = "https://github.com/bazelbuild/rules_scala/archive/%s.zip"%rules_scala_version,
    type = "zip",
    strip_prefix= "rules_scala-%s" % rules_scala_version,
)

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_repositories")

scala_repositories(("2.12.6", {
    "scala_compiler": "3023b07cc02f2b0217b2c04f8e636b396130b3a8544a8dfad498a19c3e57a863",
    "scala_library": "f81d7144f0ce1b8123335b72ba39003c4be2870767aca15dd0888ba3dab65e98",
    "scala_reflect": "ffa70d522fc9f9deec14358aa674e6dd75c9dfa39d4668ef15bb52f002ce99fa"
}))

load("@io_bazel_rules_scala//scala:toolchains.bzl", "scala_register_toolchains")
scala_register_toolchains()
