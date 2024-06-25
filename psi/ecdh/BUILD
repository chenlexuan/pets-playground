load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "ecdh_psi",
    srcs = ["main.go"],
    deps = [
            "@org_openmined_psi//private_set_intersection/go/server",
            "@org_openmined_psi//private_set_intersection/go/client",
            "@org_openmined_psi//private_set_intersection/go/datastructure",
            "@org_openmined_psi//private_set_intersection/proto:psi_go_proto",
            "@org_golang_google_protobuf//proto:go_default_library",
            ],
)