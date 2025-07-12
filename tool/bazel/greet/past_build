load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

# go-library
go_binary(
    name = "greet",
    srcs = [
        "server/greet.go",
        "server/main.go",
        "client/greet.go",
        "client/main.go"
    ],
    deps = [
        ":greet_go_proto",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
	# prepix for package path that all srcs share
    importpath = "github.com/Cotidie/STEADY/gRPC/Basics/greet",
    visibility = ["//visibility:public"],
)

# proto-library
proto_library(
    name = "greet_proto",
    srcs = ["proto/greet.proto"],
    visibility = ["//visibility:public"],
)
# go-proto-library
go_proto_library(
    name = "greet_go_proto",
    protos = [":greet_proto"],
    importpath = "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto",
)