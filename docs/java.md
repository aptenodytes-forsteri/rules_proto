---
layout: default
title: java
permalink: examples/java
parent: examples
---


# java example

`bazel test //example/golden:java_test`


## `BUILD.bazel` (after gazelle)

~~~python
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@build_stack_rules_proto//rules:proto_compile.bzl", "proto_compile")

# "proto_rule" instantiates the proto_compile rule
# gazelle:proto_rule proto_compile implementation stackb:rules_proto:proto_compile

# "proto_plugin" instantiates the builtin java plugin
# gazelle:proto_plugin java implementation builtin:java

# "proto_language" binds the rule(s) and plugin(s) together
# gazelle:proto_language builtins plugin java

proto_library(
    name = "example_proto",
    srcs = ["example.proto"],
    visibility = ["//visibility:public"],
)

proto_compile(
    name = "example_builtins_compile",
    outputs = [
        "example.srcjar",
    ],
    plugins = ["@build_stack_rules_proto//plugin/builtin:java"],
    proto = "example_proto",
)
~~~


## `BUILD.bazel` (before gazelle)

~~~python
# "proto_rule" instantiates the proto_compile rule
# gazelle:proto_rule proto_compile implementation stackb:rules_proto:proto_compile

# "proto_plugin" instantiates the builtin java plugin
# gazelle:proto_plugin java implementation builtin:java

# "proto_language" binds the rule(s) and plugin(s) together
# gazelle:proto_language builtins plugin java
~~~


## `WORKSPACE`

~~~python
local_repository(
    name = "build_stack_rules_proto",
    path = "../build_stack_rules_proto",
)

register_toolchains("@build_stack_rules_proto//toolchain")

# == Externals ==

load("@build_stack_rules_proto//deps:core_deps.bzl", "core_deps")

core_deps()

# == Go ==

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.16.2")

# == Gazelle ==

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# == Protobuf ==

load("@build_stack_rules_proto//deps:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
~~~
