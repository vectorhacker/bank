package(default_visibility=["//visibility:public"])

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_binary")

scala_binary(
    name = "Projections",
    srcs = glob(["src/main/java/com/bank/service/projections/*.java"]),
    deps = [
        "@com_google_protobuf_protobuf_java//jar",
        "@joda_time_joda_time//jar",
        "@org_joda_joda_convert//jar",
        "@org_reactivestreams_reactive_streams//jar",
        "@com_typesafe_akka_akka_parsing//jar",
        "@com_typesafe_akka_akka_http_core//jar",
        "@com_typesafe_akka_akka_http//jar",
        "@com_typesafe_akka_akka_actor//jar",
        "@com_typesafe_akka_akka_stream//jar",
        "@com_typesafe_config//jar",
        "@com_geteventstore_eventstore_client//jar",
        "@org_json_json//jar",
        "@org_apache_kafka_kafka_clients//jar",
        "@org_slf4j_slf4j_api//jar",
        "@org_slf4j_slf4j_simple//jar",
    ],
    main_class = "com.bank.service.projections.Projections",
)
