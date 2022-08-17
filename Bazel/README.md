# Bazel
![Bazel](https://ik.imagekit.io/pibjyepn7p9/Bazel_Build_Process_2x_BYDtPQkdR.png?updatedAt=1633191206254)  
Bazel is Google's own build tool for the automation of building and testing of large software repositories. Source files are organized in a nested hierarchy of packages. Bazel makes use of existing [make variables](https://bazel.build/reference/be/make-variables#predefined_genrule_variables) such as ```$@```, ```$<```.

## Terms
### Workspace
Workspace is a root directory of pacakges and is considered as a unit of applications. Directories containing ```WORKSPACE``` file are considered the root of a workspace. It specifies external dependencies for building a workspace.
- Bazel ignores subdirectory trees containing a ```WORKSPACE``` file.
- ```WORKSPACE.bazel``` is prioritized thans ```WORKSPACE``` file 
- The root directory(repository) is called ```@```

### Package
Package is the subdirectory of a workspace containing ```BUILD``` file. It defines **targets** which are files and rules. Rules acts as a header file in cpp and defines required input/output files. Rules can be used in other pacakges.

### Labels
The name of a target(rules, files) is label. 
- labels can be omitted if the name matches with last component(directory)
```bazel
# Execute rule 'greet' in root/greet's BUILD file
//greet:greet
//greet
```

## BUILD
Bazel creates pre-defined folders such as ```bazel-Bazel```, ```bazel-bin```, ```bazel-out```, ```bazel-testlogs```. Output files are located in bazel-bin with its corresponding folders. 
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

## Golang
### Snippet
#### WORKSPACE
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

#### BUILD
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

### Compilation
#### go_library
- go_library makes a package into a libarary that can be used as a dependency 
```make
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "greet_library",
    srcs = [
        "server/greet.go",
        "server/main.go",
        "client/greet.go",
        "client/main.go"
    ],
	# prepix for package path that all srcs share
    importpath = "github.com/Cotidie/STEADY/gRPC/Basics/greet",
    visibility = ["//visibility:public"],
)
```

#### go_binary

#### go_repository

#### go_test

### Protocol Buffer, gRPC
#### proto_library
Defines proto files to be used in various packages. Once declared, it can be listed in the deps cluase of supported rules such as go_proto_library.
```make
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    
)
```
#### go_proto_library

