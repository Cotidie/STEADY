# Bazel
![Bazel](https://ik.imagekit.io/pibjyepn7p9/Bazel_Build_Process_2x_BYDtPQkdR.png?updatedAt=1633191206254)  
Bazel is Google's own build tool for the automation of building and testing of large software repositories. Source files are organized in a nested hierarchy of packages. Bazel makes use of existing [make variables](https://bazel.build/reference/be/make-variables#predefined_genrule_variables) such as ```$@```, ```$<```, ```$n```.

## Table of Contents
  - [Terms](#terms)
    - [Workspace](#workspace)
    - [Package](#package)
    - [Labels](#labels)
  - [General Rules](#general-rules)
    - [Load](#load)
    - [Genrule](#genrule)
  - [Gazelle](#gazelle)
    - [WORKSPACE](#workspace-1)
    - [BUILD](#build)
  - [Golang](#golang)
    - [WORKSPACE](#workspace-2)
    - [BUILD](#build-1)
      - [go_library](#go_library)
      - [go_binary](#go_binary)
      - [go_repository](#go_repository)
    - [Protocol Buffer, gRPC](#protocol-buffer-grpc)
      - [proto_library](#proto_library)
      - [go_proto_library](#go_proto_library)

## Terms
### Workspace
Workspace is a root directory of pacakges and is considered as a unit of applications. Directories containing ```WORKSPACE``` file are considered the root of a workspace. It specifies external dependencies for building a workspace.
- Bazel ignores subdirectory trees containing a ```WORKSPACE``` file.
- ```WORKSPACE.bazel``` is prioritized than ```WORKSPACE``` file 
- The root directory(repository) is called ```@```
- external repositories are defined here and can be used with ```@```

### Package
Package is the subdirectory of a workspace containing ```BUILD``` file. It defines **targets** which are files and rules. Rules acts as a header file in cpp and defines required input/output files. Rules can be used in other pacakges.
- libraries can be embedded if importpath(package) is the same, which creates a single archive file.

### Labels
The name of a target(rules, files) is label. Labels can be omitted if the name matches with last component(directory).
```bazel
# Execute rule 'greet' in root/greet's BUILD file
# below two are identical
bazel run @//greet:greet
bazel run //greet
```

## General Rules
Bazel creates pre-defined folders(symlinks) such as ```bazel-Bazel```, ```bazel-bin```, ```bazel-out```, ```bazel-testlogs```. Output files are located in bazel-bin with its corresponding folders. 
### Load
```load``` statement imports a new symbol from an ```.bzl``` extension.
```make
# imports http.bzl as http_archive from @bazel_tools repository
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# use newly added symbol
http_archive(
	...
)
```

### Genrule
| https://bazel.build/reference/be/general  
 Genereates one or more files using a Bash command
```make
# directory: root/greet/BUILD
genrule(
    name = "greet",
    srcs = [],
    outs = ["greet2.txt"],
	# $(locations) means 'bazel-bin/greet/' in this context
    cmd = "sleep 2 && echo 'Greet BUILD' > $(locations greet2.txt)",
)
```

## Gazelle
Gazelle creates BUILD files automatically in each package folder. It detects source files to define packages.
### WORKSPACE
```make
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "685052b498b6ddfe562ca7a97736741d87916fe536623afb7da2824c0211c369",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
    ],
)


load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.18.3")

gazelle_dependencies()
```

### BUILD
```make
load("@bazel_gazelle//:def.bzl", "gazelle")

# make packages
gazelle(
    name = "gazelle",
    prefix = "<package/module prefix>"
)

# detect external dependencies
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```


## Golang
### WORKSPACE
| https://github.com/bazelbuild/rules_go#setup  
Add the snippet below to use Go toolchains and gazelle. Gazelle is for automatic management over dependencies of source codes. WORKSPACE defines dependencies for packages in a project.
- For protocol buffer dependencies: [see](https://github.com/bazelbuild/rules_proto)
- For gRPC dependencies: [see](https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#grpc-dependencies)
```make
# //WORKSPACE
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "16e9fca53ed6bd4ff4ad76facc9b7b651a89db1689a2877d6fd7b82aa824e366",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.18.2")

gazelle_dependencies()
```

### BUILD
BUILD files define rules related to compilation or bazel commands. Example below is for auto dependency management with gazelle. it updates or creates deps.bzl macro.
```make
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/example/project
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```

#### go_library
```go_library``` makes a package into a libarary that can be used as a dependency 
```make
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "server_lib",
    srcs = [
        "greet.go",
        "main.go",
    ],
    importpath = "github.com/Cotidie/STEADY/gRPC/Basics/greet/server",
    visibility = ["//visibility:private"],
    deps = [
        "//greet/proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
```

#### go_binary
```go_binary``` is the final destination with libraries as its dependencies. 
```make

```
#### go_repository

### Protocol Buffer, gRPC
#### proto_library
Defines proto files to be used in various packages. Once declared, it can be listed in the deps cluase of supported rules such as go_proto_library.
```make
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "proto_proto",
    srcs = ["greet.proto"],
    visibility = ["//visibility:public"],
)

```
#### go_proto_library
Compiles .proto files into .go files. if .proto has service declaration, gRPC compiler should be specified. This creates archive files, hence can be imported in libraries.
```make
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "proto_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto",
    proto = ":proto_proto",
    visibility = ["//visibility:public"],
)
```

## Troubleshoot
### \<filename\>: undefined: \<member\>
```bash
# Error Message
tracing/tracer.go:29:27: undefined: Option
tracing/tracer.go:30:12: undefined: newConfig
tracing/tracer.go:78:25: undefined: Config
tracing/tracer.go:83:7: undefined: Stdout
tracing/tracer.go:98:25: undefined: Config
```
- Context: When referencing a file in the same pacakge
- Solution: Adds all the files in the same pacakge in `srcs` field in BUILD file

